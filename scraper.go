// scraper.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"time"

	// import Colly
	"github.com/gocolly/colly"
)

// define some data structures
// to store the scraped data
type Bus struct {
	StartAt, EndAt time.Time
	IsFull         bool
}

func main() {
	// initialize the struct slices
	var buses []Bus

	// initialize the Collector
	c := colly.NewCollector()

	// set a valid User-Agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"

	// iterating over the list of industry card
	// HTML elements
	c.OnHTML(".elementor-element-6b05593c .section_cases__item", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		image := e.ChildAttr(".elementor-image-box-img img", "data-lazy-src")
		name := e.ChildText(".elementor-image-box-content .elementor-image-box-title")
		// filter out unwanted data
		if url != "" && image != "" && name != "" {
			// initialize a new Industry instance
			industry := Industry{
				Url:   url,
				Image: image,
				Name:  name,
			}
			// add the industry instance to the list
			// of scraped industries
			industries = append(industries, industry)
		}
	})

	// connect to the target site
	c.Visit("https://www.highwaybus.com/gp/reservation/rsvPlanList?mode=search&route=1&lineId=110&nearCheckOnStation=on&nearCheckOffStation=&onStationCd=001&offStationCd=059&bordingDate=20250221&danseiNum=1&zyoseiNum=0&adultMen=1&adultWomen=0&childMen=0&childWomen=0&handicapAdultMen=0&handicapAdultWomen=0&handicapChildMen=0&handicapChildWomen=0#infomationArea")

	// --- export to CSV ---

	// open the output CSV file
	csvFile, csvErr := os.Create("industries.csv")
	// if the file creation fails
	if csvErr != nil {
		log.Fatalln("Failed to create the output CSV file", csvErr)
	}
	// release the resource allocated to handle
	// the file before ending the execution
	defer csvFile.Close()

	// create a CSV file writer
	writer := csv.NewWriter(csvFile)
	// release the resources associated with the
	// file writer before ending the execution
	defer writer.Flush()

	// add the header row to the CSV
	headers := []string{
		"url",
		"image",
		"name",
	}
	writer.Write(headers)

	// store each Industry product in the
	// output CSV file
	for _, industry := range industries {
		// convert the Industry instance to
		// a slice of strings
		record := []string{
			industry.Url,
			industry.Image,
			industry.Name,
		}
		// add a new CSV record
		writer.Write(record)
	}

	// --- export to JSON ---

	// open the output JSON file
	jsonFile, jsonErr := os.Create("industries.json")
	if jsonErr != nil {
		log.Fatalln("Failed to create the output JSON file", jsonErr)
	}
	defer jsonFile.Close()
	// convert industries to an indented JSON string
	jsonString, _ := json.MarshalIndent(industries, " ", " ")

	// write the JSON string to file
	jsonFile.Write(jsonString)
}
