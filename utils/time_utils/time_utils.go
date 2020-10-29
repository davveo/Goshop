package time_utils

import (
	"Goshop/global/consts"
	"fmt"
	"time"

	"gitee.com/go-package/carbon"
)

/*
  更多参考: https://gitee.com/go-package/carbon
*/

func GetStartTimeAndEndTime(tp string) (start, end string) {
	if tp == consts.YEAR {
		start = carbon.CreateFromTimestamp(time.Now().Unix()).StartOfYear().ToDateTimeString()
		end = carbon.CreateFromTimestamp(time.Now().Unix()).EndOfYear().ToDateTimeString()
	} else if tp == consts.MONTH {
		start = carbon.CreateFromTimestamp(time.Now().Unix()).StartOfMonth().ToDateTimeString()
		end = carbon.CreateFromTimestamp(time.Now().Unix()).AddMonth().StartOfMonth().ToDateTimeString()
	} else {
		// pass
	}
	return
}

func GetToDayOfStart() string {
	return carbon.CreateFromTimestamp(time.Now().Unix()).StartOfDay().Format(consts.TimeFormatStyleV1)
}

func GetToDayOfEnd() string {
	return carbon.CreateFromTimestamp(time.Now().Unix()).EndOfDay().Format(consts.TimeFormatStyleV1)
}

func GetDayOfStart(timestamp int64) string {
	return carbon.CreateFromTimestamp(timestamp).StartOfDay().ToDateTimeString()
}

func GetDayOfEnd(timestamp int64) string {
	return carbon.CreateFromTimestamp(timestamp).EndOfDay().ToDateTimeString()
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

func StartOfDay() int64 {
	return carbon.CreateFromTimestamp(time.Now().Unix()).StartOfDay().ToTimestamp()
}

func EndOfDay() int64 {
	return carbon.CreateFromTimestamp(time.Now().Unix()).EndOfDay().ToTimestamp()
}
