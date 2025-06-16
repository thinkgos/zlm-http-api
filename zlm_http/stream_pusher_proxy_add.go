package zlm_http

import "context"

//* 添加rtsp/rtmp主动推流代理(把本服务器的直播流推送到其他服务器去)

type AddStreamPusherProxyRequest struct {
	Secret     string   `json:"secret,omitempty"`      // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost      string   `json:"vhost"`                 // M, 添加的流的虚拟主机, 例如__defaultVhost__
	Schema     string   `json:"schema"`                // M, 流协议, 例如: rtsp或rtmp
	App        string   `json:"app"`                   // M, 添加的流的应用名, 例如 live
	Stream     string   `json:"stream"`                // M, 添加的流的id名, 例如 test
	DstUrl     string   `json:"dst_url"`               // M, 目标转推url, 带参数需要自行url转义, 需要与schema字段协议一致
	RetryCount *int     `json:"retry_count,omitempty"` // O, 转推失败重试次数, 默认无限重试 传<=0时则无限重试
	TimeoutSec *float64 `json:"timeout_sec,omitempty"` // O, 推流超时时间, 单位秒, float类型
	RtpType    *int     `json:"rtp_type,omitempty"`    // O, rtsp推流时, 推流方式, 0：tcp, 1：udp, 2：组播
	Latency    *int     `json:"latency,omitempty"`     // O, srt延时, 单位毫秒
	Passphrase *string  `json:"passphrase,omitempty"`  // O, srt推流的密码
}
type AddStreamPusherProxyReply struct {
	BaseResult
	Data AddStreamPusherProxyData `json:"data"`
}
type AddStreamPusherProxyData struct {
	Key string `json:"key"` // 唯一key
}

func (c *Client) AddStreamPusherProxy(ctx context.Context, req *AddStreamPusherProxyRequest, opts ...CallOption) (*AddStreamPusherProxyReply, error) {
	return GenericPost[AddStreamPusherProxyRequest, AddStreamPusherProxyReply](c, "/index/api/addStreamPusherProxy", ctx, req, opts...)
}
