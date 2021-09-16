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

func CreateInfoStruct(path string) (Info, error) {
	date := time.Now()
	scripture, err := scripture.GetScripture(date, path)
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

	if info.Date.Hour() < 4 {
		timeStr = "早安"
	} else {
		timeStr = "晚安"
	}

	timeStr = fmt.Sprintf("%s，今天是%d月%d日\n", timeStr, info.Date.Month(), info.Date.Day())

	if info.Scripture == "" {
		contentStr = fmt.Sprintf("%s今天沒有經文\n\n%s", timeStr, info.Scripture, info.url)
	} else {
		contentStr = fmt.Sprintf("%s今天的經文是：%s章\n\n%s", timeStr, info.Scripture, info.url)
	}

	return contentStr
}
