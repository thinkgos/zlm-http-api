package zlm_http

import "context"

//* 获取 rtp 代理时的某路 ssrc rtp 信息
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_18%E3%80%81-index-api-getrtpinfo

type GetRtpInfoRequest struct {
	Secret   string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	StreamId string `json:"stream_id"`        // M, RTP 的 ssrc，16 进制字符串或者是流的 id(openRtpServer 接口指定)
}
type GetRtpInfoReply struct {
	BaseResult
	Exist     bool   `json:"exist"`      // 是否存在
	LocalIp   string `json:"local_ip"`   // 本地监听的ip
	LocalPort int    `json:"local_port"` // 本地监听的端口
	PeerIp    string `json:"peer_ip"`    // 推流客户端ip
	PeerPort  int    `json:"peer_port"`  // 推流客户端端口号
}

func (c *Client) GetRtpInfo(ctx context.Context, req *GetRtpInfoRequest, opts ...CallOption) (*GetRtpInfoReply, error) {
	return GenericPost[GetRtpInfoRequest, GetRtpInfoReply](c, "/index/api/getRtpInfo", ctx, req, opts...)
}
