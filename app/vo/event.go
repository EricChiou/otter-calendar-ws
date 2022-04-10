package vo

// Request VO
type AddEventReqVO struct {
	Name           string `json:"name" req:"true"`
	Type           string `json:"type" req:"true"`
	StartTime      int    `json:"startTime" req:"true"`
	RepeatUnit     string `json:"repeatUnit"`
	RepeatInterval int    `json:"repeatInterval"`
	RepeatTime     int    `json:"repeatTime"`
	Remark         string `json:"remark"`
}
