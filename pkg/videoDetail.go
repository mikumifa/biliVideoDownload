package pkg

// Dash的格式，目前只使用Dash
type VideoDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Dash              struct {
			Duration       int     `json:"duration"`
			MinBufferTime  float64 `json:"minBufferTime"`
			MinBufferTime1 float64 `json:"min_buffer_time"`
			Video          []struct {
				Id            int      `json:"id"`
				BaseUrl       string   `json:"baseUrl"`
				BaseUrl1      string   `json:"base_url"`
				BackupUrl     []string `json:"backupUrl"`
				BackupUrl1    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType1     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate1    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap1 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase1 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"video"`
			Audio []struct {
				Id            int      `json:"id"`
				BaseUrl       string   `json:"baseUrl"`
				BaseUrl1      string   `json:"base_url"`
				BackupUrl     []string `json:"backupUrl"`
				BackupUrl1    []string `json:"backup_url"`
				Bandwidth     int      `json:"bandwidth"`
				MimeType      string   `json:"mimeType"`
				MimeType1     string   `json:"mime_type"`
				Codecs        string   `json:"codecs"`
				Width         int      `json:"width"`
				Height        int      `json:"height"`
				FrameRate     string   `json:"frameRate"`
				FrameRate1    string   `json:"frame_rate"`
				Sar           string   `json:"sar"`
				StartWithSap  int      `json:"startWithSap"`
				StartWithSap1 int      `json:"start_with_sap"`
				SegmentBase   struct {
					Initialization string `json:"Initialization"`
					IndexRange     string `json:"indexRange"`
				} `json:"SegmentBase"`
				SegmentBase1 struct {
					Initialization string `json:"initialization"`
					IndexRange     string `json:"index_range"`
				} `json:"segment_base"`
				Codecid int `json:"codecid"`
			} `json:"audio"`
			Dolby struct {
				Type  int         `json:"type"`
				Audio interface{} `json:"audio"`
			} `json:"dolby"`
			Flac interface{} `json:"flac"`
		} `json:"dash"`
		SupportFormats []struct {
			Quality        int      `json:"quality"`
			Format         string   `json:"format"`
			NewDescription string   `json:"new_description"`
			DisplayDesc    string   `json:"display_desc"`
			Superscript    string   `json:"superscript"`
			Codecs         []string `json:"codecs"`
		} `json:"support_formats"`
		HighFormat   interface{} `json:"high_format"`
		LastPlayTime int         `json:"last_play_time"`
		LastPlayCid  int         `json:"last_play_cid"`
		ViewInfo     interface{} `json:"view_info"`
	} `json:"data"`
}
