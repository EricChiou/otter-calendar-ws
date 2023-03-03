package delivery

import (
	"otter-calendar/app/model/event"
	"otter-calendar/app/types"
	"otter-calendar/app/usecase"
	"otter-calendar/app/vo"
	"otter-calendar/http/middleware"
	"otter-calendar/http/paramhandler"
	"otter-calendar/http/response"
	"strconv"
)

var Event = eventDelivery{}

type eventDelivery struct{}

func (d eventDelivery) AddEvent(webInput middleware.WebInput) {
	var addEventVO vo.AddEventVO
	if err := paramhandler.Set(webInput.Context, &addEventVO); err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	if addEventVO.StartTime < 0 {
		response.Send.FormatError(webInput.Ctx, "事件開始時間不正確", nil)
		return
	}

	if !types.CheckEventType(addEventVO.Type) {
		response.Send.FormatError(webInput.Ctx, "事件類型不正確", nil)
		return
	}

	if !types.CheckEventCalType(addEventVO.CalType) {
		response.Send.FormatError(webInput.Ctx, "事件計算類型不正確", nil)
		return
	}

	if addEventVO.Type == string(types.RepeatEvent) {
		if !types.CheckRepeatUnit(addEventVO.RepeatUnit) {
			response.Send.FormatError(webInput.Ctx, "執行間隔單位不正確", nil)
			return
		}
		if addEventVO.RepeatInterval < 0 {
			response.Send.FormatError(webInput.Ctx, "執行間隔不正確", nil)
			return
		}
		if addEventVO.RepeatTime < 0 {
			response.Send.FormatError(webInput.Ctx, "執行次數不正確", nil)
			return
		}
	}

	eventEnt := event.Entity{
		Name:           addEventVO.Name,
		Type:           types.EventType(addEventVO.Type),
		CalType:        types.EventCalType(addEventVO.CalType),
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
		response.Send.OperationError(webInput.Ctx, "伺服器內部錯誤", err)
		return
	}

	response.Send.Success(webInput.Ctx, eventList)
}

func (d eventDelivery) GetEventByEventID(webInput middleware.WebInput) {
	eventID, err := strconv.Atoi(webInput.Context.PathParam("id"))
	if err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	e, err := usecase.Event.GetEventByEventID(eventID, webInput.Payload.ID)
	if err != nil {
		response.Send.OperationError(webInput.Ctx, "伺服器內部錯誤", err)
		return
	}

	response.Send.Success(webInput.Ctx, e)
}

func (d eventDelivery) UpdateEvent(webInput middleware.WebInput) {
	var updateEventVO vo.UpdateEventVO
	if err := paramhandler.Set(webInput.Context, &updateEventVO); err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	if updateEventVO.StartTime < 0 {
		response.Send.FormatError(webInput.Ctx, "事件開始時間不正確", nil)
		return
	}

	if !types.CheckEventType(updateEventVO.Type) {
		response.Send.FormatError(webInput.Ctx, "事件類型不正確", nil)
		return
	}

	if !types.CheckEventCalType(updateEventVO.CalType) {
		response.Send.FormatError(webInput.Ctx, "事件計算類型不正確", nil)
		return
	}

	if updateEventVO.Type == string(types.RepeatEvent) {
		if !types.CheckRepeatUnit(updateEventVO.RepeatUnit) {
			response.Send.FormatError(webInput.Ctx, "執行間隔單位不正確", nil)
			return
		}
		if updateEventVO.RepeatInterval < 0 {
			response.Send.FormatError(webInput.Ctx, "執行間隔不正確", nil)
			return
		}
		if updateEventVO.RepeatTime < 0 {
			response.Send.FormatError(webInput.Ctx, "執行次數不正確", nil)
			return
		}
		if updateEventVO.LastTime < 0 {
			response.Send.FormatError(webInput.Ctx, "事件執行時間不正確", nil)
			return
		}
	}

	eventEnt := event.Entity{
		ID:             updateEventVO.ID,
		Name:           updateEventVO.Name,
		Type:           types.EventType(updateEventVO.Type),
		CalType:        types.EventCalType(updateEventVO.CalType),
		StartTime:      updateEventVO.StartTime,
		RepeatUnit:     types.EventRepeatUnit(updateEventVO.RepeatUnit),
		RepeatInterval: updateEventVO.RepeatInterval,
		RepeatTime:     updateEventVO.RepeatTime,
		LastTime:       updateEventVO.LastTime,
		Remark:         updateEventVO.Remark,
		UserID:         webInput.Payload.ID,
	}

	err := usecase.Event.UpdateEvent(eventEnt)
	if err != nil {
		response.Send.OperationError(webInput.Ctx, "伺服器內部錯誤", err)
		return
	}

	response.Send.Success(webInput.Ctx, nil)
}

func (d eventDelivery) DeleteEventByEventID(webInput middleware.WebInput) {
	eventID, err := strconv.Atoi(webInput.Context.PathParam("id"))
	if err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	err = usecase.Event.DeleteEventByEventID(eventID, webInput.Payload.ID)
	if err != nil {
		response.Send.OperationError(webInput.Ctx, "伺服器內部錯誤", err)
		return
	}

	response.Send.Success(webInput.Ctx, nil)
}
