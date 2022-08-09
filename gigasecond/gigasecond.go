package gigasecond

import (
	"time"
)

const nanosecond time.Duration = time.Second * 1000000000 //(10 ^ 9)

func AddGigasecond(start time.Time) time.Time {
	return start.Add(nanosecond)
}
