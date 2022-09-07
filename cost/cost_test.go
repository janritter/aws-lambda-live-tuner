package cost

import (
	"log"
	"testing"
)

func TestCalculateCost(t *testing.T) {
	expected := 0.000021333376
	actual := Calculate(1280.0, 1024, "x86_64", "eu-central-1")

	log.Println(actual)

	if actual != expected {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}
