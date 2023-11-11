package user

import (
	"fmt"
	"random-user-generator/internal/api"
)

func NewRandomUserData(seed string, results int, namesOnly bool) *RandomUserData {
	return &RandomUserData{
		seed:      seed,
		results:   results,
		namesOnly: namesOnly,
		apiClient: api.NewClient(nil),
	}
}

type RandomUserData struct {
	seed      string
	results   int
	namesOnly bool
	apiClient *api.ApiClient
}

func (r *RandomUserData) Generate() ([]User, error) {
	resp := r.apiClient.FetchRandomUserInfo(r.seed, r.results, r.namesOnly)
	userData := &Users{}
	if err := resp.FromJson(userData); err != nil {
		return nil, fmt.Errorf("failed to generate random user data: %s", err.Error())
	}
	return userData.Results, nil
}
