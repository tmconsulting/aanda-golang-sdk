package test

import (
	"github.com/tmconsulting/zabroniryiru-sdk"
	"io/ioutil"
)

var (
	auth = zabroniryiru_sdk.Auth{
		BuyerId:  "BuyerId",
		UserId:   "UserId",
		Password: "Password",
		Language: "ru",
	}
	zApi = zabroniryiru_sdk.NewApi(auth)
)

func getJson(fileName string) (data []byte) {
	data, _ = ioutil.ReadFile("data/"+fileName)
	return data
}

