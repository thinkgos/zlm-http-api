package zlm_http

import (
	"context"
)

//* 关闭流(目前所有类型的流都支持关闭)
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_8%E3%80%81-index-api-close-streams

type CloseStreamsRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Schema string `json:"schema"`           // M, 筛选协议, 例如rtsp或 rtmp
	Vhost  string `json:"vhost"`            // M, 筛选虚拟主机, 例如__defaultVhost__
	App    string `json:"app"`              // M, 筛选应用名, 例如 live
	Stream string `json:"stream"`           // M, 筛选流id, 例如 test
	Force  *int   `json:"force,omitempty"`  // O, 是否强制关闭(有人在观看将强制关闭), 0: 不强制关闭, 1: 强制关闭
}
type CloseStreamsReply struct {
	BaseResult
	CountHit    int `json:"count_hit"`    // 筛选命中的流个数
	CountClosed int `json:"count_closed"` // 被关闭的流个数, 可能小于count_hit
}

func (c *Client) CloseStreams(ctx context.Context, req *CloseStreamsRequest, opts ...CallOption) (*CloseStreamsReply, error) {
	return GenericPost[CloseStreamsRequest, CloseStreamsReply](c, "/index/api/close_streams", ctx, req, opts...)
}
