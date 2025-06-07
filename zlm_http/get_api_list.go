package zlm_http

import "context"

//* 获取API列表
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_0%E3%80%81-index-api-getapilist

type GetApiListRequest struct{}
type GetApiListReply struct {
	BaseResult
	Data []string `json:"data"`
}

func (c *ZlmClient) GetApiList(ctx context.Context, req *GetApiListRequest) (*GetApiListReply, error) {
	var resp GetApiListReply

	err := c.Get(ctx, "/index/api/getApiList", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
