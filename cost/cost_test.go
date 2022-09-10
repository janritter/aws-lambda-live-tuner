package cost

import (
	"log"
	"testing"
)

func TestCalculateCostX86(t *testing.T) {
	expected := 0.000021333376
	actual := Calculate(1280.0, 1024, "x86_64", "eu-central-1")

	log.Println(actual)

	if actual != expected {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

func TestCalculateCostARM(t *testing.T) {
	expected := 0.0000133334
	actual := Calculate(1000.0, 1024, "arm64", "eu-west-1")

	log.Println(actual)

	if actual != expected {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

func TestCalculateCostInvalidRegion(t *testing.T) {
	expected := 0.0
	actual := Calculate(1000.0, 1024, "arm64", "unkown")

	log.Println(actual)

	if actual != expected {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}
