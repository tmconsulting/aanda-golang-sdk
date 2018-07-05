package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestHotelAmenitiesRequest_ok(t *testing.T) {
	testRequest("hotelAmenitiesRequest_answOk.txt")
	data, err := zApi.HotelAmenitiesRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 22)
}
