package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestMealTypeRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("mealTypeRequest_answOk.txt"))

	data, err := zApi.MealTypeRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 6)
	st.Expect(t, data[0].MealTypeCode, "1")
}
