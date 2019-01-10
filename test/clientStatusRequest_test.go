package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestClientStatusRequest_ok(t *testing.T) {
	testRequest("clientStatusRequest_answOk.txt")
	data, err := zApi.ClientStatusRequest(context.Background())

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 10)
}
