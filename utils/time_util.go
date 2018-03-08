package utils

import "time"

func ChinaTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func TimeStamp2UTCTime(t int64) time.Time {
	return time.Unix(t, 0).UTC()
}

func UTCTime2Timestamp(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.UTC().Unix()
}

func CurrentUTCTime() time.Time {
	return time.Now().UTC()
}
