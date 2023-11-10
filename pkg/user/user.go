package user

import "random-user-generator/internal/api"

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

func (r *RandomUserData) GetInfo() {
	api.FetchRandomUserInfo()
}
