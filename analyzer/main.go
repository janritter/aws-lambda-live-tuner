package analyzer

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

type AnalyzerAPI interface {
	CheckInvocations(memory int) error
}

type Analyzer struct {
	cloudwatch   cloudwatchlogsiface.CloudWatchLogsAPI
	waitTime     int
	functionArn  string
	logGroupName string
}

func NewAnalyzer(cloudwatch cloudwatchlogsiface.CloudWatchLogsAPI, arn string, waitTime int) *Analyzer {
	return &Analyzer{
		cloudwatch:   cloudwatch,
		waitTime:     waitTime,
		functionArn:  arn,
		logGroupName: "/aws/lambda/" + getFunctionNameFromARN(arn),
	}
}

func getFunctionNameFromARN(arn string) string {
	elements := strings.Split(arn, ":")
	return elements[len(elements)-1]
}
