package analyzer

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/janritter/aws-lambda-live-tuner/helper"
)

func (a *Analyzer) CheckInvocations(lambdaARN string, memory int) (map[string]float64, error) {
	functionName := getFunctionNameFromARN(lambdaARN)
	logGroupName := fmt.Sprintf("/aws/lambda/%s", functionName)
	startTimeDiff := getStartTimeDiff(a.waitTime)

	output, err := a.cloudwatch.StartQuery(&cloudwatchlogs.StartQueryInput{
		QueryString:  aws.String(fmt.Sprintf(`filter @type = "REPORT" and @message like "Memory Size: %d MB"`, memory)),
		LogGroupName: aws.String(logGroupName),
		StartTime:    aws.Int64(time.Now().Add(time.Duration(-1*startTimeDiff) * time.Second).Unix()),
		EndTime:      aws.Int64(time.Now().Unix()),
	})
	if err != nil {
		helper.LogError("Failed starting CloudWatch log insights query: ", err)
		return nil, err
	}

	queryID := *output.QueryId

	resultMap := make(map[string]float64)

	for {
		queryResultOutput, err := a.cloudwatch.GetQueryResults(&cloudwatchlogs.GetQueryResultsInput{
			QueryId: aws.String(queryID),
		})
		if err != nil {
			helper.LogError("Failed getting CloudWatch log insights query results: ", err)
			return nil, err
		}

		if *queryResultOutput.Status == cloudwatchlogs.QueryStatusComplete {
			results := queryResultOutput.Results

			for _, fields := range results {
				for _, field := range fields {
					if *field.Field == "@message" {
						id, duration, err := getDurationWithRequestIdFromMessage(*field.Value)
						if err != nil {
							helper.LogError("Failed to get duration with request id from message: ", err)
							return nil, err
						}
						resultMap[id] = duration
					}
				}
			}

			helper.LogNotice("CloudWatch log insights query completed successfully")
			break
		}

		if !(*queryResultOutput.Status == cloudwatchlogs.QueryStatusComplete || *queryResultOutput.Status == cloudwatchlogs.QueryStatusRunning || *queryResultOutput.Status == cloudwatchlogs.QueryStatusScheduled) {
			helper.LogError("CloudWatch log insights query is not an expected status: ", *queryResultOutput.Status)
			return nil, fmt.Errorf("CloudWatch log insights query is not an expected status: %s", *queryResultOutput.Status)
		}
	}

	return resultMap, nil
}

func getFunctionNameFromARN(arn string) string {
	elements := strings.Split(arn, ":")
	return elements[len(elements)-1]
}

func getDurationWithRequestIdFromMessage(message string) (string, float64, error) {
	requestId, err := getRequestIdFromMessage(message)
	if err != nil {
		return "", -1, err
	}

	duration, err := getDurationFromMessage(message)
	if err != nil {
		return "", -1, err
	}

	return requestId, duration, nil
}

func getRequestIdFromMessage(message string) (string, error) {
	start := strings.Index(message, "RequestId: ")
	end := strings.Index(message, "\tDuration")

	if start == -1 || end == -1 {
		return "", fmt.Errorf("Failed to parse RequestId from message: %s", message)
	}

	return (message[start+11 : end]), nil
}

func getDurationFromMessage(message string) (float64, error) {
	start := strings.Index(message, "Duration: ")
	end := strings.Index(message, " ms")

	if start == -1 || end == -1 {
		return -1, fmt.Errorf("Failed to parse duration from message: %s", message)
	}

	duration := (message[start+10 : end])

	durationFloat, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return -1, fmt.Errorf("Failed to convert duration to int: %s", duration)
	}
	return durationFloat, nil
}

func getStartTimeDiff(waitTime int) int {
	// We use the wait time multplied by 2 to not miss any invocations between checks
	wait := waitTime * 2

	// Due to delay in CloudWatch ingestion we always use the last 5 minutes as the minimum time window
	if wait <= 300 {
		wait = 300
	}

	return wait
}
