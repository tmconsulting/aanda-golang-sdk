package zabroniryiru

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiUrl = "http://api.aanda.ru/xml_gateway/"

func main() {

}
func HotelSearchRequest() []HotelSearchAnswer {
	jsonReq := `{"BuyerId":"TMCon","UserId":"tmcon","Password":"vcxq11cz!","Language":"ru","City":"2","Lat":"","Lng":"","ArrivalDate":"30.10.2017","DepartureDate":"31.10.2017","PriceFrom":"4000","PriceTo":"5000","NumberOfGuests":"1"}`
	data := url.Values{}
	data.Set("RequestType", "json")
	data.Add("RequestName", "HotelSearchRequest")
	data.Add("JSON", jsonReq)

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	jsonData := []HotelSearchAnswer{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		panic(err)
	}

	return jsonData
}
