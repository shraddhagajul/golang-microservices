package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var(
	enableMocks = false
	mocks = make(map[string]*Mock)
)
type Mock struct{
	Url string
	HttpMethod string
	Response *http.Response
	Err error
}

func getMockId( httpMethod string, url string) string{
	return fmt.Sprintf("%s_%s",httpMethod,url)
}
func StartMockups(){
	enableMocks = true
}

func StopMockups(){
	enableMocks = false
}

func AddMock(mock Mock){
	mocks[getMockId(mock.HttpMethod,mock.Url)] = &mock
}

//for every testcase, AddMock is called, increments the mocks[]. 
//Flush to start afresh 
func FlushMockUp(){
	mocks = make(map[string]*Mock)
}

func Post(url string,body interface{}, header http.Header) (*http.Response,error) {
	if enableMocks{
		mock := mocks[getMockId(http.MethodPost,url)]
		if mock == nil{
			return nil,errors.New("no mockup found for given request")
		}
		// TODO : return local mock without calling any external resource
		return mock.Response, mock.Err
	}
	jsonBytes,err := json.Marshal(body)
	if err != nil{
		return nil,err
	}
	request,err := http.NewRequest(http.MethodPost,url,bytes.NewReader(jsonBytes))
	request.Header = header

	client := http.Client{}
	return client.Do(request)
}
