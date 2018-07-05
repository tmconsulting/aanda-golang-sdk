package test

import (
	"errors"
	"testing"

	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-golang-sdk"
)

func TestHotelSearchRequest_ok(t *testing.T) {
	testRequest("hotelSearchRequest_answOk.txt")
	searchReq := aandaSdk.HotelSearchRequest{}
	data, err := zApi.HotelSearchRequest(searchReq)

	st.Expect(t, err, nil)
	st.Expect(t, data[0].CityName, "Санкт-Петербург")
}

func TestHotelSearchRequest_err(t *testing.T) {
	testRequest("hotelSearchRequest_answErr.txt")
	searchReq := aandaSdk.HotelSearchRequest{}
	_, err := zApi.HotelSearchRequest(searchReq)

	st.Expect(t, err, errors.New("Некорректная ArrivalDate[21.11.2017]"))
}
