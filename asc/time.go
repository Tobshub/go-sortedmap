package asc

import "time"

// Time is a less than comparison function for the time.Time type.
func Time(i, j time.Time) bool {
	return i.Before(j)
}
