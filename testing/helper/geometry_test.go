package helper

import (
	"testing"
)

func assertEqualWithHelper(t *testing.T, a, b int) {
	t.Helper()

	// Output that the error happens in this file and line number 24 (the calling function).
	if a != b {
		t.Errorf("%d != %d", a, b)
	}
}

func assertWithoutHelper(t *testing.T, a, b int) {
	// Output that the error happens in this file and line number 19 (inside this function).
	if a != b {
		t.Errorf("%d != %d", a, b)
	}
}

func TestHelper(t *testing.T) {
	assertEqualWithHelper(t, 6, RectArea(2, 3))
	assertWithoutHelper(t, 6, RectArea(2, 3))
}
