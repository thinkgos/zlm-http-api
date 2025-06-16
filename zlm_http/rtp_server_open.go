package zlm_http

import "context"

//* 创建GB28181 RTP 接收端口, 如果该端口接收数据超时, 则会自动被回收(不用调用 closeRtpServer 接口)
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_24%E3%80%81-index-api-openrtpserver

type OpenRtpServerRequest struct {
	Secret    string `json:"secret,omitempty"`      // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost     string `json:"vhost,omitempty"`       // O, 虚拟主机, 默认: __defaultVhost__
	App       string `json:"app,omitempty"`         // O, 流应用名, 默认: rtp
	StreamId  string `json:"stream_id"`             // M, 该端口绑定的流id, 该端口只能创建这一个流(而不是根据ssrc创建多个)
	TcpMode   int    `json:"tcp_mode"`              // M, 0: udp 模式, 1: tcp被动模式, 2: tcp 主动模式(兼容 enable_tcp为0/1)
	Port      int    `json:"port"`                  // M, 接收端口, 0则为随机端口
	OnlyTrack int    `json:"only_track,omitempty"`  // O, 流过滤, 0: 不设置(全部), 1: 单音频, 2: 单视频
	Ssrc      int    `json:"ssrc,omitempty"`        // O, ssrc, 默认为0
	LocalIp   string `json:"local_ip,omitempty"`    // O, 指定创建RTP的本地ip, ipv4可填"0.0.0.0", ipv6可填"::", 一般保持默认
	ReUsePort bool   `json:"re_use_port,omitempty"` // O, 是否重用端口, 默认为0, 非必选参数
}
type OpenRtpServerReply struct {
	BaseResult
	Port int `json:"port"` // 接收端口, 方便获取随机端口号
}

func (c *Client) OpenRtpServer(ctx context.Context, req *OpenRtpServerRequest, opts ...CallOption) (*OpenRtpServerReply, error) {
	return GenericPost[OpenRtpServerRequest, OpenRtpServerReply](c, "/index/api/openRtpServer", ctx, req, opts...)
}
