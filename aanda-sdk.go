package aandaSdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/go-playground/validator.v2"
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
	return body

}

func parseError(body []byte) error {
	// Try to parse known struct
	var aandaErr AandaError
	err := json.Unmarshal(body, &aandaErr)
	if err == nil && !aandaErr.IsEmpty() {
		return aandaErr.ToError()
	}

	var aandaErrMsg AandaErrorMsg
	err = json.Unmarshal(body, &aandaErrMsg)
	if err == nil && !aandaErrMsg.IsEmpty() {
		return aandaErrMsg.ToError()
	}

	//Try parse as JSON
	var answM map[string]interface{}
	err = json.Unmarshal(body, &answM)
	if err == nil {
		//if have only error key
		if err, ok := answM["error"].(string); ok {
			return errors.New(err)
		}
		//if have error and message key
		_, ok1 := answM["error"].(float64)
		message, ok2 := answM["message"].(string)
		if ok1 && ok2 {
			return errors.New(message)
		}
		//if have Status=Error
		st, ok3 := answM["Status"].(string)
		note, ok4 := answM["note"].(string)
		if ok3 && ok4 && st == "Error" {
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
			return errors.New(answS)
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
		"RequestName": "CountryListRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)
	jsonData := []CountryListResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return CityListResponse{}, err
		} else {
			return CityListResponse{}, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return HotelDescriptionResponse{}, err
		} else {
			return HotelDescriptionResponse{}, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
	}

	return jsonData, nil
}

func (self *Api) ServiceTypeRequest() ([]ServiceTypeResponse, error) {
	req := map[string]string{
		"RequestName": "ServiceTypeRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []ServiceTypeResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return HotelPricingResponse{}, err
		} else {
			return HotelPricingResponse{}, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			if err != nil {
				return OrderRequestResponse{}, err
			} else {
				return OrderRequestResponse{}, vErr
			}
		} else {
			return OrderRequestResponse{}, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return OrderInfoResponse{}, err
		} else {
			return OrderInfoResponse{}, respErr
		}
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
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
	}

	return jsonData, nil
}

func (self *Api) SendOrderMessageRequest(somReq SendOrderMessageRequest) (SendOrderMessageResponse, error) {
	req := map[string]string{
		"RequestName": "SendOrderMessageRequest",
		"order_id":    strconv.Itoa(somReq.OrderId),
		"Message":     somReq.Message,
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := SendOrderMessageResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return SendOrderMessageResponse{}, err
		} else {
			return SendOrderMessageResponse{}, respErr
		}
	}

	return jsonData, nil
}

func (self *Api) OrderListRequest(orderReq OrderListRequest) ([]OrderListResponse, error) {
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
	data.Add("RequestName", "OrderListRequest")
	data.Add("JSON", string(jsonReq))

	body := sendReq(data)

	jsonData := []OrderListResponse{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
	}

	return jsonData, nil
}

func (self *Api) ClientStatusRequest() ([]ClientStatusResponse, error) {
	req := map[string]string{
		"RequestName": "ClientStatusRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []ClientStatusResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
	}

	return jsonData, nil
}

func (self *Api) HotelAmenitiesRequest() ([]HotelAmenitiesResponse, error) {
	req := map[string]string{
		"RequestName": "HotelAmenitiesRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []HotelAmenitiesResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
	}

	return jsonData, nil
}

func (self *Api) RoomAmenitiesRequest() ([]RoomAmenitiesResponse, error) {
	req := map[string]string{
		"RequestName": "RoomAmenitiesRequest",
	}
	data := self.createDataReq(req)
	body := sendReq(data)

	jsonData := []RoomAmenitiesResponse{}
	err := json.Unmarshal(body, &jsonData)
	if err != nil {
		respErr := parseError(body)

		if respErr == nil {
			return nil, err
		} else {
			return nil, respErr
		}
	}

	return jsonData, nil
}
