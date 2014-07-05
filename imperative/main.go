package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Ranks, Unique Visitors, Page Views, Reach, Site, Category, Has Advertising

func main() {
	// open input file
	inputFile, err := os.Open("../input.csv")
	if err != nil {
		log.Fatal(err)
	}
	// automatically call Close() at the end of current method
	defer inputFile.Close()
	reader := csv.NewReader(inputFile)

	// open output file
	outputFile, err := os.Create("./output.csv")
	if err != nil {
		log.Fatal(err)
	}
	// automatically call Close() at the end of current method
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)

	for i := 0; i < 100; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break	// end-of-file is fitted into err
		} else if err != nil {
			log.Fatal(err)
		}
		size := getSize(record[4])
		log.Printf("%v is %v bytes", record[4], size)

		//create and write row to output
		result := append(record, fmt.Sprintf("%d", size))
		err = writer.Write(result)
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
}

func getSize(url string) int {
	resp, err := http.Get("http://" + url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return len(body)
}
