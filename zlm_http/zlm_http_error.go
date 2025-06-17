package zlm_http

import (
	"fmt"
	"net/http"
)

type ReplyError struct {
	Code   int
	Body   []byte
	Header http.Header
}

func (e *ReplyError) Error() string {
	return fmt.Sprintf("Invoke: Status Code: %d, Status Text: %s", e.Code, http.StatusText(e.Code))
}
