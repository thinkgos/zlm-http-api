package zlm_http

import (
	"context"

	"github.com/thinkgos/zlm-http-api/zlm_def"
)

//* 获取服务器配置
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_3%E3%80%81-index-api-getserverconfig

type GetServerConfigRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type GetServerConfigReply struct {
	BaseResult
	Data []zlm_def.ServerConfig `json:"data"`
}

func (c *Client) GetServerConfig(ctx context.Context, req *GetServerConfigRequest, opts ...CallOption) (*GetServerConfigReply, error) {
	return genericGet[GetServerConfigRequest, GetServerConfigReply](c, "/index/api/getServerConfig", ctx, req, opts...)
}
