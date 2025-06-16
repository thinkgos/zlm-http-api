# zlm-http-api

## http restful api

- [x] 0. `/index/api/getApiList` 获取API列表
- [x] 1. `/index/api/getThreadsLoad` 获取各 epoll(或 select)线程负载以及延时
- [x] 2. `/index/api/getWorkThreadsLoad` 获取各后台 epoll(或 select)线程负载以及延时
- [x] 3. `/index/api/getServerConfig` 获取服务器配置
- [x] 4. `/index/api/setServerConfig` 设置服务器配置
- [x] 5. `/index/api/restartServer` 重启服务器, 只有Daemon方式才能重启, 否则是直接关闭!
- [x] 6. `/index/api/getMediaList` 获取流列表, 可选筛选参数
- [ ] 7. ~~`/index/api/close_stream`~~ 关闭流(已过期, 请使用close_streams接口替换)
- [x] 8. `/index/api/close_streams` 关闭流(目前所有类型的流都支持关闭), 有人在观看是否强制关闭
- [x] 9. `/index/api/getAllSession` 获取所有TcpSession列表(获取所有tcp客户端相关信息)
- [x] 10. `/index/api/kick_session` 断开tcp连接, 比如说可以断开rtsp、rtmp播放器等
- [x] 11. `/index/api/kick_sessions` 断开tcp连接, 比如说可以断开rtsp、rtmp播放器等
- [x] 12. `/index/api/addStreamProxy` 动态添加拉流代理rtsp/rtmp/hls/http-ts/http-flv(只支持 H264/H265/aac/G711/opus 负载)
- [x] 13. `/index/api/delStreamProxy` 关闭拉流代理(流注册成功后, 也可以使用close_streams接口替代)
- [x] 13.1. `/index/api/listStreamProxy` 获取拉流代理列表
- [x] 13.2. `/index/api/getProxyInfo` 获取拉流代理信息
- [x] 14. `/index/api/addFFmpegSource` 通过 fork FFmpeg 进程的方式拉流代理, 支持任意协议
- [x] 15. `/index/api/delFFmpegSource` 关闭 ffmpeg 拉流代理(流注册成功后, 也可以使用close_streams接口替代)
- [x] 15.1 `/index/api/listFFmpegSource` 获取FFmpeg代理
- [ ] 16. ~~`/index/api/isMediaOnline`~~ 判断直播流是否在线(已过期, 请使用getMediaList接口替代)
- [ ] 17. ~~`/index/api/getMediaInfo`~~ 获取流相关信息(已过期, 请使用getMediaList接口替代)
- [x] 18. `/index/api/getRtpInfo` 获取rtp推流信息
- [x] 19. `/index/api/getMp4RecordFile` 搜索文件系统, 获取流对应的录像文件列表或日期文件夹列表
- [x] 20. `/index/api/startRecord` 开始录制hls或MP4
- [x] 21. `/index/api/stopRecord` 停止录制流
- [x] 22. `/index/api/isRecording` 获取流录制状态
- [x] 22.1. `/index/api/deleteRecordDirectory` 删除录像文件夹
- [x] 23. `/index/api/getSnap` 获取截图或生成实时截图并返回
- [x] 24. `/index/api/openRtpServer` 创建GB28181 RTP 接收端口, 如果该端口接收数据超时, 则会自动被回收(不用调用 closeRtpServer 接口)
- [x] 25. `/index/api/closeRtpServer` 关闭 GB28181 RTP接收端口
- [x] 26. `/index/api/listRtpServer` 获取 openRtpServer 接口创建的所有 RTP 服务器
- [ ] 27. `/index/api/startSendRtp` 作为GB28181客户端, 启动 ps-rtp 推流, 支持rtp/udp方式.
- [ ] 28. `/index/api/stopSendRtp` 作为GB28181客户端, 停止GB28181 ps-rtp 推流
- [x] 29. `/index/api/getStatistic` 获取主要对象个数统计, 主要用于分析内存性能.
- [x] 30. `/index/api/addStreamPusherProxy` 添加rtsp/rtmp主动推流代理(把本服务器的直播流推送到其他服务器去)
- [x] 31. `/index/api/delStreamPusherProxy` 关闭推流代理(可以使用close_streams接口关闭源直播流也可以停止推流)
- [x] 31.1 `/index/api/listStreamPusherProxy` 获取推流代理列表
- [x] 31.2 `/index/api/getProxyPusherInfo` 获取推流代理信息
- [x] 32. `index/api/version` 获取版本信息, 如分支, commit id, 编译时间
- [x] 33. `/index/api/getMediaPlayerList` 获取某个流观看者列表

