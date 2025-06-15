package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

func (w *ZlmWebhook) OnPlay(ctx context.Context, req *zlm_webhook.OnPlayRequest) (*zlm_webhook.OnPlayReply, error) {
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
		Msg("OnPlay")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
