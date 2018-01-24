package utils

import (
	"time"
)

// Retry retries a given function until it does not return an error or the timeout expires
func Retry(timeoutSeconds int, sleepSeconds int, try func() error) error {
	start, err := time.Now(), try()                       // gurantees that it tries once regardless of timeout
	time.Sleep(time.Duration(sleepSeconds) * time.Second) // all try calls are seperated by the sleep period
	for err != nil && time.Since(start) < time.Duration(timeoutSeconds)*time.Second {
		err = try()
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
	return err
}
