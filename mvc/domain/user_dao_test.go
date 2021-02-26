package domain

import (
	"net/http"
	"testing"

	//Import external libraries
	// go get library path
	//include them in imports
	"github.com/stretchr/testify/assert"
)

//go doesnt have assert as it returns as soon as an assert fails. We want to test all conditions at once
func TestGetUserNoUserFound(t *testing.T){
	user, err := UserDao.GetUser(0)
	//a and b are both same
	//a
	assert.Nil(t,user,"we were not expecting a user with id 0") 
	//b
	if user != nil{
		t.Error("we were not expecting a user with id 0")
	}
	// Err will contain some error as user is 0 does not exists.
	assert.NotNil(t,err,"we were expecting an error when user id is 0")
	if err == nil{
	t.Error("we were expecting an error when user id is 0")
	}
	//StatusCode = not found
	assert.EqualValues(t,http.StatusNotFound,err.StatusCode)
	if err.StatusCode != http.StatusNotFound{
		t.Error("we were expecting 404 when user is not found")
	}

	assert.EqualValues(t,"not_found",err.Code)
	assert.EqualValues(t,"user 0 does not exists",err.Message)
}

func TestGetUserNotError(t *testing.T){
	user,err := UserDao.GetUser(123)
	assert.Nil(t,err,"we were not ecpecting an error for user id 123")
	assert.NotNil(t,user,"we were ecxpecting user data for user id 123")
	assert.EqualValues(t,123,user.Id)
	assert.EqualValues(t,"Shraddha",user.FirstName)
	assert.EqualValues(t,"Gajul",user.LastName)
	assert.EqualValues(t,"shrads277@gmail.com",user.Email)
	}