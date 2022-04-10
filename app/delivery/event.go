package delivery

import (
	"otter-calendar/app/model/event"
	"otter-calendar/app/types"
	"otter-calendar/app/usecase"
	"otter-calendar/app/vo"
	"otter-calendar/http/middleware"
	"otter-calendar/http/paramhandler"
	"otter-calendar/http/response"
)

var Event = eventDelivery{}

type eventDelivery struct{}

func (d eventDelivery) AddEvent(webInput middleware.WebInput) {
	var addEventVO vo.AddEventReqVO
	if err := paramhandler.Set(webInput.Context, &addEventVO); err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	if !types.CheckEventType(addEventVO.Type) {
		response.Send.FormatError(webInput.Ctx, "事件類型不正確", nil)
		return
	}

	if addEventVO.Type == string(types.RepeatEvent) && !types.CheckRepeatUnit(addEventVO.RepeatUnit) {
		response.Send.FormatError(webInput.Ctx, "執行間隔單位不正確", nil)
		return
	}

	eventEnt := event.Entity{
		Name:           addEventVO.Name,
		Type:           types.EventType(addEventVO.Type),
		StartTime:      addEventVO.StartTime,
		RepeatUnit:     types.EventRepeatUnit(addEventVO.RepeatUnit),
		RepeatInterval: addEventVO.RepeatInterval,
		RepeatTime:     addEventVO.RepeatTime,
		Remark:         addEventVO.Remark,
		UserID:         webInput.Payload.ID,
	}

	err := usecase.Event.AddEvent(eventEnt)
	if err != nil {
		response.Send.OperationError(webInput.Ctx, "伺服器內部錯誤", err)
		return
	}

	response.Send.Success(webInput.Ctx, nil)
}

func (d eventDelivery) GetEventListByUserID(webInput middleware.WebInput) {
	eventList, err := usecase.Event.GetEventListByUserID(webInput.Payload.ID)
	if err != nil {
		response.Send.FormatError(webInput.Ctx, "伺服器內部錯誤", err)
		return
	}

	response.Send.Success(webInput.Ctx, eventList)
}
