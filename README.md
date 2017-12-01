# zabroniryi.ru SDK
[![CircleCI](https://circleci.com/gh/tmconsulting/zabroniryiru-sdk/tree/develop.svg?style=shield)](https://circleci.com/gh/tmconsulting/zabroniryiru-sdk)
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
package main

func main() {
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
}
 ```