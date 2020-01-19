package utils

import (
	"time"
)

func GetCurrentDateMinute() string {
	return getDateMinute(time.Now().UTC())
}

func GetDateMinuteFromEpochSeconds(
	epochSeconds int64,
) string {
	return getDateMinute(time.Unix(epochSeconds, 0).UTC())
}

func GetPartitionPeriods() (string, string) {
	now := time.Now().UTC()
	currentPeriod := getDateMinute(now)
	nextPeriod := getDateMinute(now.Add(-(time.Minute * time.Duration(15))))

	return currentPeriod, nextPeriod
}

func getDateMinute(
	date time.Time,
) string {
	return date.Format("2006.01.02 15:04")
}
