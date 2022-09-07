package cmd

import (
	"testing"
)

func TestMemorySortedList(t *testing.T) {
	expected := []int{128, 256, 384, 512}
	actual := memorySortedList(map[int]float64{
		384: 0.0,
		128: 0.0,
		512: 0.0,
		256: 0.0,
	})

	if len(actual) != len(expected) {
		t.Errorf("expected %d elements, got %d", len(expected), len(actual))
	}
	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("expected element %d, got %d", expected[i], actual[i])
		}
	}
}
