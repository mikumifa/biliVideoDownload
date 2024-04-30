package download

import (
	"biliVideoDownload/pkg/http_client"
	"biliVideoDownload/pkg/utils"
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// DefaultDownloadPartSize is the default range of bytes to get at a time when
// using Download().
const DefaultDownloadPartSize = 1024 * 1024 * 4

func getDefaultDownloadConcurrency() int {
	return runtime.NumCPU()
}

// DefaultPartBodyMaxRetries is the default number of retries to make when a part fails to download.
const DefaultPartBodyMaxRetries = 3

// 不保证里面线程安全
type Downloader struct {
	// PartSize is ignored if the Range input parameter is provided.
	PartSize int64
	// PartBodyMaxRetries is the number of retry attempts to make for failed part downloads.
	PartBodyMaxRetries int
	Concurrency        int
	Client             *http_client.HttpClient
	Url                string
}

// NewDownloader
//
//	@Description:穿件
//	@param c
//	@param options 创建的选项
//	@param url 下载的地址
//	@return *Downloader 下载器
func NewDownloader(c *http_client.HttpClient, options ...func(*Downloader)) *Downloader {

	d := &Downloader{
		Client:             c,
		PartSize:           DefaultDownloadPartSize,
		PartBodyMaxRetries: DefaultPartBodyMaxRetries,
		Concurrency:        getDefaultDownloadConcurrency(),
	}
	for _, option := range options {
		option(d)
	}

	return d
}

func WithUrl(url string) func(*Downloader) {
	return func(hc *Downloader) {
		hc.Url = url
	}
}

// Download
//
//	@Description: 会拷贝一份来保证线程安全, 各个download之间互不影响
//	@receiver d
//	@param w
//	@param options
//	@return n
//	@return err
func (d Downloader) Download(w io.WriterAt, url string, options ...func(*Downloader)) (n int64, err error) {

	impl := downloader{w: w, cfg: d}
	impl.cfg.Url = url

	for _, option := range options {
		option(&impl.cfg)
	}
	impl.partBodyMaxRetries = d.PartBodyMaxRetries

	impl.totalBytes = -1
	if impl.cfg.Concurrency == 0 {
		impl.cfg.Concurrency = getDefaultDownloadConcurrency()
	}

	if impl.cfg.PartSize == 0 {
		impl.cfg.PartSize = DefaultDownloadPartSize
	}

	return impl.download()
}
func (d Downloader) DownloadTmp(url string) (string, error) {
	tmpDir := "tmp" // 临时文件夹路径
	path := utils.CreateUniqueFileName("flv")
	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		logrus.Error(err)
		return "", err
	}
	path = filepath.Join(tmpDir, path)
	file, err := os.Create(path)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	defer file.Close()
	_, err = d.Download(file, url)
	if err != nil {
		return "", err
	}
	return path, nil
}

// downloader is the implementation structure used internally by Downloader.
type downloader struct {
	cfg Downloader

	w                  io.WriterAt
	bar                *progressbar.ProgressBar
	wg                 sync.WaitGroup
	m                  sync.Mutex
	pos                int64
	totalBytes         int64
	written            int64
	err                error
	partBodyMaxRetries int
}

// download
//
//	@Description: 当存在range参数时候，根据range去下载， downloader的witten和err作为结果
//	@receiver d
//	@return n
//	@return err
func (d *downloader) download() (n int64, err error) {
	d.getChunk()
	if total := d.getTotalBytes(); total >= 0 {
		ch := make(chan dlchunk, d.cfg.Concurrency)
		for i := 0; i < d.cfg.Concurrency; i++ {
			d.wg.Add(1)
			go d.downloadPart(ch)
		}
		for d.getErr() == nil {
			if d.pos >= total {
				break // We're finished queuing chunks
			}
			ch <- dlchunk{w: d.w, start: d.pos, size: d.cfg.PartSize}
			d.pos += d.cfg.PartSize
		}
		close(ch)
		d.wg.Wait()
	} else {
		for d.err == nil {
			d.getChunk()
		}
		var responseError interface {
			HTTPStatusCode() int
		}
		if errors.As(d.err, &responseError) {
			if responseError.HTTPStatusCode() == http.StatusRequestedRangeNotSatisfiable {
				d.err = nil
			}
		}
	}

	// Return error
	return d.written, d.err
}

func (d *downloader) downloadPart(ch chan dlchunk) {
	defer d.wg.Done()
	for {
		chunk, ok := <-ch
		if !ok {
			break
		}
		if d.getErr() != nil {
			continue
		}

		if err := d.downloadChunk(chunk); err != nil {
			d.setErr(err)
		}
	}
}

