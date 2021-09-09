package infomation

import (
	"fmt"
	"time"

	"github.com/Realive333/DailyReading/internal/scripture"
)

type Info struct {
	Date      time.Time
	Scripture string
	url       string
}

func CreateInfoStruct() (Info, error) {
	date := time.Now()
	scripture, err := scripture.GetScripture(date)
	if err != nil {
		return Info{}, err
	}
	url := fmt.Sprintf("%s%s", "https://www.biblegateway.com/passage/?version=CUVMPT&search=", scripture)

	return Info{date, scripture, url}, nil
}
