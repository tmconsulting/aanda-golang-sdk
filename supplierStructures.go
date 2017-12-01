package aandaSdk

type HotelSearchAnswer struct {
	HotelCode      string       `json:"hotel_code"`
	HotelName      string       `json:"hotel_name"`
	Address        string       `json:"address"`
	ImageUrl       string       `json:"image_url"`
	Vat            int          `json:"vat"`
	Description    string       `json:"description"`
	Amenities      string       `json:"amenities"`
	CheckInTime    string       `json:"check-in_time"`
	CheckOutTime   string       `json:"check-out_time"`
	Timezone       string       `json:"timezone"`
	CityCode       string       `json:"city_code"`
	CityName       string       `json:"city_name"`
	HotelLatitude  string       `json:"hotel_latitude"`
	HotelLongitude string       `json:"hotel_longitude"`
	CountryCode    string       `json:"country_code"`
	CountryName    string       `json:"country_name"`
	RatingCode     string       `json:"rating_code"`
	RatingName     string       `json:"rating_name"`
	StarsCode      interface{}  `json:"stars_code"` //Иногда тут пустая строка
	StarsName      string       `json:"stars_name"`
	CurrencyCode   string       `json:"currency_code"`
	CurrencyName   string       `json:"currency_name"`
	Rooms          []HotelRooms `json:"rooms"`
}
type HotelRooms struct {
	RoomCode           int           `json:"room_code"`
	RoomName           string        `json:"room_name"`
	NumberOfGuests     string        `json:"number_of_guests"`
	Price              int           `json:"price"`
	Rackrate           interface{}   `json:"rackrate"` //Иногда тут null
	Comission          RoomComission `json:"comission"`
	PenaltySize        int           `json:"penalty_size"`
	DeadlineDate       string        `json:"deadline_date"`
	DeadlineTime       string        `json:"deadline_time"`
	PenaltyInfo        string        `json:"penalty_info"`
	MealTypeCode       string        `json:"meal_type_code"`
	MealTypeName       string        `json:"meal_type_name"`
	MealCategoryCode   string        `json:"meal_category_code"`
	MealCategoryName   string        `json:"meal_category_name"`
	MealName           string        `json:"meal_name"`
	MealPrice          int           `json:"meal_price"`
	MealIsIncludedCode int           `json:"meal_is_included_code"`
	MealIsIncludedName string        `json:"meal_is_included_name"`
	AvailabilityCode   int           `json:"availability_code"`
	AvailabilityName   string        `json:"availability_name"`
	PaymentTermsCode   string        `json:"payment_terms_code"`
	PaymentTermsName   string        `json:"payment_terms_name"`
	Periods            []RoomPeriod  `json:"periods"`
}
type RoomComission struct {
	Room          interface{} `json:"room"` //
	Meal          int         `json:"meal"`
	Total         interface{} `json:"total"` //
	TotalMealfree interface{} `json:"total_mealfree"`
}
type RoomPeriod struct {
	PeriodStart     string      `json:"period_start"`
	PeriodEnd       string      `json:"period_end"`
	PeriodDays      int         `json:"period_days"`
	PeriodSummRoom  int         `json:"period_summ_room"`
	PeriodSummMeal  interface{} `json:"period_summ_meal"`
	PeriodSummTotal int         `json:"period_summ_total"`
}

type CountryListAnswer struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Cities      string `json:"cities"`
}

type CityListAnswer struct {
	Country struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"country"`
	Cities []struct {
		CityCode      string `json:"city_code"`
		CityName      string `json:"city_name"`
		Region        string `json:"region"`
		Hotels        string `json:"hotels"`
		CityLatitude  string `json:"city_latitude"`
		CityLongitude string `json:"city_longitude"`
	} `json:"cities"`
}

type HotelListAnswer struct {
	HotelCode      string      `json:"hotel_code"`
	Vat            string      `json:"vat"`
	HotelName      string      `json:"hotel_name"`
	Address        string      `json:"address"`
	Description    string      `json:"description"`
	ImageUrl       string      `json:"image_url"`
	HotelLatitude  string      `json:"hotel_latitude"`
	HotelLongitude string      `json:"hotel_longitude"`
	RatingCode     interface{} `json:"rating_code"`
	RatingName     string      `json:"rating_name"`
	StarsCode      string      `json:"stars_code"`
	StarsName      string      `json:"stars_name"`
	HotelAmenities []struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"hotel_amenities"`
}

type CurrencyListAnswer struct {
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
}
