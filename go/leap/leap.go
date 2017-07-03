package leap

import "math"

const testVersion = 3

func IsLeapYear(year int) bool {
	if math.Mod(float64(year), 400) == 0 {
		return true
	}
	if math.Mod(float64(year), 100) == 0 {
		return false
	}
	if math.Mod(float64(year), 4.0) == 0 {
		return true
	}
	return false
}
