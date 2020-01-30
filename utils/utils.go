package utils

import (
	"time"
)

func GetCurrentDateMinute() int32 {
	return GetDateMinute(time.Now().UTC())
}

func GetCurrentPartitionPeriod(
	partitionPeriodLength int,
) int32 {
	now := time.Now().UTC()

	return GetDateMinute(now.Add(-(time.Minute * time.Duration(now.Minute()%partitionPeriodLength))))
}

func GetCurrentEs() int64 {
	return time.Now().UTC().Unix()
}

func GetDateMinuteFromEpochSeconds(
	epochSeconds int64,
) int32 {
	return GetDateMinute(time.Unix(epochSeconds, 0).UTC())
}

func GetCurrentAndPreviousParitionPeriods(
	partitionPeriodLength int,
) (int32, int32) {
	now := time.Now().UTC()
	return getOffsetPartitionPeriodsFromTime(partitionPeriodLength, now.Minute()%partitionPeriodLength, now)
}

func getOffsetPartitionPeriodsFromTime(
	partitionPeriodLength int,
	minuteOffset int,
	aTime time.Time,
) (int32, int32) {
	//log.Printf("aTime: %s\n", aTime.Format("2006-01-02 15:04:05"))
	//log.Printf("minuteOffset: %d\n", minuteOffset)
	currentPeriod := GetDateMinute(aTime.Add(-(time.Minute * time.Duration(minuteOffset))))
	numMinutes := partitionPeriodLength + minuteOffset
	//log.Printf("numMinutes: %d\n", numMinutes)
	previousPeriod := GetDateMinute(aTime.Add(-(time.Minute * time.Duration(numMinutes))))

	return currentPeriod, previousPeriod
}

func GetDateMinute(
	date time.Time,
) int32 {
	year := (date.Year() - 2020) << 20
	month := int(date.Month()) << 16
	monthDate := date.Day() << 11
	hour := date.Hour() << 6
	minute := date.Minute()
	//log.Printf("minute: %d\n", minute)

	return int32(year + month + monthDate + hour + minute)
}