// getChunk grabs a chunk of data from the body.
// Not thread safe. Should only used when grabbing data on a single thread.
func (d *downloader) getChunk() {
	if d.getErr() != nil {
		return
	}

	chunk := dlchunk{w: d.w, start: d.pos, size: d.cfg.PartSize}
	d.pos += d.cfg.PartSize

	if err := d.downloadChunk(chunk); err != nil {
		d.setErr(err)
	}
}

// downloadRange
//
//	@Description: 通过string的rng下载相应的chunck，最后设置d.pos和d.written
//	@receiver d
//	@param rng 范围
func (d *downloader) downloadRange(rng string) {
	if d.getErr() != nil {
		return
	}

	chunk := dlchunk{w: d.w, start: d.pos}
	chunk.rangeStr = rng
	if err := d.downloadChunk(chunk); err != nil {
		d.setErr(err)
	}
	d.pos = d.written
}

// downloadChunk
//
//	@Description: 下载具体的块，request会进行请求,请求取决于d.cfg.url和d.headInput
//	@receiver d
//	@param chunk
//	@return error
func (d *downloader) downloadChunk(chunk dlchunk) error {
	var err error
	//添加尝试的次数
	for retry := 0; retry <= d.partBodyMaxRetries; retry++ {
		_, err = d.tryDownloadChunk(&chunk, func() (*http.Response, error) {
			return d.cfg.Client.Do(http_client.GET, d.cfg.Url, func(header *http.Header) {
				header.Set("range", chunk.ByteRange())
			}, nil, nil)
		})
		if err == nil {
			break
		}
	}
	return err
}

type FnReq func() (*http.Response, error)

// tryDownloadChunk
//
//	@Description:
//	@receiver d
//	@param w Writer自带偏移量，会在制定偏移量重写
//	@param fnReq
//	@return int64
//	@return error
func (d *downloader) tryDownloadChunk(w io.Writer, fnReq FnReq) (int64, error) {

	resp, err := fnReq()
	if err != nil {
		return 0, err
	}
	d.setTotalBytes(resp) // Set total if not yet set.
	n, _ := io.Copy(w, resp.Body)
	_ = d.bar.Add64(n)
	err = resp.Body.Close()
	return n, nil
}

// getTotalBytes is a thread-safe getter for retrieving the total byte status.
func (d *downloader) getTotalBytes() int64 {
	d.m.Lock()
	defer d.m.Unlock()
	return d.totalBytes
}

func (d *downloader) setTotalBytes(resp *http.Response) {
	d.m.Lock()
	defer d.m.Unlock()
	if d.totalBytes >= 0 {
		return
	}
	if resp.Header.Get("content-range") == "" {
		if int64(resp.ContentLength) > 0 {
			d.totalBytes = resp.ContentLength
		}
	} else {
		///Content-Range: bytes 29752736-31404094/225612510
		parts := strings.Split(resp.Header.Get("content-range"), "/")

		total := int64(-1)
		var err error
		totalStr := parts[len(parts)-1]
		if totalStr != "*" {
			total, err = strconv.ParseInt(totalStr, 10, 64)
			if err != nil {
				d.err = err
				return
			}
		}
		d.totalBytes = total
	}
	d.bar = progressbar.DefaultBytes(d.totalBytes, "downloading")

}

// getErr is a thread-safe getter for the error object
func (d *downloader) getErr() error {
	d.m.Lock()
	defer d.m.Unlock()
	return d.err
}

// setErr is a thread-safe setter for the error object
func (d *downloader) setErr(e error) {
	d.m.Lock()
	defer d.m.Unlock()
	d.err = e
}

// dlchunk represents a single chunk of data to write by the worker routine.
// This structure also implements an io.SectionReader style interface for
// io.WriterAt, effectively making it an io.SectionWriter (which does not
// exist).
type dlchunk struct {
	w     io.WriterAt
	start int64
	size  int64
	cur   int64

	rangeStr string
}

func (c *dlchunk) Write(p []byte) (n int, err error) {
	if c.cur >= c.size && len(c.rangeStr) == 0 {
		return 0, io.EOF
	}

	n, err = c.w.WriteAt(p, c.start+c.cur)
	c.cur += int64(n)
	return
}

// ByteRange returns a HTTP Byte-Range header value that should be used by the
// client to request the chunk's range.
func (c *dlchunk) ByteRange() string {
	if len(c.rangeStr) != 0 {
		return c.rangeStr
	}
	return fmt.Sprintf("bytes=%d-%d", c.start, c.start+c.size-1)
}
