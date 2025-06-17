package zlm_http

import "context"

//* 作为GB28181客户端, 启动 ps-rtp 推流, 支持rtp/udp方式.
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_27%E3%80%81-index-api-startsendrtp
//
// 该接口支持 rtsp/rtmp 等协议转 ps-rtp 推流. 第一次推流失败会直接返回错误, 成功一次后, 后续失败也将无限重试.

// TODO: fix me
type StartSendRtpRequest struct {
	Secret        string  `json:"secret,omitempty"`          // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost         string  `json:"vhost"`                     // M, 虚拟主机, 例如__defaultVhost__
	App           string  `json:"app"`                       // M, 应用名, 例如 live
	Stream        string  `json:"stream"`                    // M, 流id, 例如 test
	DstUrl        string  `json:"dst_url"`                   // M, 目标地址
	DstPort       int     `json:"dst_port,string"`           // M, 目标端口
	IsUdp         int     `json:"is_udp,string"`             // M, 0:tcp active模式, 1:udp active模式
	Ssrc          int     `json:"ssrc"`                      // M, 推流的rtp的ssrc, 指定不同的ssrc可以同时推流到多个服务器
	SrcPort       *int    `json:"src_port,string,omitempty"` // O, 指定tcp/udp客户端使用的本地端口, 0或不传:为随机端口.
	FromMp4       *int    `json:"from_mp4,omitempty"`        // O, 是否推送本地MP4录像
	Type          *int    `json:"type,omitempty"`            // O, rtp打包模式, 0: es, 1: ps, 2: ts
	Pt            *string `json:"pt,omitempty"`              // O, rtp payload type, 默认96.
	OnlyAudio     *int    `json:"only_audio,omitempty"`      // O, rtp es方式打包时, 是否只打包音频
	UdpRtpTimeout *int    `json:"udp_rtp_timeout,omitempty"` // O, udp方式推流时, 是否开启rtcp发送和rtcp接收超时判断, 开启后(默认关闭), 如果接收rr rtcp超时, 将导致主动停止rtp发送.
	RecvStreamId  *string `json:"recv_stream_id,omitempty"`  // O, 发送rtp同时接收, 一般用于双向语言对讲, 如果不为空, 说明开启接收, 值为接收流的id
}
type StartSendRtpReply struct {
	BaseResult
	LocalPort int `json:"local_port"` // 使用的本地端口号
}

func (c *Client) StartSendRtp(ctx context.Context, req *StartSendRtpRequest, opts ...CallOption) (*StartSendRtpReply, error) {
	return GenericPost[StartSendRtpRequest, StartSendRtpReply](c, "/index/api/startSendRtp", ctx, req, opts...)
}
