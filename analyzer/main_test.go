package analyzer

import (
	"testing"

	"github.com/janritter/aws-lambda-live-tuner/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetFunctionNameFromARN(t *testing.T) {
	arn := "arn:aws:lambda:us-east-1:123456789012:function:my-function"
	expected := "my-function"
	actual := getFunctionNameFromARN(arn)

	assert.Equal(t, expected, actual)
}

func TestNewLambda(t *testing.T) {
	cloudwatch := new(mocks.CloudWatchLogsAPI)

	service := NewAnalyzer(cloudwatch, "arn:aws:test", 0)

	assert.NotEmpty(t, service, "Expected to be not empty")
	assert.NotNil(t, service, "Expected not to be nil")
}
