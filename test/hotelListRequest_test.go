package test

import (
	"github.com/nbio/st"
	"testing"
)

func TestHotelListRequest_ok(t *testing.T) {
	testRequest("hotelListRequest_answOk.txt")
	data, err := zApi.HotelListRequest(1)

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 296)
	st.Expect(t, data[0].HotelCode, "12360")
	st.Expect(t, data[0].HotelName, "А-Хостел")
}
