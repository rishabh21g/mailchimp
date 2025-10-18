package main

import (
	"encoding/csv"
	"os"
)

// Producer function to load recipients from a CSV file
func LoadRecipients(filepath string, ch chan Recipient) error {

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after reading

	// Read CSV content

	read := csv.NewReader(file)
	records, err := read.ReadAll()
	if err != nil {
		return err
	}
	for _, record := range records[1:] {

		//After getting this send through channel to consumer
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}

	}
	return nil

}
