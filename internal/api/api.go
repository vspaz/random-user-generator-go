package api

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
	"random-user-generator/internal/configuration"
)

type ApiClient struct {
	HttpClient *ghttp.GoatClient
}

func NewClient(client *ghttp.GoatClient) *ApiClient {
	if client == nil {
		return &ApiClient{
			HttpClient: ghttp.NewClientBuilder().
				WithHost(configuration.ApiConfig.Host).
				WithUserAgent(configuration.ApiConfig.UserAgent).
				WithRetry(configuration.ApiConfig.Http.Retries.Count, configuration.ApiConfig.Http.Retries.OnErrors).
				WithDelay(configuration.ApiConfig.Http.Retries.Delay).
				WithConnectionTimeout(configuration.ApiConfig.Http.Timeouts.Connection).
				WithResponseTimeout(configuration.ApiConfig.Http.Timeouts.Response).
				WithLogLevel(configuration.ApiConfig.LogLevel).
				Build(),
		}
	}
	return &ApiClient{
		HttpClient: client,
	}
}

func (a *ApiClient) FetchRandomUserInfo(seed string, results uint, namesOnly bool) *ghttp.Response {
	resp, err := a.HttpClient.DoGet(configuration.ApiConfig.Endpoint, nil, nil)
	if err != nil {
		log.Fatalf("failed to fetch user info")
	}
	return resp
}
