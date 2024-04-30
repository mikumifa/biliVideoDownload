package pkg

// DownloadOption
// @Description:  下载的选项
type VideoType string

type FnvalOptions int

const (
	FLV         FnvalOptions = 1 << iota // 1：FLV 格式，已下线
	MP4                                  // 2：MP4 格式，已下线
	DASH                                 // 4：DASH 格式
	HDR                                  // 8：是否需求 HDR 视频
	FourK                                // 16：是否需求 4K 分辨率
	DolbyAudio                           // 32：是否需求杜比音频
	DolbyVision                          // 64：是否需求杜比视界
	EightK                               // 128：是否需求 8K 分辨率
	AV1                                  // 256：是否需求 AV1 编码
)

type DownloadOption struct {
	Qn    QnOptions    //视频清晰度标识
	Fnval FnvalOptions //视频流格式标识
	Fourk FourkOptions //fnval&128=128且fourk=1 实现4k
}

// func (d *DownloadOption) Check() bool {}
func (d *DownloadOption) Url() string {
	return "https://api.bilibili.com/x/player/wbi/playurl"
}

// QnOptions represents the options for qn.
type QnOptions int

type FourkOptions int

const (
	Fourk1080p = 0
	Fourk4k    = 1
)

// 这部分暂时无用
const (
	Qn240P        QnOptions = 6   // 240P 极速
	Qn360P        QnOptions = 16  // 360P 流畅
	Qn480P        QnOptions = 32  // 480P 清晰
	Qn720P        QnOptions = 64  // 720P 高清
	Qn720P60      QnOptions = 74  // 720P60 高帧率
	Qn1080P       QnOptions = 80  // 1080P 高清
	Qn1080PPlus   QnOptions = 112 // 1080P+ 高码率
	Qn1080P60     QnOptions = 116 // 1080P60 高帧率
	Qn4K          QnOptions = 120 // 4K 超清
	QnHDR         QnOptions = 125 // HDR 真彩色
	QnDolbyVision QnOptions = 126 // 杜比视界
	Qn8K          QnOptions = 127 // 8K 超高清
)

// OptionFn is a function type for setting DownloadOption values.
type OptionFn func(*DownloadOption)

// NewDownloadOption
//
//	@Description: 使用默认来提高速度
//	@param optionFn
//	@return *DownloadOption
//	@return error
func NewDownloadOption(optionFn ...func(*DownloadOption)) *DownloadOption {

	downloadOption := DownloadOption{
		Qn:    0,  //视频清晰度标识
		Fourk: 1,  //视频流格式标识
		Fnval: 16, //fnval&128=128且fourk=1 实现4k
	}

	for _, option := range optionFn {
		option(&downloadOption)
	}
	return &downloadOption
}

// WithQn sets the qn value.
func WithQn(options ...QnOptions) OptionFn {
	return func(opt *DownloadOption) {
		for _, option := range options {
			opt.Qn |= option
		}
	}
}

// WithFnval sets the fnval value using OR operation for multiple options.
func WithFnval(options ...FnvalOptions) OptionFn {
	return func(opt *DownloadOption) {
		for _, option := range options {
			opt.Fnval |= option
		}
	}
}

// WithFourk sets the fourk value.
func WithFourk(fourk FourkOptions) OptionFn {
	return func(opt *DownloadOption) {
		opt.Fourk = fourk
	}
}
