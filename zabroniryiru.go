package zabroniryiru

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiUrl = "http://api.aanda.ru/xml_gateway/"

func main() {

}
func HotelSearchRequest() /*[]HotelSearchAnswer */ {
	json := "{'BuyerId':'TMCon','UserId':'tmcon','Password':'vcxq11cz!','Language':'en','city_code':'142','Lat':'','Lng':'','arrival_date':'07.06.2017','departure_date':'08.06.2017','PriceFrom':'4000','PriceTo':'5000','number_of_guests':'1'}"
	//xml := "<HotelSearchRequest BuyerId=\"TMCon\" UserId=\"tmcon\" Password=\"vcxq11cz!\" Language=\"ru\" City=\"1\" ArrivalDate=\"20.05.2014\" DepartureDate=\"21.05.2014\" PriceFrom=\"4000\" PriceTo=\"5000\"NumberOfGuests=\"1\" />"
	data := url.Values{}
	//data.Set("RequestType", "json")
	data.Add("RequestName", "HotelSearchRequest")
	data.Add("JSON", json)
	/*data.Add("BuyerId", "TMCon")
	data.Add("UserId", "tmcon")
	data.Add("Password", "vcxq11cz!")

	data.Add("city_code", "142")
	data.Add("arrival_date", "07.06.2017")
	data.Add("departure_date", "26.01.2018")
	data.Add("PriceFrom", "0")
	data.Add("PriceTo", "20")
	data.Add("number_of_guests", "15")*/

	//

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	/*re := regexp.MustCompile(`(?U)\[(.+)\]`)
	jsonRaw := re.FindStringSubmatch(string(body))

	jsonData := []HotelSearchAnswer{}
	err = json.Unmarshal([]byte(jsonRaw[0]), &jsonData)
	if err != nil {
		panic(err)
	}

	return jsonData*/
}
