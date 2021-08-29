package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func HistogramLr(file string) {

	// Open file
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// https://pkg.go.dev/github.com/go-gota/gota/dataframe#ReadCSV
	dataframe := dataframe.ReadCSV(f)

	// Create a histogram for each column
	for _, colName := range dataframe.Names() {

		// Create a plotter.Values and fill it
		// with the values from each column in the dataframe

		plotVals := make(plotter.Values, dataframe.Nrow())
		for i, floatVal := range dataframe.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// Create the plot and title
		p := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// Create & normalize the histogram
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}
		h.Normalize(1)

		// Generate PNG output
		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}

func ScatterLr(file string) {

	// open file and create dataframe
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dataframe := dataframe.ReadCSV(f)

	// Select the y-axis from a given column in the dataset (dependant variable)
	yVals := dataframe.Col("Petal width").Float()

	// Create a scatter plot for each column in the dataset (independant variables)
	for _, colName := range dataframe.Names() {

		// Create and populate plotter
		pts := make(plotter.XYs, dataframe.Nrow())

		for i, floatVal := range dataframe.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// Create & configure the plot
		p := plot.New()
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Radius = vg.Points(3)

		// Generate PNG output
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
