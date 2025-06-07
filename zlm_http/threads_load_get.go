package zlm_http

import "context"

//* 获取各epoll(或 select)线程负载以及延时
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_1%E3%80%81-index-api-getthreadsload

type GetThreadsLoadRequest struct{}
type GetThreadsLoadReply struct {
	BaseResult
	Data []ThreadsLoadEntry `json:"data"`
}
type ThreadsLoadEntry struct {
	Delay int `json:"delay"` // 该线程延迟
	Load  int `json:"load"`  // 该线程负载, 0-100
}

func (c *ZlmClient) GetThreadsLoad(ctx context.Context, req *GetThreadsLoadRequest, opts ...CallOption) (*GetThreadsLoadReply, error) {
	var resp GetThreadsLoadReply

	err := c.Get(ctx, "/index/api/getThreadsLoad", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
