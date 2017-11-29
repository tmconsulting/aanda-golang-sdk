# zabroniryi.ru SDK
[![CircleCI](https://circleci.com/gh/tmconsulting/zabroniryiru-sdk/tree/develop.svg?style=shield)](https://circleci.com/gh/tmconsulting/zabroniryiru-sdk)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/circleci/cci-demo-react/master/LICENSE)

## Using
```golang
package main
import (
	"fmt"
	"github.com/tmconsulting/zabroniryiru-sdk"
	"strconv"
)
var (
	auth = zabroniryiru_sdk.Auth{
		BuyerId:  "BuyerId",
		UserId:   "BuyerId",
		Password: "Password",
		Language: "ru",
	}
	zApi = zabroniryiru_sdk.NewApi(auth)
)
func main() {
	searchReq := zabroniryiru.HotelSearchRequest{
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
	data, err := zApi.HotelSearchRequest(searchReq)
	if err != nil {
		panic(err)
	} else {
		//Вывод результата
		for _, d := range data {
			fmt.Println("Город - " + d.CityName)
			for _, r := range d.Rooms {
				price := strconv.Itoa(r.Price)
				fmt.Println("Номер - " + r.RoomName + " - " + price + "руб.")
			}
			fmt.Println("")
		}
	}
}
 ```

## Docs

### HotelSearchRequest(searchReq)
Принимает параметры
<pre>
BuyerId        string `json:"BuyerId"`
UserId         string `json:"UserId"`
Password       string `json:"Password"`
Language       string `json:"Language"`
City           string `json:"City"`
Lat            string `json:"Lat"`
Lng            string `json:"Lng"`
Radius         string `json:"Radius"`
ArrivalDate    string `json:"ArrivalDate"`
DepartureDate  string `json:"DepartureDate"`
PriceFrom      string `json:"PriceFrom"`
PriceTo        string `json:"PriceTo"`
NumberOfGuests string `json:"NumberOfGuests"`
</pre>


