package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestCityListRequest_ok(t *testing.T) {
	testRequest("cityListRequest_answOk.txt")
	data, err := zApi.CityListRequest(9)

	st.Expect(t, err, nil)
	st.Expect(t, data.Country.Id, "9")
	st.Expect(t, data.Country.Name, "Россия")
	st.Expect(t, data.Cities[0].CityCode, "1702")
	st.Expect(t, data.Cities[0].CityName, "Абаза")
}
