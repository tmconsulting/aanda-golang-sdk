package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOrderMessagesRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderMessagesRequest_answOk.txt"))

	data, err := zApi.OrderMessagesRequest(2213397)

	st.Expect(t, err, nil)
	st.Expect(t, len(data) > 0, true)
	st.Expect(t, data[0].OrderCode, "2213397")
}
