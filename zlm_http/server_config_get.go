package zlm_http

import (
	"context"
)

//* 获取服务器配置
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_3%E3%80%81-index-api-getserverconfig

type GetServerConfigRequest struct {
	Secret string `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
}
type GetServerConfigReply struct {
	BaseResult
	Data []ServerConfigData `json:"data"`
}

type ServerConfigData struct {
	Api_ApiDebug                        string `json:"api.apiDebug"`
	Api_DefaultSnap                     string `json:"api.defaultSnap"`
	Api_DownloadRoot                    string `json:"api.downloadRoot"`
	Api_Secret                          string `json:"api.secret"`
	Api_SnapRoot                        string `json:"api.snapRoot"`
	Cluster_OriginUrl                   string `json:"cluster.origin_url"`
	Cluster_RetryCount                  string `json:"cluster.retry_count"`
	Cluster_TimeoutSec                  string `json:"cluster.timeout_sec"`
	Ffmpeg_Bin                          string `json:"ffmpeg.bin"`
	Ffmpeg_Cmd                          string `json:"ffmpeg.cmd"`
	Ffmpeg_Log                          string `json:"ffmpeg.log"`
	Ffmpeg_RestartSec                   string `json:"ffmpeg.restart_sec"`
	Ffmpeg_Snap                         string `json:"ffmpeg.snap"`
	General_BroadcastPlayerCountChanged string `json:"general.broadcast_player_count_changed"`
	General_CheckNvidiaDev              string `json:"general.check_nvidia_dev"`
	General_EnableVhost                 string `json:"general.enableVhost"`
	General_EnableFfmpegLog             string `json:"general.enable_ffmpeg_log"`
	General_FlowThreshold               string `json:"general.flowThreshold"`
	General_ListenIP                    string `json:"general.listen_ip"`
	General_MaxStreamWaitMs             string `json:"general.maxStreamWaitMS"`
	General_MediaServerId               string `json:"general.mediaServerId"`
	General_MergeWriteMs                string `json:"general.mergeWriteMS"`
	General_ResetWhenRePlay             string `json:"general.resetWhenRePlay"`
	General_StreamNoneReaderDelayMs     string `json:"general.streamNoneReaderDelayMS"`
	General_UnreadyFrameCache           string `json:"general.unready_frame_cache"`
	General_WaitAddTrackMs              string `json:"general.wait_add_track_ms"`
	General_WaitAudioTrackDataMs        string `json:"general.wait_audio_track_data_ms"`
	General_WaitTrackReadyMs            string `json:"general.wait_track_ready_ms"`
	Hls_BroadcastRecordTs               string `json:"hls.broadcastRecordTs"`
	Hls_DeleteDelaySec                  string `json:"hls.deleteDelaySec"`
	Hls_FastRegister                    string `json:"hls.fastRegister"`
	Hls_FileBufSize                     string `json:"hls.fileBufSize"`
	Hls_SegDelay                        string `json:"hls.segDelay"`
	Hls_SegDur                          string `json:"hls.segDur"`
	Hls_SegKeep                         string `json:"hls.segKeep"`
	Hls_SegNum                          string `json:"hls.segNum"`
	Hls_SegRetain                       string `json:"hls.segRetain"`
	Hook_AliveInterval                  string `json:"hook.alive_interval"`
	Hook_Enable                         string `json:"hook.enable"`
	Hook_OnFlowReport                   string `json:"hook.on_flow_report"`
	Hook_OnHttpAccess                   string `json:"hook.on_http_access"`
	Hook_OnPlay                         string `json:"hook.on_play"`
	Hook_OnPublish                      string `json:"hook.on_publish"`
	Hook_OnRecordMp4                    string `json:"hook.on_record_mp4"`
	Hook_OnRecordTs                     string `json:"hook.on_record_ts"`
	Hook_OnRtpServerTimeout             string `json:"hook.on_rtp_server_timeout"`
	Hook_OnRtspAuth                     string `json:"hook.on_rtsp_auth"`
	Hook_OnRtspRealm                    string `json:"hook.on_rtsp_realm"`
	Hook_OnSendRtpStopped               string `json:"hook.on_send_rtp_stopped"`
	Hook_OnServerExited                 string `json:"hook.on_server_exited"`
	Hook_OnServerKeepalive              string `json:"hook.on_server_keepalive"`
	Hook_OnServerStarted                string `json:"hook.on_server_started"`
	Hook_OnShellLogin                   string `json:"hook.on_shell_login"`
	Hook_OnStreamChanged                string `json:"hook.on_stream_changed"`
	Hook_OnStreamNoneReader             string `json:"hook.on_stream_none_reader"`
	Hook_OnStreamNotFound               string `json:"hook.on_stream_not_found"`
	Hook_Retry                          string `json:"hook.retry"`
	Hook_RetryDelay                     string `json:"hook.retry_delay"`
	Hook_StreamChangedSchemas           string `json:"hook.stream_changed_schemas"`
	Hook_TimeoutSec                     string `json:"hook.timeoutSec"`
	Http_AllowCrossDomains              string `json:"http.allow_cross_domains"`
	Http_AllowIPRange                   string `json:"http.allow_ip_range"`
	Http_CharSet                        string `json:"http.charSet"`
	Http_DirMenu                        string `json:"http.dirMenu"`
	Http_ForbidCacheSuffix              string `json:"http.forbidCacheSuffix"`
	Http_ForwardedIPHeader              string `json:"http.forwarded_ip_header"`
	Http_KeepAliveSecond                string `json:"http.keepAliveSecond"`
	Http_MaxReqSize                     string `json:"http.maxReqSize"`
	Http_NotFound                       string `json:"http.notFound"`
	Http_Port                           int    `json:"http.port,string"`
	Http_RootPath                       string `json:"http.rootPath"`
	Http_SendBufSize                    string `json:"http.sendBufSize"`
	Http_Sslport                        int    `json:"http.sslport,string"`
	Http_VirtualPath                    string `json:"http.virtualPath"`
	Multicast_AddrMax                   string `json:"multicast.addrMax"`
	Multicast_AddrMin                   string `json:"multicast.addrMin"`
	Multicast_UdpTtl                    string `json:"multicast.udpTTL"`
	Protocol_AddMuteAudio               string `json:"protocol.add_mute_audio"`
	Protocol_AutoClose                  string `json:"protocol.auto_close"`
	Protocol_ContinuePushMs             string `json:"protocol.continue_push_ms"`
	Protocol_EnableAudio                string `json:"protocol.enable_audio"`
	Protocol_EnableFmp4                 string `json:"protocol.enable_fmp4"`
	Protocol_EnableHls                  string `json:"protocol.enable_hls"`
	Protocol_EnableHlsFmp4              string `json:"protocol.enable_hls_fmp4"`
	Protocol_EnableMp4                  string `json:"protocol.enable_mp4"`
	Protocol_EnableRtmp                 string `json:"protocol.enable_rtmp"`
	Protocol_EnableRtsp                 string `json:"protocol.enable_rtsp"`
	Protocol_EnableTs                   string `json:"protocol.enable_ts"`
	Protocol_Fmp4Demand                 string `json:"protocol.fmp4_demand"`
	Protocol_HlsDemand                  string `json:"protocol.hls_demand"`
	Protocol_HlsSavePath                string `json:"protocol.hls_save_path"`
	Protocol_ModifyStamp                string `json:"protocol.modify_stamp"`
	Protocol_Mp4AsPlayer                string `json:"protocol.mp4_as_player"`
	Protocol_Mp4MaxSecond               string `json:"protocol.mp4_max_second"`
	Protocol_Mp4SavePath                string `json:"protocol.mp4_save_path"`
	Protocol_PacedSenderMs              string `json:"protocol.paced_sender_ms"`
	Protocol_RtmpDemand                 string `json:"protocol.rtmp_demand"`
	Protocol_RtspDemand                 string `json:"protocol.rtsp_demand"`
	Protocol_TsDemand                   string `json:"protocol.ts_demand"`
	Record_AppName                      string `json:"record.appName"`
	Record_EnableFmp4                   string `json:"record.enableFmp4"`
	Record_FastStart                    string `json:"record.fastStart"`
	Record_FileBufSize                  string `json:"record.fileBufSize"`
	Record_FileRepeat                   string `json:"record.fileRepeat"`
	Record_SampleMS                     string `json:"record.sampleMS"`
	Rtc_Bfilter                         string `json:"rtc.bfilter"`
	Rtc_DatachannelEcho                 string `json:"rtc.datachannel_echo"`
	Rtc_ExternIP                        string `json:"rtc.externIP"`
	Rtc_MaxRtpCacheMS                   string `json:"rtc.maxRtpCacheMS"`
	Rtc_MaxRtpCacheSize                 string `json:"rtc.maxRtpCacheSize"`
	Rtc_MaxBitrate                      string `json:"rtc.max_bitrate"`
	Rtc_MinBitrate                      string `json:"rtc.min_bitrate"`
	Rtc_NackIntervalRatio               string `json:"rtc.nackIntervalRatio"`
	Rtc_NackMaxCount                    string `json:"rtc.nackMaxCount"`
	Rtc_NackMaxMS                       string `json:"rtc.nackMaxMS"`
	Rtc_NackMaxSize                     string `json:"rtc.nackMaxSize"`
	Rtc_NackRtpSize                     string `json:"rtc.nackRtpSize"`
	Rtc_Port                            string `json:"rtc.port"`
	Rtc_PreferredCodecA                 string `json:"rtc.preferredCodecA"`
	Rtc_PreferredCodecV                 string `json:"rtc.preferredCodecV"`
	Rtc_RembBitRate                     string `json:"rtc.rembBitRate"`
	Rtc_StartBitrate                    string `json:"rtc.start_bitrate"`
	Rtc_TCPPort                         string `json:"rtc.tcpPort"`
	Rtc_TimeoutSec                      string `json:"rtc.timeoutSec"`
	Rtmp_DirectProxy                    string `json:"rtmp.directProxy"`
	Rtmp_Enhanced                       string `json:"rtmp.enhanced"`
	Rtmp_HandshakeSecond                string `json:"rtmp.handshakeSecond"`
	Rtmp_KeepAliveSecond                string `json:"rtmp.keepAliveSecond"`
	Rtmp_Port                           int    `json:"rtmp.port,string"`
	Rtmp_Sslport                        int    `json:"rtmp.sslport,string"`
	Rtp_AudioMtuSize                    string `json:"rtp.audioMtuSize"`
	Rtp_H264StapA                       string `json:"rtp.h264_stap_a"`
	Rtp_LowLatency                      string `json:"rtp.lowLatency"`
	Rtp_RtpMaxSize                      string `json:"rtp.rtpMaxSize"`
	Rtp_VideoMtuSize                    string `json:"rtp.videoMtuSize"`
	RtpProxy_DumpDir                    string `json:"rtp_proxy.dumpDir"`
	RtpProxy_GopCache                   string `json:"rtp_proxy.gop_cache"`
	RtpProxy_H264Pt                     string `json:"rtp_proxy.h264_pt"`
	RtpProxy_H265Pt                     string `json:"rtp_proxy.h265_pt"`
	RtpProxy_OpusPt                     string `json:"rtp_proxy.opus_pt"`
	RtpProxy_Port                       int    `json:"rtp_proxy.port,string"`
	RtpProxy_PortRange                  string `json:"rtp_proxy.port_range"`
	RtpProxy_PsPt                       string `json:"rtp_proxy.ps_pt"`
	RtpProxy_RtpG711DurMs               string `json:"rtp_proxy.rtp_g711_dur_ms"`
	RtpProxy_TimeoutSec                 string `json:"rtp_proxy.timeoutSec"`
	RtpProxy_UDPRecvSocketBuffer        string `json:"rtp_proxy.udp_recv_socket_buffer"`
	Rtsp_AuthBasic                      string `json:"rtsp.authBasic"`
	Rtsp_DirectProxy                    string `json:"rtsp.directProxy"`
	Rtsp_HandshakeSecond                string `json:"rtsp.handshakeSecond"`
	Rtsp_KeepAliveSecond                string `json:"rtsp.keepAliveSecond"`
	Rtsp_LowLatency                     string `json:"rtsp.lowLatency"`
	Rtsp_Port                           int    `json:"rtsp.port,string"`
	Rtsp_RtpTransportType               string `json:"rtsp.rtpTransportType"`
	Rtsp_Sslport                        int    `json:"rtsp.sslport,string"`
	Shell_MaxReqSize                    string `json:"shell.maxReqSize"`
	Shell_Port                          string `json:"shell.port"`
	Srt_LatencyMul                      string `json:"srt.latencyMul"`
	Srt_PassPhrase                      string `json:"srt.passPhrase"`
	Srt_PktBufSize                      string `json:"srt.pktBufSize"`
	Srt_Port                            string `json:"srt.port"`
	Srt_TimeoutSec                      string `json:"srt.timeoutSec"`
}

func (c *Client) GetServerConfig(ctx context.Context, req *GetServerConfigRequest, opts ...CallOption) (*GetServerConfigReply, error) {
	var resp GetServerConfigReply

	err := c.Get(ctx, "/index/api/getServerConfig", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
