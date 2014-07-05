package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

//Ranks, Unique Visitors,	Page Views,	Reach, Site, Category, Has Advertising

func main() {

	//tell go we can use all CPU's
	/* runtime.GOMAXPROCS(runtime.NumCPU()) */
	fmt.Printf("%v", runtime.GOMAXPROCS(runtime.NumCPU()))

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

	processed := make(chan []string)

	for i := 0; i < 100; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break // end-of-file is fitted into err
		} else if err != nil {
			log.Fatal(err)
		}
		go getSize(record, processed)
	}

	for i := 0; i < 100; i++ {
		result := <-processed
		err = writer.Write(result)
		if err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
}

func getSize(record []string, processed chan []string) {
	resp, err := http.Get("http://" + record[4])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	size := len(body)
	result := append(record, fmt.Sprintf("%d", size))
	processed <- result
}
