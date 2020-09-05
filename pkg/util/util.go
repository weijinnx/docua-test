package util

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/weijinnx/docua-test/pkg/errors"
)

// AddJob to channel
func AddJob(jobs chan string, t time.Time, msg string) {
	timer := time.NewTimer(t.Sub(time.Now()))
	defer timer.Stop()
	<-timer.C

	// add message to print
	jobs <- msg
}

// ParseQuery params to get needed data
func ParseQuery(values url.Values) (time.Time, string, error) {
	timeStr := values.Get("time")
	if timeStr == "" {
		return time.Time{}, "", errors.EmptyTimeError
	}
	t, err := ParseTime(timeStr)
	if err != nil {
		return time.Time{}, "", err
	}
	msg := values.Get("message")
	if msg == "" {
		return time.Time{}, "", errors.EmptyMessageError
	}

	return t, msg, nil
}

// ParseTime from string in RFC3339 format
func ParseTime(tstr string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, tstr)
	if err != nil {
		log.Printf("parse error: %v", err)
		return time.Time{}, fmt.Errorf("%s: %v", errors.TimeParseError, err)
	}
	return t, nil
}

// SliceContains checks that provided string
// exists in slice or not
func SliceContains(slice []string, s string) bool {
	for _, str := range slice {
		if str == s {
			return true
		}
	}
	return false
}
