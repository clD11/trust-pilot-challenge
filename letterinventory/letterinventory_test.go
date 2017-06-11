package letterinventory

import "testing"

func TestCreateLetterInventory(t *testing.T) {
	letterInventory := CreateLetterInventory("aabcd")

	expected := 5
	actual := letterInventory.GetSize()

	if actual != expected {
		t.Error("Test failed: expected ", expected, "actual", actual)
	}
}

func TestSubtract(t *testing.T) {
	letterInventory := CreateLetterInventory("aabcd")
	letterInventory.Subtract("aabcd")

	expected := 0
	actual := letterInventory.GetSize()

	if actual != expected {
		t.Error("Test failed: expected ", expected, "actual", actual)
	}
}

func TestContainsTrue(t *testing.T) {
	letterInventory := CreateLetterInventory("aabcd")

	expected := true
	actual := letterInventory.Contains("aabcd")

	if actual != expected {
		t.Error("Test failed: expected ", expected, "actual", actual)
	}
}

func TestContainsFalse(t *testing.T) {
	letterInventory := CreateLetterInventory("aabcd")

	expected := false
	actual := letterInventory.Contains("aabcdee")

	if actual != expected {
		t.Error("Test failed: expected ", expected, "actual", actual)
	}
}
