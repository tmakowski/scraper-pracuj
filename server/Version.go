package server

import "fmt"

// Structure holing information about the program's version.
type version struct {
	main int
	major int
	minor int
}

// Function returns string representing given version.
func (v version) toString() string {
	return fmt.Sprintf("%d.%d.%d", v.main, v.major, v.minor)
}

func CurrentVersion() version {
	return version{0, 1, 0}
}