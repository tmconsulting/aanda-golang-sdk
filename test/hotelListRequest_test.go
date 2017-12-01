package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestHotelListRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("hotelListRequest_answOk.txt"))

	data, err := zApi.HotelListRequest(1)

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 296)
	st.Expect(t, data[0].HotelCode, "12360")
	st.Expect(t, data[0].HotelName, "А-Хостел")
}
