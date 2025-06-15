package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnHttpAccess implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnHttpAccess(ctx context.Context, req *zlm_webhook.OnHttpAccessRequest) (*zlm_webhook.OnHttpAccessReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			Ip(req.Ip),
			Port(req.Port),
		).
		Any("payload", req).
		Msg("OnHttpAccess")
	return &zlm_webhook.OnHttpAccessReply{
		Code:          0,
		Err:           "",
		MediaServerId: req.MediaServerId,
		Path:          req.Path,
		Second:        3600,
	}, nil
}
