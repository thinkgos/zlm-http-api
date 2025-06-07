package zlm_http

import "context"

//* 设置服务器配置
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_4%E3%80%81-index-api-setserverconfig

type SetServerConfigRequest struct {
	// TODO: ...
}
type SetServerConfigReply struct {
	BaseResult
	Changed int `json:"changed"` // 配置项变更个数
}

func (c *ZlmClient) SetServerConfig(ctx context.Context, req *SetServerConfigRequest) (*SetServerConfigReply, error) {
	var resp SetServerConfigReply

	err := c.Post(ctx, "/index/api/setServerConfig", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
