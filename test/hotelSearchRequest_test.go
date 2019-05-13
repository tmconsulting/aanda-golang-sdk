package test

import (
	"context"
	"errors"
	"testing"

	"github.com/nbio/st"
	aandaSdk "github.com/tmconsulting/aanda-golang-sdk"
)

func TestHotelSearchRequest_ok(t *testing.T) {
	testRequest("hotelSearchRequest_answOk.txt")
	searchReq := aandaSdk.HotelSearchRequest{}
	data, err := zApi.HotelSearchRequest(context.Background(), searchReq)

	st.Expect(t, err, nil)
	st.Expect(t, data[0].CityName, "Санкт-Петербург")
}

func TestHotelSearchRequest_err(t *testing.T) {
	testRequest("hotelSearchRequest_answErr.txt")
	searchReq := aandaSdk.HotelSearchRequest{}
	_, err := zApi.HotelSearchRequest(context.Background(), searchReq)

	st.Expect(t, err.Error(), errors.New("Некорректная ArrivalDate[21.11.2017]").Error())
}
