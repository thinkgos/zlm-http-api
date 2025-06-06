package zlm_http

const (
	// baseUrl = "http://127.0.0.1"
	baseUrl = "http://10.110.18.131"
	secret  = "ynaHOQRQaXCBCZn50hreQ8xHzqdDr8eh"
)

var client = func() *ZlmClient {
	c := NewClient(WithCallOption(WithCoNoAuth()))
	c.Deref().SetBaseURL(baseUrl).
		SetQueryParam("secret", secret)
	return c
}()
