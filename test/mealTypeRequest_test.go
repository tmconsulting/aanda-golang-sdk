package test

import (
	"github.com/nbio/st"
	"testing"
)

func TestMealTypeRequest_ok(t *testing.T) {
	testRequest("mealTypeRequest_answOk.txt")
	data, err := zApi.MealTypeRequest()

	st.Expect(t, err, nil)
	st.Expect(t, len(data), 6)
	st.Expect(t, data[0].MealTypeCode, "1")
}
