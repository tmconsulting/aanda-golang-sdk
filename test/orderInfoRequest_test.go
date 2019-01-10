package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestOrderInfoRequest_ok(t *testing.T) {
	testRequest("orderInfoRequest_answOk.txt")
	data, err := zApi.OrderInfoRequest(context.Background(), 2213397)

	st.Expect(t, err, nil)
	st.Expect(t, data.OrderId, "2213397")
}
