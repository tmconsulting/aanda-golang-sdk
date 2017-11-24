package test

import "testing"

func TestExample(t *testing.T) {
	actualResult := "Hello"
	var expectedResult = "Hello"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}