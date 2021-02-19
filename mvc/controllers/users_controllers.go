package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shraddhagajul/golang-microservices/mvc/services"
	"github.com/shraddhagajul/golang-microservices/mvc/utils"
)


func GetUser(resp http.ResponseWriter,  req *http.Request) {
	userId, err:= strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	if err != nil{
		appErr := &utils.ApplicationError{
			Message: "user_id not a number",
			StatusCode: http.StatusBadRequest,
			Code : "bad request",
		}
		jsonValue,_ := json.Marshal(appErr)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonValue)
		//Return Bad Request to user
		return
	}

	user, appErr := services.GetUser(userId)
	
	if appErr != nil {
		jsonValue,_ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	//return json to Client
	jsonValue,_ := json.Marshal(user)
	resp.Write(jsonValue)
	
}