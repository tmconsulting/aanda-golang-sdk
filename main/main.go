package main

import (
	"fmt"
	"github.com/versoul/zabroniryi-ru-connector-go"
)

func main() {
	data := zabroniryiru.HotelSearchRequest()
	for _, d := range data {
		fmt.Println(d.CurrencyName)
	}
	zabroniryiru.HotelSearchRequest()
}
