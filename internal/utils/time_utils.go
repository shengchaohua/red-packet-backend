package utils

import "time"

const (
	Second           = 1
	SecondsPerMinute = 60 * Second
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
)

func GetCurrentTime() uint32 {
	return uint32(time.Now().Unix())
}
