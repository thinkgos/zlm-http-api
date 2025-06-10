package zlm_http

import "context"

//* 断开tcp连接, 比如说可以断开rtsp、rtmp播放器等
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_11%E3%80%81-index-api-kick-sessions

type KickSessionsRequest struct {
	Secret    string `json:"secret,omitempty"`     // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	LocalPort int    `json:"local_port,omitempty"` // O, 筛选本机端口, 例如筛选rtsp链接: 554
	PeerIp    string `json:"peer_ip,omitempty"`    // O, 筛选客户端ip
}
type KickSessionsReply struct {
	BaseResult
	CountHit int `json:"count_hit"` // 筛选命中客户端个数
}

func (c *Client) KickSessions(ctx context.Context, req *KickSessionsRequest, opts ...CallOption) (*KickSessionsReply, error) {
	return genericPost[KickSessionsRequest, KickSessionsReply](c, "/index/api/kick_sessions", ctx, req, opts...)
}
