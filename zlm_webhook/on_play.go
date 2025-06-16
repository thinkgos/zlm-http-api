package zlm_webhook

//* 播放器鉴权事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_6%E3%80%81on-play
//
//  rtsp/rtmp/http-flv/ws-flv/hls的播放都将触发此鉴权事件;
// 如果流不存在, 那么先触发on_play事件然后触发on_stream_not_found事件.
// 播放rtsp流时, 如果该流启动了rtsp专属鉴权(on_rtsp_realm)那么将不再触发on_play事件.

type OnPlayRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Schema        string `json:"schema"`        // 播放的协议, 可能是rtsp、rtmp、http
	Vhost         string `json:"vhost"`         // 流虚拟主机
	App           string `json:"app"`           // 流应用名
	Stream        string `json:"stream"`        // 流id
	Params        string `json:"params"`        // 播放url参数
	Protocol      string `json:"protocol"`      // 流协议
	Id            string `json:"id"`            // tcp链接唯一id
	Ip            string `json:"ip"`            // 播放器ip
	Port          int    `json:"port"`          // 播放器端口号
}

type OnPlayReply = DefaultReply
