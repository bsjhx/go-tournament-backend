package tournament

import "testing"

func Test_RoundRobin(t *testing.T) {
	// Given
	n := 14

	GenerateSchedule(n)
	GenerateSchedule(3, true)

}
