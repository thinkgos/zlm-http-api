package zlm_def

// OriginSock 客户端和服务器网络信息
type OriginSock struct {
	Identifier string `json:"identifier"` // 识别码
	LocalIp    string `json:"local_ip"`   // 本机ip
	LocalPort  int    `json:"local_port"` // 本机端口
	PeerIp     string `json:"peer_ip"`    // 对端ip
	PeerPort   int    `json:"peer_port"`  // 对端端口
}

// Tracks 音视频轨道
type Tracks struct {
	CodecId       int     `json:"codec_id"`                  // H264 = 0, H265 = 1, AAC = 2, G711A = 3, G711U = 4
	CodecIdName   string  `json:"codec_id_name"`             // 编码类型名称
	CodecType     int     `json:"codec_type"`                // Video = 0, Audio = 1
	Ready         bool    `json:"ready"`                     // 轨道是否准备就绪
	Channels      int     `json:"channels,omitempty"`        // 音频通道数
	SampleBit     int     `json:"sample_bit,omitempty"`      // 音频采样位数
	SampleRate    int     `json:"sample_rate,omitempty"`     // 音频采样率
	Fps           float64 `json:"fps,omitempty"`             // 视频fps
	Height        int     `json:"height,omitempty"`          // 视频高
	Width         int     `json:"width,omitempty"`           // 视频宽
	GopIntervalMs int     `json:"gop_interval_ms,omitempty"` // gop大小, 单位帧数
	GopSize       int     `json:"gop_size,omitempty"`        // gop大小, 单位帧数
	KeyFrames     int     `json:"key_frames,omitempty"`      // 累计接收关键帧数
}
