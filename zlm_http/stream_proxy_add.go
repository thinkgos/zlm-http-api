package zlm_http

import "context"

//* 动态添加拉流代理rtsp/rtmp/hls/http-ts/http-flv(只支持 H264/H265/aac/G711/opus 负载)
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_12%E3%80%81-index-api-addstreamproxy

type AddStreamProxyRequest struct {
	Secret        string   `json:"secret,omitempty"`          // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Vhost         string   `json:"vhost"`                     // M, 添加的流的虚拟主机, 例如__defaultVhost__
	App           string   `json:"app"`                       // M, 添加的流的应用名, 例如 live
	Stream        string   `json:"stream"`                    // M, 添加的流的id名, 例如 test
	Url           string   `json:"url"`                       // M, 拉流地址, 例如: rtmp://live.hkstv.hk.lxdns.com/live/hks2
	RetryCount    *int     `json:"retry_count,omitempty"`     // O, 拉流重试次数, 默认为-1无限重试, 传<=0时则无限重试
	TimeoutSec    *float64 `json:"timeout_sec,omitempty"`     // O, 拉流超时时间, 单位秒, float类型
	RtpType       *int     `json:"rtp_type,omitempty"`        // O, rtsp拉流时, 拉流方式, 0：tcp, 1：udp, 2：组播
	EnableHls     *bool    `json:"enable_hls,omitempty"`      // O, 是否转 hls-mpegts 协议
	EnableHlsFmp4 *bool    `json:"enable_hls_fmp4,omitempty"` // O, 是否转 hls-fmp4 协议
	EnableMp4     *bool    `json:"enable_mp4,omitempty"`      // O, 是否 mp4 录制
	EnableRtsp    *bool    `json:"enable_rtsp,omitempty"`     // O, 是否转 rtsp/webrtc 协议
	EnableRtmp    *bool    `json:"enable_rtmp,omitempty"`     // O, 是否转 rtmp/flv 协议
	EnableTs      *bool    `json:"enable_ts,omitempty"`       // O, 是否转 http-ts/ws-ts 协议
	EnableFmp4    *bool    `json:"enable_fmp4,omitempty"`     // O, 是否转 http-fmp4/ws-fmp4 协议
	HlsDemand     *bool    `json:"hls_demand,omitempty"`      // O, 该协议是否有人观看才生成hls
	RtspDemand    *bool    `json:"rtsp_demand,omitempty"`     // O, 该协议是否有人观看才生成rtsp
	RtmpDemand    *bool    `json:"rtmp_demand,omitempty"`     // O, 该协议是否有人观看才生成rtmp
	TsDemand      *bool    `json:"ts_demand,omitempty"`       // O, 该协议是否有人观看才生成ts
	Fmp4Demand    *bool    `json:"fmp4_demand,omitempty"`     // O, 该协议是否有人观看才生成fmp4
	EnableAudio   *bool    `json:"enable_audio,omitempty"`    // O, 转协议时, 是否开启音频
	AddMuteAudio  *bool    `json:"add_mute_audio,omitempty"`  // O, 转协议时, 无音频是否添加静音aac音频
	Mp4SavePath   *string  `json:"mp4_save_path,omitempty"`   // O, mp4 录制文件保存根目录, 置空使用默认
	Mp4MaxSecond  *int     `json:"mp4_max_second,omitempty"`  // O, mp4 录制切片大小, 单位秒
	Mp4AsPlayer   *bool    `json:"mp4_as_player,omitempty"`   // O, mp4 录制是否当作观看者参与播放人数计数
	HlsSavePath   *string  `json:"hls_save_path,omitempty"`   // O, hls 文件保存保存根目录, 置空使用默认
	ModifyStamp   *int     `json:"modify_stamp,omitempty"`    // O, 是否开启时间戳覆盖, 默认值2(0:源视频绝对时间戳,不做任何改变,1:采用ZLMediaKit接收数据时的系统时间戳(有平滑处理);,2:.采用源视频流时间戳相对时间戳(增长量), 有做时间戳跳跃和回退矫正)
	AutoClose     *bool    `json:"auto_close,omitempty"`      // O, 无人观看是否自动关闭流(不触发无人观看on_none_reader)
	Latency       *int     `json:"latency,omitempty"`         // O, srt延时, 单位毫秒
	Passphrase    *string  `json:"passphrase,omitempty"`      // O, srt拉流的密码
}
type AddStreamProxyReply struct {
	BaseResult
	Data AddStreamProxyData `json:"data"`
}
type AddStreamProxyData struct {
	Key string `json:"key"` // 唯一key
}

func (c *Client) AddStreamProxy(ctx context.Context, req *AddStreamProxyRequest, opts ...CallOption) (*AddStreamProxyReply, error) {
	return GenericPost[AddStreamProxyRequest, AddStreamProxyReply](c, "/index/api/addStreamProxy", ctx, req, opts...)
}
