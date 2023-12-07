package desc

import "time"

// Time is a greater than comparison function for the time.Time type.
func Time(i, j time.Time) bool {
	return i.After(j)
}
