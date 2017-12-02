# zabroniryi.ru SDK
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

### Example CountryListRequest
```golang
package main

func main() {
	data, err := aApi.CountryListRequest()
	if err == nil {
		//Work with data
	}
}
 ```

### Example CityListRequest
```golang
package main

func main() {
	data, err := aApi.CityListRequest(9)//9 is CountryCode
	if err == nil {
		//Work with data
	}
}
 ```

### Example HotelListRequest
```golang
package main

func main() {
	data, err := aApi.HotelListRequest(1)//1 is CityCode == Moscow
	if err == nil {
		//Work with data
	}
}
 ```

 ### Example HotelDescriptionRequest
```golang
package main

func main() {
	data, err := aApi.HotelDescriptionRequest(2150)//2150 is HotelCode
	if err == nil {
		//Work with data
	}
}
 ```

### Example CurrencyListRequest
```golang
package main

func main() {
	data, err := aApi.CurrencyListRequest()
	if err == nil {
		//Work with data
	}
}
 ```

 ### Example MealTypeRequest
```golang
package main

func main() {
	data, err := aApi.MealTypeRequest()
	if err == nil {
		//Work with data
	}
}
 ```

 ### Example MealCategoryeRequest
```golang
package main

func main() {
	data, err := aApi.MealCategoryRequest()
	if err == nil {
		//Work with data
	}
}
 ```