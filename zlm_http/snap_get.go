package zlm_http

import (
	"context"
	"net/http"
)

//* 获取截图或生成实时截图并返回
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_23%E3%80%81-index-api-getsnap

type GetSnapRequest struct {
	Secret     string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Url        string `json:"url"`              // M, 需要截图的url, 可以是本机的, 也可以是远程主机的.
	TimeoutSec int    `json:"timeout_sec"`      // M, 截图失败超时时间, 防止FFmpeg一直等待截图
	ExpireSec  int    `json:"expire_sec"`       // M, 截图的过期时间, 该时间内产生的截图都会作为缓存返回
	Filename   string `json:"-"`                // M, 本地保存的位置
}

func (c *Client) GetSnap(ctx context.Context, req *GetSnapRequest, opts ...CallOption) error {
	return c.Download(ctx, http.MethodGet, "/index/api/getSnap", req, req.Filename, opts...)
}
