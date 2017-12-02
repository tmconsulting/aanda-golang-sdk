package aandaSdk

type Auth struct {
	BuyerId  string
	UserId   string
	Password string
	Language string
}

type HotelSearchRequest struct {
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
}

type HotelPricingRequest struct {
	BuyerId        string `json:"BuyerId"`
	UserId         string `json:"UserId"`
	Password       string `json:"Password"`
	Language       string `json:"Language"`
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
