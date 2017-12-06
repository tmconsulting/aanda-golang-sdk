package test

import (
	"github.com/nbio/st"
	"testing"
)

func TestCurrencyListRequest_ok(t *testing.T) {
	testRequest("currencyListRequest_answOk.txt")
	data, err := zApi.CurrencyListRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 4)
	st.Expect(t, data[0].CurrencyName, "RUR")
}
