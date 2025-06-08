package zlm_http

import "context"

//* 重启服务器, 只有Daemon方式才能重启, 否则是直接关闭!
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_5%E3%80%81-index-api-restartserver

type RestartServerRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type RestartServerReply struct {
	BaseResult
}

func (c *Client) RestartServer(ctx context.Context, req *RestartServerRequest, opts ...CallOption) (*RestartServerReply, error) {
	var resp RestartServerReply

	err := c.Post(ctx, "/index/api/restartServer", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
