package aandaSdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const apiUrl = "http://api.aanda.ru/xml_gateway/"

const (
	TestOrderId = 2213397
)

type Api struct {
	EventListener

	BuyerId  string
	UserId   string
	Password string
	Language string
	url      string
}

func NewApi(auth Auth) *Api {
	return (&Api{
		BuyerId:  auth.BuyerId,
		UserId:   auth.UserId,
		Password: auth.Password,
		Language: auth.Language,
		url:      apiUrl,
	}).init()
}

func (self *Api) init() *Api {
	self.EventListener.Init()

	return self
}

func (self *Api) createReq(requestName string, params map[string]string) url.Values {
	data := url.Values{}

	data.Set("RequestName", requestName)
	data.Set("RequestType", "json")
	data.Add("CompanyId", self.BuyerId)
	data.Add("UserId", self.UserId)
	data.Add("Password", self.Password)
	data.Add("Language", self.Language)

	for key, value := range params {
		data.Add(key, value)
	}

	return data
}

func (self *Api) createReqWithAuth(requestName string, requestObj interface{}, auth *Auth, params map[string]string) (url.Values, error) {
	auth.BuyerId = self.BuyerId
	auth.UserId = self.UserId
	auth.Password = self.Password
	auth.Language = self.Language

	jsonReq, err := json.Marshal(requestObj)
	if err != nil {
		return nil, err
	}

	data := url.Values{}
	data.Set("RequestName", requestName)
	data.Set("RequestType", "json")
	data.Add("JSON", string(jsonReq))

	for key, value := range params {
		data.Add(key, value)
	}

	return data, nil
}

func (self *Api) sendReq(requestName string, params url.Values) ([]byte, error) {
	reqContentType := "application/x-www-form-urlencoded"
	reqData := []byte(params.Encode())

	self.EventListener.raiseEvent(BeforeRequestSend, requestName, reqContentType, reqData)

	req, err := http.NewRequest("POST", self.url, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", reqContentType)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		respContentType := resp.Header.Get("Content-Type")
		if len(respContentType) == 0 {
			respContentType = "text/json"
		} else {
			parts := strings.Split(respContentType, "; ")
			respContentType = strings.TrimSpace(parts[0])
		}

		self.EventListener.raiseEvent(AfterResponseReceive, requestName, respContentType, respData)
	}

	return respData, err
}

func (self *Api) sendReqParams(requestName string, params map[string]string) ([]byte, error) {
	return self.sendReq(requestName, self.createReq(requestName, params))
}

func (self *Api) sendReqWithAuth(requestName string, requestObj interface{}, auth *Auth, params map[string]string) ([]byte, error) {
	p, err := self.createReqWithAuth(requestName, requestObj, auth, params)
	if err != nil {
		return nil, err
	}

	return self.sendReq(requestName, p)
}

