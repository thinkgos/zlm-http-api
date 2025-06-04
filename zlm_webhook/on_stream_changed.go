package zlm_webhook

//* rtsp/rtmp 流注册或注销时触发此事件; 此事件对回复不敏感.
// https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html#_12%E3%80%81on-stream-changed

// 注销
//	{
//		"mediaServerId" : "your_server_id",
//		"app" : "live",
//		"regist" : false,
//		"schema" : "rtsp",
//		"stream" : "obs",
//		"vhost" : "__defaultVhost__"
//	}

// 注册
//
//	{
//	    "regist" : true,
//	    "aliveSecond": 0, #存活时间, 单位秒
//	    "app": "live", # 应用名
//	    "bytesSpeed": 0, #数据产生速度, 单位byte/s
//	    "createStamp": 1617956908,  #GMT unix系统时间戳, 单位秒
//	    "mediaServerId": "your_server_id", # 服务器id
//	    "originSock": {
//	        "identifier": "000001C257D35E40",
//	        "local_ip": "172.26.20.112", # 本机ip
//	        "local_port": 50166, # 本机端口
//	        "peer_ip": "172.26.20.112", # 对端ip
//	        "peer_port": 50155 # 对端port
//	    },
//	    "originType": 8,  # 产生源类型, 包括 unknown = 0,rtmp_push=1,rtsp_push=2,rtp_push=3,pull=4,ffmpeg_pull=5,mp4_vod=6,device_chn=7,rtc_push=8
//	    "originTypeStr": "rtc_push",
//	    "originUrl": "", #产生源的url
//	    "readerCount": 0, # 本协议观看人数
//	    "schema": "rtsp", # 协议
//	    "stream": "test",  # 流id
//	    "totalReaderCount": 0, # 观看总人数, 包括hls/rtsp/rtmp/http-flv/ws-flv/rtc
//	    "tracks": [{
//	       "channels" : 1, # 音频通道数
//	       "codec_id" : 2, # H264 = 0, H265 = 1, AAC = 2, G711A = 3, G711U = 4
//	       "codec_id_name" : "CodecAAC", # 编码类型名称
//	       "codec_type" : 1, # Video = 0, Audio = 1
//	       "ready" : true, # 轨道是否准备就绪
//	       "sample_bit" : 16, # 音频采样位数
//	       "sample_rate" : 8000 # 音频采样率
//	    },
//	    {
//	       "codec_id" : 0, # H264 = 0, H265 = 1, AAC = 2, G711A = 3, G711U = 4
//	       "codec_id_name" : "CodecH264", # 编码类型名称
//	       "codec_type" : 0, # Video = 0, Audio = 1
//	       "fps" : 59,  # 视频fps
//	       "height" : 720, # 视频高
//	       "ready" : true,  # 轨道是否准备就绪
//	       "width" : 1280 # 视频宽
//	    }],
//	    "vhost": "__defaultVhost__"
//	}

type OnStreamChangedRequest struct {
	Regist           bool       `json:"regist"`
	AliveSecond      int        `json:"aliveSecond"`
	App              string     `json:"app"`
	BytesSpeed       int        `json:"bytesSpeed"`
	CreateStamp      int        `json:"createStamp"`
	MediaServerID    string     `json:"mediaServerId"`
	OriginSock       OriginSock `json:"originSock"`
	OriginType       int        `json:"originType"`
	OriginTypeStr    string     `json:"originTypeStr"`
	OriginURL        string     `json:"originUrl"`
	ReaderCount      int        `json:"readerCount"`
	Schema           string     `json:"schema"`
	Stream           string     `json:"stream"`
	TotalReaderCount int        `json:"totalReaderCount"`
	Tracks           []Tracks   `json:"tracks"`
	Vhost            string     `json:"vhost"`
}

type OriginSock struct {
	Identifier string `json:"identifier"`
	LocalIP    string `json:"local_ip"`
	LocalPort  int    `json:"local_port"`
	PeerIP     string `json:"peer_ip"`
	PeerPort   int    `json:"peer_port"`
}
type Tracks struct {
	Channels    int     `json:"channels,omitempty"`
	CodecID     int     `json:"codec_id"`
	CodecIDName string  `json:"codec_id_name"`
	CodecType   int     `json:"codec_type"`
	Ready       bool    `json:"ready"`
	SampleBit   int     `json:"sample_bit,omitempty"`
	SampleRate  int     `json:"sample_rate,omitempty"`
	Fps         float32 `json:"fps,omitempty"`
	Height      int     `json:"height,omitempty"`
	Width       int     `json:"width,omitempty"`
}

type OnStreamChangedReply = DefaultReply
