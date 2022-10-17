package cost

import "github.com/janritter/aws-lambda-live-tuner/helper"

var standardGBSecondX86 = map[string]float64{
	"us-east-2":      0.0000166667,
	"us-gov-west-1":  0.0000166667,
	"us-west-1":      0.0000166667,
	"ap-southeast-2": 0.0000166667,
	"ap-southeast-1": 0.0000166667,
	"eu-north-1":     0.0000166667,
	"eu-west-3":      0.0000166667,
	"me-central-1":   0.0000206667,
	"ap-southeast-3": 0.0000166667,
	"ap-northeast-2": 0.0000166667,
	"ap-northeast-1": 0.0000166667,
	"ap-northeast-3": 0.0000166667,
	"us-gov-east-1":  0.0000166667,
	"sa-east-1":      0.0000166667,
	"ca-central-1":   0.0000166667,
	"me-south-1":     0.0000206667,
	"ap-south-1":     0.0000166667,
	"eu-west-2":      0.0000166667,
	"eu-central-1":   0.0000166667,
	"eu-west-1":      0.0000166667,
	"eu-south-1":     0.0000195172,
	"us-east-1":      0.0000166667,
	"us-west-2":      0.0000166667,
	"ap-east-1":      0.0000229200,
	"af-south-1":     0.0000221000,
}

var standardGBSecondARM = map[string]float64{
	"af-south-1":     0.0000176800,
	"ap-south-1":     0.0000133334,
	"eu-central-1":   0.0000133334,
	"us-east-1":      0.0000133334,
	"ap-east-1":      0.0000183000,
	"ap-northeast-3": 0.0000133334,
	"us-east-2":      0.0000133334,
	"sa-east-1":      0.0000133334,
	"ap-southeast-2": 0.0000133334,
	"ap-southeast-1": 0.0000133334,
	"eu-west-1":      0.0000133334,
	"ap-southeast-3": 0.0000133334,
	"me-south-1":     0.0000165334,
	"eu-west-3":      0.0000133334,
	"eu-south-1":     0.0000156138,
	"us-west-2":      0.0000133334,
	"eu-west-2":      0.0000133334,
	"eu-north-1":     0.0000133334,
	"ca-central-1":   0.0000133334,
	"us-west-1":      0.0000133334,
	"ap-northeast-2": 0.0000133334,
	"ap-northeast-1": 0.0000133334,
}

func getGBSecondPriceForArchitectureRegion(architecture, region string) float64 {
	switch architecture {
	case "arm64":
		if price, ok := standardGBSecondARM[region]; ok {
			return price
		}
	case "x86_64":
		if price, ok := standardGBSecondX86[region]; ok {
			return price
		}
	}

	helper.LogError("No price found for architecture %s and region %s", architecture, region)
	return 0.0
}

func Calculate(duration float64, memory int, architecture, region string) float64 {
	gbSecond := getGBSecondPriceForArchitectureRegion(architecture, region)
	costForMemoryInMilliseconds := (gbSecond / 1024 * float64(memory)) / 1000

	return costForMemoryInMilliseconds * duration
}
