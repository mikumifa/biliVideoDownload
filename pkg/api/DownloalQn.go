package api

type DownloadQn uint32

// Qn 7 bit
// Audio 15 bit
// 设为0表示不需要
// AUTO 表示选最新的
const (
	AUTOQn             DownloadQn = 7
	Qn240PFirst        DownloadQn = 6                         // 240P 极速 110
	Qn360PFirst        DownloadQn = 1 << 4                    // 360P 流畅 10000
	Qn480PFirst        DownloadQn = 1 << 5                    // 480P 清晰
	Qn720PFirst        DownloadQn = 1 << 6                    // 720P 高清
	Qn720P60First      DownloadQn = 1<<6 + 1<<3 + 1<<1        // 720P60 高帧率
	Qn1080PFirst       DownloadQn = 1<<6 + 1<<4               // 1080P 高清
	Qn1080PPlusFirst   DownloadQn = 1<<6 + 1<<5 + 1<<4        // 1080P+ 高码率
	Qn1080P60First     DownloadQn = 1<<6 + 1<<5 + 1<<4 + 1<<2 // 1080P60 高帧率
	Qn4KFirst          DownloadQn = 120                       // 4K 超清
	QnHDRFirst         DownloadQn = 125                       // HDR 真彩色
	QnDolbyVisionFirst DownloadQn = 126                       // 杜比视界
	Qn8KFirst          DownloadQn = 127                       // 8K 超高清
	AUTOAudio          DownloadQn = 7
	Audio64k           DownloadQn = 2<<13 + 2<<12 + 2<<11 + 2<<9 + 2<<8 + 2<<2 //64K
	Audio132K          DownloadQn = 30232                                      //132K
	AudioDubi          DownloadQn = 30250                                      //杜比全景声
	AudioHiRes         DownloadQn = 30251                                      //Hi-Res无损
	Audio192K          DownloadQn = 30280                                      //192K
)

func MakeQn(audio DownloadQn, video DownloadQn) DownloadQn {
	return audio<<7 | video
}
func GetVideo(way DownloadQn) DownloadQn {
	return way & 0x0000007f
}
func GetAudio(way DownloadQn) DownloadQn {
	return way >> 7
}
