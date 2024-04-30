package api

import (
	. "biliVideoDownload/pkg"
	"biliVideoDownload/pkg/config"
	"biliVideoDownload/pkg/download"
	. "biliVideoDownload/pkg/http_client"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"strconv"
)

type RequestWay int
type FnChooseWay func(detail *VideoDetail) (audioUrl string, videoUrl string, err error)

const (
	Part RequestWay = iota
	Single
)

func videoInfo(client *HttpClient, bvid string) (*VideoInfo, error) {
	url := "https://api.bilibili.com/x/web-interface/view"
	res, err := client.Do(GET, url, nil, map[string]string{
		"bvid": bvid,
	}, nil)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	videoInfo := &VideoInfo{}
	err = json.Unmarshal(body, videoInfo)
	if err != nil {
		logrus.Error("json 格式错误", body, err)
		return nil, err
	}
	return videoInfo, nil
}

// DownloadVideo
// @Description: 根据 DownloadOption 下载视频
// @param v 视频的info
// @param dlOpt  下载的配置
// @return []byte 视频的数据
// @return error
func downloadVideo(client *HttpClient, v *VideoInfo, dlOpt *DownloadOption) (*VideoDetail, error) {
	url := dlOpt.Url()
	queryMap := map[string]string{
		"bvid":  v.Data.Bvid,
		"qn":    strconv.Itoa(int(dlOpt.Qn)),
		"fnval": strconv.Itoa(int(dlOpt.Fnval)),
		"fourk": strconv.Itoa(int(dlOpt.Fourk)),
		"cid":   strconv.Itoa(v.Data.Cid),
	}
	res, err := client.Do(GET, url, nil, queryMap, nil)
	if err != nil {
		logrus.Error("获取视频分段失败", url, err)
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	detail := VideoDetail{}
	err = json.Unmarshal(body, &detail)
	if err != nil {
		logrus.Error("获取视频分段失败", url, err)
		return nil, err
	}
	return &detail, nil
}

func Download(bvid string, filename string, way RequestWay, qn DownloadQn, fnChooseWay FnChooseWay) {
	httpClient := NewHttpClient(WithCookie(config.Get().HttpConfig.Cookie))
	info, err := videoInfo(httpClient, bvid)
	if err != nil {
		logrus.Error(err)
	}
	video, err := downloadVideo(httpClient, info, NewDownloadOption())
	if err != nil {
		logrus.Error(err)
	}
	audioUrl, videoUrl, err := chooseUrl(video, qn, fnChooseWay)
	if err != nil {
		logrus.Error(err)
		return
	}

	switch way {
	case Part:
		downloader := download.NewDownloader(httpClient)

		logrus.Infof("Part Video download...")
		videoPath, err := downloader.DownloadTmp(videoUrl)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Infof("Part Audio download...")
		audioPath, err := downloader.DownloadTmp(audioUrl)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Infof(audioPath, videoPath)
		err = Merge(audioPath, videoPath, filename)
		if err != nil {
			logrus.Error(err)
			return
		}
		break
	case Single:
		//TODO
		//logrus.Infof("Single download...")
		//resp, _ := httpClient.Do(GET, videoUrl, nil, nil, nil)
		//bar := progressbar.DefaultBytes(
		//	resp.ContentLength,
		//	"downloading",
		//)
		//_, _ = io.Copy(io.MultiWriter(file, bar), resp.Body)
		//resp, _ = httpClient.Do(GET, audioUrl, nil, nil, nil)
		//bar = progressbar.DefaultBytes(
		//	resp.ContentLength,
		//	"downloading",
		//)
		//_, _ = io.Copy(io.MultiWriter(file, bar), resp.Body)
		break

	}
	if err != nil {
		logrus.Error(err)
		return
	}

}

// chooseUrl
//
//	@Description: 没找到切制定fn时候回去调用选择函数，否则没找到就会去下载最新的
//	@param detail
//	@param qn
//	@param fnChooseWay
//	@return audioUrl
//	@return videoUrl
//	@return err
func chooseUrl(detail *VideoDetail, qn DownloadQn, fnChooseWay FnChooseWay) (audioUrl string, videoUrl string, err error) {
	audio := GetAudio(qn)
	video := GetVideo(qn)
	isNeedLastFilled := false
	if video != 0 {
		found := false
		for _, v := range detail.Data.Dash.Video {
			if v.Id == int(video) {
				videoUrl = v.BaseUrl
				found = true
			}
		}
		isNeedLastFilled = isNeedLastFilled && found
	}
	if audio != 0 {
		found := false
		for _, v := range detail.Data.Dash.Audio {
			if v.Id == int(audio) {
				audioUrl = v.BaseUrl
				found = true
			}
		}
		isNeedLastFilled = isNeedLastFilled && found
	}
	// 全是0 根据情况来定
	if audio == 0 && video == 0 || isNeedLastFilled {
		if fnChooseWay == nil {
			audioUrl = detail.Data.Dash.Audio[0].BaseUrl
			videoUrl = detail.Data.Dash.Video[0].BaseUrl
			return
		} else {
			return fnChooseWay(detail)
		}
	}
	return
}

func Merge(video, audio, output string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", video, "-i", audio, "-c", "copy", output)
	err := cmd.Run()
	if err != nil {
		return err
	}
	logrus.Infof("Merged video")
	err = os.Remove(video)
	if err != nil {
		return err
	}
	err = os.Remove(audio)
	if err != nil {
		return err
	}
	return nil
}
