package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnRecordMp4 implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnRecordMp4(ctx context.Context, req *zlm_webhook.OnRecordMp4Request) (*zlm_webhook.OnRecordMp4Reply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Vhost(req.Vhost),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnRecordMp4")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
