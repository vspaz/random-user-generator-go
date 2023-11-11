package api

import (
	"github.com/vspaz/goat/pkg/ghttp"
	"log"
	"random-user-generator/internal/configuration"
	"strconv"
)

type ApiClient struct {
	HttpClient *ghttp.GoatClient
	Config     *configuration.Config
}

func NewClient(client *ghttp.GoatClient) *ApiClient {
	config := configuration.GetConfig()
	apiClient := &ApiClient{Config: config}
	if client == nil {
		apiClient.HttpClient = ghttp.NewClientBuilder().
			WithHost(config.Host).
			WithUserAgent(config.UserAgent).
			WithRetry(config.Http.Retries.Count, config.Http.Retries.OnErrors).
			WithDelay(config.Http.Retries.Delay).
			WithConnectionTimeout(config.Http.Timeouts.Connection).
			WithResponseTimeout(config.Http.Timeouts.Response).
			WithLogLevel(config.LogLevel).
			Build()
	} else {
		apiClient.HttpClient = client
	}
	return apiClient
}

func (a *ApiClient) FetchRandomUserInfo(seed string, results int, namesOnly bool) *ghttp.Response {
	queryParams := map[string]string{"seed": seed, "results": strconv.Itoa(results)}
	if namesOnly {
		queryParams["inc"] = "name"
	}
	resp, err := a.HttpClient.DoGet(a.Config.Endpoint, nil, queryParams)
	if err != nil {
		log.Fatalf("failed to fetch user info")
	}
	return resp
}
