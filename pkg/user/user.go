package user

import (
	"fmt"
	"random-user-generator/internal/api"
)

func NewRandomUserData(seed string, results int, namesOnly bool, latinOnly bool) *RandomUserData {
	return &RandomUserData{
		seed:      seed,
		results:   results,
		namesOnly: namesOnly,
		latinOnly: latinOnly,
		apiClient: api.NewClient(nil),
	}
}

type RandomUserData struct {
	seed      string
	results   int
	namesOnly bool
	latinOnly bool
	apiClient *api.ApiClient
}

func (r *RandomUserData) Generate() ([]User, error) {
	resp, err := r.apiClient.FetchRandomUserInfo(r.seed, r.results, r.namesOnly, r.latinOnly)
	if err != nil {
		return nil, err
	}
	userData := &Users{}
	if err = resp.FromJson(userData); err != nil {
		return nil, fmt.Errorf("failed to generate random user data: %s", err.Error())
	}
	return userData.Results, nil
}
