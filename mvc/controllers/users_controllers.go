package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shraddhagajul/golang-microservices/mvc/services"
	"github.com/shraddhagajul/golang-microservices/mvc/utils"
)

// func GetUser(resp http.ResponseWriter,  req *http.Request) {
	func GetUser(c *gin.Context){
	// userId, err:= strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	userId, err:= strconv.ParseInt(c.Param("user_id"),10,64)
	if err != nil{
		appErr := &utils.ApplicationError{
			Message: "user_id not a number",
			StatusCode: http.StatusBadRequest,
			Code : "bad request",
		}
	utils.RespondError(c,appErr)

		// jsonValue,_ := json.Marshal(appErr)
		// resp.WriteHeader(http.StatusBadRequest)
		// resp.Write(jsonValue)

		//Return Bad Request to user
		return
	}

	user, appErr := services.UsersService.GetUser(userId)

	if appErr != nil {
		utils.RespondError(c,appErr)
		// jsonValue,_ := json.Marshal(appErr)
		// resp.WriteHeader(appErr.StatusCode)
		// resp.Write(jsonValue)
		return
	}
	// //return json to Client
	// jsonValue,_ := json.Marshal(user)
	// resp.Write(jsonValue)

	utils.Respond(c, http.StatusOK,user)
	
}