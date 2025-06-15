package zlm_http

const (
	// testBaseUrl = "http://127.0.0.1"
	testBaseUrl = "http://10.110.18.131"
	testSecret  = "ynaHOQRQaXCBCZn50hreQ8xHzqdDr8eh"
	testSchema  = "rtsp"
	testStream  = "chibo01"
	testApp     = "live"
)

var client = func() *Client {
	c := NewClient()
	c.Deref().
		SetBaseURL(testBaseUrl).
		SetQueryParam("secret", testSecret)
	return c
}()
