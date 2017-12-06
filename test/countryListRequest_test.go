package test

import (
	"errors"
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"testing"
)

func TestCountryListRequest_ok(t *testing.T) {
	testRequest("countryListRequest_answOk.txt")
	data, err := zApi.CountryListRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 5)
	st.Expect(t, data[0].CountryName, "Армения")
}

func TestCountryListRequest_err(t *testing.T) {
	testRequest("countryListRequest_answErr.txt")
	searchReq := aandaSdk.HotelSearchRequest{}
	_, err := zApi.HotelSearchRequest(searchReq)

	st.Expect(t, err, errors.New("Ошибка авторазиции"))
}
