package util

import "time"

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TimeToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}

	s := t.Format("2006-01-02 15:04:05")
	return &s
}