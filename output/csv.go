package output

import (
	"encoding/csv"
	"os"

	"github.com/janritter/aws-lambda-live-tuner/helper"
)

func WriteCSV(filename string, records [][]string) error {
	helper.LogInfo("Writing results to %s.csv", filename)

	f, err := os.Create(filename + ".csv") // creates a file at current directory
	if err != nil {
		helper.LogError("Failed to create file: %s", err)
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		helper.LogError("Failed writing to csv file: %s", err)
		return err
	}

	return nil
}
