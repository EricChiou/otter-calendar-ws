package types

func CheckEventType(unit string) bool {
	switch unit {
	case string(SingleEvent), string(RepeatEvent):
		return true
	}
	return false
}

func CheckRepeatUnit(unit string) bool {
	switch unit {
	case string(Day), string(Week), string(Month), string(Year):
		return true
	}
	return false
}
