package utils

import (
	"time"
)

func GetDateHour() string {
	return time.Now().Format("2006010215")
}

func GetDateHourFromEpochSeconds(
	epochSeconds int64,
) string {
	return time.Unix(epochSeconds, 0).UTC().Format("2006010215")
}
