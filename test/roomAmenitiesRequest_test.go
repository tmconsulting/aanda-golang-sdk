package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestRoomAmenitiesRequest_ok(t *testing.T) {
	testRequest("roomAmenitiesRequest_answOk.txt")
	data, err := zApi.HotelAmenitiesRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 44)
}

