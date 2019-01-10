package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestRoomAmenitiesRequest_ok(t *testing.T) {
	testRequest("roomAmenitiesRequest_answOk.txt")
	data, err := zApi.HotelAmenitiesRequest(context.Background())

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 44)
}
