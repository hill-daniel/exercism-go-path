// Provides functionality to handle gigaseconds.
package gigasecond

import (
	"time"
)

const gigasecond = 1000000000 * time.Second

// Adds a gigasecond to given input.
// A gigasecond is 10^9 (1,000,000,000) seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(gigasecond))
}
