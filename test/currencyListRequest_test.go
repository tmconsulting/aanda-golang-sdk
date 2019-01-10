package test

import (
	"context"
	"testing"

	"github.com/nbio/st"
)

func TestCurrencyListRequest_ok(t *testing.T) {
	testRequest("currencyListRequest_answOk.txt")
	data, err := zApi.CurrencyListRequest(context.Background())

	st.Expect(t, err, nil)
	st.Expect(t, data[0].CurrencyName, "RUR")
}
