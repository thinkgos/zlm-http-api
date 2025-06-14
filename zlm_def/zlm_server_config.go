package zlm_def

//* 服务器配置
// https://docs.zlmediakit.com/zh/guide/media_server/config_file.html
// https://github.com/zlmediakit/ZLMediaKit/blob/master/conf/config.ini

// ServerConfig 服务器配置
type ServerConfig struct {
	MediaServerId string `json:"mediaServerId"`
	// api域

	// 是否调试http api,启用调试后, 会打印每次http请求的内容和回复
	Api_ApiDebug int `json:"api.apiDebug,string"`
	// 一些比较敏感的http api在访问时需要提供secret, 否则无权限调用
	Api_Secret string `json:"api.secret"`
	// 默认截图图片, 在启动FFmpeg截图后但是截图还未生成时, 可以返回默认的预设图片
	Api_DefaultSnap string `json:"api.defaultSnap"`
	// 截图保存路径根目录, 截图通过http api(/index/api/getSnap)生成和获取
	Api_SnapRoot string `json:"api.snapRoot"`
	// downloadFile http接口可访问文件的根目录, 支持多个目录, 不同目录通过分号(;)分隔
	Api_DownloadRoot string `json:"api.downloadRoot"`

	// FFmpeg域

	// FFmpeg可执行程序路径,支持相对路径/绝对路径
	Ffmpeg_Bin string `json:"ffmpeg.bin"`
	// FFmpeg拉流再推流的命令模板, 通过该模板可以设置再编码的一些参数
	Ffmpeg_Cmd string `json:"ffmpeg.cmd"`
	// FFmpeg生成截图的命令, 可以通过修改该配置改变截图分辨率或质量
	Ffmpeg_Snap string `json:"ffmpeg.snap"`
	// FFmpeg日志的路径, 如果置空则不生成FFmpeg日志, 可以为相对(相对于本可执行程序目录)或绝对路径
	Ffmpeg_Log string `json:"ffmpeg.log"`
	// 自动重启的时间(秒), 默认为0, 也就是不自动重启. 主要是为了避免长时间ffmpeg拉流导致的不同步现象
	Ffmpeg_RestartSec int `json:"ffmpeg.restart_sec,string"`

	// protocol域

	// 转协议时, 是否开启帧级时间戳覆盖
	// 0:采用源视频流绝对时间戳, 不做任何改变
	// 1:采用zlmediakit接收数据时的系统时间戳(有平滑处理)
	// 2:采用源视频流时间戳相对时间戳(增长量), 有做时间戳跳跃和回退矫正
	Protocol_ModifyStamp int `json:"protocol.modify_stamp,string"`
	// 转协议是否开启音频
	Protocol_EnableAudio int `json:"protocol.enable_audio,string"`
	// 添加acc静音音频, 在关闭音频时, 此开关无效
	Protocol_AddMuteAudio int `json:"protocol.add_mute_audio,string"`
	// 无人观看时, 是否直接关闭(而不是通过on_none_reader hook返回close)
	// 此配置置1时, 此流如果无人观看, 将不触发on_none_reader hook回调, 而是将直接关闭流
	Protocol_AutoClose int `json:"protocol.auto_close,string"`
	// 推流断开后可以在超时时间内重新连接上继续推流, 这样播放器会接着播放。
	// 置0关闭此特性(推流断开会导致立即断开播放器)
	// 此参数不应大于播放器超时时间;单位毫秒
	Protocol_ContinuePushMs int `json:"protocol.continue_push_ms,string"`

	// 平滑发送定时器间隔, 单位毫秒, 置0则关闭；开启后影响cpu性能同时增加内存
	// 该配置开启后可以解决一些流发送不平滑导致zlmediakit转发也不平滑的问题
	Protocol_PacedSenderMs int `json:"protocol.paced_sender_ms,string"`

	// 是否开启转换为hls(mpegts)
	Protocol_EnableHls int `json:"protocol.enable_hls,string"`
	// 是否开启转换为hls(fmp4)
	Protocol_EnableHlsFmp4 int `json:"protocol.enable_hls_fmp4,string"`
	// 是否开启MP4录制
	Protocol_EnableMp4 int `json:"protocol.enable_mp4,string"`
	// 是否开启转换为rtsp/webrtc
	Protocol_EnableRtsp int `json:"protocol.enable_rtsp,string"`
	// 是否开启转换为rtmp/flv
	Protocol_EnableRtmp int `json:"protocol.enable_rtmp,string"`
	// 是否开启转换为http-ts/ws-ts
	Protocol_EnableTs int `json:"protocol.enable_ts,string"`
	// 是否开启转换为http-fmp4/ws-fmp4
	Protocol_EnableFmp4 int `json:"protocol.enable_fmp4,string"`
	// 是否将mp4录制当做观看者
	Protocol_Mp4AsPlayer int `json:"protocol.mp4_as_player,string"`
	// mp4切片大小, 单位秒
	Protocol_Mp4MaxSecond int `json:"protocol.mp4_max_second,string"`
	// mp4录制保存路径
	Protocol_Mp4SavePath string `json:"protocol.mp4_save_path"`
	// hls录制保存路径
	Protocol_HlsSavePath string `json:"protocol.hls_save_path"`

	// 以下是按需转协议的开关, 在测试ZLMediaKit的接收推流性能时, 请把下面开关置1
	// 如果某种协议你用不到, 你可以把以下开关置1以便节省资源(但是还是可以播放, 只是第一个播放者体验稍微差点),
	// 如果某种协议你想获取最好的用户体验, 请置0(第一个播放者可以秒开, 且不花屏)

	// hls协议是否按需生成, 如果hls.segNum配置为0(意味着hls录制), 那么hls将一直生成(不管此开关)
	Protocol_HlsDemand int `json:"protocol.hls_demand,string"`
	// rtsp[s]协议是否按需生成
	Protocol_RtspDemand int `json:"protocol.rtsp_demand,string"`
	// rtmp[s]、http[s]-flv、ws[s]-flv协议是否按需生成
	Protocol_RtmpDemand int `json:"protocol.rtmp_demand,string"`
	// http[s]-ts协议是否按需生成
	Protocol_TsDemand int `json:"protocol.ts_demand,string"`
	// http[s]-fmp4、ws[s]-fmp4协议是否按需生成
	Protocol_Fmp4Demand int `json:"protocol.fmp4_demand,string"`

	// general域

	// 是否启用虚拟主机
	General_EnableVhost int `json:"general.enableVhost,string"`
	// 播放器或推流器在断开后会触发hook.on_flow_report事件(使用多少流量事件),
	// flowThreshold参数控制触发hook.on_flow_report事件阈值, 使用流量超过该阈值后才触发, 单位KB
	General_FlowThreshold int `json:"general.flowThreshold,string"`
	// 播放最多等待时间, 单位毫秒
	// 播放在播放某个流时, 如果该流不存在,
	// ZLMediaKit会最多让播放器等待maxStreamWaitMS毫秒
	// 如果在这个时间内, 该流注册成功, 那么会立即返回播放器播放成功
	// 否则返回播放器未找到该流, 该机制的目的是可以先播放再推流
	General_MaxStreamWaitMs int `json:"general.maxStreamWaitMS,string"`
	// 某个流无人观看时, 触发hook.on_stream_none_reader事件的最大等待时间, 单位毫秒
	// 在配合hook.on_stream_none_reader事件时, 可以做到无人观看自动停止拉流或停止接收推流
	General_StreamNoneReaderDelayMs int `json:"general.streamNoneReaderDelayMS,string"`
	// 拉流代理时如果断流再重连成功是否删除前一次的媒体流数据, 如果删除将重新开始,
	// 如果不删除将会接着上一次的数据继续写(录制hls/mp4时会继续在前一个文件后面写)
	General_ResetWhenRePlay int `json:"general.resetWhenRePlay,string"`
	// 合并写缓存大小(单位毫秒), 合并写指服务器缓存一定的数据后才会一次性写入socket, 这样能提高性能, 但是会提高延时
	// 开启后会同时关闭TCP_NODELAY并开启MSG_MORE
	General_MergeWriteMs int `json:"general.mergeWriteMS,string"`
	// 服务器唯一id, 用于触发hook时区别是哪台服务器
	General_MediaServerId string `json:"general.mediaServerId"`
	// 最多等待未初始化的Track时间, 单位毫秒, 超时之后会忽略未初始化的Track
	General_WaitTrackReadyMs int `json:"general.wait_track_ready_ms,string"`
	// 最多等待音频Track收到数据时间, 单位毫秒, 超时且完全没收到音频数据, 忽略音频Track
	// 加快某些带封装的流metadata说明有音频, 但是实际上没有的流ready时间（比如很多厂商的GB28181 PS）
	General_WaitAudioTrackDataMs int `json:"general.wait_audio_track_data_ms,string"`
	// 如果流只有单Track, 最多等待若干毫秒, 超时后未收到其他Track的数据, 则认为是单Track
	// 如果协议元数据有声明特定track数, 那么无此等待时间
	General_WaitAddTrackMs int `json:"general.wait_add_track_ms,string"`
	// 如果track未就绪, 我们先缓存帧数据, 但是有最大个数限制, 防止内存溢出
	General_UnreadyFrameCache int `json:"general.unready_frame_cache,string"`
	// 是否启用观看人数变化事件广播, 置1则启用, 置0则关闭
	General_BroadcastPlayerCountChanged int `json:"general.broadcast_player_count_changed,string"`
	// 绑定的本地网卡ip
	General_ListenIP string `json:"general.listen_ip"`
	// FIXME: 目前看来已弃用
	General_CheckNvidiaDev string `json:"general.check_nvidia_dev"`
	// FIXME: 目前看来已弃用
	General_EnableFfmpegLog string `json:"general.enable_ffmpeg_log"`
	// hls域

	// hls写文件的buf大小, 调整参数可以提高文件io性能
	Hls_FileBufSize int `json:"hls.fileBufSize,string"`
	// hls最大切片时间
	Hls_SegDur int `json:"hls.segDur,string"`
	// m3u8索引中,hls保留切片个数(实际保留切片个数+segRetain个)
	// 如果设置为0, 则不删除切片且m3u8文件全量记录切片列表
	Hls_SegNum int `json:"hls.segNum,string"`
	// HLS切片延迟个数, 大于0将生成hls_delay.m3u8文件, 0则不生成
	Hls_SegDelay int `json:"hls.segDelay,string"`
	// HLS切片从m3u8文件中移除后, 继续保留在磁盘上的个数
	Hls_SegRetain int `json:"hls.segRetain,string"`
	// 是否广播 hls切片(ts/fmp4)完成通知(on_record_ts)
	Hls_BroadcastRecordTs int `json:"hls.broadcastRecordTs,string"`
	// 直播hls文件删除延时, 单位秒, issue: #913
	Hls_DeleteDelaySec int `json:"hls.deleteDelaySec,string"`
	// 此选项开启后m3u8文件还是表现为直播, 但是切片文件会被全部保留为点播用
	// segDur设置为0或segKeep设置为1的情况下, 每个切片文件夹下会生成一个vod.m3u8文件用于点播该时间段的录像
	Hls_SegKeep int `json:"hls.segKeep,string"`
	// 如果设置为1, 则第一个切片长度强制设置为1个GOP。当GOP小于segDur, 可以提高首屏速度
	Hls_FastRegister int `json:"hls.fastRegister,string"`

	// hook域

	// 是否启用hook事件, 启用后, 推拉流都将进行鉴权
	Hook_Enable          int    `json:"hook.enable,string"`
	Hook_OnFlowReport    string `json:"hook.on_flow_report"`
	Hook_OnHttpAccess    string `json:"hook.on_http_access"`
	Hook_OnPlay          string `json:"hook.on_play"`
	Hook_OnPublish       string `json:"hook.on_publish"`
	Hook_OnRecordMp4     string `json:"hook.on_record_mp4"`
	Hook_OnRecordTs      string `json:"hook.on_record_ts"`
	Hook_OnRtspAuth      string `json:"hook.on_rtsp_auth"`
	Hook_OnRtspRealm     string `json:"hook.on_rtsp_realm"`
	Hook_OnShellLogin    string `json:"hook.on_shell_login"`
	Hook_OnStreamChanged string `json:"hook.on_stream_changed"`
	// 过滤on_stream_changed hook的协议类型, 可以选择只监听某些感兴趣的协议；置空则不过滤协议
	Hook_StreamChangedSchemas string `json:"hook.stream_changed_schemas"`
	Hook_OnStreamNoneReader   string `json:"hook.on_stream_none_reader"`
	Hook_OnStreamNotFound     string `json:"hook.on_stream_not_found"`
	Hook_OnServerStarted      string `json:"hook.on_server_started"`
	Hook_OnServerExited       string `json:"hook.on_server_exited"`
	Hook_OnServerKeepalive    string `json:"hook.on_server_keepalive"`
	Hook_OnSendRtpStopped     string `json:"hook.on_send_rtp_stopped"`
	Hook_OnRtpServerTimeout   string `json:"hook.on_rtp_server_timeout"`
	// hook api最大等待回复时间, 单位秒
	Hook_TimeoutSec string `json:"hook.timeoutSec"`
	// keepalive hook触发间隔,单位秒, float类型
	Hook_AliveInterval float64 `json:"hook.alive_interval,string"`
	// hook通知失败重试次数,正整数。为0不重试, 1时重试一次, 以此类推
	Hook_Retry int `json:"hook.retry,string"`
	// #hook通知失败重试延时, 单位秒, float型
	Hook_RetryDelay float64 `json:"hook.retry_delay,string"`

	// cluster域

	// 设置源站拉流url模板, 格式跟printf类似, 第一个%s指定app,第二个%s指定stream_id,
	// 开启集群模式后, on_stream_not_found和on_stream_none_reader hook将无效.
	// 溯源模式支持以下类型:
	// rtmp方式: rtmp://127.0.0.1:1935/%s/%s
	// rtsp方式: rtsp://127.0.0.1:554/%s/%s
	// hls方式: http://127.0.0.1:80/%s/%s/hls.m3u8
	// http-ts方式: http://127.0.0.1:80/%s/%s.live.ts
	// 支持多个源站, 不同源站通过分号(;)分隔
	Cluster_OriginUrl string `json:"cluster.origin_url"`
	// 溯源总超时时长, 单位秒, float型；假如源站有3个, 那么单次溯源超时时间为timeout_sec除以3
	// 单次溯源超时时间不要超过general.maxStreamWaitMS配置
	Cluster_TimeoutSec string `json:"cluster.timeout_sec"`
	// 溯源失败尝试次数, -1时永久尝试
	Cluster_RetryCount string `json:"cluster.retry_count"`

	// http域

	// http服务器字符编码集
	Http_CharSet string `json:"http.charSet"`
	// http链接超时时间
	Http_KeepAliveSecond int `json:"http.keepAliveSecond,string"`
	// http请求体最大字节数, 如果post的body太大, 则不适合缓存body在内存
	Http_MaxReqSize int `json:"http.maxReqSize,string"`
	// 404网页内容, 用户可以自定义404网页
	Http_NotFound string `json:"http.notFound"`
	// http服务器监听端口
	Http_Port int `json:"http.port,string"`
	// https服务器监听端口
	Http_Sslport int `json:"http.sslport,string"`
	// http文件服务器根目录
	// 可以为相对(相对于本可执行程序目录)或绝对路径
	Http_RootPath string `json:"http.rootPath"`
	// http文件服务器读文件缓存大小, 单位BYTE, 调整该参数可以优化文件io性能
	Http_SendBufSize int `json:"http.sendBufSize,string"`
	// 是否显示文件夹菜单, 开启后可以浏览文件夹
	Http_DirMenu int `json:"http.dirMenu,string"`
	// 虚拟目录, 虚拟目录名和文件路径使用","隔开, 多个配置路径间用";"隔开
	// 例如赋值为 app_a,/path/to/a;app_b,/path/to/b 那么
	// 访问 http://127.0.0.1/app_a/file_a 对应的文件路径为 /path/to/a/file_a
	// 访问 http://127.0.0.1/app_b/file_b 对应的文件路径为 /path/to/b/file_b
	// 访问其他http路径,对应的文件路径还是在rootPath内
	Http_VirtualPath string `json:"http.virtualPath"`
	// 禁止后缀的文件使用mmap缓存, 使用“,”隔开
	// 例如赋值为 .mp4,.flv
	// 那么访问后缀为.mp4与.flv 的文件不缓存
	Http_ForbidCacheSuffix string `json:"http.forbidCacheSuffix"`
	// 可以把http代理前真实客户端ip放在http头中：https://github.com/ZLMediaKit/ZLMediaKit/issues/1388
	// 切勿暴露此key, 否则可能导致伪造客户端ip
	Http_ForwardedIPHeader string `json:"http.forwarded_ip_header"`
	// 默认允许所有跨域请求
	Http_AllowCrossDomains int `json:"http.allow_cross_domains,string"`
	// 允许访问http api和http文件索引的ip地址范围白名单, 置空情况下不做限制
	Http_AllowIpRange string `json:"http.allow_ip_range"`

	// multicast 域

	// rtp组播截止组播ip地址
	Multicast_AddrMax string `json:"multicast.addrMax"`
	// rtp组播起始组播ip地址
	Multicast_AddrMin string `json:"multicast.addrMin"`
	// 组播udp ttl
	Multicast_UdpTtl int `json:"multicast.udpTTL,string"`

	// record域

	// mp4录制或mp4点播的应用名, 通过限制应用名, 可以防止随意点播
	// 点播的文件必须放置在此文件夹下
	Record_AppName string `json:"record.appName"`
	// mp4录制写文件缓存, 单位BYTE,调整参数可以提高文件io性能
	Record_FileBufSize int `json:"record.fileBufSize,string"`
	// mp4点播每次流化数据量, 单位毫秒,
	// 减少该值可以让点播数据发送量更平滑, 增大该值则更节省cpu资源
	Record_SampleMs int `json:"record.sampleMS,string"`
	// mp4录制完成后是否进行二次关键帧索引写入头部
	Record_FastStart int `json:"record.fastStart,string"`
	// MP4点播(rtsp/rtmp/http-flv/ws-flv)是否循环播放文件
	Record_FileRepeat int `json:"record.fileRepeat,string"`
	// MP4录制写文件格式是否采用fmp4, 启用的话, 断电未完成录制的文件也能正常打开
	Record_EnableFmp4 int `json:"record.enableFmp4,string"`

	// rtmp域

	// rtmp必须在此时间内完成握手, 否则服务器会断开链接, 单位秒
	Rtmp_HandshakeSecond int `json:"rtmp.handshakeSecond,string"`
	// rtmp超时时间, 如果该时间内未收到客户端的数据,
	// 或者tcp发送缓存超过这个时间, 则会断开连接, 单位秒
	Rtmp_KeepAliveSecond int `json:"rtmp.keepAliveSecond,string"`
	// rtmp服务器监听端口
	Rtmp_Port int `json:"rtmp.port,string"`
	// rtmps服务器监听地址
	Rtmp_Sslport int `json:"rtmp.sslport,string"`
	//  rtmp是否直接代理模式
	Rtmp_DirectProxy int `json:"rtmp.directProxy,string"`
	// h265 rtmp打包采用增强型rtmp标准还是国内拓展标准
	Rtmp_Enhanced int `json:"rtmp.enhanced,string"`

	// rtp域

	// 音频mtu大小, 该参数限制rtp最大字节数, 推荐不要超过1400, 加大该值会明显增加直播延时
	Rtp_AudioMtuSize int `json:"rtp.audioMtuSize,string"`
	// 视频mtu大小, 该参数限制rtp最大字节数, 推荐不要超过1400
	Rtp_VideoMtuSize int `json:"rtp.videoMtuSize,string"`
	// tp包最大长度限制, 单位KB,主要用于识别TCP上下文破坏时, 获取到错误的rtp
	Rtp_RtpMaxSize int `json:"rtp.rtpMaxSize,string"`
	// rtp 打包时, 低延迟开关, 默认关闭（为0）, h264存在一帧多个slice（NAL）的情况, 在这种情况下, 如果开启可能会导致画面花屏
	Rtp_LowLatency int `json:"rtp.lowLatency,string"`
	// H264 rtp打包模式是否采用stap-a模式(为了在老版本浏览器上兼容webrtc)还是采用Single NAL unit packet per H.264 模式
	// 有些老的rtsp设备不支持stap-a rtp, 设置此配置为0可提高兼容性
	Rtp_H264StapA int `json:"rtp.h264_stap_a,string"`

	// rtp proxy域

	// 导出调试数据(包括rtp/ps/h264)至该目录,置空则关闭数据导出
	RtpProxy_DumpDir string `json:"rtp_proxy.dumpDir"`
	// udp和tcp代理服务器, 支持rtp(必须是ts或ps类型)代理
	RtpProxy_Port int `json:"rtp_proxy.port,string"`
	// rtp超时时间, 单位秒
	RtpProxy_TimeoutSec int `json:"rtp_proxy.timeoutSec,string"`
	// 随机端口范围, 最少确保36个端口
	// 该范围同时限制rtsp服务器udp端口范围
	RtpProxy_PortRange string `json:"rtp_proxy.port_range"`
	// rtp h264 负载的pt
	RtpProxy_H264Pt string `json:"rtp_proxy.h264_pt"`
	// rtp h265 负载的pt
	RtpProxy_H265Pt string `json:"rtp_proxy.h265_pt"`
	// rtp ps 负载的pt
	RtpProxy_PsPt string `json:"rtp_proxy.ps_pt"`
	// rtp opus 负载的pt
	RtpProxy_OpusPt string `json:"rtp_proxy.opus_pt"`
	// startSendRtp、startRecord相关功能是否提前开启gop缓存优化级联秒开体验, 默认开启, 并缓存1个GOP
	// 如果不调用startSendRtp、startRecord后相关接口, 可以置0节省内存；如果缓存多个gop, 可以加大该参数
	RtpProxy_GopCache int `json:"rtp_proxy.gop_cache,string"`
	// 国标发送g711 rtp 打包时, 每个包的语音时长是多少, 默认是100 ms, 范围为20~180ms(gb28181-2016, c.2.4规定),
	// 最好为20 的倍数, 程序自动向20的倍数取整
	RtpProxy_RtpG711DurMs int `json:"rtp_proxy.rtp_g711_dur_ms,string"`
	// udp接收数据socket buffer大小配置
	// 4*1024*1024=4196304
	RtpProxy_UdpRecvSocketBuffer int `json:"rtp_proxy.udp_recv_socket_buffer,string"`

	// rtc域

	// rtc播放推流、播放超时时间
	Rtc_TimeoutSec int `json:"rtc.timeoutSec,string"`
	// 本机对rtc客户端的可见ip, 作为服务器时一般为公网ip, 可有多个, 用','分开, 当置空时, 会自动获取网卡ip
	// 同时支持环境变量, 以$开头, 如"$EXTERN_IP"; 请参考：https://github.com/ZLMediaKit/ZLMediaKit/pull/1786
	Rtc_ExternIP string `json:"rtc.externIP"`
	// rtc udp服务器监听端口号, 所有rtc客户端将通过该端口传输stun/dtls/srtp/srtcp数据,
	// 该端口是多线程的, 同时支持客户端网络切换导致的连接迁移
	// 需要注意的是, 如果服务器在nat内, 需要做端口映射时, 必须确保外网映射端口跟该端口一致
	Rtc_Port int `json:"rtc.port,string"`
	// rtc tcp服务器监听端口号, 在udp 不通的情况下, 会使用tcp传输数据
	// 该端口是多线程的, 同时支持客户端网络切换导致的连接迁移
	// 需要注意的是, 如果服务器在nat内, 需要做端口映射时, 必须确保外网映射端口跟该端口一致
	Rtc_TcpPort int `json:"rtc.tcpPort,string"`
	// 设置remb比特率, 非0时关闭twcc并开启remb。该设置在rtc推流时有效, 可以控制推流画质
	// 目前已经实现twcc自动调整码率, 关闭remb根据真实网络状况调整码率
	Rtc_RembBitRate int `json:"rtc.rembBitRate,string"`
	// rtc支持的音频codec类型,在前面的优先级更高
	// 以下范例为所有支持的音频codec
	Rtc_PreferredCodecA string `json:"rtc.preferredCodecA"`
	// rtc支持的视频codec类型,在前面的优先级更高
	// 以下范例为所有支持的视频codec
	Rtc_PreferredCodecV string `json:"rtc.preferredCodecV"`

	// webrtc比特率设置
	Rtc_StartBitrate int `json:"rtc.start_bitrate,string"`
	Rtc_MaxBitrate   int `json:"rtc.max_bitrate,string"`
	Rtc_MinBitrate   int `json:"rtc.min_bitrate,string"`
	// nack接收端, rtp发送端, zlm发送rtc流
	// rtp重发缓存列队最大长度, 单位毫秒
	Rtc_MaxRtpCacheMS int `json:"rtc.maxRtpCacheMS,string"`
	// rtp重发缓存列队最大长度, 单位个数
	Rtc_MaxRtpCacheSize int `json:"rtc.maxRtpCacheSize,string"`

	// nack发送端, rtp接收端, zlm接收rtc推流
	// 最大保留的rtp丢包状态个数
	Rtc_NackMaxSize int `json:"rtc.nackMaxSize,string"`
	// rtp丢包状态最长保留时间
	Rtc_NackMaxMs int `json:"rtc.nackMaxMS,string"`
	// nack最多请求重传次数
	Rtc_NackMaxCount int `json:"rtc.nackMaxCount,string"`
	// nack重传频率, rtt的倍数
	Rtc_NackIntervalRatio float64 `json:"rtc.nackIntervalRatio,string"`
	// nack包中rtp个数, 减小此值可以让nack包响应更灵敏
	Rtc_NackRtpSize int `json:"rtc.nackRtpSize,string"`
	// 是否尝试过滤 b帧
	Rtc_Bfilter int `json:"rtc.bfilter,string"`
	// FIXME: 目前看来已弃用
	Rtc_DatachannelEcho string `json:"rtc.datachannel_echo"`

	// srt域

	// srt播放推流、播放超时时间,单位秒
	Srt_TimeoutSec int `json:"srt.timeoutSec,string"`
	// srt udp服务器监听端口号, 所有srt客户端将通过该端口传输srt数据,
	// 该端口是多线程的, 同时支持客户端网络切换导致的连接迁移
	Srt_Port int `json:"srt.port,string"`
	// srt 协议中延迟缓存的估算参数, 在握手阶段估算rtt, 然后latencyMul*rtt 为最大缓存时长, 此参数越大, 表示等待重传的时长就越大
	Srt_LatencyMul int `json:"srt.latencyMul,string"`
	// 包缓存的大小
	Srt_PktBufSize int `json:"srt.pktBufSize,string"`
	// srt udp服务器的密码,为空表示不加密
	Srt_PassPhrase string `json:"srt.passPhrase"`

	// rtsp域

	// rtsp专有鉴权方式是采用base64还是md5方式
	Rtsp_AuthBasic int `json:"rtsp.authBasic,string"`
	// rtsp拉流、推流代理是否是直接代理模式
	// 直接代理后支持任意编码格式, 但是会导致GOP缓存无法定位到I帧, 可能会导致开播花屏
	// 并且如果是tcp方式拉流, 如果rtp大于mtu会导致无法使用udp方式代理
	// 假定您的拉流源地址不是264或265或AAC, 那么你可以使用直接代理的方式来支持rtsp代理
	// 如果你是rtsp推拉流, 但是webrtc播放, 也建议关闭直接代理模式,
	// 因为直接代理时, rtp中可能没有sps pps,会导致webrtc无法播放; 另外webrtc也不支持Single NAL Unit Packets类型rtp
	// 默认开启rtsp直接代理, rtmp由于没有这些问题, 是强制开启直接代理的
	Rtsp_DirectProxy int `json:"rtsp.directProxy,string"`
	// rtsp必须在此时间内完成握手, 否则服务器会断开链接, 单位秒
	Rtsp_HandshakeSecond int `json:"rtsp.handshakeSecond,string"`
	// rtsp超时时间, 如果该时间内未收到客户端的数据,
	// 或者tcp发送缓存超过这个时间, 则会断开连接, 单位秒
	Rtsp_KeepAliveSecond int `json:"rtsp.keepAliveSecond,string"`
	// rtsp服务器监听地址
	Rtsp_Port int `json:"rtsp.port,string"`
	// rtsps服务器监听地址
	Rtsp_Sslport int `json:"rtsp.sslport,string"`
	// rtsp 转发是否使用低延迟模式, 当开启时, 不会缓存rtp包, 来提高并发, 可以降低一帧的延迟
	Rtsp_LowLatency int `json:"rtsp.lowLatency,string"`
	// 强制协商rtp传输方式(0:TCP,1:UDP,2:MULTICAST,-1:不限制)
	// 当客户端发起RTSP SETUP的时候如果传输类型和此配置不一致则返回461 Unsupported transport
	// 迫使客户端重新SETUP并切换到对应协议。目前支持FFMPEG和VLC
	Rtsp_RtpTransportType int `json:"rtsp.rtpTransportType,string"`

	// shell域

	// 调试telnet服务器接受最大bufffer大小
	Shell_MaxReqSize int `json:"shell.maxReqSize,string"`
	// 调试telnet服务器监听端口
	Shell_Port int `json:"shell.port,string"`
}
