package zlm_http

import (
	"context"

	"github.com/thinkgos/zlm-http-api/zlm_def"
)

//* 取流列表, 可选筛选参数
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_6%E3%80%81-index-api-getmedialist

type GetMediaListRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Schema string `json:"schema,omitempty"` // O, 筛选协议, 例如rtsp或 rtmp
	Vhost  string `json:"vhost,omitempty"`  // O, 筛选虚拟主机, 例如__defaultVhost__
	App    string `json:"app,omitempty"`    // O, 筛选应用名, 例如 live
	Stream string `json:"stream,omitempty"` // O, 筛选流id, 例如 test
}
type GetMediaListReply struct {
	BaseResult
	Data []MediaEntry `json:"data"`
}

type MediaEntry struct {
	App              string              `json:"app"`              // 流应用名
	Schema           string              `json:"schema"`           // 协议
	Stream           string              `json:"stream"`           // 流id
	Vhost            string              `json:"vhost"`            // 流虚拟主机
	ReaderCount      int                 `json:"readerCount"`      // 本协议观看人数
	TotalReaderCount int                 `json:"totalReaderCount"` // 观看总人数，包括hls/rtsp/rtmp/http-flv/ws-flv/rtc
	OriginSock       *zlm_def.OriginSock `json:"originSock"`       // 客户端和服务器网络信息, 可能为null类型
	OriginType       int                 `json:"originType"`       // 产生源类型，包括 unknown = 0,rtmp_push=1,rtsp_push=2,rtp_push=3,pull=4,ffmpeg_pull=5,mp4_vod=6,device_chn=7,rtc_push=8
	OriginTypeStr    string              `json:"originTypeStr"`    // 产生源类型字符串
	OriginUrl        string              `json:"originUrl"`        // 产生源的url
	Params           string              `json:"params"`           // 流url参数
	TotalBytes       int                 `json:"totalBytes"`       // 总字节数
	IsRecordingHls   bool                `json:"isRecordingHLS"`   // 是否记录HLS
	IsRecordingMp4   bool                `json:"isRecordingMP4"`   // 是否记录MP4
	AliveSecond      int                 `json:"aliveSecond"`      // 存活时间, 单位秒
	BytesSpeed       int                 `json:"bytesSpeed"`       // 数据产生速率, 单位byte/s
	CreateStamp      int                 `json:"createStamp"`      // GMT unix系统时间戳, 单位秒
	Tracks           []zlm_def.Tracks    `json:"tracks"`           // 追踪

}

func (c *Client) GetMediaList(ctx context.Context, req *GetMediaListRequest, opts ...CallOption) (*GetMediaListReply, error) {
	return GenericGet[GetMediaListRequest, GetMediaListReply](c, "/index/api/getMediaList", ctx, req, opts...)
}
