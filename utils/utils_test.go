package utils

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRetryTimeout(t *testing.T) {
	timeoutSeconds, sleepSeconds, try := 5, 0, func() error {
		return errors.New("error")
	}
	start, err := time.Now(), Retry(timeoutSeconds, sleepSeconds, try)
	if err == nil || time.Since(start) < time.Duration(timeoutSeconds)*time.Second {
		t.Fatalf("Retry(...) falsely returning a nil error or returned before timeout")
	}
}

func TestRetrySleep(t *testing.T) {
	timesCalled := 0
	timeoutSeconds, sleepSeconds, try := 1, 2, func() error {
		timesCalled++
		return errors.New("error")
	}
	Retry(timeoutSeconds, sleepSeconds, try)
	if timesCalled != 1 {
		log.Print(timesCalled)
		t.Fatalf("Retry(...) is sleeping too little/much")
	}
}

func TestRetryNoError(t *testing.T) {
	timeoutSeconds, sleepSeconds, try := 1, 0, func() error {
		return nil
	}
	if Retry(timeoutSeconds, sleepSeconds, try) != nil {
		t.Fatalf("Retry(...) is returning an error but it's not supposed to")
	}
}
