package aandaSdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"gopkg.in/go-playground/validator.v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const apiUrl = "http://api.aanda.ru/xml_gateway/"

type Api struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

func sendReq(data url.Values) []byte {
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	return body

}
func parseError(body []byte) error {
	//Try parse as JSON
	var answM map[string]interface{}
	err := json.Unmarshal(body, &answM)
	if err == nil {
		//if have only error key
		if err, ok := answM["error"].(string); ok {
			return errors.New(err)
		}
		//if have Status=Error
		st, ok1 := answM["Status"].(string)
		note, ok2 := answM["note"].(string)
		if ok1 && ok2 && st == "Error" {
			return errors.New(note)
		}
	}
	//may be simple string, try parse it
	var answS string
	err = json.Unmarshal(body, &answS)
	if err == nil {
		if strings.Index(answS, "Ошибка авторазации") == -1 {
			return errors.New("Authorization error")
		} else {
			errors.New(answS)
		}
	}
	//may be part of XML, try parse in by regexp
	re := regexp.MustCompile(`Note="(.*)"`)
	res := re.FindStringSubmatch(string(body))
	if len(res) != 0 {
		return errors.New(res[1])
	}

	return nil
}

func NewApi(auth Auth) *Api {
	return &Api{
		BuyerId:  auth.BuyerId,
		UserId:   auth.UserId,
		Password: auth.Password,
		Language: auth.Language,
	}
}

func (self *Api) createDataReq(req map[string]string) url.Values {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)
	for key, value := range req {
		data.Add(key, value)
	}

	return data
}

func (self *Api) CountryListRequest() ([]CountryListResponse, error) {
	req := map[string]string{
		"RequestName": "CounttryListRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)
	jsonData := []CountryListResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) CityListRequest(countryCode int) (CityListResponse, error) {
	req := map[string]string{
		"RequestName": "CityListRequest",
		"CountryCode": strconv.Itoa(countryCode),
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	//FIX BUG of API (space in key)
	bodyStr := strings.Replace(string(body), "city _code", "city_code", -1)
	jsonData := CityListResponse{}

	err := json.Unmarshal([]byte(bodyStr), &jsonData)
	if err != nil {
		return CityListResponse{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelListRequest(cityCode int) ([]HotelListResponse, error) {
	req := map[string]string{
		"RequestName": "HotelListRequest",
		"CityCode":    strconv.Itoa(cityCode),
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []HotelListResponse{}

	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelDescriptionRequest(hotelCode int) (HotelDescriptionResponse, error) {
	req := map[string]string{
		"RequestName": "HotelDescriptionRequest",
		"HotelCode":   strconv.Itoa(hotelCode),
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := HotelDescriptionResponse{}

	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return HotelDescriptionResponse{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) CurrencyListRequest() ([]CurrencyListResponse, error) {
	req := map[string]string{
		"RequestName": "CurrencyListRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []CurrencyListResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) MealTypeRequest() ([]MealTypeResponse, error) {
	req := map[string]string{
		"RequestName": "MealTypeRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []MealTypeResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) MealCategoryRequest() ([]MealCategoryResponse, error) {
	req := map[string]string{
		"RequestName": "MealCategoryRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []MealCategoryResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelPricingRequest(priceReq HotelPricingRequest) (HotelPricingResponse, error) {
	priceReq.BuyerId = self.BuyerId
	priceReq.UserId = self.UserId
	priceReq.Password = self.Password
	priceReq.Language = self.Language

	jsonReq, err := json.Marshal(priceReq)
	if err != nil {
		panic(err)
	}

	req := map[string]string{
		"RequestName": "HotelSearchRequest",
		"JSON":        string(jsonReq),
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := HotelPricingResponse{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return HotelPricingResponse{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelSearchRequest(searchReq HotelSearchRequest) ([]HotelSearchResponse, error) {
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

	body := sendReq(data)
	jsonData := []HotelSearchResponse{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) OrderRequest(orderReq OrderRequest) (OrderRequestResponse, error) {
	orderReq.BuyerId = self.BuyerId
	orderReq.UserId = self.UserId
	orderReq.Password = self.Password
	orderReq.Language = self.Language

	jsonReq, err := json.Marshal(orderReq)
	if err != nil {
		panic(err)
	}

	req := map[string]string{
		"RequestName": "OrderRequest",
		"JSON":        string(jsonReq),
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := OrderRequestResponse{}
	err = json.Unmarshal(body, &jsonData)
	vErr := validator.ValidateStruct(jsonData)
	if err != nil || vErr != nil {
		return OrderRequestResponse{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) OrderInfoRequest(id int) (OrderInfoResponse, error) {
	orderInfReq := OrderInfoRequest{}
	orderInfReq.BuyerId = self.BuyerId
	orderInfReq.UserId = self.UserId
	orderInfReq.Password = self.Password
	orderInfReq.Language = self.Language
	orderInfReq.Id = strconv.Itoa(id)

	jsonReq, err := json.Marshal(orderInfReq)
	if err != nil {
		panic(err)
	}

	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "OrderInfoRequest")
	data.Add("JSON", string(jsonReq))

	body := sendReq(data)

	jsonData := OrderInfoResponse{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return OrderInfoResponse{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) OrderMessagesRequest(orderId int) ([]OrderMessagesResponse, error) {
	req := map[string]string{
		"RequestName": "OrderMessagesRequest",
		"order_id":    strconv.Itoa(2213397),
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []OrderMessagesResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) SendOrderMessageRequest(somReq SendOrderMessageRequest) (SendOrderMessageResponse, error) {
	req := map[string]string{
		"RequestName": "OrderMessagesRequest",
		"order_id":    strconv.Itoa(somReq.OrderId),
		"Message":     somReq.Message,
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := SendOrderMessageResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return SendOrderMessageResponse{}, parseError(body)
	}

	return jsonData, nil
}
