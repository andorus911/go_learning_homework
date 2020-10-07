package models

import "time"

type Event struct {
	Id          string // todo uuid
	Title       string
	Description string
	Owner       string
	StartTime   time.Time
	EndTime     time.Time // mb I should do it as a pointer? for nil
}

type EditEvent struct {
	Title       *string // should i use flags?
	Description *string
	StartTime   *time.Time
	EndTime     *time.Time
}
