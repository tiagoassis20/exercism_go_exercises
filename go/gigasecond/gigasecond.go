// Package Gigasecond manage time operation using gigasecond (10^9 seconds)
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond Add one Gigasecond to the specified time
func AddGigasecond(t time.Time) time.Time {
	var gs time.Duration = 1000000000 * time.Second
	return t.Add(gs)
}
