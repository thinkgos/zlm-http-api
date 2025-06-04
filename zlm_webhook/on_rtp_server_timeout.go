package zlm_webhook

type OnRtpServerTimeoutRequest struct {
	MediaServerID string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	LocalPort     int    `json:"local_port"`    // openRtpServer 输入的参数
	ReUsePort     bool   `json:"re_use_port"`   // openRtpServer 输入的参数
	SSRC          uint32 `json:"ssrc"`          // openRtpServer 输入的参数
	StreamID      string `json:"stream_id"`     // openRtpServer 输入的参数
	TCPMode       int    `json:"tcp_mode"`      // openRtpServer 输入的参数

}
type OnRtpServerTimeoutReply = DefaultReply
