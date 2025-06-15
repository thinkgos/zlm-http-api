package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

func (w *ZlmWebhook) OnRtpServerTimeout(ctx context.Context, req *zlm_webhook.OnRtpServerTimeoutRequest) (*zlm_webhook.OnRtpServerTimeoutReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			StreamId(req.StreamId),
		).
		Any("payload", req).
		Msg("OnRtpServerTimeout")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
