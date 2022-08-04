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
