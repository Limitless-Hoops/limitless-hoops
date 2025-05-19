package testutil

import "time"

func StrPtr(s string) *string {
	return &s
}

func DatePtr(t time.Time) *time.Time {
	return &t
}
