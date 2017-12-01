package test

import (
	"github.com/tmconsulting/aanda-sdk"
	"io/ioutil"
)

var (
	auth = aandaSdk.Auth{
		BuyerId:  "BuyerId",
		UserId:   "UserId",
		Password: "Password",
		Language: "ru",
	}
	zApi = aandaSdk.NewApi(auth)
)

func getJson(fileName string) (data []byte) {
	data, _ = ioutil.ReadFile("data/" + fileName)
	return data
}
