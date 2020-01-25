package utils

import (
	"time"
)

func GetCurrentDateMinute() int32 {
	return getDateMinute(time.Now().UTC())
}
func GetCurrentEs() int64 {
	return time.Now().UTC().Unix()
}

func GetDateMinuteFromEpochSeconds(
	epochSeconds int64,
) int32 {
	return getDateMinute(time.Unix(epochSeconds, 0).UTC())
}

func GetPartitionPeriods(
	partitionPeriod int,
) (int32, int32) {
	return GetOffsetPartitionPeriods(partitionPeriod, 0)
}

func GetOffsetPartitionPeriods(
	partitionPeriod int,
	minuteOffset int,
) (int32, int32) {
	now := time.Now().UTC()
	currentPeriod := getDateMinute(now.Add(-(time.Minute * time.Duration(minuteOffset))))
	numMinutes := partitionPeriod + minuteOffset
	nextPeriod := getDateMinute(now.Add(-(time.Minute * time.Duration(numMinutes))))

	return currentPeriod, nextPeriod
}

func getDateMinute(
	date time.Time,
) int32 {
	year := (date.Year() - 2020) << 20
	month := date.Month() << 16
	monthDate := date.Date() << 11
	hour := date.Hour() << 6
	minute := date.Minute()

	return year + month + monDate + hour + minute
}
