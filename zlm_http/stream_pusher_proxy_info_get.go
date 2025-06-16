package zlm_http

import "context"

//* 获取推流代理信息

type GetStreamPusherProxyInfoRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Key    string `json:"key"`              // M, addStreamPusherProxy接口返回的key
}
type GetStreamPusherProxyInfoReply struct {
	BaseResult
	Data *StreamPusherProxyEntry `json:"data"`
}

func (c *Client) GetStreamPusherProxyInfo(ctx context.Context, req *GetStreamPusherProxyInfoRequest, opts ...CallOption) (*GetStreamPusherProxyInfoReply, error) {
	resp, err := GenericPost[GetStreamPusherProxyInfoRequest, GetStreamPusherProxyInfoReply](c, "/index/api/getProxyPusherInfo", ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if resp.Data != nil {
		resp.Data.Key = req.Key
	}
	return resp, nil
}
