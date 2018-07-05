package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestMealCategoryRequest_ok(t *testing.T) {
	testRequest("mealCategoryRequest_answOk.txt")
	data, err := zApi.MealCategoryRequest()

	st.Expect(t, err, nil)
	st.Expect(t, data[0].MealCategoryCode, "1")
}
