package aandaSdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
	var answM map[string]interface{}
	err := json.Unmarshal(body, &answM)
	if err == nil {
		if err, ok := answM["error"].(string); ok {
			return errors.New(err)
		}

		st, ok1 := answM["Status"].(string)
		note, ok2 := answM["note"].(string)
		if ok1 && ok2 && st == "Error" {
			return errors.New(note)
		}
	}

	var answS string
	err = json.Unmarshal(body, &answS)
	if err == nil {
		if strings.Index(answS, "Ошибка авторазации") == -1 {
			return errors.New("Ошибка авторазиции")
		}
	}

	re := regexp.MustCompile(`Note="(.*)"`)
	res := re.FindStringSubmatch(string(body))
	if len(res) != 0 {
		return errors.New(res[1])
	}

	return nil
}

func main() {
	//validate = validator.New()
	//validate.RegisterStructValidation(UserStructLevelValidation, User{})
}

func NewApi(auth Auth) *Api {
	return &Api{
		BuyerId:  auth.BuyerId,
		UserId:   auth.UserId,
		Password: auth.Password,
		Language: auth.Language,
	}
}

func (self *Api) CountryListRequest() ([]CountryListAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "CountryListRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)

	body := sendReq(data)
	jsonData := []CountryListAnswer{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) CityListRequest(countryCode int) (CityListAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "CityListRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)
	data.Add("CountryCode", strconv.Itoa(countryCode))

	body := sendReq(data)

	//FIX BUG of API
	bodyStr := strings.Replace(string(body), "city _code", "city_code", -1)
	jsonData := CityListAnswer{}

	err := json.Unmarshal([]byte(bodyStr), &jsonData)
	if err != nil {
		return CityListAnswer{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelListRequest(cityCode int) ([]HotelListAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "HotelListRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)
	data.Add("CityCode", strconv.Itoa(cityCode))

	body := sendReq(data)

	jsonData := []HotelListAnswer{}

	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelDescriptionRequest(hotelCode int) (HotelDescriptionAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "HotelDescriptionRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)
	data.Add("HotelCode", strconv.Itoa(hotelCode))

	body := sendReq(data)

	jsonData := HotelDescriptionAnswer{}

	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return HotelDescriptionAnswer{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) CurrencyListRequest() ([]CurrencyListAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "CurrencyListRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)

	body := sendReq(data)

	jsonData := []CurrencyListAnswer{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) MealTypeRequest() ([]MealTypeAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "MealTypeRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)

	body := sendReq(data)

	jsonData := []MealTypeAnswer{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) MealCategoryRequest() ([]MealCategoryAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "MealCategoryRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)

	body := sendReq(data)

	jsonData := []MealCategoryAnswer{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) HotelPricingRequest(priceReq HotelPricingRequest) (HotelPricingAnswer, error) {
	priceReq.BuyerId = self.BuyerId
	priceReq.UserId = self.UserId
	priceReq.Password = self.Password
	priceReq.Language = self.Language

	jsonReq, err := json.Marshal(priceReq)
	if err != nil {
		panic(err)
	}

	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "HotelPricingRequest")
	data.Add("JSON", string(jsonReq))

	body := sendReq(data)

	jsonData := HotelPricingAnswer{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return HotelPricingAnswer{}, parseError(body)
	}

	return jsonData, nil
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

	body := sendReq(data)
	jsonData := []HotelSearchAnswer{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) OrderRequest(orderReq OrderRequest) (OrderRequestAnswer, error) {
	orderReq.BuyerId = self.BuyerId
	orderReq.UserId = self.UserId
	orderReq.Password = self.Password
	orderReq.Language = self.Language

	jsonReq, err := json.Marshal(orderReq)
	if err != nil {
		panic(err)
	}

	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "OrderRequest")
	data.Add("JSON", string(jsonReq))

	body := sendReq(data)

	jsonData := OrderRequestAnswer{}
	err = json.Unmarshal(body, &jsonData)
	vErr := validator.ValidateStruct(jsonData)
	if err != nil || vErr != nil {
		return OrderRequestAnswer{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) OrderInfoRequest(id int) (OrderInfoAnswer, error) {
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

	jsonData := OrderInfoAnswer{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return OrderInfoAnswer{}, parseError(body)
	}

	return jsonData, nil
}

func (self *Api) OrderMessagesRequest(orderId int) ([]MealCategoryAnswer, error) {
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "OrderMessagesRequest")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)
	data.Add("order_id", "2213397")

	body := sendReq(data)
	fmt.Println("HERE")
	fmt.Println(string(body))
	jsonData := []MealCategoryAnswer{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, parseError(body)
	}

	return jsonData, nil
}