- [x] 100. `/index/api/downloadFile` 下载文件, 直接下载, 建议加鉴权, 避免任意文件被下载.
  - `file_path`(必选参数): 文件绝对路径, 根据文件名生成Content-Type该接口将触发on_http_access hook.
  - `save_name`(可选参数): 浏览器下载文件后保存文件名.

## WebHook api

Link: [WebHook api](https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html)

- [x] `on_flow_report` 流量统计事件
- [x] `on_http_access` 访问http文件服务器上hls之外的文件时触发
- [x] `on_play` 播放器鉴权事件(如果rtsp启用专有鉴权, 将不会触发该事件)
- [x] `on_publish` 推流鉴权事件
- [x] `on_record_mp4` 录制mp4完成后通知事件; 此事件对回复不敏感.
- [x] `on_record_ts` 录制hls完成后通知事件; 此事件对回复不敏感.
- [x] `on_rtp_server_timeout` 调用openRtpServer接口, rtp server长时间未收到数据, 执行此 web hook, 对回复不敏感.
- [ ] `on_send_rtp_stopped` 发送rtp(startSendRtp)被动关闭时回调
- [x] `on_rtsp_auth` rtsp专用的鉴权事件, 先触发on_rtsp_realm事件然后才会触发on_rtsp_auth事件.
- [x] `on_rtsp_realm` 该rtsp流是否开启rtsp专用方式的鉴权事件, 开启后才会触发on_rtsp_auth事件
- [x] `on_shell_login` shell登录鉴权, telnet调试方式
- [x] `on_stream_changed` rtsp/rtmp 流注册或注销时触发此事件; 此事件对回复不敏感.
- [x] `on_stream_none_reader` 流无人观看时事件
- [x] `on_stream_not_found` 流未找到事件
- [x] `on_server_started` 服务器启动事件
- [x] `on_server_exited` 服务器退出事件
- [x] `on_server_keepalive` 服务器定时上报, 上报间隔可配置, 默认10s上报一次

## 播放url规则

详细查看原文档: [播放url规则](https://docs.zlmediakit.com/zh/guide/media_server/play_url_rules.html)

### 1. url的组成部分

以`rtsp://somedomain.com:554/live/0?token=abcdefg&field=value`为例,该`url`分为以下几个部分:

- 协议(`schema`): rtsp 协议, 默认端口 554
- 虚拟主机(`vhost`): `somedomain.com`, 该字段既可以是域名也可以是ip, 如果是ip则对应的虚拟主机为`__defaultVhost__`
- 服务端口号(`port`): 554, 如果不指定端口号, 则使用协议默认端口号
- 应用名(`app`): live
- 流id(`streamid`): 0
- 参数(`args`): `token=abcdefg&field=value`

### 2. ZLMediaKit 中的流媒体源

在`ZLMediaKit`中, 流媒体源是一种可以被用于直播转发、推流转发等功能的数据对象, 在本项目中被称作为`MediaSource`, 目前支持 5 种类型的流媒体源, 分别是`RtspMediaSource`、`RtmpMediaSource`、`HlsMediaSource`、`TSMediaSource`、`FMP4MediaSource`.

定位一个流媒体源, 主要通过4个元素(我们后续称其为4元组), 分别是:

- 协议(`scheam`)
- 虚拟主机(`vhost`)
- 应用名(`app`)
- 流id(`streamid`)

`RtspMediaSource`支持 rtsp 播放、rtsp 推流、webrtc 播放、webrtc 推流.
`RtmpMediaSource`支持 rtmp 推流/播放、http-flv 播放、ws-flv 播放.
`HlsMediaSource`支持 hls 播放.
`TSMediaSource` 支持 http-ts 播放、ws-ts 播放.
`FMP4MediaSource`支持 http-fmp4 播放、ws-fmp4 播放.

## License

[MIT](LICENSE)
