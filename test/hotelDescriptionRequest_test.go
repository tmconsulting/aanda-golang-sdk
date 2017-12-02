package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestHotelDescriptionRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("hotelDescriptionRequest_answOk.txt"))

	data, err := zApi.HotelDescriptionRequest(2150)

	st.Expect(t, err, nil)
	st.Expect(t, data.HotelName, "Ярославская")
}
