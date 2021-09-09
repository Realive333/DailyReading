package scripture

import (
	"fmt"
	"time"

	"github.com/Realive333/DailyReading/pkg/filereader"
)

func GetScripture(date time.Time) (scripture string, err error) {
	items, err := filereader.ReadCsvFile("document/scripture.csv")
	if err != nil {
		return
	}

	_, month, day := date.Date()
	strDate := fmt.Sprintf("%d-%d", month, day)

	for _, item := range items {
		if strDate == item[1] {
			scripture = fmt.Sprintf("%s%s", item[2], item[3])
			return
		}
	}
	return "", fmt.Errorf("Unexpected error: did not find corresponding scripture, date: %v", date)
}
