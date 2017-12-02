package test

import (
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-sdk"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestSendOrderMessagesRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("sendOrderMessageRequest_answOk.txt"))

	somReq := aandaSdk.SendOrderMessageRequest{}
	data, err := zApi.SendOrderMessageRequest(somReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.MessageCode, "1269154")
}
