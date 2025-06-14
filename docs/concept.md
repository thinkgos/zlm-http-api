# concept

- `RTSP`(Real Time Streaming Protocol): 实时流协议, [RFC2326](https://datatracker.ietf.org/doc/html/rfc2326)
- `RTP` (A Transport Protocol for Real-Time Applications): 实时传输协议 [RFC3550](https://datatracker.ietf.org/doc/html/rfc3550)
- `RTCP`(RTP Control Protocol): RTP控制协议, [RFC3550第6章](https://datatracker.ietf.org/doc/html/rfc3550#section-6)
- `SDP`(Session Description Protocol): 会话描述协议, [RFC2327](https://datatracker.ietf.org/doc/html/rfc4566)
- `RTMP`(Real Time Messaging Protocol): 实时消息协议 [Adobe RTMP Specification](https://rtmp.veriskope.com/docs/spec/)
- `H.264`(Advanced video coding for generic audiovisual services): 通用视听业务的先进的视频编码, [H.264](https://www.itu.int/itu-t/recommendations/rec.aspx?rec=15935&lang=zh)

`RTSP`负责控制命令的交互, 发起和终止音视频数据的传输; `RTP`负责传输音视频数据; `RTCP`负责RTP的服务质量反馈.

## vlc/ffplay 通过RTSP连接IPC流程

1. `vlc/ffplay` ----> `IPC`: 向`IPC`发送RTSP OPTIONS请求, 查询支持的RTSP方法
2. `vlc/ffplay` <---- `IPC`: `IPC`回复200 Reply
3. `vlc/ffplay` ----> `IPC`: (有鉴权)向`IPC`发送RTSP DESCRIBE请求, 请求获取媒体描述(SDP格式)
4. `vlc/ffplay` <---- `IPC`: (有鉴权)`IPC`回复401 Reply
5. `vlc/ffplay` ----> `IPC`: 向`IPC`发送RTSP DESCRIBE请求, 请求获取媒体描述(SDP格式)
6. `vlc/ffplay` <---- `IPC`: `IPC`回复200 Reply且包含SDP信息
7. `vlc/ffplay` ----> `IPC`: 向`IPC`发送RTSP SETUP视频请求(为每个媒体流)
8. `vlc/ffplay` <---- `IPC`: `IPC`回复200 Reply
9. `vlc/ffplay` ----> `IPC`: 向`IPC`发送RTSP SETUP音频请求(为每个媒体流)
10. `vlc/ffplay` <---- `IPC`: `IPC`回复200 Reply
11. `vlc/ffplay` ----> `IPC`: 向`IPC`发送RTSP PLAY请求, 指定播放范围等信息
12. `vlc/ffplay` <---- `IPC`: `IPC`回复200 Reply
13. `vlc/ffplay` <---- `IPC`: `IPC`发送RTP媒体数据包
14. `vlc/ffplay` <---- `IPC`: `IPC`发送RTP媒体数据包
15. `vlc/ffplay` ----> `IPC`:  向`IPC`发送RTCP传输控制信息
16. `vlc/ffplay` ----> `IPC`:  向`IPC`发送RTCP传输控制信息
17. `vlc/ffplay` ----> `IPC`:  向`IPC`发送RTSP PAUSE(暂停流)/TEARDOWN(终端会话)
18. `vlc/ffplay` <---- `IPC`: `IPC`回复200 Reply
19. 释放所有资源, 关闭网络连接

## 参考

- [直播协议详解 RTMP、HLS、HTTP-FLV、WebRTC、RTSP 的区别](https://www.cnblogs.com/eddyz/p/17869403.html)
