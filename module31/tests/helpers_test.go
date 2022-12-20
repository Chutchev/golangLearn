package tests

import (
	"module31/pkg/helpers"
	"strconv"
	"testing"
)

func TestRemoveElementsFromArray(t *testing.T) {
	array := []string{"1", "2", "3", "4", "5"}
	expectedArray := make([]string, 4)
	for i := 0; i < len(array)-1; i++ {
		expectedArray[i] = strconv.Itoa(i + 1)
	}
	expectedArray[3] = "5"
	resultArray := helpers.RemoveElementFromArray(array, "4")
	if len(resultArray) != len(expectedArray) {
		t.Log("Длина не совпадает: ", len(resultArray), "!=", len(expectedArray))
		t.Fatal("Длина не совпадает: ", len(resultArray), "!=", len(expectedArray))
	}
	for i := 0; i < len(resultArray); i++ {
		if resultArray[i] != expectedArray[i] {
			t.Logf("Элемент %v != %v", resultArray[i], expectedArray[i])
			t.Fatalf("Элемент %v != %v", resultArray[i], expectedArray[i])
		}
	}
}
