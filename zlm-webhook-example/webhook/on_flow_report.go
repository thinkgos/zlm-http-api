package webhook

import (
	"context"

	"github.com/thinkgos/logger"
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

// OnFlowReport implements zlm_webhook.Webhook.
func (w *ZlmWebhook) OnFlowReport(ctx context.Context, req *zlm_webhook.OnFlowReportRequest) (*zlm_webhook.OnFlowReportReply, error) {
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
		Msg("OnFlowReport")
	return zlm_webhook.NewDefaultSuccessReply(), nil
}
