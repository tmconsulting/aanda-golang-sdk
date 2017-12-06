# Aanda SDK for Reservation Hotels zabroniryi.ru
[![CircleCI](https://circleci.com/gh/tmconsulting/aanda-sdk/tree/develop.svg?style=shield)](https://circleci.com/gh/tmconsulting/aanda-sdk)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/circleci/cci-demo-react/master/LICENSE)

### Install
```
go get github.com/tmconsulting/aanda-sdk
```

### Import
```golang
import "github.com/tmconsulting/aanda-sdk"
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
)
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
data, err := aApi.HotelSearchRequest(searchReq)
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
data, err := aApi.HotelPricingRequest(priceReq)
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
data, err := aApi.OrderRequest(orderReq)
if err == nil {
	//Work with data
}
 ```

### Example OrderInfoRequest
```golang
data, err := aApi.OrderInfoRequest(2213397) // 2213397 is OrderId
if err == nil {
	//Work with data
}
 ```

 ### Example OrderListRequest
```golang
orderReq := aandaSdk.OrderRequest{}
data, err := aApi.OrderListRequest(orderReq)
if err == nil {
	//Work with data
}
 ```

  ### Example SendOrderMessageRequest
```golang
somReq := aandaSdk.SendOrderMessageRequest{
	OrderId: 2213397,
	Message: "test message 2 3 4 22 22",
}
data, err := aApi.SendOrderMessageRequest(somReq)
if err == nil {
	//Work with data
}
 ```

 ### Example OrderMessagesRequest
```golang
data, err := aApi.OrderMessagesRequest(2213397) // 2213397 is OrderId
if err == nil {
	//Work with data
}
 ```

### Example CountryListRequest
```golang
data, err := aApi.CountryListRequest()
if err == nil {
	//Work with data
}
 ```

### Example CityListRequest
```golang
data, err := aApi.CityListRequest(9)//9 is CountryCode
if err == nil {
	//Work with data
}
 ```

### Example HotelListRequest
```golang
data, err := aApi.HotelListRequest(1)//1 is CityCode == Moscow
if err == nil {
	//Work with data
}
 ```

 ### Example HotelDescriptionRequest
```golang
data, err := aApi.HotelDescriptionRequest(2150)//2150 is HotelCode
if err == nil {
	//Work with data
}
 ```

### Example CurrencyListRequest
```golang
data, err := aApi.CurrencyListRequest()
if err == nil {
	//Work with data
}
 ```

 ### Example MealTypeRequest
```golang
data, err := aApi.MealTypeRequest()
if err == nil {
	//Work with data
}
 ```

 ### Example MealCategoryeRequest
```golang
data, err := aApi.MealCategoryRequest()
if err == nil {
	//Work with data
}
 ```