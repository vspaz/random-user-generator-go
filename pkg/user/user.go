package user

import (
	"log"
	"random-user-generator/internal/api"
)

func NewRandomUserData(seed string, results uint, namesOnly bool) RandomUserData {
	return RandomUserData{
		seed:      seed,
		results:   results,
		namesOnly: namesOnly,
	}
}

type RandomUserData struct {
	seed      string
	results   uint
	namesOnly bool
}

func (r *RandomUserData) Generate() *[]User {
	resp := api.FetchRandomUserInfo(r.seed, r.results, r.namesOnly)
	userData := &Users{}
	err := resp.FromJson(userData)
	if err != nil {
		log.Fatalf("failed to generate random user data")
		return nil
	}
	return &userData.Results
}
