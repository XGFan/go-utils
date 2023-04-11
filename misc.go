package utils

import (
	"errors"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RaceResultWithError[T any, R any](works []T, workFunc func(T) (R, error), timeout time.Duration) (R, error) {
	result := make(chan R)
	timer := time.NewTimer(timeout)
	for _, w := range works {
		go func(arg T) {
			r, e := workFunc(arg)
			if e == nil {
				select {
				case result <- r:
				case <-timer.C:
				}
			}
		}(w)
	}

	for {
		select {
		case ret := <-result:
			return ret, nil
		case <-timer.C:
			return *new(R), errors.New("all attempt fail")
		}
	}
}

func RaceResult[T any, R any](works []T, workFunc func(T) R, timeout time.Duration) (R, error) {
	return RaceResultWithError[T, R](works, func(t T) (R, error) {
		return workFunc(t), nil
	}, timeout)
}
