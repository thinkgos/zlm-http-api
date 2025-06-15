package zlm_http

import "context"

//* 获取 openRtpServer 接口创建的所有 RTP 服务器
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_26%E3%80%81-index-api-listrtpserver

type ListRtpServerRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type ListRtpServerReply struct {
	BaseResult
	Data []*RtpServerEntry `json:"data"`
}
type RtpServerEntry struct {
	Port     int    `json:"port"`      // 绑定的端口号
	StreamId string `json:"stream_id"` //绑定的流id
}

func (c *Client) ListRtpServer(ctx context.Context, req *ListRtpServerRequest, opts ...CallOption) (*ListRtpServerReply, error) {
	return GenericGet[ListRtpServerRequest, ListRtpServerReply](c, "/index/api/listRtpServer", ctx, req, opts...)
}
