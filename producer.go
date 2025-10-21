package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Receiver struct {
	Name  string
	Email string
}

func loadRecipient(filePath string, ch chan Receiver) error {
	// Close channel after sending receivers is done
	defer close(ch)
	// Open CSV File
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file")
	}
	defer f.Close()

	// Create Reader for that CSV File
	r := csv.NewReader(f)

	// Read CSV File
	records, err := r.ReadAll()
	if err != nil {
		return fmt.Errorf("unable to read CSV file")
	}

	// Read the Records of CSV File
	if len(records) > 1 {
		for _, record := range records[1:] {
			ch <- Receiver{
				Name: record[0],
				Email: record[1],
			}
		}
	} else {
		return fmt.Errorf("CSV file is empty or only contains a header.")
	}
	return nil
}
