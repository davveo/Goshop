package time_utils

import (
	"fmt"
	"Eshop/global/consts"
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
