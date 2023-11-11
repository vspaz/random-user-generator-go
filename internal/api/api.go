package api

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
	"random-user-generator/internal/configuration"
)

type ApiClient struct {
	HttpClient *ghttp.GoatClient
	Config     *configuration.Config
}

func NewClient(client *ghttp.GoatClient) *ApiClient {
	config := configuration.GetConfig()
	if client == nil {
		return &ApiClient{
			HttpClient: ghttp.NewClientBuilder().
				WithHost(config.Host).
				WithUserAgent(config.UserAgent).
				WithRetry(config.Http.Retries.Count, config.Http.Retries.OnErrors).
				WithDelay(config.Http.Retries.Delay).
				WithConnectionTimeout(config.Http.Timeouts.Connection).
				WithResponseTimeout(config.Http.Timeouts.Response).
				WithLogLevel(config.LogLevel).
				Build(),
			Config: config,
		}
	}
	return &ApiClient{
		HttpClient: client,
		Config:     config,
	}
}

func (a *ApiClient) FetchRandomUserInfo(seed string, results uint, namesOnly bool) *ghttp.Response {
	resp, err := a.HttpClient.DoGet(a.Config.Endpoint, nil, nil)
	if err != nil {
		log.Fatalf("failed to fetch user info")
	}
	return resp
}
