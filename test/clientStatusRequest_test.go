package test

import (
	"testing"
	"github.com/nbio/st"
)

func TestClientStatusRequest_ok(t *testing.T) {
	testRequest("clientStatusRequest_answOk.txt")
	data, err := zApi.ClientStatusRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 10)
}

