// Package clock provides functionality for handling times without dates.
package clock

import (
	"fmt"
)

const (
	minuteCap = 60
	hourCap   = 24
	totalCap  = minuteCap * hourCap
)

// Clock handles times without dates.
type Clock struct {
	minutes int
}

// New creates a clock
func New(hours, minutes int) Clock {
	c := Clock{(minuteCap*hours + minutes) % totalCap}
	return set(c)
}

// Add adds the given number of minutes to the time of the clock.
func (c Clock) Add(minutes int) Clock {
	c.minutes = (c.minutes + minutes) % totalCap
	return set(c)
}

// Subtract subtracts the given number of minutes from the time of the clock.
func (c Clock) Subtract(minutes int) Clock {
	c.minutes = (c.minutes - minutes) % totalCap
	return set(c)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

func set(c Clock) Clock {
	if c.minutes < 0 {
		c.minutes += totalCap
	}
	return c
}
