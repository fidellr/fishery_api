package utils

import (
	"strconv"
	"time"
)

func GetByWeekBeforeTimestamp(timestampStr string, weekAgo int) (string, error) {
	timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return "", err
	}

	timestamp := time.Unix(0, timestampInt*int64(time.Millisecond))
	oneWeekAgo := timestamp.AddDate(0, 0, -(weekAgo * 7))

	return oneWeekAgo.Format("2006-01-02T15:04:05Z"), nil
}

// func FormatDateToTimestamp()
