package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnRecordTs implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnRecordTs(ctx context.Context, req *zlm_webhook.OnRecordTsRequest) (*zlm_webhook.OnRecordTsReply, error) {
	logger.OnDebugContext(ctx).
		// Dict("media",
		// 	MediaServerId(req.MediaServerId),
		// 	App(req.App),
		// 	Vhost(req.Vhost),
		// 	StreamId(req.Stream),
		// ).
		Any("payload", req).
		Msg("OnRecordTs")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
