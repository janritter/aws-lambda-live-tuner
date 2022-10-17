package helper

import (
	"log"
	"os"
	"strings"
)

func ValidateWaitTime(waitTime int) {
	if waitTime < 30 {
		LogError("Wait time must be at least 30 seconds")
		os.Exit(1)
	}
}

func ValidateLambdaARN(lambdaARN string) {
	if !strings.HasPrefix(lambdaARN, "arn:aws:lambda:") {
		LogError("Lambda ARN must be in the format arn:aws:lambda:<region>:<account-id>:function:<function-name>")
		os.Exit(1)
	}
}

func ValidateMemoryMinValue(memoryMin int) {
	if memoryMin < 128 {
		LogError("Memory min value must be greater than or equal to 128")
		os.Exit(1)
	}
}

func ValidateMemoryMaxValue(memoryMax int, memoryMin int) {
	if memoryMax <= memoryMin {
		LogError("Memory max value must be greater than the min value")
		os.Exit(1)
	}
	if memoryMax > 10240 {
		LogError("Memory max value must be less than or equal to 10240")
		os.Exit(1)
	}
}

func ValidateMemoryIncrement(memoryIncrement int) {
	if memoryIncrement < 1 {
		LogError("Memory increment must be greater than 0")
		os.Exit(1)
	}
}

func ValidateMinRequests(minRequests int) {
	if minRequests < 1 {
		log.Println("Minimum number of requests must be greater than 0")
		os.Exit(1)
	}
}

func ValidateGraphConifg(graphConfig bool, csvFilename string) {
	if graphConfig && csvFilename == "" {
		log.Println("csv-output flag must be set when enabling graph-config")
		os.Exit(1)
	}
}
