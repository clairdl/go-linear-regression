package main

import (
	"bufio"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func TrainTestSplit() {

	// Open file & create dataframe
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dataframe := dataframe.ReadCSV(f)

	trainingNum := (4 * dataframe.Nrow()) / 5
	testNum := dataframe.Nrow() / 5

	if trainingNum+testNum < dataframe.Nrow() {
		trainingNum++
	}

	trainingIdx := make([]int, trainingNum)
	testIdx := make([]int, testNum)

	for i := 0; i < trainingNum; i++ {
		trainingIdx[i] = i
	}

	// Enumerate the test indices.

	for i := 0; i < testNum; i++ {
		testIdx[i] = trainingNum + i
	}

	// Create the subset dataframes
	trainingDF := dataframe.Subset(trainingIdx)
	testDF := dataframe.Subset(testIdx)

	setMap := map[int]dataframe.DataFrame{
		0: trainingDF,
		1: testDF,
	}

	for idx, setName := range []string{"training.csv", "test.csv"} {

		// Save the filtered dataset file.
		f, err := os.Create(setName)
		if err != nil {
			log.Fatal(err)
		}

		// Create a buffered writer
		w := bufio.NewWriter(f)

		// Write the dataframe out as a CSV.
		if err := setMap[idx].WriteCSV(w); err != nil {
			log.Fatal(err)
		}

	}

}
