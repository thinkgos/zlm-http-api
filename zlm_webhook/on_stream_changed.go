package zlm_webhook

import "github.com/thinkgos/zlm-http-api/zlm_def"

//* rtsp/rtmp 流注册或注销时触发此事件; 此事件对回复不敏感.
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_12%E3%80%81on-stream-changed

type OnStreamChangedRequest struct {
	HookIndex     int    `json:"hook_index"`
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	Schema        string `json:"schema"`        // 协议
	Vhost         string `json:"vhost"`         // 流虚拟主机
	App           string `json:"app"`           // 流应用名
	Stream        string `json:"stream"`        // 流id
	Params        string `json:"params"`        // 流url参数
	Regist        bool   `json:"regist"`        // 是否注册
	// 以上注册/注销都有
	// 以下仅注册时有效(regist=true)
	OriginType       int                `json:"originType"`       // 产生源类型，包括 unknown = 0,rtmp_push=1,rtsp_push=2,rtp_push=3,pull=4,ffmpeg_pull=5,mp4_vod=6,device_chn=7,rtc_push=8
	OriginTypeStr    string             `json:"originTypeStr"`    // 产生源类型字符串
	OriginUrl        string             `json:"originUrl"`        // 产生源的url
	OriginSock       zlm_def.OriginSock `json:"originSock"`       // 产生源信息
	TotalBytes       int                `json:"totalBytes"`       // 总字节数
	IsRecordingHls   bool               `json:"isRecordingHLS"`   // 是否记录HLS
	IsRecordingMp4   bool               `json:"isRecordingMP4"`   // 是否记录MP4
	ReaderCount      int                `json:"readerCount"`      // 本协议观看人数
	TotalReaderCount int                `json:"totalReaderCount"` // 观看总人数，包括hls/rtsp/rtmp/http-flv/ws-flv/rtc
	AliveSecond      int                `json:"aliveSecond"`      // 存活时间, 单位秒
	BytesSpeed       int                `json:"bytesSpeed"`       // 数据产生速率, 单位byte/s
	CreateStamp      int                `json:"createStamp"`      // GMT unix系统时间戳, 单位秒
	Tracks           []zlm_def.Tracks   `json:"tracks"`           // 追踪
}
type OnStreamChangedReply = DefaultReply
