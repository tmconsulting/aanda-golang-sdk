package test

import (
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestOrderListRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("orderListRequest_answOk.txt"))

	searchReq := aandaSdk.OrderListRequest{}
	data, err := zApi.OrderListRequest(searchReq)

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 17)
	st.Expect(t, data[0].OrderId, "2216689")
}
