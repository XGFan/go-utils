package utils

import (
	"testing"
	"time"
)

func TestRaceResult(t *testing.T) {
	t.Run("simple racing", func(t *testing.T) {
		got, err := RaceResult[int, int]([]int{3, 4, 5, 1, 2}, func(i int) int {
			time.Sleep(time.Duration(i) * time.Second)
			return i
		}, 1100*time.Millisecond)
		if got != 1 || err != nil {
			t.Errorf("should return 1, but got %d", got)
		}
	})
	t.Run("all timeout", func(t *testing.T) {
		_, err := RaceResult[int, int]([]int{3, 4, 5, 1, 2}, func(i int) int {
			time.Sleep(time.Duration(i) * time.Second)
			return i
		}, 500*time.Millisecond)
		if err == nil {
			t.Errorf("should get error")
		}
	})
}
