package test

import "testing"

func TestReverseToReturnReversedInputString(t *testing.T) {
	actualResult := "Hello"
	var expectedResult = "Hello1"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}