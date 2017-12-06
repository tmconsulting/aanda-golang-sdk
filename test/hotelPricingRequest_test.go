package test

import (
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"testing"
)

func TestHotelPricingRequest_ok(t *testing.T) {
	testRequest("hotelPricingRequest_answOk.txt")
	priceReq := aandaSdk.HotelPricingRequest{}
	data, err := zApi.HotelPricingRequest(priceReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.HotelCode, "2150")
	st.Expect(t, data.HotelName, "Ярославская")
}
