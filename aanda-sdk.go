package aandaSdk

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const apiUrl = "http://api.aanda.ru/xml_gateway/"

const (
	TestOrderId = "2213397"
)

type Api struct {
	EventListener

	BuyerId  string
	UserId   string
	Password string
	Language string
	url      string
	urlInfo  *url.URL
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

	self.urlInfo, _ = url.Parse(self.url)

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

func (self *Api) sendReq(ctx context.Context, requestName string, params url.Values) ([]byte, error) {
	reqContentType := "application/x-www-form-urlencoded"
	reqData := []byte(params.Encode())

	req, err := http.NewRequest("POST", self.url, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", reqContentType)

	self.EventListener.raiseEvent(BeforeRequestSend, ctx, requestName, self.urlInfo.RawQuery, reqContentType, reqData)

	resp, err := (&http.Client{}).Do(req.WithContext(ctx))
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

		self.EventListener.raiseEvent(AfterResponseReceive, ctx, requestName, self.urlInfo.RawQuery, respContentType, respData)
	}

	return respData, err
}

func (self *Api) sendReqParams(ctx context.Context, requestName string, params map[string]string) ([]byte, error) {
	return self.sendReq(ctx, requestName, self.createReq(requestName, params))
}

func (self *Api) sendReqWithAuth(ctx context.Context, requestName string, requestObj interface{}, auth *Auth, params map[string]string) ([]byte, error) {
	p, err := self.createReqWithAuth(requestName, requestObj, auth, params)
	if err != nil {
		return nil, err
	}

	return self.sendReq(ctx, requestName, p)
}

func (self *Api) CountryListRequest(ctx context.Context) ([]CountryListResponse, error) {
	body, err := self.sendReqParams(ctx, "CountryListRequest", nil)

	resp := []CountryListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) CityListRequest(ctx context.Context, countryCode int) (CityListResponse, error) {
	body, err := self.sendReqParams(ctx, "CityListRequest", map[string]string{
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

func (self *Api) HotelListRequest(ctx context.Context, cityCode int) ([]HotelListResponse, error) {
	body, err := self.sendReqParams(ctx, "HotelListRequest", map[string]string{
		"CityCode": strconv.Itoa(cityCode),
	})

	resp := []HotelListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelDescriptionRequest(ctx context.Context, hotelCode int) (HotelDescriptionResponse, error) {
	body, err := self.sendReqParams(ctx, "HotelDescriptionRequest", map[string]string{
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

func (self *Api) CurrencyListRequest(ctx context.Context) ([]CurrencyListResponse, error) {
	body, err := self.sendReqParams(ctx, "CurrencyListRequest", nil)

	resp := []CurrencyListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) MealTypeRequest(ctx context.Context) ([]MealTypeResponse, error) {
	body, err := self.sendReqParams(ctx, "MealTypeRequest", nil)
	resp := []MealTypeResponse{}

	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) MealCategoryRequest(ctx context.Context) ([]MealCategoryResponse, error) {
	body, err := self.sendReqParams(ctx, "MealCategoryRequest", nil)

	resp := []MealCategoryResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) ServiceTypeRequest(ctx context.Context) ([]ServiceTypeResponse, error) {
	body, err := self.sendReqParams(ctx, "ServiceTypeRequest", nil)

	resp := []ServiceTypeResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelPricingRequest(ctx context.Context, priceReq HotelPricingRequest) (HotelPricingResponse, error) {
	body, err := self.sendReqWithAuth(ctx, "HotelSearchRequest", &priceReq, &priceReq.Auth, nil)

	resp := HotelPricingResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelSearchRequest(ctx context.Context, searchReq HotelSearchRequest) ([]HotelSearchResponse, error) {
	body, err := self.sendReqWithAuth(ctx, "HotelSearchRequest", &searchReq, &searchReq.Auth, nil)

	resp := []HotelSearchResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) OrderRequest(ctx context.Context, orderReq OrderRequest) (OrderRequestResponse, error) {
	body, err := self.sendReqWithAuth(ctx, "OrderRequest", &orderReq, &orderReq.Auth, nil)

	resp := OrderRequestResponse{}
	if err == nil {
		err = parseResp(body, &resp, validateStruct, nil)
	}

	return resp, err
}

func (self *Api) OrderInfoRequest(ctx context.Context, orderId string) (OrderInfoResponse, error) {
	orderInfoReq := &OrderInfoRequest{
		Id: orderId,
	}

	body, err := self.sendReqWithAuth(ctx, "OrderInfoRequest", orderInfoReq, &orderInfoReq.Auth, nil)

	resp := OrderInfoResponse{}
	if err == nil {
		if err = parseResp(body, &resp, validateStruct, nil); err != nil {
			resp = OrderInfoResponse{}
		}
	}

	return resp, err
}

func (self *Api) OrderMessagesRequest(ctx context.Context, orderId string) ([]OrderMessagesResponse, error) {
	body, err := self.sendReqParams(ctx, "OrderMessagesRequest", map[string]string{
		"order_id": orderId,
	})

	resp := []OrderMessagesResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) SendOrderMessageRequest(ctx context.Context, somReq SendOrderMessageRequest) (SendOrderMessageResponse, error) {
	body, err := self.sendReqParams(ctx, "SendOrderMessageRequest", map[string]string{
		"order_id": somReq.OrderId,
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

func (self *Api) OrderListRequest(ctx context.Context, orderReq OrderListRequest) ([]OrderListResponse, error) {
	body, err := self.sendReqWithAuth(ctx, "OrderListRequest", &orderReq, &orderReq.Auth, nil)

	resp := []OrderListResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) ClientStatusRequest(ctx context.Context) ([]ClientStatusResponse, error) {
	body, err := self.sendReqParams(ctx, "ClientStatusRequest", nil)

	resp := []ClientStatusResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) HotelAmenitiesRequest(ctx context.Context) ([]HotelAmenitiesResponse, error) {
	body, err := self.sendReqParams(ctx, "HotelAmenitiesRequest", nil)

	resp := []HotelAmenitiesResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}

func (self *Api) RoomAmenitiesRequest(ctx context.Context) ([]RoomAmenitiesResponse, error) {
	body, err := self.sendReqParams(ctx, "RoomAmenitiesRequest", nil)

	resp := []RoomAmenitiesResponse{}
	if err == nil {
		err = parseResp(body, &resp, nil, nil)
	}

	return resp, err
}
