package zlm_http

import "context"

//* 关闭 GB28181 RTP 接收端口
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_25%E3%80%81-index-api-closertpserver

type CloseRtpServerRequest struct {
	Secret   string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost    string `json:"vhost,omitempty"`  // O, 虚拟主机, 默认: __defaultVhost__
	App      string `json:"app,omitempty"`    // O, 流应用名, 默认: rtp
	StreamId string `json:"stream_id"`        // M, 该端口绑定的流id, 该端口只能创建这一个流(而不是根据ssrc创建多个)
}
type CloseRtpServerReply struct {
	BaseResult
	Hit int `json:"hit"` // 是否找到记录并关闭
}

func (c *Client) CloseRtpServer(ctx context.Context, req *CloseRtpServerRequest, opts ...CallOption) (*CloseRtpServerReply, error) {
	return GenericPost[CloseRtpServerRequest, CloseRtpServerReply](c, "/index/api/closeRtpServer", ctx, req, opts...)
}
