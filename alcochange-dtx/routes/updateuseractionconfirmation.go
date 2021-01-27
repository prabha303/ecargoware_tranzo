package routes

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/services/useractionconfirmation"
	"ecargoware/alcochange-dtx/sentryaccounts"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func updateUserActionConfirmation(router *httprouter.Router) {
	router.POST("/v1/action/confirm", UpdateUserActionConfirmation)
}

func UpdateUserActionConfirmation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	ua := useractionconfirmation.NewUserActionConfirm(rd.l, rd.dbConn)
	reqBody := dtos.UserActionConfirmationReq{}

	if !parseJSON(w, r.Body, &reqBody) {
		return
	}

	res, errW := ua.UserActionConfirmation(reqBody)
	if errW != nil {
		rd.l.Errorf("issue with GetWarningLabel ", errW.Error())
		sentryaccounts.SentryLogExceptions(errW)
		writeJSONMessage(errW.Error(), ERR_MSG, http.StatusBadRequest, rd)
		return
	}
	writeJSONStruct(res, http.StatusOK, rd)
}
