package date

import "time"

const apiDateLayout = "2006-01-02T15:04:05Z"

func GetDate() string {
	return time.Now().UTC().Format(apiDateLayout)
}
