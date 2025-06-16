package zlm_http

import "context"

//* 获取拉流代理列表

type ListStreamProxyRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type ListStreamProxyReply struct {
	BaseResult
	Data []StreamProxyEntry `json:"data"`
}
type StreamProxyEntry struct {
	Key              string         `json:"key"`              // 唯一码, 格式: {vhost}/{app}/{stream}
	Url              string         `json:"url"`              // 拉流url
	Src              StreamProxySrc `json:"src"`              // 拉流的源信息
	LiveSecs         int            `json:"liveSecs"`         // 拉流时长
	RePullCount      int            `json:"rePullCount"`      // 重试次数
	BytesSpeed       int            `json:"bytesSpeed"`       // 字节速度
	TotalBytes       int            `json:"totalBytes"`       // 总字节数
	TotalReaderCount int            `json:"totalReaderCount"` // 观看人数
	Status           int            `json:"status"`           // 流状态
}
type StreamProxySrc struct {
	Vhost  string `json:"vhost"`  // 流的虚拟主机
	App    string `json:"app"`    // 流的应用名
	Stream string `json:"stream"` // 流id
	Params string `json:"params"` // 流参数
}

func (c *Client) ListStreamProxy(ctx context.Context, req *ListStreamProxyRequest, opts ...CallOption) (*ListStreamProxyReply, error) {
	return GenericPost[ListStreamProxyRequest, ListStreamProxyReply](c, "/index/api/listStreamProxy", ctx, req, opts...)
}
