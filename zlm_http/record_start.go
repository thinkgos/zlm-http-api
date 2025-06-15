package zlm_http

import "context"

//* 开始录制 hls 或 MP4
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_20%E3%80%81-index-api-startrecord

type StartRecordRequest struct {
	Secret         string `json:"secret,omitempty"`          // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost          string `json:"vhost"`                     // M, 流的虚拟主机, 例如__defaultVhost__
	App            string `json:"app"`                       // M, 流的应用名, 例如 live
	Stream         string `json:"stream"`                    // M, 流id, 例如 test
	Type           int    `json:"type,string"`               // M, 0: hls, 1: mp4
	CustomizedPath string `json:"customized_path,omitempty"` // O, 录像保存目录
	MaxSecond      int    `json:"max_second,string"`         // O, mp4 录像切片时间大小, 单位秒, 置0则采用配置项
}
type StartRecordReply struct {
	BaseResult
	Result bool `json:"result"` // 成功与否
}

func (c *Client) StartRecord(ctx context.Context, req *StartRecordRequest, opts ...CallOption) (*StartRecordReply, error) {
	return GenericPost[StartRecordRequest, StartRecordReply](c, "/index/api/startRecord", ctx, req, opts...)
}
