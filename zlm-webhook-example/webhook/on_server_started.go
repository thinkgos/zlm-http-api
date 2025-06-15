package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnServerStarted implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnServerStarted(ctx context.Context, req *zlm_webhook.OnServerStartedRequest) (*zlm_webhook.OnServerStartedReply, error) {
	logger.OnDebugContext(ctx).
		Any("payload", req).
		Msg("OnServerStarted")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
