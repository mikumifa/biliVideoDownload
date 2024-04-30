package pkg

type VideoInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizId   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State     int `json:"state"`
		Duration  int `json:"duration"`
		MissionId int `json:"mission_id"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		ArgueInfo struct {
			ArgueMsg  string `json:"argue_msg"`
			ArgueType int    `json:"argue_type"`
			ArgueLink string `json:"argue_link"`
		} `json:"argue_info"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonId           int         `json:"season_id"`
		Premiere           interface{} `json:"premiere"`
		TeenageMode        int         `json:"teenage_mode"`
		IsChargeableSeason bool        `json:"is_chargeable_season"`
		IsStory            bool        `json:"is_story"`
		IsUpowerExclusive  bool        `json:"is_upower_exclusive"`
		IsUpowerPlay       bool        `json:"is_upower_play"`
		IsUpowerPreview    bool        `json:"is_upower_preview"`
		EnableVt           int         `json:"enable_vt"`
		VtDisplay          string      `json:"vt_display"`
		NoCache            bool        `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool `json:"allow_submit"`
			List        []struct {
				Id          int64  `json:"id"`
				Lan         string `json:"lan"`
				LanDoc      string `json:"lan_doc"`
				IsLock      bool   `json:"is_lock"`
				SubtitleUrl string `json:"subtitle_url"`
				Type        int    `json:"type"`
				IdStr       string `json:"id_str"`
				AiType      int    `json:"ai_type"`
				AiStatus    int    `json:"ai_status"`
				Author      struct {
					Mid            int    `json:"mid"`
					Name           string `json:"name"`
					Sex            string `json:"sex"`
					Face           string `json:"face"`
					Sign           string `json:"sign"`
					Rank           int    `json:"rank"`
					Birthday       int    `json:"birthday"`
					IsFakeAccount  int    `json:"is_fake_account"`
					IsDeleted      int    `json:"is_deleted"`
					InRegAudit     int    `json:"in_reg_audit"`
					IsSeniorMember int    `json:"is_senior_member"`
				} `json:"author"`
			} `json:"list"`
		} `json:"subtitle"`
		UgcSeason struct {
			Id        int    `json:"id"`
			Title     string `json:"title"`
			Cover     string `json:"cover"`
			Mid       int    `json:"mid"`
			Intro     string `json:"intro"`
			SignState int    `json:"sign_state"`
			Attribute int    `json:"attribute"`
			Sections  []struct {
				SeasonId int    `json:"season_id"`
				Id       int    `json:"id"`
				Title    string `json:"title"`
				Type     int    `json:"type"`
				Episodes []struct {
					SeasonId  int    `json:"season_id"`
					SectionId int    `json:"section_id"`
					Id        int    `json:"id"`
					Aid       int    `json:"aid"`
					Cid       int    `json:"cid"`
					Title     string `json:"title"`
					Attribute int    `json:"attribute"`
					Arc       struct {
						Aid       int    `json:"aid"`
						Videos    int    `json:"videos"`
						TypeId    int    `json:"type_id"`
						TypeName  string `json:"type_name"`
						Copyright int    `json:"copyright"`
						Pic       string `json:"pic"`
						Title     string `json:"title"`
						Pubdate   int    `json:"pubdate"`
						Ctime     int    `json:"ctime"`
						Desc      string `json:"desc"`
						State     int    `json:"state"`
						Duration  int    `json:"duration"`
						Rights    struct {
							Bp            int `json:"bp"`
							Elec          int `json:"elec"`
							Download      int `json:"download"`
							Movie         int `json:"movie"`
							Pay           int `json:"pay"`
							Hd5           int `json:"hd5"`
							NoReprint     int `json:"no_reprint"`
							Autoplay      int `json:"autoplay"`
							UgcPay        int `json:"ugc_pay"`
							IsCooperation int `json:"is_cooperation"`
							UgcPayPreview int `json:"ugc_pay_preview"`
							ArcPay        int `json:"arc_pay"`
							FreeWatch     int `json:"free_watch"`
						} `json:"rights"`
						Author struct {
							Mid  int    `json:"mid"`
							Name string `json:"name"`
							Face string `json:"face"`
						} `json:"author"`
						Stat struct {
							Aid        int    `json:"aid"`
							View       int    `json:"view"`
							Danmaku    int    `json:"danmaku"`
							Reply      int    `json:"reply"`
							Fav        int    `json:"fav"`
							Coin       int    `json:"coin"`
							Share      int    `json:"share"`
							NowRank    int    `json:"now_rank"`
							HisRank    int    `json:"his_rank"`
							Like       int    `json:"like"`
							Dislike    int    `json:"dislike"`
							Evaluation string `json:"evaluation"`
							ArgueMsg   string `json:"argue_msg"`
							Vt         int    `json:"vt"`
							Vv         int    `json:"vv"`
						} `json:"stat"`
						Dynamic   string `json:"dynamic"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						DescV2             interface{} `json:"desc_v2"`
						IsChargeableSeason bool        `json:"is_chargeable_season"`
						IsBlooper          bool        `json:"is_blooper"`
						EnableVt           int         `json:"enable_vt"`
						VtDisplay          string      `json:"vt_display"`
					} `json:"arc"`
					Page struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"page"`
					Bvid string `json:"bvid"`
				} `json:"episodes"`
			} `json:"sections"`
			Stat struct {
				SeasonId int `json:"season_id"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Fav      int `json:"fav"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			EpCount     int  `json:"ep_count"`
			SeasonType  int  `json:"season_type"`
			IsPaySeason bool `json:"is_pay_season"`
			EnableVt    int  `json:"enable_vt"`
		} `json:"ugc_season"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
		} `json:"honor_reply"`
		LikeIcon          string `json:"like_icon"`
		NeedJumpBv        bool   `json:"need_jump_bv"`
		DisableShowUpInfo bool   `json:"disable_show_up_info"`
		IsStoryPlay       int    `json:"is_story_play"`
	} `json:"data"`
}
