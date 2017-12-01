package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestCityListRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("cityListRequest_answOk.txt"))

	data, err := zApi.CityListRequest(9)

	st.Expect(t, err, nil)
	st.Expect(t, data.Country.Id, "9")
	st.Expect(t, data.Country.Name, "Россия")
	st.Expect(t, data.Cities[0].CityCode, "1702")
	st.Expect(t, data.Cities[0].CityName, "Абаза")
}
