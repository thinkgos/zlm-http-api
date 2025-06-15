package zlm_http

import "context"

//* 获取各后台epoll(或 select)线程负载以及延时
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_2%E3%80%81-index-api-getworkthreadsload

type GetWorkThreadsLoadRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type GetWorkThreadsLoadReply struct {
	BaseResult
	Data []WorkThreadsLoadEntry `json:"data"`
}
type WorkThreadsLoadEntry struct {
	Delay int `json:"delay"` // 该线程延迟
	Load  int `json:"load"`  // 该线程负载, 0-100
}

func (c *Client) GetWorkThreadsLoad(ctx context.Context, req *GetWorkThreadsLoadRequest, opts ...CallOption) (*GetWorkThreadsLoadReply, error) {
	return GenericGet[GetWorkThreadsLoadRequest, GetWorkThreadsLoadReply](c, "/index/api/getWorkThreadsLoad", ctx, req, opts...)
}
