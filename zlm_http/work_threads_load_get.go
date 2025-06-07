package zlm_http

import "context"

//* 获取各后台epoll(或 select)线程负载以及延时
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_2%E3%80%81-index-api-getworkthreadsload

type GetWorkThreadsLoadRequest struct{}
type GetWorkThreadsLoadReply struct {
	BaseResult
	Data []WorkThreadsLoadEntry `json:"data"`
}
type WorkThreadsLoadEntry struct {
	Delay int `json:"delay"` // 该线程延迟
	Load  int `json:"load"`  // 该线程负载, 0-100
}

func (c *ZlmClient) GetWorkThreadsLoad(ctx context.Context, req *GetWorkThreadsLoadRequest) (*GetWorkThreadsLoadReply, error) {
	var resp GetWorkThreadsLoadReply

	err := c.Get(ctx, "/index/api/getWorkThreadsLoad", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
