// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import (
	"fmt"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	Hour   int
	Minute int
}

func New(hour, minute int) Clock {
	if minute >= 60 {
		hour = hour + minute/60
		minute = minute % 60
	}
	if minute < 0 {
		hour -= 1 + minute%60
		minute = 60 + minute%60
	}
	if hour >= 24 {
		hour = hour % 24
	}
	if hour < 0 {
		hour = 24 + hour%24
	}
	return Clock{
		Hour:   hour,
		Minute: minute,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%2.2d:%2.2d", c.Hour, c.Minute)
}

func (Clock) Add(minutes int) Clock {
	return New(1, 1)
}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