func (self *Api) CountryListRequest() ([]CountryListResponse, error) {
	body, err := self.sendReqParams("CountryListRequest", nil)

	resp := []CountryListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) CityListRequest(countryCode int) (CityListResponse, error) {
	body, err := self.sendReqParams("CityListRequest", map[string]string{
		"CountryCode": strconv.Itoa(countryCode),
	})

	resp := CityListResponse{}
	if err == nil {
		//FIX BUG of API (space in key)
		body = []byte(strings.Replace(string(body), "city _code", "city_code", -1))

		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelListRequest(cityCode int) ([]HotelListResponse, error) {
	body, err := self.sendReqParams("HotelListRequest", map[string]string{
		"CityCode": strconv.Itoa(cityCode),
	})

	resp := []HotelListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelDescriptionRequest(hotelCode int) (HotelDescriptionResponse, error) {
	body, err := self.sendReqParams("HotelDescriptionRequest", map[string]string{
		"HotelCode": strconv.Itoa(hotelCode),
	})

	resp := HotelDescriptionResponse{}
	if err == nil {
		if err = parseResp(body, &resp, nil, func() *string { return resp.Status }); err != nil {
			resp = HotelDescriptionResponse{}
		}
	}

	return resp, err
}

func (self *Api) CurrencyListRequest() ([]CurrencyListResponse, error) {
	body, err := self.sendReqParams("CurrencyListRequest", nil)

	resp := []CurrencyListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) MealTypeRequest() ([]MealTypeResponse, error) {
	body, err := self.sendReqParams("MealTypeRequest", nil)
	resp := []MealTypeResponse{}

	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) MealCategoryRequest() ([]MealCategoryResponse, error) {
	body, err := self.sendReqParams("MealCategoryRequest", nil)

	resp := []MealCategoryResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) ServiceTypeRequest() ([]ServiceTypeResponse, error) {
	body, err := self.sendReqParams("ServiceTypeRequest", nil)

	resp := []ServiceTypeResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelPricingRequest(priceReq HotelPricingRequest) (HotelPricingResponse, error) {
	body, err := self.sendReqWithAuth("HotelSearchRequest", &priceReq, &priceReq.Auth, nil)

	resp := HotelPricingResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelSearchRequest(searchReq HotelSearchRequest) ([]HotelSearchResponse, error) {
	body, err := self.sendReqWithAuth("HotelSearchRequest", &searchReq, &searchReq.Auth, nil)

	resp := []HotelSearchResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) OrderRequest(orderReq OrderRequest) (OrderRequestResponse, error) {
	body, err := self.sendReqWithAuth("OrderRequest", &orderReq, &orderReq.Auth, nil)

	resp := OrderRequestResponse{}
	if err == nil {
		err = parseResp(body, &resp, validateStruct, nil)
	}

	return resp, err
}

func (self *Api) OrderInfoRequest(id int) (OrderInfoResponse, error) {
	orderInfoReq := &OrderInfoRequest{
		Id: strconv.Itoa(id),
	}

	body, err := self.sendReqWithAuth("OrderInfoRequest", orderInfoReq, &orderInfoReq.Auth, nil)

	resp := OrderInfoResponse{}
	if err == nil {
		if err = parseResp(body, &resp, validateStruct, nil); err != nil {
			resp = OrderInfoResponse{}
		}
	}

	return resp, err
}

func (self *Api) OrderMessagesRequest(orderId int) ([]OrderMessagesResponse, error) {
	body, err := self.sendReqParams("OrderMessagesRequest", map[string]string{
		"order_id": strconv.Itoa(orderId),
	})

	resp := []OrderMessagesResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) SendOrderMessageRequest(somReq SendOrderMessageRequest) (SendOrderMessageResponse, error) {
	body, err := self.sendReqParams("SendOrderMessageRequest", map[string]string{
		"order_id": strconv.Itoa(somReq.OrderId),
		"Message":  somReq.Message,
	})

	resp := SendOrderMessageResponse{}
	if err == nil {
		if err = parseResp(body, &resp, nil, nil); err != nil {
			resp = SendOrderMessageResponse{}
		}
	}

	return resp, err
}

func (self *Api) OrderListRequest(orderReq OrderListRequest) ([]OrderListResponse, error) {
	body, err := self.sendReqWithAuth("OrderListRequest", &orderReq, &orderReq.Auth, nil)

	resp := []OrderListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) ClientStatusRequest() ([]ClientStatusResponse, error) {
	body, err := self.sendReqParams("ClientStatusRequest", nil)

	resp := []ClientStatusResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelAmenitiesRequest() ([]HotelAmenitiesResponse, error) {
	body, err := self.sendReqParams("HotelAmenitiesRequest", nil)

	resp := []HotelAmenitiesResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) RoomAmenitiesRequest() ([]RoomAmenitiesResponse, error) {
	body, err := self.sendReqParams("RoomAmenitiesRequest", nil)

	resp := []RoomAmenitiesResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}
