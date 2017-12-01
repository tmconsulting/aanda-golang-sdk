package test

import (
	"github.com/nbio/st"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

func TestMealCategoryRequest_ok(t *testing.T) {
	defer gock.Off()
	gock.New("http://api.aanda.ru").
		Post("/xml_gateway/").
		Reply(200).
		JSON(getJson("mealCategoryRequest_answOk.txt"))

	data, err := zApi.MealCategoryRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 18)
	st.Expect(t, data[0].MealCategoryCode, "1")
}
