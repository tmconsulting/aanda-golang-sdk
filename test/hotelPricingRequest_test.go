package test

import (
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestHotelPricingRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("hotelPricingRequest_answOk.txt"))

	priceReq := aandaSdk.HotelPricingRequest{}
	data, err := zApi.HotelPricingRequest(priceReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.HotelCode, "2150")
	st.Expect(t, data.HotelName, "Ярославская")
}
