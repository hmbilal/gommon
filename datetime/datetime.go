package datetime

import "time"

const DefaultFormat = time.RFC3339

func TimeToString(t time.Time, format string) string {
	if len(format) == 0 {
		return t.Format(DefaultFormat)
	}

	return t.Format(format)
}

func TimestampToString(t int64, format string) string {
	if len(format) == 0 {
		return time.Unix(t, 0).Format(DefaultFormat)
	}

	return time.Unix(t, 0).Format(format)
}

func GetWeekNumber(t time.Time) int {
	tn := time.Unix(t.UTC().Unix(), 0)
	_, week := tn.ISOWeek()
	return week
}
