package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCreateRepoRequestAsJson(t *testing.T){
	request := CreateRepoRequest{
	Name:     "golang introduction",
	Description: "a golang introduction repository",
	Homepage:   "https://github.com",
	Private:    true,
	HasIssues:  true,
	HasProjects: true,
	HasWiki:    true,	
	}
// Marshal - attempts to create a valid json string based on our input interface
	bytes,err := json.Marshal(request)
	assert.Nil(t,err)
	assert.NotNil(t,bytes)

	fmt.Println(string(bytes))

	assert.EqualValues(t,`{"name":"golang introduction","description":"a golang introduction repository","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`,string(bytes))
		
	var target CreateRepoRequest
// Unmarshal takes a byte array and a *pointer* that we are trying to fill using json
// type CreateRepoRequest struct {
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Homepage    string `json:"homepage"`
// 	Private     bool   `json:"private"`
// 	HasIssues   bool   `json:"has_issues"`
// 	HasProjects bool   `json:"has_projects"`
// 	HasWiki     bool   `json:"has_wiki"`
// }

// string(bytes) = {"name":"golang introduction","description":"a golang introduction repository","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`
// Unmarshal -> put value "golang introduction" in Name field of struct CreateRepoRequest -> target 	
	err = json.Unmarshal(bytes,&target)
	assert.Nil(t,err)
	assert.EqualValues(t,target.Name, request.Name)
	assert.EqualValues(t,target.HasIssues,request.HasIssues)
}