package zlm_def

// bool定义
const (
	False = 0
	True  = 1
)

// schema定义
const (
	Schema_RTMP  = "rtmp"
	Schema_RTMPS = "rtmps"
	Schema_RTSP  = "rtsp"
	Schema_RTSPS = "rtsps"
	Schema_RTP   = "rtp"
	Schema_HTTP  = "http"
	Schema_HTTPS = "https"
	Schema_WS    = "ws"
	Schema_WSS   = "wss"
)

// 虚拟主机
const (
	Vhost_Default = "__defaultVhost__"
)

// 编码id
const (
	CodecId_H264  = 0 // H264
	CodecId_H265  = 1 // H265
	CodecId_AAC   = 2 // AAC
	CodecId_G711A = 3 // G711A
	CodecId_G711U = 4 // G711U
)

func CodecIdString(codecId int) string {
	switch codecId {
	case CodecId_H264:
		return "H264"
	case CodecId_H265:
		return "H265"
	case CodecId_AAC:
		return "AAC"
	case CodecId_G711A:
		return "G711A"
	case CodecId_G711U:
		return "G711U"
	}
	return "undefined"
}

// 编解码类型
const (
	CodecType_Video = 0 // Video
	CodecType_Audio = 1 // Audio
)

// rtp的TcpMode
const (
	Rtp_TcpMode_Udp        = 0 // udp模式
	Rtp_TcpMode_TcpPassive = 1 // tcp被动模式
	Rtp_TcpMode_TcpActive  = 2 // tcp主动模式
)

// rtsp的推/拉流时的方式
const (
	Rtp_Type_Tcp       = 0 // tcp
	Rtp_Type_Udp       = 1 // udp
	Rtp_Type_Multicast = 2 // 组播
)

// 录制类型
const (
	Record_Type_Hls = 0 // hls
	Record_Type_Mp4 = 1 // mp4
)

const (
	OriginType_Unknown    = 0 // unknown
	OriginType_RtmpPush   = 1 // rtmp_push
	OriginType_RtspPush   = 2 // rtsp_push
	OriginType_RtpPush    = 3 // rtp_push
	OriginType_Pull       = 4 // pull
	OriginType_FfmpegPull = 5 // ffmpeg_pull
	OriginType_Mp4Vod     = 6 // mp4_vod
	OriginType_DeviceChn  = 7 // device_chn
	OriginType_RtcPush    = 8 // rtc_push
)
