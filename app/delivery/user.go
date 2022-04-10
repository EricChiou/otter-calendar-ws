package delivery

import (
	"otter-calendar/app/usecase"
	"otter-calendar/app/vo"
	"otter-calendar/http/middleware"
	"otter-calendar/http/paramhandler"
	"otter-calendar/http/response"
)

var User = userDelivery{}

type userDelivery struct{}

func (d userDelivery) SignUp(webInput middleware.WebInput) {
	var signupVO vo.SignUpReqVO
	if err := paramhandler.Set(webInput.Context, &signupVO); err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	err := usecase.User.SignUp(signupVO.Account, signupVO.Password)
	if err != nil {
		response.Send.OperationError(webInput.Ctx, "帳號重複", err)
		return
	}

	response.Send.Success(webInput.Ctx, nil)
}

func (d userDelivery) Login(webInput middleware.WebInput) {
	var loginReqVO vo.LoginReqVO
	if err := paramhandler.Set(webInput.Context, &loginReqVO); err != nil {
		response.Send.FormatError(webInput.Ctx, "資料格式不正確", err)
		return
	}

	token, err := usecase.User.Login(loginReqVO.Account, loginReqVO.Password)
	if err != nil {
		response.Send.FormatError(webInput.Ctx, "帳號或密碼錯誤", err)
		return
	}

	result := make(map[string]string)
	result["token"] = token
	response.Send.Success(webInput.Ctx, result)
}
