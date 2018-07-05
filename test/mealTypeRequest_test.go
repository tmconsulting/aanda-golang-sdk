package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestMealTypeRequest_ok(t *testing.T) {
	testRequest("mealTypeRequest_answOk.txt")
	data, err := zApi.MealTypeRequest()

	st.Expect(t, err, nil)
	st.Expect(t, data[0].MealTypeCode, "1")
}
