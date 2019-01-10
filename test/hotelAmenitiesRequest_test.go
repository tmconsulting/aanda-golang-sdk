package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestHotelAmenitiesRequest_ok(t *testing.T) {
	testRequest("hotelAmenitiesRequest_answOk.txt")
	data, err := zApi.HotelAmenitiesRequest(context.Background())

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 22)
}
