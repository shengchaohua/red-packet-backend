package timeutils

import "time"

func GetTimeStamp() uint64 {
	return uint64(time.Now().Unix())
}
