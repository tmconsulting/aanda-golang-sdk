package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestOrderMessagesRequest_ok(t *testing.T) {
	testRequest("orderMessagesRequest_answOk.txt")
	data, err := zApi.OrderMessagesRequest(context.Background(), "2213397")

	st.Expect(t, err, nil)
	st.Expect(t, len(data) > 0, true)
	st.Expect(t, data[0].OrderCode, "2213397")
}
