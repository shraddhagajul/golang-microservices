package github_provider

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/shraddhagajul/golang-microservices/src/api/client/restclient"
	"github.com/shraddhagajul/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

// entry point of every test suite
func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}
func TestConstants(t *testing.T){
	assert.EqualValues(t,"Authorization",headerAuthorization)
	assert.EqualValues(t,"token %s",headerAuthorizationFormat)
	assert.EqualValues(t,"https://api.github.com/user/repos",urlCreateRepo)
}
func TestGetAuthorizationHeader(t *testing.T){
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t,"token abc123",header)
}

func TestDefer(t *testing.T){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("function body")
}

// O/p
// function body
// 3
// 2
// 1


func TestCreateRepoErrorRestClient(t *testing.T){
	restclient.FlushMockUp()
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err : errors.New("invalid restclient response"),
		})
	response, err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,response)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	assert.EqualValues(t,"invalid restclient response",err.Message)


	// restclient.StopMockups()

	// response, err = CreateRepo("",github.CreateRepoRequest{})
	// assert.Nil(t,response)
	// assert.NotNil(t,err)


}

func TestCreateRepoInvalidResponseBody(t *testing.T){
	restclient.FlushMockUp()
	invalidCloser,_ := os.Open("-asf3")
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: invalidCloser,

		},
		})
	response, err := CreateRepo("",github.CreateRepoRequest{})
	assert.Nil(t,response)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	assert.EqualValues(t,"invalid response body",err.Message)

	}

func TestCreateRepoInvalidErrorInterface(t *testing.T){
	restclient.FlushMockUp()
	
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body : ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
		})
	response, err := CreateRepo("",github.CreateRepoRequest{})
	 assert.Nil(t,response)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	assert.EqualValues(t,"invalid json response body",err.Message)

	}

func TestCreateRepoUnauthorized(t *testing.T){
	restclient.FlushMockUp()
	
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body : ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authorization"}`)),
		},
		})
	response, err := CreateRepo("",github.CreateRepoRequest{})
	 assert.Nil(t,response)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusUnauthorized,err.StatusCode)
	assert.EqualValues(t,"Requires Authorization",err.Message)

	}

	func TestCreateRepoInvalidResponse(t *testing.T){
	restclient.FlushMockUp()
	
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
		StatusCode: http.StatusCreated,
		Body : ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
		})
	response, err := CreateRepo("",github.CreateRepoRequest{})
	
	 assert.Nil(t,response)
	assert.NotNil(t,err)
	assert.EqualValues(t,http.StatusInternalServerError,err.StatusCode)
	assert.EqualValues(t,"error while trying to unmarshalling github create repo response",err.Message)

	}

func TestCreateRepoNoError(t *testing.T){
	restclient.FlushMockUp()
	restclient.AddMock(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
		StatusCode: http.StatusCreated,
		Body : ioutil.NopCloser(strings.NewReader(`{"id": 123 }`)), //how to pass owner 
		},
		})
	response, err := CreateRepo("",github.CreateRepoRequest{})
	
	assert.NotNil(t,response)
	assert.Nil(t,err)
	assert.EqualValues(t,123,response.Id)
	}



