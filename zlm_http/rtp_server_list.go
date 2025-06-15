package zlm_http

import "context"

//* 获取 openRtpServer 接口创建的所有 RTP 服务器
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_26%E3%80%81-index-api-listrtpserver

type ListRtpServerRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type ListRtpServerReply struct {
	BaseResult
	Data []RtpServerEntry `json:"data"`
}
type RtpServerEntry struct {
	Vhost     string `json:"vhost"`      // 虚拟主机
	App       string `json:"app"`        // 流应用名
	StreamId  string `json:"stream_id"`  // 流id
	TcpMode   int    `json:"tcp_mode"`   // 0: udp 模式, 1: tcp被动模式, 2: tcp 主动模式(兼容 enable_tcp为0/1)
	Port      int    `json:"port"`       // 端口号
	OnlyTrack int    `json:"only_track"` // 流过滤, 0: 全部, 1: 只音频, 2: 只视频
	Ssrc      int    `json:"ssrc"`       // ssrc
}

func (c *Client) ListRtpServer(ctx context.Context, req *ListRtpServerRequest, opts ...CallOption) (*ListRtpServerReply, error) {
	return GenericGet[ListRtpServerRequest, ListRtpServerReply](c, "/index/api/listRtpServer", ctx, req, opts...)
}
