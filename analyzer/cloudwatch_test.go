package analyzer

import "testing"

func TestGetFunctionNameFromARN(t *testing.T) {
	arn := "arn:aws:lambda:us-east-1:123456789012:function:my-function"
	expected := "my-function"
	actual := getFunctionNameFromARN(arn)
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestGetDurationFromMessage(t *testing.T) {
	message := "REPORT RequestId: 7239847892374-7238947823947823-234234324\tDuration: 189.12 ms\tBilled Duration: 190 ms\tMemory Size: 256 MB\tMax Memory Used: 67 MB\t\n"
	expected := 189.12
	actual, err := getDurationFromMessage(message)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	if actual != expected {
		t.Errorf("expected %f, got %f", expected, actual)
	}
}

func TestGetRequestIdFromMessage(t *testing.T) {
	message := "REPORT RequestId: 7239847892374-7238947823947823-234234324\tDuration: 189.12 ms\tBilled Duration: 190 ms\tMemory Size: 256 MB\tMax Memory Used: 67 MB\t\n"
	expected := "7239847892374-7238947823947823-234234324"
	actual, err := getRequestIdFromMessage(message)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
