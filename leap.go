// Utility package for juggling leapseconds
package leap

import "time"

// A list of all two-second long POSIX timestamps that cross a leap second.
var Seconds = []int64{
	time.Date(2015, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(2012, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(2008, 12, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(2005, 12, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1998, 12, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1997, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1995, 12, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1994, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1993, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1992, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1990, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1989, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1987, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1985, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1983, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1982, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1981, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1976, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1976, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1976, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1976, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1975, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1974, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1973, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1972, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
	time.Date(1972, 06, 30, 23, 59, 59, 0, time.UTC).Unix(),
}

// Get the number of leap seconds that occured before this time.
func NumLeaps(t time.Time) int {
	u := t.Unix()
	for i, s := range Seconds {
		if u > s {
			return len(Seconds) - i
		}
	}
	return 0
}

// Get the number of leap seconds that occured between two times.
func LeapDiff(t1 time.Time, t2 time.Time) int {
	n := LeapCount(t1) - LeapCount(t2)
	if n < 0 {
		n = -n
	}
	return n
}
