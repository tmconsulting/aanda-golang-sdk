package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-golang-sdk"
)

func TestHotelPricingRequest_ok(t *testing.T) {
	testRequest("hotelPricingRequest_answOk.txt")
	priceReq := aandaSdk.HotelPricingRequest{}
	data, err := zApi.HotelPricingRequest(context.Background(), priceReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.HotelCode, "2150")
	st.Expect(t, data.HotelName, "Ярославская")
}
