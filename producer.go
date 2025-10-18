package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Reciever struct {
	Name  string
	Email string
}

func producer() error {
	// Open CSV File
	f, err := os.Open("mails.csv")
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
		// Read the Records of CSV File (skipping the header at index 0)
		for _, record := range records[1:] {
			// ! record is a []string where each element is a column value
			fmt.Printf("Name: %s, Email: %s\n", record[0], record[1])
			// TODO: In a real application, you'd populate the Reciever struct here:
			// r := Reciever{Name: record[0], Email: record[1]}
		}
	} else {
		fmt.Println("CSV file is empty or only contains a header.")
	}
	return nil
}
