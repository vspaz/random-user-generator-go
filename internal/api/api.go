package api

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
)

type ApiClient struct {
	HttpClient *ghttp.GoatClient
}

func NewClient(client *ghttp.GoatClient) *ApiClient {
	if client == nil {
		return &ApiClient{
			HttpClient: ghttp.NewClientBuilder().
				WithHost("https://randomuser.me").
				WithUserAgent("randomuser").
				WithRetry(3, []int{500, 503}).
				WithDelay(0.5).
				WithResponseTimeout(10.0).
				WithConnectionTimeout(2.0).
				WithLogLevel("info").
				Build(),
		}
	}
	return &ApiClient{
		HttpClient: client,
	}
}

func (a *ApiClient) FetchRandomUserInfo(seed string, results uint, namesOnly bool) *ghttp.Response {
	resp, err := a.HttpClient.DoGet("/api", nil, nil)
	if err != nil {
		log.Fatalf("failed to fetch user info")
	}
	return resp
}
