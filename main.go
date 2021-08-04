package main

import (
	"os"
	"dataframe"
	"fmt"
	"log"
)

func main() {

advertFile, err := os.Open("Advertising.csv")

if err != nil { 
	log.Fatal(err)
} 

defer advertFile.Close()

// Create a dataframe from the CSV file.
 advertDF := dataframe.ReadCSV(advertFile)

 // use the describe method to calculate summary statistics
 // for all of the columns in one shot
 advertSummary := advertDF.describe()

 // output the summary statistics to stdout
 fmt.Println(advertSummary)
}