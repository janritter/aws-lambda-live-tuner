package cost

var standardGBSecondX86 = map[string]float64{
	"eu-central-1": 0.0000166667,
}

var standardGBSecondARM = map[string]float64{
	"eu-central-1": 0.0000133334,
}

func getGBSecondPriceForArchitectureRegion(architecture, region string) float64 {
	if architecture == "arm64" {
		return standardGBSecondARM[region]
	}

	// default is x86
	return standardGBSecondX86[region]
}

func Calculate(duration float64, memory int, architecture, region string) float64 {
	gbSecond := getGBSecondPriceForArchitectureRegion(architecture, region)
	costForMemoryInMilliseconds := (gbSecond / 1024 * float64(memory)) / 1000

	return costForMemoryInMilliseconds * duration
}
