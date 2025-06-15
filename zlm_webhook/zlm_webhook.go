package zlm_webhook

import "context"

// doc: https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html

const (
	// url path prefix
	UrlPathPrefix = "/index/hook"
	// url path
	UrlPath_OnFlowReport       = "/on_flow_report"
	UrlPath_OnHttpAccess       = "/on_http_access"
	UrlPath_OnPlay             = "/on_play"
	UrlPath_OnPublish          = "/on_publish"
	UrlPath_OnRecordMp4        = "/on_record_mp4"
	UrlPath_OnRecordTs         = "/on_record_ts"
	UrlPath_OnRtpServerTimeout = "/on_rtp_server_timeout"
	UrlPath_OnSendRtpStopped   = "/on_send_rtp_stopped"
	UrlPath_OnRtspAuth         = "/on_rtsp_auth"
	UrlPath_OnRtspRealm        = "/on_rtsp_realm"
	UrlPath_OnShellLogin       = "/on_shell_login"
	UrlPath_OnStreamChanged    = "/on_stream_changed"
	UrlPath_OnStreamNoneReader = "/on_stream_none_reader"
	UrlPath_OnStreamNotFound   = "/on_stream_not_found"
	UrlPath_OnServerStarted    = "/on_server_started"
	UrlPath_OnServerExited     = "/on_server_exited"
	UrlPath_OnServerKeepalive  = "/on_server_keepalive"
)

type Webhook interface {
	// 流量统计事件
	OnFlowReport(ctx context.Context, req *OnFlowReportRequest) (*OnFlowReportReply, error)
	// 访问http文件服务器上hls之外的文件时触发
	OnHttpAccess(ctx context.Context, req *OnHttpAccessRequest) (*OnHttpAccessReply, error)
	// 播放器鉴权事件
	OnPlay(ctx context.Context, req *OnPlayRequest) (*OnPlayReply, error)
	// rtsp/rtmp/rtp 推流鉴权事件
	OnPublish(ctx context.Context, req *OnPublishRequest) (*OnPublishReply, error)
	// 录制mp4完成后通知事件
	OnRecordMp4(ctx context.Context, req *OnRecordMp4Request) (*OnRecordMp4Reply, error)
	// 录制ts完成后通知事件
	OnRecordTs(ctx context.Context, req *OnRecordTsRequest) (*OnRecordTsReply, error)
	// 调用openRtpServer接口, rtp server长时间未收到数据, 执行此 web hook, 对回复不敏感.
	OnRtpServerTimeout(ctx context.Context, req *OnRtpServerTimeoutRequest) (*OnRtpServerTimeoutReply, error)
	// 发送rtp(startSendRtp)被动关闭时回调
	OnSendRtpStopped(ctx context.Context, req *OnSendRtpStoppedRequest) (*OnSendRtpStoppedReply, error)
	// rtsp专用的鉴权事件
	// 先触发on_rtsp_realm事件然后才会触发on_rtsp_auth事件。
	OnRtspAuth(ctx context.Context, req *OnRtspAuthRequest) (*OnRtspAuthReply, error)
	// 该rtsp流是否开启rtsp专用方式的鉴权事件
	// 开启后才会触发on_rtsp_auth事件
	OnRtspRealm(ctx context.Context, req *OnRtspRealmRequest) (*OnRtspRealmReply, error)
	// 服务器定时上报时间, 上报间隔可配置, 默认10s上报一次
	OnServerKeepalive(ctx context.Context, req *OnServerKeepaliveRequest) (*OnServerKeepaliveReply, error)
	// 服务器启动事件
	OnServerStarted(ctx context.Context, req *OnServerStartedRequest) (*OnServerStartedReply, error)
	// 服务器退出事件
	OnServerExited(ctx context.Context, req *OnServerExitedRequest) (*OnServerExitedReply, error)
	// shell登录鉴权, telnet调试方式
	OnShellLogin(ctx context.Context, req *OnShellLoginRequest) (*OnShellLoginReply, error)
	// rtsp/rtmp 流注册或注销时触发此事件; 此事件对回复不敏感.
	OnStreamChanged(ctx context.Context, req *OnStreamChangedRequest) (*OnStreamChangedReply, error)
	// 流无人观看时事件
	OnStreamNoneReader(ctx context.Context, req *OnStreamNoneReaderRequest) (*OnStreamNoneReaderReply, error)
	// 流未找到事件
	OnStreamNotFound(ctx context.Context, req *OnStreamNotFoundRequest) (*OnStreamNotFoundReply, error)
}
