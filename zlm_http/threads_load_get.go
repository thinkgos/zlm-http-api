package zlm_http

import "context"

//* 获取各epoll(或 select)线程负载以及延时
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_1%E3%80%81-index-api-getthreadsload

type GetThreadsLoadRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type GetThreadsLoadReply struct {
	BaseResult
	Data []ThreadsLoadEntry `json:"data"`
}
type ThreadsLoadEntry struct {
	Delay int `json:"delay"` // 该线程延迟
	Load  int `json:"load"`  // 该线程负载, 0-100
}

func (c *Client) GetThreadsLoad(ctx context.Context, req *GetThreadsLoadRequest, opts ...CallOption) (*GetThreadsLoadReply, error) {
	return genericGet[GetThreadsLoadRequest, GetThreadsLoadReply](c, "/index/api/getThreadsLoad", ctx, req, opts...)
}
