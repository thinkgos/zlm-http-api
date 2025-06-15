package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

func (w *ZlmWebhook) OnServerKeepalive(ctx context.Context, req *zlm_webhook.OnServerKeepaliveRequest) (*zlm_webhook.OnServerKeepaliveReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
		).
		Any("payload", req).
		Msg("OnServerKeepalive")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
