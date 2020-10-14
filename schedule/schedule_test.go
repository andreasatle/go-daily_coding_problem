package schedule_test

import (
	"fmt"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/schedule"
	"github.com/stretchr/testify/assert"
)

// goodBye in an example Func
func goodBye() {
	fmt.Println("Good bye!")
}

// x2 is an example FuncIntToInt
func x2(x int) int {
	return x * x
}

func TestSchedule(t *testing.T) {
	// Call function to get coverage, it's hard to test timings...
	schedule.Job(func() {}, 200)()
	assert.Equal(t, 25, schedule.JobIntToInt(x2, 1200)(5))
}
