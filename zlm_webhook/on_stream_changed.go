package zlm_webhook

//* rtsp/rtmp 流注册或注销时触发此事件; 此事件对回复不敏感.
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_12%E3%80%81on-stream-changed

type OnStreamChangedRequest struct {
	MediaServerId string `json:"mediaServerId"` // 服务器id, 通过配置文件设置
	App           string `json:"app"`           // 流应用名
	Schema        string `json:"schema"`        // 协议
	Stream        string `json:"stream"`        // 流id
	Vhost         string `json:"vhost"`         // 流虚拟主机
	Regist        bool   `json:"regist"`        // 是否注册
	// 以上注册/注销都有
	// 以下仅注册时有效(regist=true)
	AliveSecond      int        `json:"aliveSecond"`      // 存活时间, 单位秒
	BytesSpeed       int        `json:"bytesSpeed"`       // 数据产生速率, 单位byte/s
	CreateStamp      int        `json:"createStamp"`      // GMT unix系统时间戳, 单位秒
	OriginSock       OriginSock `json:"originSock"`       // 产生源信息
	OriginType       int        `json:"originType"`       // 产生源类型，包括 unknown = 0,rtmp_push=1,rtsp_push=2,rtp_push=3,pull=4,ffmpeg_pull=5,mp4_vod=6,device_chn=7,rtc_push=8
	OriginTypeStr    string     `json:"originTypeStr"`    // 产生源类型字符串
	OriginUrl        string     `json:"originUrl"`        // 产生源的url
	ReaderCount      int        `json:"readerCount"`      // 本协议观看人数
	TotalReaderCount int        `json:"totalReaderCount"` // 观看总人数，包括hls/rtsp/rtmp/http-flv/ws-flv/rtc
	Tracks           []Tracks   `json:"tracks"`           // 追踪
}

type OriginSock struct {
	Identifier string `json:"identifier"` // 识别码
	LocalIP    string `json:"local_ip"`   // 本机ip
	LocalPort  int    `json:"local_port"` // 本机端口
	PeerIP     string `json:"peer_ip"`    // 对端ip
	PeerPort   int    `json:"peer_port"`  // 对端端口
}
type Tracks struct {
	CodecId     int     `json:"codec_id"`              // H264 = 0, H265 = 1, AAC = 2, G711A = 3, G711U = 4
	CodecIdName string  `json:"codec_id_name"`         // 编码类型名称
	CodecType   int     `json:"codec_type"`            // Video = 0, Audio = 1
	Ready       bool    `json:"ready"`                 // 轨道是否准备就绪
	Channels    int     `json:"channels,omitempty"`    // 音频通道数
	SampleBit   int     `json:"sample_bit,omitempty"`  // 音频采样位数
	SampleRate  int     `json:"sample_rate,omitempty"` // 音频采样率
	Fps         float64 `json:"fps,omitempty"`         // 视频fps
	Height      int     `json:"height,omitempty"`      // 视频高
	Width       int     `json:"width,omitempty"`       // 视频宽
}

type OnStreamChangedReply = DefaultReply
