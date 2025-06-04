package zlm_webhook

//* rtsp/rtmp/rtp 推流鉴权事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_7%E3%80%81on-publish

type OnPublishRequest struct {
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	App           string `json:"app"`           // 流应用名
	Id            string `json:"id"`            // tcp链接唯一id
	Ip            string `json:"ip"`            // 推流器ip
	Port          int    `json:"port"`          // 推流器端口号
	Params        string `json:"params"`        // 推流url参数
	Schema        string `json:"schema"`        // 推流的协议, 可能是 rtsp、rtmp
	Stream        string `json:"stream"`        // 流id
	Vhost         string `json:"vhost"`         // 流虚拟主机
}
type OnPublishReply struct {
	Code           int     `json:"code"`                       // 错误代码, 0: 成功
	Msg            string  `json:"msg,omitempty"`              // 错误提示
	EnableHls      *bool   `json:"enable_hls,omitempty"`       // 是否转换成 hls-mpegts 协议
	EnableHlsFmp4  *bool   `json:"enable_hls_fmp4,omitempty"`  // 是否转换成 hls-fmp4 协议
	EnableMp4      *bool   `json:"enable_mp4,omitempty"`       // 是否允许 mp4 录制
	EnableFmp4     *bool   `json:"enable_fmp4,omitempty"`      // 是否转 http-fmp4/ws-fmp4 协议
	EnableRtmp     *bool   `json:"enable_rtmp,omitempty"`      // 是否转 rtmp/flv 协议
	EnableRtsp     *bool   `json:"enable_rtsp,omitempty"`      // 是否转 rtsp 协议
	EnableTs       *bool   `json:"enable_ts,omitempty"`        // 是否转 http-ts/ws-ts 协议
	EnableAudio    *bool   `json:"enable_audio,omitempty"`     // 转协议时是否开启音频
	AddMuteAudio   *bool   `json:"add_mute_audio,omitempty"`   // 转协议时, 无音频是否添加静音aac音频
	HlsDemand      bool    `json:"hls_demand,omitempty"`       // 该协议是否有人观看才生成
	Fmp4Demand     bool    `json:"fmp4_demand,omitempty"`      // 该协议是否有人观看才生成
	RtspDemand     bool    `json:"rtsp_demand,omitempty"`      // 该协议是否有人观看才生成
	RtmpDemand     bool    `json:"rtmp_demand,omitempty"`      // 该协议是否有人观看才生成
	TsDemand       bool    `json:"ts_demand,omitempty"`        // 该协议是否有人观看才生成
	ContinuePushMs *int    `json:"continue_push_ms,omitempty"` // 断连续推延时, 单位毫秒, 置空使用配置文件默认值
	HlsSavePath    *string `json:"hls_save_path,omitempty"`    // hls 文件保存保存根目录, 置空使用默认
	ModifyStamp    *int    `json:"modify_stamp,omitempty"`     // 该流是否开启时间戳覆盖(0:绝对时间戳/1:系统时间戳/2:相对时间戳)
	Mp4AsPlayer    *bool   `json:"mp4_as_player,omitempty"`    // mp4 录制是否当作观看者参与播放人数计数
	Mp4MaxSecond   *int    `json:"mp4_max_second,omitempty"`   // mp4 录制切片大小, 单位秒
	Mp4SavePath    *string `json:"mp4_save_path,omitempty"`    // mp4 录制文件保存根目录, 置空使用默认
	AutoClose      *bool   `json:"auto_close,omitempty"`       // 无人观看是否自动关闭流(不触发无人观看 hook)
	StreamReplace  *string `json:"stream_replace,omitempty"`   // 是否修改流id, 通过此参数可以自定义流 id(譬如替换 ssrc)
}
