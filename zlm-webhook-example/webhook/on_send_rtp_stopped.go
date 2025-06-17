package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnSendRtpStopped implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnSendRtpStopped(ctx context.Context, req *zlm_webhook.OnSendRtpStoppedRequest) (*zlm_webhook.OnSendRtpStoppedReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Vhost(req.Vhost),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnSendRtpStopped")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
