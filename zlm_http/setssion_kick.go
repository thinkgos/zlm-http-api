package zlm_http

import "context"

//* 断开tcp连接, 比如说可以断开rtsp、rtmp播放器等
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_10%E3%80%81-index-api-kick-session

type KickSessionRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Id     string `json:"id"`               // M, 客户端唯一 id, 可以通过 KickSession 接口获取
}
type KickSessionReply = BaseResult

func (c *Client) KickSession(ctx context.Context, req *KickSessionRequest, opts ...CallOption) (*KickSessionReply, error) {
	return genericPost[KickSessionRequest, KickSessionReply](c, "/index/api/kick_session", ctx, req, opts...)
}
