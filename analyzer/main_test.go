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
