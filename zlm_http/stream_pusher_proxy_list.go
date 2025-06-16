package zlm_http

import "context"

//* 获取推流代理列表

type ListStreamPusherProxyRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type ListStreamPusherProxyReply struct {
	BaseResult
	Data []StreamPusherProxyEntry `json:"data"`
}
type StreamPusherProxyEntry struct {
	Key            string               `json:"key"`            // 唯一码, , 格式: {schema}/{vhost}/{app}/{stream}/{uniqueKey}
	Url            string               `json:"url"`            // 目标转推url
	Src            StreamPusherProxySrc `json:"src"`            // 转推流的源信息
	LiveSecs       int                  `json:"liveSecs"`       // 转推流时长
	RePublishCount int                  `json:"rePublishCount"` // 重试次数
	BytesSpeed     int                  `json:"bytesSpeed"`     // 字节速度
	TotalBytes     int                  `json:"totalBytes"`     // 总字节数
	Status         int                  `json:"status"`         // 流状态
}
type StreamPusherProxySrc struct {
	Vhost  string `json:"vhost"`  // 流的虚拟主机
	App    string `json:"app"`    // 流的应用名
	Stream string `json:"stream"` // 流id
	Params string `json:"params"` // 流参数
}

func (c *Client) ListStreamPusherProxy(ctx context.Context, req *ListStreamPusherProxyRequest, opts ...CallOption) (*ListStreamPusherProxyReply, error) {
	return GenericPost[ListStreamPusherProxyRequest, ListStreamPusherProxyReply](c, "/index/api/listStreamPusherProxy", ctx, req, opts...)
}
