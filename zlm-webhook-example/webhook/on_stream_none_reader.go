package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

func (w *ZlmWebhook) OnStreamNoneReader(ctx context.Context, req *zlm_webhook.OnStreamNoneReaderRequest) (*zlm_webhook.OnStreamNoneReaderReply, error) {
	// rtmp 无人观看时，也允许推流
	// 存在录像计划时，不关闭流
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Schema(req.Schema),
			Vhost(req.Vhost),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnStreamNoneReader")
	return &zlm_webhook.OnStreamNoneReaderReply{
		Code:  0,
		Close: false,
	}, nil
}
