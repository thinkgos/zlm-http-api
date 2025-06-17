package zlm_webhook

//* 发送rtp(startSendRtp)被动关闭时回调

type OnSendRtpStoppedRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Vhost         string `json:"vhost"`         // 流虚拟主机
	App           string `json:"app"`           // 流应用名
	Stream        string `json:"stream"`        // 流id
	OriginType    int    `json:"originType"`    // 产生源类型
	OriginTypeStr string `json:"originTypeStr"` // 产生源类型字符串
	OriginUrl     string `json:"originUrl"`     // 产生源信息
	Params        string `json:"params"`        // 参数
	Ssrc          int    `json:"ssrc,string"`   // 对应ssrc
	Err           int    `json:"err"`           // 发生错误, >0: 错误
	Msg           string `json:"msg"`           // 错误信息
}
type OnSendRtpStoppedReply = DefaultReply
