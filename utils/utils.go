package utils

import (
	"time"
)

func GetCurrentDateMinute() int32 {
	return GetDateMinute(time.Now().UTC())
}
func GetCurrentEs() int64 {
	return time.Now().UTC().Unix()
}

func GetDateMinuteFromEpochSeconds(
	epochSeconds int64,
) int32 {
	return GetDateMinute(time.Unix(epochSeconds, 0).UTC())
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
	currentPeriod := GetDateMinute(now.Add(-(time.Minute * time.Duration(minuteOffset))))
	numMinutes := partitionPeriod + minuteOffset
	nextPeriod := GetDateMinute(now.Add(-(time.Minute * time.Duration(numMinutes))))

	return currentPeriod, nextPeriod
}

func GetDateMinute(
	date time.Time,
) int32 {
	year := (date.Year() - 2020) << 20
	month := int(date.Month()) << 16
	monthDate := date.Day() << 11
	hour := date.Hour() << 6
	minute := date.Minute()

	return int32(year + month + monthDate + hour + minute)
}
