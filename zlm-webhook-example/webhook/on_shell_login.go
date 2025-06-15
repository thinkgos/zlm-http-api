package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnShellLogin implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnShellLogin(ctx context.Context, req *zlm_webhook.OnShellLoginRequest) (*zlm_webhook.OnShellLoginReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			Ip(req.Ip),
			Port(req.Port),
		).
		Any("payload", req).
		Msg("OnShellLogin")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
