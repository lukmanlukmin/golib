package parser

import "time"

func ValidateTime(t time.Time) time.Time {
	res, _ := time.Parse(time.RFC3339, t.Format(time.RFC3339))
	return res
}
