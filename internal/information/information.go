package information

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

func (info Info) GetMessageStr() string {
	var (
		timeStr    string
		contentStr string
	)

	if info.Date.Hour() < 12 {
		timeStr = "早安，"
	} else {
		timeStr = "晚安，"
	}

	contentStr = fmt.Sprintf("%s\n\n今天的經文是：%s章\n\n%s", timeStr, info.Scripture, info.url)

	return contentStr
}
