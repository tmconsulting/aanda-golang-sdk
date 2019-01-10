package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestHotelDescriptionRequest_ok(t *testing.T) {
	testRequest("hotelDescriptionRequest_answOk.txt")
	data, err := zApi.HotelDescriptionRequest(context.Background(), 2150)

	st.Expect(t, err, nil)
	st.Expect(t, data.HotelName, "Ярославская")
}
