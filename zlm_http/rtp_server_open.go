package zlm_http

import "context"

//* 创建GB28181 RTP 接收端口, 如果该端口接收数据超时, 则会自动被回收(不用调用 closeRtpServer 接口)
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_24%E3%80%81-index-api-openrtpserver

type OpenRtpServerRequest struct {
	Secret   string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Port     int    `json:"port"`             // M, 接收端口, 0则为随机端口
	TcpMode  int8   `json:"tcp_mode"`         // M, 0: udp 模式, 1: tcp被动模式, 2: tcp 主动模式(兼容 enable_tcp为0/1)
	StreamId string `json:"stream_id"`        // M, 该端口绑定的流id, 该端口只能创建这一个流(而不是根据ssrc创建多个)
}
type OpenRtpServerReply struct {
	BaseResult
	Port int `json:"port"` // 接收端口, 方便获取随机端口号
}

func (c *ZlmClient) OpenRtpServer(ctx context.Context, req *OpenRtpServerRequest, opts ...CallOption) (*OpenRtpServerReply, error) {
	var resp OpenRtpServerReply

	err := c.Post(ctx, "/index/api/openRtpServer", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
