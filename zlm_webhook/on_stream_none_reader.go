package zlm_webhook

//* 流无人观看时事件
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_13%E3%80%81on-stream-none-reader
//
// 用户可以通过此事件选择是否关闭无人看的流.
// 一个直播流注册上线了, 如果一直没人观看也会触发一次无人观看事件,
// 触发时的协议schema是随机的, 看哪种协议最晚注册(一般为hls).
// 后续从有人观看转为无人观看, 触发协议schema为最后一名观看者使用何种协议.
// 目前mp4/hls录制不当做观看人数
// (mp4录制可以通过配置文件mp4_as_player控制, 但是 rtsp/rtmp/rtp 转推算观看人数, 也会触发该事件.)

type OnStreamNoneReaderRequest struct {
	MediaServerId string `json:"mediaServerId"` // 服务器 id, 通过配置文件设置
	App           string `json:"app"`           // 流应用名
	Schema        string `json:"schema"`        // rtsp或rtmp
	Stream        string `json:"stream"`        // 流id
	Vhost         string `json:"vhost"`         // 流虚拟主机
}

type OnStreamNoneReaderReply struct {
	Code  int  `json:"code"`  // 请固定返回 0
	Close bool `json:"close"` // 否关闭推流或拉流
}
