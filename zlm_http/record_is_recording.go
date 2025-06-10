package zlm_http

import "context"

//* 获取流录制状态
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_22%E3%80%81-index-api-isrecording

type IsRecordingRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost  string `json:"vhost"`            // M, 流的虚拟主机, 例如__defaultVhost__
	App    string `json:"app"`              // M, 流的应用名, 例如 live
	Stream string `json:"stream"`           // M, 流id, 例如 test
	Type   string `json:"type"`             // M, 0: hls, 1: mp4
}
type IsRecordingReply struct {
	BaseResult
	Status bool `json:"result"` // false:未录制, true:正在录制
}

func (c *Client) IsRecording(ctx context.Context, req *IsRecordingRequest, opts ...CallOption) (*IsRecordingReply, error) {
	return genericPost[IsRecordingRequest, IsRecordingReply](c, "/index/api/isRecording", ctx, req, opts...)
}
