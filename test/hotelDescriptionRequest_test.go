package test

import (
	"github.com/nbio/st"
	"testing"
)

func TestHotelDescriptionRequest_ok(t *testing.T) {
	testRequest("hotelDescriptionRequest_answOk.txt")
	data, err := zApi.HotelDescriptionRequest(2150)

	st.Expect(t, err, nil)
	st.Expect(t, data.HotelName, "Ярославская")
}
