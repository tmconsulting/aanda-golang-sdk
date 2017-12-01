package test

import (
	"errors"
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestCountryListRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("countryListRequest_answOk.txt"))

	data, err := zApi.CountryListRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 5)
	st.Expect(t, data[0].CountryName, "Армения")
}

func TestCountryListRequest_err(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("countryListRequest_answErr.txt"))

	searchReq := aandaSdk.HotelSearchRequest{}
	_, err := zApi.HotelSearchRequest(searchReq)

	st.Expect(t, err, errors.New("Ошибка авторазиции"))
}
