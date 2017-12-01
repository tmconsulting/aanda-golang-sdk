package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestCurrencyListRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("currencyListRequest_answOk.txt"))

	data, err := zApi.CurrencyListRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 4)
	st.Expect(t, data[0].CurrencyName, "RUR")
}
