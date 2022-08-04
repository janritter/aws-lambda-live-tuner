package analyzer

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
)

type AnalyzerAPI interface {
	CheckInvocations(lambdaARN string, memory int) error
}

type Analyzer struct {
	cloudwatch cloudwatchlogsiface.CloudWatchLogsAPI
	waitTime   int
}

func NewAnalyzer(cloudwatch cloudwatchlogsiface.CloudWatchLogsAPI, waitTime int) *Analyzer {
	return &Analyzer{
		cloudwatch: cloudwatch,
		waitTime:   waitTime,
	}
}
