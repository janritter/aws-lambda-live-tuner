package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDurationWithRequestIdFromMessage(t *testing.T) {
	message := "REPORT RequestId: 7239847892374-7238947823947823-234234324\tDuration: 189.12 ms\tBilled Duration: 190 ms\tMemory Size: 256 MB\tMax Memory Used: 67 MB\t\n"
	expectedDuration := 189.12
	expectedRequestId := "7239847892374-7238947823947823-234234324"

	requestId, duration, err := getDurationWithRequestIdFromMessage(message)
	assert.Nil(t, err)
	assert.Equal(t, expectedDuration, duration)
	assert.Equal(t, expectedRequestId, requestId)
}

func TestGetDurationWithRequestIdFromMessageNoDuration(t *testing.T) {
	message := "REPORT RequestId: 7239847892374-7238947823947823-234234324\tDuration \tMemory Size: 256 MB\tMax Memory Used: 67 MB\t\n"

	requestId, duration, err := getDurationWithRequestIdFromMessage(message)
	assert.Error(t, err)
	assert.Equal(t, -1.0, duration)
	assert.Equal(t, "", requestId)
}

func TestGetDurationWithRequestIdFromMessageNoRequestId(t *testing.T) {
	message := "REPORT Duration: 189.12 ms\tBilled Duration: 190 ms\tMemory Size: 256 MB\tMax Memory Used: 67 MB\t\n"

	requestId, duration, err := getDurationWithRequestIdFromMessage(message)
	assert.Error(t, err)
	assert.Equal(t, -1.0, duration)
	assert.Equal(t, "", requestId)
}

func TestGetStartTimeDiff(t *testing.T) {
	assert.Equal(t, 300, getStartTimeDiff(150))

	input := 301
	expected := input * 2
	assert.Equal(t, expected, getStartTimeDiff(input))
}
