package main

import (
	"fmt"
	"github.com/versoul/zabroniryi-ru-connector-go"
	"strconv"
)

var (
	auth = zabroniryiru.Auth{
		BuyerId:  "TMCon",
		UserId:   "tmcon",
		Password: "vcxq11cz!",
		Language: "ru",
	}
	zApi = zabroniryiru.NewApi(auth)
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
