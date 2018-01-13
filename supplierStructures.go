package aandaSdk

type HotelSearchResponse struct {
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

type CountryListResponse struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Cities      string `json:"cities"`
}

type CityListResponse struct {
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

type HotelListResponse struct {
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

type HotelDescriptionResponse struct {
	HotelCode       string `json:"hotel_code"`
	HotelName       string `json:"hotel_name"`
	Vat             string `json:"vat"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	Description     string `json:"description"`
	FullDescription string `json:"full_description"`
	FullAddress     struct {
		Zip            string `json:"zip"`
		Region         string `json:"region"`
		City           string `json:"city"`
		Addinfo        string `json:"addinfo"`
		Strtype        string `json:"strtype"`
		Strname        string `json:"strname"`
		Proptype       string `json:"proptype"`
		Propnumber     string `json:"propnumber"`
		Buildingattr   string `json:"buildingattr"`
		Buildingnumber string `json:"buildingnumber"`
		Addinfo2       string `json:"addinfo2"`
		CityCode       string `json:"city_code"`
		CityName       string `json:"city_name"`
		CityLatitude   string `json:"city_latitude"`
		CityLongitude  string `json:"city_longitude"`
		HotelLatitude  string `json:"hotel_latitude"`
		HotelLongitude string `json:"hotel_longitude"`
		CountryCode    string `json:"country_code"`
		CountryName    string `json:"country_name"`
	} `json:"full_address"`
	RatingCode string `json:"rating_code"`
	RatingName string `json:"rating_name"`
	StarsCode  string `json:"stars_code"`
	StarsName  string `json:"stars_name"`
	Images     []struct {
		URL  string `json:"Url"`
		Desc string `json:"desc"`
	} `json:"images"`
	CurrencyCode   string `json:"currency_code"`
	CurrencyName   string `json:"currency_name"`
	HotelAmenities []struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"hotel_amenities"`
	Rooms []struct {
		RoomCode        string `json:"room_code"`
		RoomName        string `json:"room_name"`
		NumberOfGuests  string `json:"number_of_guests"`
		RoomDescription string `json:"room_description"`
		Images          string `json:"images"`
		RoomAmenities   []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"room_amenities"`
	} `json:"rooms"`
	Conference []interface{} `json:"conference"`
	Group      struct {
		Qty      string `json:"qty"`
		Type     string `json:"type"`
		Typename string `json:"typename"`
		Note     string `json:"note"`
		Rule     string `json:"rule"`
	} `json:"group"`
}

type CurrencyListResponse struct {
	CurrencyCode string `json:"currency_code"`
	CurrencyName string `json:"currency_name"`
}

type MealTypeResponse struct {
	MealTypeCode string `json:"meal_type_code"`
	MealTypeName string `json:"meal_type_name"`
}

type MealCategoryResponse struct {
	MealCategoryCode string `json:"meal_category_code"`
	MealCategoryName string `json:"meal_category_name"`
}

type HotelPricingResponse struct {
	HotelCode          string `json:"hotel_code"`
	HotelName          string `json:"hotel_name"`
	NumberOfGuests     int    `json:"number_of_guests"`
	CheckInTime        string `json:"check-in_time"`
	CheckOutTime       string `json:"check-out_time"`
	ArrivalDate        string `json:"arrival_date"`
	DepartureDate      string `json:"departure_date"`
	Vat                int    `json:"vat"`
	TimeZone           string `json:"time_zone"`
	CountryCode        string `json:"country_code"`
	CountryName        string `json:"country_name"`
	CityCode           string `json:"city_code"`
	CityName           string `json:"city_name"`
	CityLatitude       string `json:"city_latitude"`
	CityLongitude      string `json:"city_longitude"`
	HotelLatitude      string `json:"hotel_latitude"`
	HotelLongitude     string `json:"hotel_longitude"`
	AllowEarlyCheckIn  string `json:"allow_early_check-in"`
	AllowEarlyCheckOut string `json:"allow_early_check-out"`
	CurrencyCode       string `json:"currency_code"`
	CurrencyName       string `json:"currency_name"`
	Rooms              []struct {
		RoomCode            int           `json:"room_code"`
		RoomName            string        `json:"room_name"`
		RateName            string        `json:"rate_name"`
		AllowEarlierCheckin bool          `json:"allow_earlier_checkin"`
		AllowLateCheckout   bool          `json:"allow_late_checkout"`
		Checkins            []interface{} `json:"checkins"`
		Checkouts           []interface{} `json:"checkouts"`
		Price               int           `json:"price"`
		Rackrate            int           `json:"rackrate"`
		Comission           struct {
			Room          int `json:"room"`
			Meal          int `json:"meal"`
			Total         int `json:"total"`
			Checkin       int `json:"checkin"`
			Checkout      int `json:"checkout"`
			TotalMealfree int `json:"total_mealfree"`
		} `json:"comission"`
		PenaltySize struct {
			Room  int `json:"room"`
			Total int `json:"total"`
		} `json:"penalty_size"`
		DeadlineDate       string `json:"deadline_date"`
		DeadlineTime       string `json:"deadline_time"`
		PenaltyInfo        string `json:"penalty_info"`
		MaxGuests          string `json:"max_guests"`
		MealTypeCode       string `json:"meal_type_code"`
		MealTypeName       string `json:"meal_type_name"`
		MealCategoryCode   string `json:"meal_category_code"`
		MealCategoryName   string `json:"meal_category_name"`
		MealName           string `json:"meal_name"`
		MealPrice          int    `json:"meal_price"`
		MealIsIncludedCode int    `json:"meal_is_included_code"`
		PaymentTermsCode   string `json:"payment_terms_code"`
		PaymentTermsName   string `json:"payment_terms_name"`
		AvailabilityCode   int    `json:"availability_code"`
		AvailabilityName   string `json:"availability_name"`
		RoomsAvailable     int    `json:"rooms_available"`
		RoomAmenities      []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"room_amenities"`
	} `json:"rooms"`
	Group struct {
		Qty      string `json:"qty"`
		Type     string `json:"type"`
		Typename string `json:"typename"`
		Note     string `json:"note"`
		Rule     string `json:"rule"`
	} `json:"group"`
	HotelAmenities []struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"hotel_amenities"`
}

type OrderRequestResponse struct {
	Status  string  `json:"Status" validate:"required"`
	OrderId string  `json:"order_id" validate:"required"`
	Time    float64 `json:"Time" validate:"required"`
}

type OrderInfoResponse struct {
	OrderId           string `json:"order_id"`
	ReferenceNumber   string `json:"reference_number"`
	Created           string `json:"created"`
	DeadlineDate      string `json:"deadline_date"`
	DeadlineTime      string `json:"deadline_time"`
	TotalPrice        int    `json:"total_price"`
	Comission         string `json:"comission"`
	ArrivalDate       string `json:"arrival_date"`
	DepartureDate     string `json:"departure_date"`
	CurrencyCode      string `json:"currency_code"`
	CurrencyName      string `json:"currency_name"`
	PaymentTermsCode  string `json:"payment_terms_code"`
	PaymentTermsName  string `json:"payment_terms_name"`
	StatusCode        string `json:"status_code"`
	StatusName        string `json:"status_name"`
	ContactpersonName string `json:"contactperson_name"`
	PersonPhone       string `json:"person_phone"`
	PersonFax         string `json:"person_fax"`
	PersonEmail       string `json:"person_email"`
	Rooms             []struct {
		RoomCode           string `json:"room_code"`
		ArrivalDate        string `json:"arrival_date"`
		ArrivalTime        string `json:"arrival_time"`
		DepartureDate      string `json:"departure_date"`
		DepartureTime      string `json:"departure_time"`
		NumberOfNights     string `json:"number_of_nights"`
		NumberOfRooms      string `json:"number_of_rooms"`
		NumberOfGuests     string `json:"number_of_guests"`
		AdditionalInfo     string `json:"additional_info"`
		SupplierInfo       string `json:"supplier_info"`
		ConfirmationNumber string `json:"confirmation_number"`
		Price              int    `json:"price"`
		RoomPrice          string `json:"room_price"`
		Comission          int    `json:"comission"`
		Penalty            string `json:"penalty"`
		PenaltyNote        string `json:"penalty_note"`
		DeadlineDate       string `json:"deadline_date"`
		PossiblePenalty    string `json:"possible_penalty"`
		CancelledCode      string `json:"cancelled_code"`
		CancelledName      string `json:"cancelled_name"`
		ChangeCode         string `json:"change_code"`
		ChangeName         string `json:"change_name"`
		HotelCode          string `json:"hotel_code"`
		Vat                string `json:"vat"`
		TimeZone           string `json:"time_zone"`
		CountryCode        string `json:"country_code"`
		CountryName        string `json:"country_name"`
		CityCode           string `json:"city_code"`
		CityName           string `json:"city_name"`
		StatusCode         string `json:"status_code"`
		StatusName         string `json:"status_name"`
		AllowCancelCode    int    `json:"allow_cancel_code"`
		AllowCancelName    string `json:"allow_cancel_name"`
		AllowChangeCode    int    `json:"allow_change_code"`
		AllowChangeName    string `json:"allow_change_name"`
		RoomName           string `json:"room_name"`
		MealCode           int    `json:"meal_code"`
		MealName           string `json:"meal_name"`
		Persons            []struct {
			Lastname  string `json:"lastname"`
			Firstname string `json:"firstname"`
		} `json:"persons"`
	} `json:"rooms"`
	ServiceList []struct {
		Id          string      `json:"id"`
		Name        string      `json:"name"`
		Price       string      `json:"price"`
		ServiceType string      `json:"service_type"`
		ServiceName string      `json:"service_name"`
		StartDate   string      `json:"start_date"`
		EndDate     string      `json:"end_date"`
		InvoiceId   interface{} `json:"invoice_id"`
		PersonId    string      `json:"person_id"`
	} `json:"service_list"`
	GroupInfo []interface{} `json:"group_info"`
}

type OrderListResponse struct {
	OrderId           string `json:"order_id"`
	ReferenceNumber   string `json:"reference_number"`
	Created           string `json:"created"`
	DeadlineDate      string `json:"deadline_date"`
	TotalPrice        string `json:"total_price"`
	Comission         int    `json:"comission"`
	Penalty           string `json:"penalty"`
	ArrivalDate       string `json:"arrival_date"`
	DepartureDate     string `json:"departure_date"`
	CurrencyCode      string `json:"currency_code"`
	CurrencyName      string `json:"currency_name"`
	StatusCode        string `json:"status_code"`
	StatusName        string `json:"status_name"`
	ContactpersonName string `json:"contactperson_name"`
	PersonPhone       string `json:"person_phone"`
	PersonFax         string `json:"person_fax"`
	PersonEmail       string `json:"person_email"`
}

type OrderMessagesResponse struct {
	OrderCode   string `json:"order_code"`
	MessageCode string `json:"message_code"`
	Created     string `json:"created"`
	From        string `json:"from"`
	Message     string `json:"message"`
}

type SendOrderMessageResponse struct {
	MessageCode string `json:"message_code"`
}

type ClientStatusResponse struct {
	ClientStatusCode        string `json:"client_status_code"`
	ClientStatusName        string `json:"client_status_name"`
	ClientStatusExplanation string `json:"client_status_explanation"`
}
