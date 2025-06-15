package zlm_http

import "context"

//* 停止录制流
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_21%E3%80%81-index-api-stoprecord

type StopRecordRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost  string `json:"vhost"`            // M, 流的虚拟主机, 例如__defaultVhost__
	App    string `json:"app"`              // M, 流的应用名, 例如 live
	Stream string `json:"stream"`           // M, 流id, 例如 test
	Type   int    `json:"type,string"`      // M, 0: hls, 1: mp4
}
type StopRecordReply struct {
	BaseResult
	Result bool `json:"result"` // 成功与否
}

func (c *Client) StopRecord(ctx context.Context, req *StopRecordRequest, opts ...CallOption) (*StopRecordReply, error) {
	return GenericPost[StopRecordRequest, StopRecordReply](c, "/index/api/stopRecord", ctx, req, opts...)
}
