package api

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
)

func FetchRandomUserInfo() *ghttp.Response {
	client := ghttp.NewClientBuilder().
		WithHost("https://randomuser.me/api").
		WithUserAgent("randomuser").
		WithRetry(3, []int{500, 503}).
		WithDelay(0.5).
		WithResponseTimeout(10.0).
		WithConnectionTimeout(2.0).
		WithHeadersReadTimeout(2.0).
		WithLogLevel("info").
		Build()

	resp, err := client.DoGet("/get", nil, nil)
	if err != nil {
		log.Fatalf("failed to fetch user info")
	}
	return resp
}
