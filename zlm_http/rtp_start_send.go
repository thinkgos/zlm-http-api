package zlm_http

import "context"

//* 作为GB28181客户端, 启动 ps-rtp 推流, 支持rtp/udp方式.
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_27%E3%80%81-index-api-startsendrtp
//
// 该接口支持 rtsp/rtmp 等协议转 ps-rtp 推流. 第一次推流失败会直接返回错误, 成功一次后, 后续失败也将无限重试.

// TODO: fix me
type StartSendRtpRequest struct {
	Secret  string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost   string `json:"vhost"`            // M, 虚拟主机, 例如__defaultVhost__
	App     string `json:"app"`              // M, 应用名, 例如 live
	Stream  string `json:"stream"`           // M, 流id, 例如 test
	Ssrc    string `json:"ssrc"`             // O, 推流的rtp的ssrc, 指定不同的ssrc可以同时推流到多个服务器
	DstUrl  string `json:"dst_url"`
	DstPort int    `json:"dst_port,string"`
	IsUdp   int    `json:"is_udp,string"`
	SrcPort int    `json:"src_port,string"`
}
type StartSendRtpReply struct {
	BaseResult
	LocalPort int `json:"local_port"` // 使用的本地端口号
}

func (c *Client) StartSendRtp(ctx context.Context, req *StartSendRtpRequest, opts ...CallOption) (*StartSendRtpReply, error) {
	return genericPost[StartSendRtpRequest, StartSendRtpReply](c, "/index/api/startSendRtp", ctx, req, opts...)
}
