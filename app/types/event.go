package types

type EventType string

const (
	SingleEvent EventType = "single"
	RepeatEvent EventType = "repeat"
)

type EventCalType string

const (
	ByStart EventCalType = "by_start"
	ByLast  EventCalType = "by_last"
)

type EventRepeatUnit string

const (
	Day   EventRepeatUnit = "day"
	Week  EventRepeatUnit = "week"
	Month EventRepeatUnit = "month"
	Year  EventRepeatUnit = "year"
)

var RepeatUnits []EventRepeatUnit = []EventRepeatUnit{Day, Week, Month, Year}
