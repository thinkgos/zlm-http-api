package zlm_http

import "context"

type GetAPIListParams struct{}

type GetAPIListReply struct {
	BaseResult
	Data []string `json:"data"`
}

func (c *ZlmClient) GetApiList(ctx context.Context, req *GetAPIListParams) (*GetAPIListReply, error) {
	var resp GetAPIListReply

	err := c.Get(ctx, "/index/api/getApiList", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
