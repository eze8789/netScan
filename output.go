package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// toCSV store the output to a csv file
func toCSV(ipAddress, status string, n int) error {
	f, err := os.OpenFile("out.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()
	data := []string{ipAddress, strconv.Itoa(n), status}
	err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// toJSON store the output to a csv file
// TODO: implement toJSON function
func toJSON(ipAddress string, n int) error {
	return nil
}

// writeOutput decide which method choose to store/present the data
func writeOutput(output, ipAddress, status string, p int) {
	switch output {
	case "csv":
		err := toCSV(ipAddress, status, p)
		if err != nil {
			log.Printf("unable to store csv output in file: %s", err)
		}
	case "json":
		err := toJSON(ipAddress, p)
		if err != nil {
			log.Printf("unable to store json output in file: %s", err)
		}
	default:
		fmt.Printf("port is %s: %d\n", status, p)
	}
}
