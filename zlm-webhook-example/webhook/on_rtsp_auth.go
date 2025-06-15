package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnRtspAuth implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnRtspAuth(ctx context.Context, req *zlm_webhook.OnRtspAuthRequest) (*zlm_webhook.OnRtspAuthReply, error) {
	logger.OnDebugContext(ctx).
		Dict("media",
			MediaServerId(req.MediaServerId),
			App(req.App),
			Schema(req.Schema),
			Vhost(req.Vhost),
			Ip(req.Ip),
			Port(req.Port),
			StreamId(req.Stream),
		).
		Any("payload", req).
		Msg("OnRtspAuth")
	return &zlm_webhook.OnRtspAuthReply{
		Code:      0,
		Msg:       "",
		Encrypted: false,
		Passwd:    testRtspPassword,
	}, nil
}
