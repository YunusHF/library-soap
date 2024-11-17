package pkgtime

import (
	"time"
)

const (
	YYYYMMDDFormat = "2006-01-02"
)

type Time interface {
	Now() time.Time
	GetYMDFormat(token string) (string, error)
}

var _ Time = (*realClock)(nil)

type realClock struct {
}

func NewTime() Time {
	return &realClock{}
}

func (rc *realClock) Now() time.Time {
	return time.Now()
}

func (rc *realClock) GetYMDFormat(token string) (string, error) {
	date, err := time.Parse(YYYYMMDDFormat, token)

	if err != nil {
		return "", err
	}

	return date.String(), nil
}
