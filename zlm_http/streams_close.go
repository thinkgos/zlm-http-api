package zlm_http

import (
	"context"
)

//* 关闭流(目前所有类型的流都支持关闭)
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_8%E3%80%81-index-api-close-streams

type CloseStreamsRequest struct {
	Schema string `json:"schema"` // O, 筛选协议, 例如rtsp或 rtmp
	Vhost  string `json:"vhost"`  // O, 筛选虚拟主机, 例如__defaultVhost__
	App    string `json:"app"`    // O, 筛选应用名, 例如 live
	Stream string `json:"stream"` // O, 筛选流id, 例如 test
	Force  int    `json:"force"`  // O, 是否强制关闭
}
type CloseStreamsReply struct {
	BaseResult
	CountHit    int `json:"count_hit"`    // 筛选命中的流个数
	CountClosed int `json:"count_closed"` // 被关闭的流个数, 可能小于count_hit
}

func (c *ZlmClient) CloseStreams(ctx context.Context, req *CloseStreamsRequest, opts ...CallOption) (*CloseStreamsReply, error) {
	var resp CloseStreamsReply

	err := c.Post(ctx, "/index/api/close_streams", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
