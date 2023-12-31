package user

import (
	"fmt"
	"random-user-generator/internal/api"
)

func NewRandomUserData(seed string, userCount int, nameOnly bool, latinOnly bool) *RandomUserData {
	return &RandomUserData{
		seed:      seed,
		userCount: userCount,
		nameOnly:  nameOnly,
		latinOnly: latinOnly,
		apiClient: api.NewClient(nil),
	}
}

type RandomUserData struct {
	seed      string
	userCount int
	nameOnly  bool
	latinOnly bool
	apiClient *api.ApiClient
}

func (r *RandomUserData) Generate() ([]User, error) {
	resp, err := r.apiClient.FetchRandomUserInfo(r.seed, r.userCount, r.nameOnly, r.latinOnly)
	if err != nil {
		return nil, err
	}
	userData := &Users{}
	if err = resp.FromJson(userData); err != nil {
		return nil, fmt.Errorf("failed to generate random user data: %s", err.Error())
	}
	return userData.Results, nil
}
