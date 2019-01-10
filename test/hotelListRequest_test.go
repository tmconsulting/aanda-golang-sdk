package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestHotelListRequest_ok(t *testing.T) {
	testRequest("hotelListRequest_answOk.txt")
	data, err := zApi.HotelListRequest(context.Background(), 1)

	st.Expect(t, err, nil)
	st.Expect(t, data[0].HotelCode, "12360")
	st.Expect(t, data[0].HotelName, "А-Хостел")
}
