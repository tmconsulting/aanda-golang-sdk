package test

import (
	"errors"
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestHotelSearchRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("hotelSearchRequest_answOk.txt"))

	searchReq := aandaSdk.HotelSearchRequest{}
	data, err := zApi.HotelSearchRequest(searchReq)

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 108)
	st.Expect(t, data[0].CityName, "Санкт-Петербург")
}

func TestHotelSearchRequest_err(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("hotelSearchRequest_answErr.txt"))

	searchReq := aandaSdk.HotelSearchRequest{}
	_, err := zApi.HotelSearchRequest(searchReq)

	st.Expect(t, err, errors.New("Некорректная ArrivalDate[21.11.2017]"))
}
