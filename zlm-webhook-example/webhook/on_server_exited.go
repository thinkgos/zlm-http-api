package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnServerExited implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnServerExited(ctx context.Context, req *zlm_webhook.OnServerExitedRequest) (*zlm_webhook.OnServerExitedReply, error) {
	logger.OnDebugContext(ctx).
		Any("payload", req).
		Msg("OnServerExited")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
