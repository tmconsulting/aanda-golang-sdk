package zabroniryiru_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiUrl = "http://api.aanda.ru/xml_gateway/"

type Api struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

func NewApi(auth Auth) *Api {
	return &Api{
		BuyerId:  auth.BuyerId,
		UserId:   auth.UserId,
		Password: auth.Password,
		Language: auth.Language,
	}
}
func (self *Api) HotelSearchRequest(searchReq HotelSearchRequest) ([]HotelSearchAnswer, error) {
	searchReq.BuyerId = self.BuyerId
	searchReq.UserId = self.UserId
	searchReq.Password = self.Password
	searchReq.Language = self.Language

	jsonReq, err := json.Marshal(searchReq)
	if err != nil {
		panic(err)
	}

	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "HotelSearchRequest")
	data.Add("JSON", string(jsonReq))
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	jsonData := []HotelSearchAnswer{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, errors.New("Ошибка запроса АПИ zabronirui.ru \n" + string(body))
	}
	return jsonData, nil
}