package zlm_http

const (
	testZLMediaKitIp = "10.110.18.250"
	// testZLMediaKitIp = "127.0.0.1"
	testBaseUrl = "http://" + testZLMediaKitIp
	testSecret  = "ynaHOQRQaXCBCZn50hreQ8xHzqdDr8eh"
	testSchema  = "rtsp"
	testStream  = "chibo01"
	testStream2 = "chibo02"
	testApp     = "live"
)

var client = func() *Client {
	c := NewClient()
	c.Deref().
		SetBaseURL(testBaseUrl).
		SetQueryParam("secret", testSecret).
		SetDebug(true).
		SetGenerateCurlCmd(true).
		SetDebugLogCurlCmd(true)
	return c
}()
