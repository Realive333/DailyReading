package test

import (
	"testing"

	"github.com/Realive333/DailyReading/pkg/filereader"
)

func Test_ReadCsvFile(t *testing.T) {
	actual, err := filereader.ReadCsvFile("document/scripture.csv")
	if err != nil {
		t.Errorf("Test ReadCsvFile Error: %v", err.Error())
	}

	if len(actual) != 31 {
		t.Errorf("Test ReadCsvFile: length should be %d; is %d", 31, len(actual))
	}

	if actual[0][0] != "1" {
		t.Errorf("Test ReadCsvFile: first item unexpected; is %v", actual[0])
	}
}
