package webhook

import (
	"github.com/thinkgos/zlm-http-api/zlm_webhook"
)

var _ zlm_webhook.Webhook = (*ZlmWebhook)(nil)

type ZlmWebhook struct{}

func NewZlmHook() zlm_webhook.Webhook {
	return &ZlmWebhook{}
}
