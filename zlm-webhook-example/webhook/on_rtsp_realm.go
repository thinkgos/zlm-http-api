package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnRtspRealm implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnRtspRealm(ctx context.Context, req *zlm_webhook.OnRtspRealmRequest) (*zlm_webhook.OnRtspRealmReply, error) {
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
		Msg("OnRtspRealm")
	return &zlm_webhook.OnRtspRealmReply{
		Code:  0,
		Realm: testRtspRealm,
	}, nil
}
