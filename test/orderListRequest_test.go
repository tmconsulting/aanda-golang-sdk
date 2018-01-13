package test

import (
	"github.com/nbio/st"
	"github.com/tmconsulting/aanda-golang-sdk"
	"testing"
)

func TestOrderListRequest_ok(t *testing.T) {
	testRequest("orderListRequest_answOk.txt")
	searchReq := aandaSdk.OrderListRequest{}
	data, err := zApi.OrderListRequest(searchReq)

	st.Expect(t, err, nil)
	st.Expect(t, data[0].OrderId, "2216689")
}
