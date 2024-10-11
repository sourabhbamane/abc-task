package repository

import (
	"encoding/csv"
	"healthclub/entity"
	"os"
	"strconv"
	"time"
)

// Date format (YYYY-MM-DD)
const format = "2006-01-02"

// write data in csv file
func AppendToCSV(filePath string, record []string) error {
	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err // Return the error if opening the file fails
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the record to the file
	if err := writer.Write(record); err != nil {
		return err // Return the error if writing the record fails
	}

	return nil
}

// reads all the classes from a CSV file
func ReadClassesFromCSV(filePath string) ([]entity.Class, error) {
	var classes []entity.Class

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return classes, nil // Return an empty array if file doesn't exist yet
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {

		startDate, err := time.Parse(format, record[1])
		if err != nil {
			return nil, err
		}

		endDate, err := time.Parse(format, record[2])
		if err != nil {
			return nil, err
		}

		capacity, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, err
		}

		class := entity.Class{
			Name:      record[0],
			StartDate: startDate,
			EndDate:   endDate,
			Capacity:  capacity,
		}

		classes = append(classes, class)
	}

	return classes, nil
}
