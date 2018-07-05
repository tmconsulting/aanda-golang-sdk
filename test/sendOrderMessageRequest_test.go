package test

import (
	"testing"

	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-golang-sdk"
)

func TestSendOrderMessagesRequest_ok(t *testing.T) {
	testRequest("sendOrderMessageRequest_answOk.txt")
	somReq := aandaSdk.SendOrderMessageRequest{}
	data, err := zApi.SendOrderMessageRequest(somReq)

	st.Expect(t, err, nil)
	st.Expect(t, data.MessageCode, "1269154")
}
