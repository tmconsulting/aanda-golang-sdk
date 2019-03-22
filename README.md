# Aanda SDK for Reservation Hotels zabroniryi.ru
[![CircleCI](https://circleci.com/gh/tmconsulting/aanda-golang-sdk/tree/master.svg?style=shield)](https://circleci.com/gh/tmconsulting/aanda-golang-sdk)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/tmconsulting/aanda-sdk/blob/master/LICENSE)

### Install
```
go get github.com/tmconsulting/aanda-golang-sdk
```

### Import
```golang
import (
    "context"
    "log"
    
    "github.com/tmconsulting/aanda-golang-sdk"
)
```

### Example init variables
```golang
var (
	auth = aandaSdk.Auth{
		BuyerId:  "BuyerId",
		UserId:   "BuyerId",
		Password: "Password",
		Language: "ru",
	}
	aApi = aandaSdk.NewApi(auth)
	
	ctx = context.WithValue(context.Backround(), "requestId", "b5cd6a4a-efee-4146-8bec-bf5457558750")
)
```

### Example init hooks/handlers
```golang
	aApi.RegisterEventHandler(aandaSdk.BeforeRequestSend, func(ctx context.Context, methodName, mimeType string, data []byte) {
	    requestId := ctx.Value("requestId").(string)
	
            log.Println("request: ", requestId, methodName, mimeType, string(data))
	}).RegisterEventHandler(aandaSdk.AfterResponseReceive, func(ctx context.Context, methodName, mimeType string, data []byte) {
	    requestId := ctx.Value("requestId").(string)

            log.Println("response: ", requestId, methodName, mimeType, string(data))
	})
```

### Example HotelSearchRequest
```golang
searchReq := aandaSdk.HotelSearchRequest{
	City:           "2",
	Lat:            "",
	Lng:            "",
	Radius:         "30",
	ArrivalDate:    "09.11.2017",
	DepartureDate:  "10.11.2017",
	PriceFrom:      "2000",
	PriceTo:        "3000",
	NumberOfGuests: "1",
}
data, err := aApi.HotelSearchRequest(ctx, searchReq)
if err == nil {
	//Work with data
}
 ```

 ### Example HotelPricingRequest
```golang
priceReq := aandaSdk.HotelPricingRequest{
	Hotel:          "2150",
	ProductCode:    "",
	Currency:       "1",
	WhereToPay:     "1",
	ArrivalDate:    "05.12.2017",
	DepartureDate:  "06.12.2017",
	ArrivalTime:    "2000",
	DepartureTime:  "3000",
	NumberOfGuests: "1",
}
data, err := aApi.HotelPricingRequest(ctx, priceReq)
if err == nil {
	//Work with data
}
 ```

### Example OrderRequest
```golang
orderReq := aandaSdk.OrderRequest{
	ArrivalDate:   "02.01.2018",
	DepartureDate: "03.01.2018",
	AddInfo:       "ТЕСТОВЫЙ ЗАКАЗ",
	HotelCode:     "2150",
	RoomCode:      "32078",
	Meal:           "",
	ArrivalTime:    "14:00",
	DepartureTime:  "12:00",
	NumberOfGuests: "1",
	Person: []aandaSdk.Person{aandaSdk.Person{
		FirstName: "Name",
		LastName:  "Name",
	}},
}
data, err := aApi.OrderRequest(ctx, orderReq)
if err == nil {
	//Work with data
}
 ```

### Example OrderInfoRequest
```golang
data, err := aApi.OrderInfoRequest(ctx, "2213397") // 2213397 is OrderId
if err == nil {
	//Work with data
}
 ```

 ### Example OrderListRequest
```golang
orderReq := aandaSdk.OrderListRequest{
	//ArrivalDateFrom: "24.02.2018",
	//ArrivalDateTo:   "25.02.2018",
	//DepartureDateFrom: "25.02.2018",
	//DepartureDateTo:   "26.02.2018",
	//RegistrationDateFrom: "23.02.2018",
	//RegistrationDateTo: "22.02.2018",
	}
data, err := aApi.OrderListRequest(ctx, orderReq)
if err == nil {
	//Work with data
}
 ```

### Example SendOrderMessageRequest
```golang
somReq := aandaSdk.SendOrderMessageRequest{
	OrderId: "2213397",
	Message: "test message 2 3 4 22 22",
}
data, err := aApi.SendOrderMessageRequest(ctx, somReq)
if err == nil {
	//Work with data
}
 ```

 ### Example OrderMessagesRequest
```golang
data, err := aApi.OrderMessagesRequest(ctx, "2213397") // 2213397 is OrderId
if err == nil {
	//Work with data
}
 ```

### Example CountryListRequest
```golang
data, err := aApi.CountryListRequest(ctx)
if err == nil {
	//Work with data
}
 ```

### Example CityListRequest
```golang
data, err := aApi.CityListRequest(ctx, 9)//9 is CountryCode
if err == nil {
	//Work with data
}
 ```

### Example HotelListRequest
```golang
data, err := aApi.HotelListRequest(ctx, 1)//1 is CityCode == Moscow
if err == nil {
	//Work with data
}
 ```

 ### Example HotelDescriptionRequest
```golang
data, err := aApi.HotelDescriptionRequest(ctx, 2150)//2150 is HotelCode
if err == nil {
	//Work with data
}
 ```

### Example CurrencyListRequest
```golang
data, err := aApi.CurrencyListRequest(ctx)
if err == nil {
	//Work with data
}
 ```

 ### Example MealTypeRequest
```golang
data, err := aApi.MealTypeRequest(ctx)
if err == nil {
	//Work with data
}
 ```

 ### Example MealCategoryeRequest
```golang
data, err := aApi.MealCategoryRequest(ctx)
if err == nil {
	//Work with data
}
 ```

  ### Example ClientStatusRequest
 ```golang
 data, err := aApi.ClientStatusRequest(ctx)
 if err == nil {
 	//Work with data
 }
  ```
