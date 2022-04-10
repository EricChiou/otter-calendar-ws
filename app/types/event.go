package types

type EventType string

const (
	SingleEvent EventType = "single"
	RepeatEvent EventType = "repeat"
)

type EventRepeatUnit string

const (
	Day   EventRepeatUnit = "day"
	Week  EventRepeatUnit = "week"
	Month EventRepeatUnit = "month"
	Year  EventRepeatUnit = "year"
)

var RepeatUnits []EventRepeatUnit = []EventRepeatUnit{Day, Week, Month, Year}
