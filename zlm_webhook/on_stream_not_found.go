package zlm_webhook

//* 流未找到事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_14%E3%80%81on-stream-not-found
//
// 用户可以在此事件触发时, 立即去拉流, 这样可以实现按需拉流; 此事件对回复不敏感.

type OnStreamNotFoundRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Schema        string `json:"schema"`        // 播放的协议, 可能是rtsp、rtmp、http
	Vhost         string `json:"vhost"`         // 流虚拟主机
	App           string `json:"app"`           // 流应用名
	Stream        string `json:"stream"`        // 流id
	Params        string `json:"params"`        // 播放url参数
	Id            string `json:"id"`            // tcp链接唯一id
	Ip            string `json:"ip"`            // 播放器ip
	Port          int    `json:"port"`          // 播放器端口号
}
type OnStreamNotFoundReply = DefaultReply
