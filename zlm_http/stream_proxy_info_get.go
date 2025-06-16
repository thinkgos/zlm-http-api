package zlm_http

import "context"

//* 获取拉流代理信息

type GetStreamProxyInfoRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Key    string `json:"key"`              // M, addStreamProxy接口返回的key
}
type GetStreamProxyInfoReply struct {
	BaseResult
	Data *StreamProxyEntry `json:"data"`
}

func (c *Client) GetStreamProxyInfo(ctx context.Context, req *GetStreamProxyInfoRequest, opts ...CallOption) (*GetStreamProxyInfoReply, error) {
	return GenericPost[GetStreamProxyInfoRequest, GetStreamProxyInfoReply](c, "/index/api/getProxyInfo", ctx, req, opts...)
}
