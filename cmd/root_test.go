package cmd

import (
	"log"
	"testing"
)

func TestCalculateCost(t *testing.T) {
	expected := 0.000021333376
	actual := calculateCost(1280.0, 1024)

	log.Println(actual)

	if actual != expected {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

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
