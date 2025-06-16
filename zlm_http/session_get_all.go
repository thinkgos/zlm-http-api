package zlm_http

import "context"

//* 获取所有 TcpSession 列表(获取所有 tcp 客户端相关信息)
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_9%E3%80%81-index-api-getallsession

type GetAllSessionRequest struct {
	Secret    string `json:"secret,omitempty"`     // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	LocalPort int    `json:"local_port,omitempty"` // O, 筛选本机端口, 例如筛选rtsp链接: 554
	PeerIp    string `json:"peer_ip,omitempty"`    // O, 筛选客户端ip
}
type GetAllSessionReply struct {
	BaseResult
	Data []SessionEntry `json:"data"`
}

type SessionEntry struct {
	Id         string `json:"id"`         // 该tcp链接唯一id
	Identifier string `json:"identifier"` // 识别码
	LocalIP    string `json:"local_ip"`   // 本机网卡ip
	LocalPort  int    `json:"local_port"` // 本机端口号(这是个rtmp播放器或推流器)
	PeerIp     string `json:"peer_ip"`    // 客户端ip
	PeerPort   int    `json:"peer_port"`  // 客户端端口号
	Typeid     string `json:"typeid"`     // 客户端TCPSession typeid
}

func (c *Client) GetAllSession(ctx context.Context, req *GetAllSessionRequest, opts ...CallOption) (*GetAllSessionReply, error) {
	return GenericGet[GetAllSessionRequest, GetAllSessionReply](c, "/index/api/getAllSession", ctx, req, opts...)
}
