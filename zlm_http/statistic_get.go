package zlm_http

import "context"

//* 获取主要对象个数统计, 主要用于分析内存性能.
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_29%E3%80%81-index-api-getstatistic

type GetStatisticRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type GetStatisticReply struct {
	BaseResult
	Data *StatisticData `json:"data"`
}
type StatisticData struct {
	Buffer                int `json:"Buffer"`
	BufferLikeString      int `json:"BufferLikeString"`
	BufferList            int `json:"BufferList"`
	BufferRaw             int `json:"BufferRaw"`
	Frame                 int `json:"Frame"`
	FrameImp              int `json:"FrameImp"`
	MediaSource           int `json:"MediaSource"`
	MultiMediaSourceMuxer int `json:"MultiMediaSourceMuxer"`
	Socket                int `json:"Socket"`
	TcpClient             int `json:"TcpClient"`
	TcpServer             int `json:"TcpServer"`
	TcpSession            int `json:"TcpSession"`
	UdpServer             int `json:"UdpServer"`
	UdpSession            int `json:"UdpSession"`
	RtmpPacket            int `json:"RtmpPacket"`
	RtpPacket             int `json:"RtpPacket"`
}

func (c *Client) GetStatistic(ctx context.Context, req *GetStatisticRequest, opts ...CallOption) (*GetStatisticReply, error) {
	return GenericGet[GetStatisticRequest, GetStatisticReply](c, "/index/api/getStatistic", ctx, req, opts...)
}
