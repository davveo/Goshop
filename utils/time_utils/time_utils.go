package time_utils

import (
	"Goshop/global/consts"
	"fmt"
	"time"
)

func GetStartTimeAndEndTime(tp string) (start, end string) {
	var (
		year      = time.Now().Format("2006")
		month     = time.Now().Format("01")
		nextMonth = time.Now().AddDate(0, 1, 0).Format("01")
	)
	if tp == consts.YEAR {
		start = fmt.Sprintf("%s-01-01 00:00:00", year)
		end = fmt.Sprintf("%s-12-31 23:59:59", year)
	} else if tp == consts.MONTH {
		start = fmt.Sprintf("%s-%s-01 00:00:00", year, month)
		end = fmt.Sprintf("%s-%s-01 00:00:00", year, nextMonth)
	} else {
		// pass
	}
	return
}

func GetToDayOfStart() string {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(),
		currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	return fmt.Sprintf(startTime.Format(consts.TimeFormatStyleV1))
}

func GetToDayOfEnd() string {
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(),
		currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
	return fmt.Sprintf(endTime.Format(consts.TimeFormatStyleV1))
}

func GetDateStr(style string) string {
	currDate := time.Now()
	if style == consts.TimeFormatStyleV1 {
		return fmt.Sprintf(currDate.Format(consts.TimeFormatStyleV1))
	} else if style == consts.TimeFormatStyleV2 {
		return fmt.Sprintf(currDate.Format(consts.TimeFormatStyleV2))
	}
	return ""
}
