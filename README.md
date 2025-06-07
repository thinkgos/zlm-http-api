# zlm-http-api

## http restful api

- [x] 0. `/index/api/getApiList` 获取API列表
- [x] 1. `/index/api/getThreadsLoad` 获取各 epoll(或 select)线程负载以及延时
- [x] 2. `/index/api/getWorkThreadsLoad` 获取各后台 epoll(或 select)线程负载以及延时
- [ ] 3. `/index/api/getServerConfig` 获取服务器配置
- [ ] 4. `/index/api/setServerConfig` 设置服务器配置
- [ ] 5. `/index/api/restartServer` 重启服务器, 只有Daemon方式才能重启, 否则是直接关闭!
- [x] 6. `/index/api/getMediaList` 获取流列表, 可选筛选参数
- [ ] 7. `/index/api/close_stream` 关闭流(已过期, 请使用close_streams接口替换)
- [ ] 8. `/index/api/close_streams` 关闭流(目前所有类型的流都支持关闭)
- [x] 23. `/index/api/getSnap` 获取截图或生成实时截图并返回
- [ ] 24. `/index/api/openRtpServer` 创建GB28181 RTP 接收端口, 如果该端口接收数据超时, 则会自动被回收(不用调用 closeRtpServer 接口)
- [ ] 25. `/index/api/closeRtpServer` 关闭 GB28181 RTP接收端口
- [ ] 26. `/index/api/listRtpServer` 获取 openRtpServer 接口创建的所有 RTP 服务器
- [x] 32. `index/api/version` 获取版本信息, 如分支, commit id, 编译时间

## WebHook api

Link: [WebHook api](https://docs.zlmediakit.com/zh/guide/media_server/web_hook_api.html)

- [x] `on_flow_report`
- [x] `on_http_access`
- [x] `on_play`
- [x] `on_publish`
- [x] `on_record_mp4`
- [x] `on_rtsp_auth`
- [x] `on_rtsp_realm`
- [x] `on_shell_login`
- [x] `on_stream_changed`
- [x] `on_stream_none_reader`
- [x] `on_stream_not_found`
- [x] `on_server_started`
- [x] `on_server_keepalive`
- [x] `on_rtp_server_timeout`

## License
