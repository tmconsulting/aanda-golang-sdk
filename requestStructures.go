package aandaSdk

type Auth struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

type HotelSearchRequest struct {
	Auth
	City           string `json:"City"`
	Lat            string `json:"Lat"`
	Lng            string `json:"Lng"`
	Radius         string `json:"Radius"`
	ArrivalDate    string `json:"ArrivalDate"`
	DepartureDate  string `json:"DepartureDate"`
	PriceFrom      string `json:"PriceFrom"`
	PriceTo        string `json:"PriceTo"`
	NumberOfGuests string `json:"NumberOfGuests"`
}

type HotelPricingRequest struct {
	Auth
	Hotel          string `json:"Hotel"`
	ProductCode    string `json:"ProductCode"`
	Currency       string `json:"Currency"`
	WhereToPay     string `json:"WhereToPay"`
	ArrivalDate    string `json:"ArrivalDate"`
	DepartureDate  string `json:"DepartureDate"`
	ArrivalTime    string `json:"ArrivalTime"`
	DepartureTime  string `json:"DepartureTime"`
	NumberOfGuests string `json:"NumberOfGuests"`
}

type OrderRequest struct {
	Auth
	OrderId        string   `json:"order_id"`
	ArrivalDate    string   `json:"arrival_date"`
	DepartureDate  string   `json:"departure_date"`
	AddInfo        string   `json:"add_info"`
	ChangeCode     string   `json:"change_code"`
	CancelCode     string   `json:"cancel_code"`
	HotelCode      string   `json:"hotel_code"`
	RoomCode       string   `json:"room_code"`
	Meal           string   `json:"Meal"`
	ArrivalTime    string   `json:"arrival_time"`
	DepartureTime  string   `json:"departure_time"`
	NumberOfGuests string   `json:"NumberOfGuests"`
	Person         []Person `json:"Person"`
}

type Person struct {
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type OrderInfoRequest struct {
	Auth
	Id string `json:"order_id"`
}

type SendOrderMessageRequest struct {
	OrderId int
	Message string
}

type OrderListRequest struct {
	BuyerId              string `json:"BuyerId"`
	UserId               string `json:"UserId"`
	Password             string `json:"Password"`
	Language             string `json:"Language"`
	LastName             string `json:"LastName"` //Доп параметры
	ArrivalDateFrom      string `json:"ArrivalDateFrom"`
	ArrivalDateTo        string `json:"ArrivalDateTo"`
	DepartureDateFrom    string `json:"DepartureDateFrom"`
	DepartureDateTo      string `json:"DepartureDateTo"`
	RegistrationDateFrom string `json:"RegistrationDateFrom"`
	RegistrationDateTo   string `json:"RegistrationDateTo"`
	ChangeDateFrom       string `json:"ChangeDateFrom"`
	ChangeDateTo         string `json:"ChangeDateTo"`
}
