package supplement

import "time"

func DiffNano(startTime time.Time) int64 {
	return int64(time.Since(startTime))
}
