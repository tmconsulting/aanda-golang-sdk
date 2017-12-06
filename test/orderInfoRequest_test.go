package test

import (
	"github.com/nbio/st"
	"testing"
)

func TestOrderInfoRequest_ok(t *testing.T) {
	testRequest("orderInfoRequest_answOk.txt")
	data, err := zApi.OrderInfoRequest(2213397)

	st.Expect(t, err, nil)
	st.Expect(t, data.OrderId, "2213397")
}
