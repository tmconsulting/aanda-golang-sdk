package zabroniryiru_sdk

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