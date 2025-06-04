package zlm_webhook

//* doc: https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html

const (
	// url path prefix
	UrlPathPrefix = "/index/hook"
	// url path
	UrlPath_OnFlowReport       = "/on_flow_report"
	UrlPath_OnHttpAccess       = "/on_http_access"
	UrlPath_OnPlay             = "/on_play"
	UrlPath_OnPublish          = "/on_publish"
	UrlPath_OnRecordMp4        = "/on_record_mp4"
	UrlPath_OnRtspAuth         = "/on_rtsp_auth"
	UrlPath_OnRtspRealm        = "/on_rtsp_realm"
	UrlPath_OnShellLogin       = "/on_shell_login"
	UrlPath_OnStreamChanged    = "/on_stream_changed"
	UrlPath_OnStreamNoneReader = "/on_stream_none_reader"
	UrlPath_OnStreamNotFound   = "/on_stream_not_found"
	UrlPath_OnServerStarted    = "/on_server_started"
	UrlPath_OnServerKeepalive  = "/on_server_keepalive"
	UrlPath_OnRtpServerTimeout = "/on_rtp_server_timeout"
)
