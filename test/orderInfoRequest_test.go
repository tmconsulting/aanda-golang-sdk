package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOrderInfoRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderInfoRequest_answOk.txt"))

	data, err := zApi.OrderInfoRequest(2213397)

	st.Expect(t, err, nil)
	st.Expect(t, data.OrderId, "2213397")
}
