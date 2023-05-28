package utils

import "time"

const (
	timeLayout = "2006-01-02T15:04:05Z"
)

var countryTz = map[string]string{
	"Indonesia": "Asia/Jakarta",
}

func TimeIn(date time.Time, name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return date.In(loc)
}

func ParseTime(date time.Time) string {
	formattedTime := TimeIn(date, "Indonesia").Format(timeLayout)
	return formattedTime
}
