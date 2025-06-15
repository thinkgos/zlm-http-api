package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

func (w *ZlmWebhook) OnStreamChanged(ctx context.Context, req *zlm_webhook.OnStreamChangedRequest) (*zlm_webhook.OnStreamChangedReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Schema(req.Schema),
			Vhost(req.Vhost),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnStreamChanged")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
