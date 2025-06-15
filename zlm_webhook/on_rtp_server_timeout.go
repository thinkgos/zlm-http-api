package zlm_webhook

//* 调用openRtpServer接口, rtp server长时间未收到数据, 执行此 web hook, 对回复不敏感.
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_17%E3%80%81on-rtp-server-timeout

type OnRtpServerTimeoutRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	App           string `json:"app"`           // openRtpServer 流应用名
	Vhost         string `json:"vhost"`         // openRtpServer 流虚拟主机
	StreamId      string `json:"stream_id"`     // openRtpServer 输入的参数
	LocalPort     int    `json:"local_port"`    // openRtpServer 输入的参数
	ReUsePort     bool   `json:"re_use_port"`   // openRtpServer 输入的参数
	Ssrc          uint32 `json:"ssrc"`          // openRtpServer 输入的参数
	TcpMode       int    `json:"tcp_mode"`      // openRtpServer 输入的参数

}
type OnRtpServerTimeoutReply = DefaultReply
