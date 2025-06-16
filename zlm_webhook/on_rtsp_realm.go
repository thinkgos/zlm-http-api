package zlm_webhook

//* 该rtsp流是否开启rtsp专用方式的鉴权事件, 开启后才会触发on_rtsp_auth事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_9%E3%80%81on-rtsp-realm
//
// 需要指出的是rtsp也支持url参数鉴权, 它支持两种方式鉴权.

type OnRtspRealmRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Schema        string `json:"schema"`        // rtsp 或 rtsps
	Vhost         string `json:"vhost"`         // 流虚拟主机
	App           string `json:"app"`           // 录制的流应用名
	Stream        string `json:"stream"`        // 流id
	Params        string `json:"params"`        // 播放rtsp url参数
	Protocol      string `json:"protocol"`      // 流协议
	Id            string `json:"id"`            // tcp链接唯一id
	Ip            string `json:"ip"`            // rtsp播放器ip
	Port          int    `json:"port"`          // rtsp播放器端口号
}

type OnRtspRealmReply struct {
	Code  int    `json:"code"`  // 请固定返回 0
	Realm string `json:"realm"` // 该rtsp流是否需要rtsp专有鉴权, 空字符串代码不需要鉴权
}
