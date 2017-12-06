package test

import (
	"github.com/nbio/st"
	"testing"
)

func TestOrderMessagesRequest_ok(t *testing.T) {
	testRequest("orderMessagesRequest_answOk.txt")
	data, err := zApi.OrderMessagesRequest(2213397)

	st.Expect(t, err, nil)
	st.Expect(t, len(data) > 0, true)
	st.Expect(t, data[0].OrderCode, "2213397")
}
