package zabroniryiru

type HotelSearchAnswer struct {
	HotelCode   string `json:"hotel_code"`
	HotelName   string `json:"hotel_name"`
	Address     string `json:"address"`
	ImageUrl    string `json:"image_url"`
	Vat         string `json:"vat"`
	Description string `json:"description"`
	Amenities   string `json:"amenities"`
	//CheckInTime    string `json:"check-in_timee"`
	//CheckOutTime   string `json:"check-out_time"`
	//Timezone       string `json:"timezone"`
	//CityCode       string `json:"city_code"`
	//CityName       string `json:"city_name"`
	HotelLatitude  string `json:"hotel_latitude"`
	HotelLongitude string `json:"hotel_longitude"`
	//CountryCode    string `json:"country_code"`
	//CountryName    string `json:"country_name"`
	RatingCode string `json:"rating_code"`
	RatingName string `json:"rating_name"`
	StarsCode  string `json:"stars_code"`
	StarsName  string `json:"stars_name"`
	//CurrencyCode   string `json:"currency_code"`
	//CurrencyName   string `json:"currency_name"`
}
