package retry

import (
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	cases := []int{0, 1, 2, 5, 20, 100}

	for i := 1; i < len(cases); i++ {
		a, b := cases[i-1], cases[i]
		tryCase(t, New([]Strategy{
			&CountStrategy{Tries: a},
			&CountStrategy{Tries: b},
		}), testCase{
			name:     []int{a, b},
			attempts: b + 1,
			minimum:  a,
			maximum:  a,
		})
	}

}

func TestIterationsAndTime(t *testing.T) {
	for _, test := range []struct {
		attempts   int
		iterations int
		duration   time.Duration
	}{
		{15, 1, time.Millisecond},
		{10, 30, time.Millisecond},
	} {
		tryCase(t, New([]Strategy{
			&CountStrategy{Tries: test.iterations},
			&MaximumTimeStrategy{Duration: test.duration},
		}), testCase{
			name:        test,
			attempts:    test.attempts,
			maximum:     test.iterations,
			maxDuration: test.duration,
			step:        time.Millisecond / 10,
		})
	}

}
