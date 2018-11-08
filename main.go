package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	csvFileName = flag.String("filename", "test.csv", "csv file name")
)

// BalanceResponse represents balance file response record
type BalanceResponse struct {
	Acct         string `json:"account_number"`
	Responsetime uint8  `json:"response_time"`
	Raw          string `json:"raw,omitempty"`
}

func main() {
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// this is how to get a csv reader
	in := csv.NewReader(bufio.NewReader(file))
	// slice for accumulating records (for json marshalling
	var recs []BalanceResponse
	for {
		line, err := in.Read()
		if err != nil {
			// is there a more idiomatic pattern here?
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		responseTime, err := strconv.ParseInt(line[1], 10, 8)
		if err != nil {
			log.Println(err)
		}
		recs = append(recs, BalanceResponse{line[0], uint8(responseTime), line[2]})
	}
	jsn, err := json.Marshal(recs)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsn))
}
