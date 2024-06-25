package utils

import "time"

func ConvertTimeToText(t time.Time) string {

	r, _ := t.UTC().MarshalText()

	return string(r)
}
