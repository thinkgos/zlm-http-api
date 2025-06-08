package zlm_http

import "context"

//* 获取API列表
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_0%E3%80%81-index-api-getapilist

type GetApiListRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type GetApiListReply struct {
	BaseResult
	Data []string `json:"data"`
}

func (c *Client) GetApiList(ctx context.Context, req *GetApiListRequest, opts ...CallOption) (*GetApiListReply, error) {
	var resp GetApiListReply

	err := c.Get(ctx, "/index/api/getApiList", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError();err != nil {
		return nil, err
	}
	return &resp, nil
}
