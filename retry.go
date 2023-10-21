package utils

import (
	"io"
	"log"
	"time"
)

func RetryEnhance[T any](max int, f func() (T, error), handler func(error)) (T, error) {
	result, err := f()
	if err != nil {
		if max-1 == 0 {
			return result, err
		} else {
			handler(err)
			return RetryEnhance(max-1, f, handler)
		}
	}
	return result, err
}

func RetryOrPanic[T any](max int, f func() (T, error)) T {
	result, err := RetryEnhance(max, f, func(err error) {
	})
	PanicIfErr(err)
	return result
}

func Retry[T any](max int, f func() (T, error)) (T, error) {
	return RetryEnhance(max, f, func(err error) {
	})
}

func RetryWithSleep[T any](max int, f func() (T, error), duration time.Duration) (T, error) {
	return RetryEnhance(max, f, func(err error) {
		time.Sleep(duration)
	})
}

func PanicIfErr(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func CloseSilent(c io.Closer) {
	_ = c.Close()
}
