package filereader

import (
	"encoding/csv"
	"os"
)

// ReadFile reading .csv file and return it as slice of string
func ReadCsvFile(filePath string) (items [][]string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	items, err = csvReader.ReadAll()
	if err != nil {
		return
	}

	return
}
