package helper

import (
	"sort"
	"strings"
)

func GetRegionFromARN(arn string) string {
	elements := strings.Split(arn, ":")
	return elements[3]
}

func MapAverage(data map[string]float64) float64 {
	var total float64 = 0.0
	for _, value := range data {
		total += value
	}
	return total / float64(len(data))
}

func SortByMemory(results map[int]float64) []int {
	keys := make([]int, 0, len(results))
	for key := range results {
		keys = append(keys, key)
	}

	sort.IntSlice(keys).Sort()

	return keys
}
