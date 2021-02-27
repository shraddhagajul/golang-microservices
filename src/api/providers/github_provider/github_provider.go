package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/shraddhagajul/golang-microservices/src/api/client/restclient"
	"github.com/shraddhagajul/golang-microservices/src/api/domain/github"
)

// "f14763a2fe92a0f45129fd94e711d2c3c293b2cf"

const(
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string{
	return fmt.Sprintf(headerAuthorizationFormat,accessToken)
}
 func CreateRepo(accessToken string,request github.CreateRepoRequest) (*github.CreateRepoResponse,*github.GithubErrorResponse){
	headers :=http.Header{}
	headers.Set(headerAuthorization,getAuthorizationHeader(accessToken)) 
	response ,err := restclient.Post(urlCreateRepo,request,headers)
	if err != nil{
		log.Println(fmt.Sprintf("error while trying to create new repo in github: %s",err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	bytes,err := ioutil.ReadAll(response.Body)
	if err != nil{
		return nil, &github.GithubErrorResponse{
			StatusCode:  http.StatusInternalServerError,
			Message: "invalid response body",
		}
	}
// Closes while return is executed for this func 
	defer response.Body.Close()
	if response.StatusCode > 299{
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes,&errResponse); err != nil{
			return nil, &github.GithubErrorResponse{
			StatusCode:  http.StatusInternalServerError,
			Message: "invalid json response body",
		}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	 
	var result github.CreateRepoResponse
	if err:= json.Unmarshal(bytes,&result); err != nil{
	log.Println(fmt.Sprintf("error while trying to unmarshal create repo successful response: %s",err.Error()))
	return nil, &github.GithubErrorResponse{
			StatusCode:  http.StatusInternalServerError,
			Message: "error while trying to unmarshalling github create repo response",
		}
	}
	return &result,nil
}