package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

func (h *ZlmWebhook) OnPublish(ctx context.Context, req *zlm_webhook.OnPublishRequest) (*zlm_webhook.OnPublishReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Schema(req.Schema),
			Vhost(req.Vhost),
			Ip(req.Ip),
			Port(req.Port),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnPublish")
	return &zlm_webhook.OnPublishReply{Code: zlm_webhook.Code_Success, Msg: zlm_webhook.Code_Success_Msg}, nil
}
