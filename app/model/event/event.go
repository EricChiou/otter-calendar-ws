package event

import "otter-calendar/app/types"

const (
	Table string = "public.event"
	PK    string = "id"
)

const (
	ID             string = "id"
	Name           string = "name"
	Type           string = "type"
	StartTime      string = "start_time"
	RepeatUnit     string = "repeat_unit"
	RepeatInterval string = "repeat_interval"
	RepeatTime     string = "repeat_time"
	LastTime       string = "last_time"
	Remark         string = "remark"
	UserID         string = "user_id"
)

type Entity struct {
	ID             int                   `json:"id,omitempty"`
	Name           string                `json:"name,omitempty"`
	Type           types.EventType       `json:"type,omitempty"`
	StartTime      int                   `json:"startTime,omitempty"`
	RepeatUnit     types.EventRepeatUnit `json:"repeatUnit,omitempty"`
	RepeatInterval int                   `json:"repeatInterval,omitempty"`
	RepeatTime     int                   `json:"repeatTime,omitempty"`
	LastTime       int                   `json:"lastTime,omitempty"`
	Remark         string                `json:"remark,omitempty"`
	UserID         int                   `json:"userID,omitempty"`
}
