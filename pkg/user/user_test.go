package user

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vspaz/goat/pkg/ghttp"
	"net/http"
	"net/http/httptest"
	"random-user-generator/internal/api"
	"random-user-generator/internal/configuration"
	"testing"
)

const (
	seed      = "foobar"
	userCount = 1
)

var endpoint = fmt.Sprintf("/api?inc=name&results=%d&seed=%s", userCount, seed)

func startTestServer(t *testing.T) *httptest.Server {
	payload := Users{
		Results: []User{
			{
				Name: Name{
					Title: "Mr",
					First: "John",
					Last:  "Doe",
				}},
		},
	}
	return httptest.NewServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				assert.Equal(t, endpoint, request.URL.String())
				writer.WriteHeader(http.StatusOK)
				encodedBody, _ := json.Marshal(payload)
				if _, err := writer.Write(encodedBody); err != nil {
					t.Fatal(err)
				}
			},
		),
	)
}

func TestRandomUserData_Generate(t *testing.T) {
	server := startTestServer(t)
	defer server.Close()
	apiClient := &api.ApiClient{
		HttpClient: ghttp.NewClientBuilder().WithHost(server.URL).Build(),
		Config:     configuration.GetConfig(),
	}
	randomUser := &RandomUserData{
		seed:      seed,
		userCount: userCount,
		latinOnly: false,
		namesOnly: true,
		apiClient: apiClient,
	}
	randomData, err := randomUser.Generate()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(randomData))
	assert.Equal(t, "John", randomData[0].Name.First)
	assert.Equal(t, "Doe", randomData[0].Name.Last)
}
