package zlm_webhook

//* rtsp专用的鉴权事件, 先触发on_rtsp_realm事件然后才会触发on_rtsp_auth事件。
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_10%E3%80%81on-rtsp-auth

type OnRtspAuthRequest struct {
	MediaServerId string `json:"mediaServerId"`   // 服务器id, 通过配置文件设置
	App           string `json:"app"`             // 录制的流应用名
	Id            string `json:"id"`              // tcp链接唯一id
	Ip            string `json:"ip"`              // rtsp播放器ip
	Port          int    `json:"port"`            // rtsp 播放器端口号
	Params        string `json:"params"`          // 播放rtsp url参数
	Schema        string `json:"schema"`          // rtsp 或 rtsps
	Stream        string `json:"stream"`          // 流id
	Vhost         string `json:"vhost"`           // 流虚拟主机
	MustNoEncrypt bool   `json:"must_no_encrypt"` // 请求的密码是否必须为明文(base64鉴权需要明文密码)
	Realm         string `json:"realm"`           // rtsp播放鉴权加密realm
	UserName      string `json:"user_name"`       // 播放用户名
}

type OnRtspAuthReply struct {
	Code      int    `json:"code"`          // 代码, 0代表允许播放
	Msg       string `json:"msg,omitempty"` // 播放鉴权失败时的错误提示
	Encrypted bool   `json:"encrypted"`     // 用户密码是明文还是摘要
	Passwd    string `json:"passwd"`        // 用户密码明文或摘要(md5(username:realm:password))
}
