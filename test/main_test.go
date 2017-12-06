package test

import (
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
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

func main() {
	defer gock.Off()
}

func getJson(fileName string) (data []byte) {
	data, _ = ioutil.ReadFile("data/" + fileName)
	return data
}

func testRequest(filename string) {
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson(filename))
}
