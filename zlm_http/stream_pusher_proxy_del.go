package zlm_http

import "context"

//* 删除推流代理

type DelStreamPusherProxyRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Key    string `json:"key"`              // M, addStreamPusherProxy接口返回的key
}
type DelStreamPusherProxyReply struct {
	BaseResult
	Data DelStreamPusherProxyData `json:"data"`
}
type DelStreamPusherProxyData struct {
	Flag bool `json:"flag"` // 成功与否
}

func (c *Client) DelStreamPusherProxy(ctx context.Context, req *DelStreamPusherProxyRequest, opts ...CallOption) (*DelStreamPusherProxyReply, error) {
	return GenericPost[DelStreamPusherProxyRequest, DelStreamPusherProxyReply](c, "/index/api/delStreamPusherProxy", ctx, req, opts...)
}
