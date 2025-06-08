package zlm_http

import "context"

//* 设置服务器配置
// https://docs.zlmediakit.com/zh/guide/media_server/restful_api.html#_4%E3%80%81-index-api-setserverconfig

type SetServerConfigRequest struct {
	Secret                              string  `json:"secret,omitempty"` // O, api操作密钥(配置文件配置), 未填则忽略, 可设置全局参数或token来统一传.
	Api_ApiDebug                        *string `json:"api.apiDebug,omitempty"`
	Api_DefaultSnap                     *string `json:"api.defaultSnap,omitempty"`
	Api_DownloadRoot                    *string `json:"api.downloadRoot,omitempty"`
	Api_Secret                          *string `json:"api.secret,omitempty"`
	Api_SnapRoot                        *string `json:"api.snapRoot,omitempty"`
	Cluster_OriginUrl                   *string `json:"cluster.origin_url,omitempty"`
	Cluster_RetryCount                  *string `json:"cluster.retry_count,omitempty"`
	Cluster_TimeoutSec                  *string `json:"cluster.timeout_sec,omitempty"`
	Ffmpeg_Bin                          *string `json:"ffmpeg.bin,omitempty"`
	Ffmpeg_Cmd                          *string `json:"ffmpeg.cmd,omitempty"`
	Ffmpeg_Log                          *string `json:"ffmpeg.log,omitempty"`
	Ffmpeg_RestartSec                   *string `json:"ffmpeg.restart_sec,omitempty"`
	Ffmpeg_Snap                         *string `json:"ffmpeg.snap,omitempty"`
	General_BroadcastPlayerCountChanged *string `json:"general.broadcast_player_count_changed,omitempty"`
	General_CheckNvidiaDev              *string `json:"general.check_nvidia_dev,omitempty"`
	General_EnableVhost                 *string `json:"general.enableVhost,omitempty"`
	General_EnableFfmpegLog             *string `json:"general.enable_ffmpeg_log,omitempty"`
	General_FlowThreshold               *string `json:"general.flowThreshold,omitempty"`
	General_ListenIP                    *string `json:"general.listen_ip,omitempty"`
	General_MaxStreamWaitMs             *string `json:"general.maxStreamWaitMS,omitempty"`
	General_MediaServerId               *string `json:"general.mediaServerId,omitempty"`
	General_MergeWriteMs                *string `json:"general.mergeWriteMS,omitempty"`
	General_ResetWhenRePlay             *string `json:"general.resetWhenRePlay,omitempty"`
	General_StreamNoneReaderDelayMs     *string `json:"general.streamNoneReaderDelayMS,omitempty"`
	General_UnreadyFrameCache           *string `json:"general.unready_frame_cache,omitempty"`
	General_WaitAddTrackMs              *string `json:"general.wait_add_track_ms,omitempty"`
	General_WaitAudioTrackDataMs        *string `json:"general.wait_audio_track_data_ms,omitempty"`
	General_WaitTrackReadyMs            *string `json:"general.wait_track_ready_ms,omitempty"`
	Hls_BroadcastRecordTs               *string `json:"hls.broadcastRecordTs,omitempty"`
	Hls_DeleteDelaySec                  *string `json:"hls.deleteDelaySec,omitempty"`
	Hls_FastRegister                    *string `json:"hls.fastRegister,omitempty"`
	Hls_FileBufSize                     *string `json:"hls.fileBufSize,omitempty"`
	Hls_SegDelay                        *string `json:"hls.segDelay,omitempty"`
	Hls_SegDur                          *string `json:"hls.segDur,omitempty"`
	Hls_SegKeep                         *string `json:"hls.segKeep,omitempty"`
	Hls_SegNum                          *string `json:"hls.segNum,omitempty"`
	Hls_SegRetain                       *string `json:"hls.segRetain,omitempty"`
	Hook_AliveInterval                  *string `json:"hook.alive_interval,omitempty"`
	Hook_Enable                         *string `json:"hook.enable,omitempty"`
	Hook_OnFlowReport                   *string `json:"hook.on_flow_report,omitempty"`
	Hook_OnHttpAccess                   *string `json:"hook.on_http_access,omitempty"`
	Hook_OnPlay                         *string `json:"hook.on_play,omitempty"`
	Hook_OnPublish                      *string `json:"hook.on_publish,omitempty"`
	Hook_OnRecordMp4                    *string `json:"hook.on_record_mp4,omitempty"`
	Hook_OnRecordTs                     *string `json:"hook.on_record_ts,omitempty"`
	Hook_OnRtpServerTimeout             *string `json:"hook.on_rtp_server_timeout,omitempty"`
	Hook_OnRtspAuth                     *string `json:"hook.on_rtsp_auth,omitempty"`
	Hook_OnRtspRealm                    *string `json:"hook.on_rtsp_realm,omitempty"`
	Hook_OnSendRtpStopped               *string `json:"hook.on_send_rtp_stopped,omitempty"`
	Hook_OnServerExited                 *string `json:"hook.on_server_exited,omitempty"`
	Hook_OnServerKeepalive              *string `json:"hook.on_server_keepalive,omitempty"`
	Hook_OnServerStarted                *string `json:"hook.on_server_started,omitempty"`
	Hook_OnShellLogin                   *string `json:"hook.on_shell_login,omitempty"`
	Hook_OnStreamChanged                *string `json:"hook.on_stream_changed,omitempty"`
	Hook_OnStreamNoneReader             *string `json:"hook.on_stream_none_reader,omitempty"`
	Hook_OnStreamNotFound               *string `json:"hook.on_stream_not_found,omitempty"`
	Hook_Retry                          *string `json:"hook.retry,omitempty"`
	Hook_RetryDelay                     *string `json:"hook.retry_delay,omitempty"`
	Hook_StreamChangedSchemas           *string `json:"hook.stream_changed_schemas,omitempty"`
	Hook_TimeoutSec                     *string `json:"hook.timeoutSec,omitempty"`
	Http_AllowCrossDomains              *string `json:"http.allow_cross_domains,omitempty"`
	Http_AllowIPRange                   *string `json:"http.allow_ip_range,omitempty"`
	Http_CharSet                        *string `json:"http.charSet,omitempty"`
	Http_DirMenu                        *string `json:"http.dirMenu,omitempty"`
	Http_ForbidCacheSuffix              *string `json:"http.forbidCacheSuffix,omitempty"`
	Http_ForwardedIPHeader              *string `json:"http.forwarded_ip_header,omitempty"`
	Http_KeepAliveSecond                *string `json:"http.keepAliveSecond,omitempty"`
	Http_MaxReqSize                     *string `json:"http.maxReqSize,omitempty"`
	Http_NotFound                       *string `json:"http.notFound,omitempty"`
	Http_Port                           *string `json:"http.port,omitempty"`
	Http_RootPath                       *string `json:"http.rootPath,omitempty"`
	Http_SendBufSize                    *string `json:"http.sendBufSize,omitempty"`
	Http_Sslport                        *string `json:"http.sslport,omitempty"`
	Http_VirtualPath                    *string `json:"http.virtualPath,omitempty"`
	Multicast_AddrMax                   *string `json:"multicast.addrMax,omitempty"`
	Multicast_AddrMin                   *string `json:"multicast.addrMin,omitempty"`
	Multicast_UdpTtl                    *string `json:"multicast.udpTTL,omitempty"`
	Protocol_AddMuteAudio               *string `json:"protocol.add_mute_audio,omitempty"`
	Protocol_AutoClose                  *string `json:"protocol.auto_close,omitempty"`
	Protocol_ContinuePushMs             *string `json:"protocol.continue_push_ms,omitempty"`
	Protocol_EnableAudio                *string `json:"protocol.enable_audio,omitempty"`
	Protocol_EnableFmp4                 *string `json:"protocol.enable_fmp4,omitempty"`
	Protocol_EnableHls                  *string `json:"protocol.enable_hls,omitempty"`
	Protocol_EnableHlsFmp4              *string `json:"protocol.enable_hls_fmp4,omitempty"`
	Protocol_EnableMp4                  *string `json:"protocol.enable_mp4,omitempty"`
	Protocol_EnableRtmp                 *string `json:"protocol.enable_rtmp,omitempty"`
	Protocol_EnableRtsp                 *string `json:"protocol.enable_rtsp,omitempty"`
	Protocol_EnableTs                   *string `json:"protocol.enable_ts,omitempty"`
	Protocol_Fmp4Demand                 *string `json:"protocol.fmp4_demand,omitempty"`
	Protocol_HlsDemand                  *string `json:"protocol.hls_demand,omitempty"`
	Protocol_HlsSavePath                *string `json:"protocol.hls_save_path,omitempty"`
	Protocol_ModifyStamp                *string `json:"protocol.modify_stamp,omitempty"`
	Protocol_Mp4AsPlayer                *string `json:"protocol.mp4_as_player,omitempty"`
	Protocol_Mp4MaxSecond               *string `json:"protocol.mp4_max_second,omitempty"`
	Protocol_Mp4SavePath                *string `json:"protocol.mp4_save_path,omitempty"`
	Protocol_PacedSenderMs              *string `json:"protocol.paced_sender_ms,omitempty"`
	Protocol_RtmpDemand                 *string `json:"protocol.rtmp_demand,omitempty"`
	Protocol_RtspDemand                 *string `json:"protocol.rtsp_demand,omitempty"`
	Protocol_TsDemand                   *string `json:"protocol.ts_demand,omitempty"`
	Record_AppName                      *string `json:"record.appName,omitempty"`
	Record_EnableFmp4                   *string `json:"record.enableFmp4,omitempty"`
	Record_FastStart                    *string `json:"record.fastStart,omitempty"`
	Record_FileBufSize                  *string `json:"record.fileBufSize,omitempty"`
	Record_FileRepeat                   *string `json:"record.fileRepeat,omitempty"`
	Record_SampleMS                     *string `json:"record.sampleMS,omitempty"`
	Rtc_Bfilter                         *string `json:"rtc.bfilter,omitempty"`
	Rtc_DatachannelEcho                 *string `json:"rtc.datachannel_echo,omitempty"`
	Rtc_ExternIP                        *string `json:"rtc.externIP,omitempty"`
	Rtc_MaxRtpCacheMS                   *string `json:"rtc.maxRtpCacheMS,omitempty"`
	Rtc_MaxRtpCacheSize                 *string `json:"rtc.maxRtpCacheSize,omitempty"`
	Rtc_MaxBitrate                      *string `json:"rtc.max_bitrate,omitempty"`
	Rtc_MinBitrate                      *string `json:"rtc.min_bitrate,omitempty"`
	Rtc_NackIntervalRatio               *string `json:"rtc.nackIntervalRatio,omitempty"`
	Rtc_NackMaxCount                    *string `json:"rtc.nackMaxCount,omitempty"`
	Rtc_NackMaxMS                       *string `json:"rtc.nackMaxMS,omitempty"`
	Rtc_NackMaxSize                     *string `json:"rtc.nackMaxSize,omitempty"`
	Rtc_NackRtpSize                     *string `json:"rtc.nackRtpSize,omitempty"`
	Rtc_Port                            *string `json:"rtc.port,omitempty"`
	Rtc_PreferredCodecA                 *string `json:"rtc.preferredCodecA,omitempty"`
	Rtc_PreferredCodecV                 *string `json:"rtc.preferredCodecV,omitempty"`
	Rtc_RembBitRate                     *string `json:"rtc.rembBitRate,omitempty"`
	Rtc_StartBitrate                    *string `json:"rtc.start_bitrate,omitempty"`
	Rtc_TCPPort                         *string `json:"rtc.tcpPort,omitempty"`
	Rtc_TimeoutSec                      *string `json:"rtc.timeoutSec,omitempty"`
	Rtmp_DirectProxy                    *string `json:"rtmp.directProxy,omitempty"`
	Rtmp_Enhanced                       *string `json:"rtmp.enhanced,omitempty"`
	Rtmp_HandshakeSecond                *string `json:"rtmp.handshakeSecond,omitempty"`
	Rtmp_KeepAliveSecond                *string `json:"rtmp.keepAliveSecond,omitempty"`
	Rtmp_Port                           *string `json:"rtmp.port,omitempty"`
	Rtmp_Sslport                        *string `json:"rtmp.sslport,omitempty"`
	Rtp_AudioMtuSize                    *string `json:"rtp.audioMtuSize,omitempty"`
	Rtp_H264StapA                       *string `json:"rtp.h264_stap_a,omitempty"`
	Rtp_LowLatency                      *string `json:"rtp.lowLatency,omitempty"`
	Rtp_RtpMaxSize                      *string `json:"rtp.rtpMaxSize,omitempty"`
	Rtp_VideoMtuSize                    *string `json:"rtp.videoMtuSize,omitempty"`
	RtpProxy_DumpDir                    *string `json:"rtp_proxy.dumpDir,omitempty"`
	RtpProxy_GopCache                   *string `json:"rtp_proxy.gop_cache,omitempty"`
	RtpProxy_H264Pt                     *string `json:"rtp_proxy.h264_pt,omitempty"`
	RtpProxy_H265Pt                     *string `json:"rtp_proxy.h265_pt,omitempty"`
	RtpProxy_OpusPt                     *string `json:"rtp_proxy.opus_pt,omitempty"`
	RtpProxy_Port                       *string `json:"rtp_proxy.port,omitempty"`
	RtpProxy_PortRange                  *string `json:"rtp_proxy.port_range,omitempty"`
	RtpProxy_PsPt                       *string `json:"rtp_proxy.ps_pt,omitempty"`
	RtpProxy_RtpG711DurMs               *string `json:"rtp_proxy.rtp_g711_dur_ms,omitempty"`
	RtpProxy_TimeoutSec                 *string `json:"rtp_proxy.timeoutSec,omitempty"`
	RtpProxy_UdpRecvSocketBuffer        *string `json:"rtp_proxy.udp_recv_socket_buffer,omitempty"`
	Rtsp_AuthBasic                      *string `json:"rtsp.authBasic,omitempty"`
	Rtsp_DirectProxy                    *string `json:"rtsp.directProxy,omitempty"`
	Rtsp_HandshakeSecond                *string `json:"rtsp.handshakeSecond,omitempty"`
	Rtsp_KeepAliveSecond                *string `json:"rtsp.keepAliveSecond,omitempty"`
	Rtsp_LowLatency                     *string `json:"rtsp.lowLatency,omitempty"`
	Rtsp_Port                           *string `json:"rtsp.port,omitempty"`
	Rtsp_RtpTransportType               *string `json:"rtsp.rtpTransportType,omitempty"`
	Rtsp_Sslport                        *string `json:"rtsp.sslport,omitempty"`
	Shell_MaxReqSize                    *string `json:"shell.maxReqSize,omitempty"`
	Shell_Port                          *string `json:"shell.port,omitempty"`
	Srt_LatencyMul                      *string `json:"srt.latencyMul,omitempty"`
	Srt_PassPhrase                      *string `json:"srt.passPhrase,omitempty"`
	Srt_PktBufSize                      *string `json:"srt.pktBufSize,omitempty"`
	Srt_Port                            *string `json:"srt.port,omitempty"`
	Srt_TimeoutSec                      *string `json:"srt.timeoutSec,omitempty"`
}
type SetServerConfigReply struct {
	BaseResult
	Changed int `json:"changed"` // 配置项变更个数
}

func (c *Client) SetServerConfig(ctx context.Context, req *SetServerConfigRequest, opts ...CallOption) (*SetServerConfigReply, error) {
	var resp SetServerConfigReply

	err := c.Post(ctx, "/index/api/setServerConfig", req, &resp, opts...)
	if err != nil {
		return nil, err
	}
	if err = resp.inspectError(); err != nil {
		return nil, err
	}
	return &resp, nil
}
