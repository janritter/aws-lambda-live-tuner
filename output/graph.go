package output

import (
	"os"
	"strconv"

	"github.com/janritter/aws-lambda-live-tuner/helper"
)

func WriteGnuplotConfig(filename string, memoryIncrement, elements int) error {
	helper.LogInfo("Writing Gnuplot config to %s.gp", filename)

	f, err := os.Create(filename + ".gp")
	if err != nil {
		helper.LogError("Failed to create Gnuplot config file: %s", err)
		return err
	}
	defer f.Close()

	width := strconv.Itoa(getGnuplotWidth(elements))
	increment := strconv.Itoa(memoryIncrement)

	_, err = f.WriteString(`set datafile separator ','
set ylabel "Duration in ms" 
set xlabel 'Memory in MB'
set xtics ` + increment + `
set y2tics
set y2label "Cost in USD"
set format y2 "%10.9f"

set key autotitle columnhead
set grid
show grid

set terminal png size ` + width + `,1500 enhanced font "Arial,20"
set output '` + filename + `.png'

plot "` + filename + `.csv" using 1:2 with lp pt 7 lw 3 ps 1, '' using 1:3 with lp pt 7 lw 3 ps 1 axis x1y2`)

	if err != nil {
		helper.LogError("Failed to create Gnuplot config file: %s", err)
		return err
	}

	helper.LogInfo("Exeute the following command to generate the graph: gnuplot -p " + filename + ".gp")

	return nil
}

func getGnuplotWidth(elements int) int {
	if elements <= 15 {
		return 2600
	}

	return ((elements - 15) * 50) + 2600
}
