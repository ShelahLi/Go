package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	// 时间段
	TimeOut time.Duration
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever: {UserAgent=%s, TimeOut=%d}", r.UserAgent, r.TimeOut)
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}

	return string(result)
}
