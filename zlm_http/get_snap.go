package zlm_http

import "context"

type GetSnapRequest struct {
	Url        string `json:"url"`         // 需要截图的url, 可以是本机的, 也可以是远程主机的.
	TimeoutSec int    `json:"timeout_sec"` // 截图失败超时时间, 防止FFmpeg一直等待截图
	ExpireSec  int    `json:"expire_sec"`  // 截图的过期时间, 该时间内产生的截图都会作为缓存返回
}

// TODO: Implement me
func (c *ZlmClient) GetSnap(ctx context.Context, req *GetApiListRequest) (*GetApiListReply, error) {
	var resp GetApiListReply

	err := c.Get(ctx, "/index/api/getSnap", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
