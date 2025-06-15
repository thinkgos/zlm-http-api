package zlm_http

import "context"

//* 停止 GB28181 ps-rtp 推流
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_28%E3%80%81-index-api-stopsendrtp

type StopSendRtpRequest struct {
	Secret string  `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost  string  `json:"vhost"`            // M, 虚拟主机, 例如__defaultVhost__
	App    string  `json:"app"`              // M, 应用名, 例如 live
	Stream string  `json:"stream"`           // M, 流id, 例如 test
	Ssrc   *string `json:"ssrc,omitempty"`   // O, 根据ssrc关停某路rtp推流, 置空时关闭所有流
}
type StopSendRtpReply = BaseResult

func (c *Client) StopSendRtp(ctx context.Context, req *StopSendRtpRequest, opts ...CallOption) (*StopSendRtpReply, error) {
	return GenericPost[StopSendRtpRequest, StopSendRtpReply](c, "/index/api/stopSendRtp", ctx, req, opts...)
}
