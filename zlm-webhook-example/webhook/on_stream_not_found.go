package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnStreamNotFound implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnStreamNotFound(ctx context.Context, req *zlm_webhook.OnStreamNotFoundRequest) (*zlm_webhook.OnStreamNotFoundReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Schema(req.Schema),
			Vhost(req.Vhost),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnStreamNotFound")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
