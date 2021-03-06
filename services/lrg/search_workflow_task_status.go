package lrg

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SearchWorkflowTaskStatus invokes the lrg.SearchWorkflowTaskStatus API synchronously
// api document: https://help.aliyun.com/api/lrg/searchworkflowtaskstatus.html
func (client *Client) SearchWorkflowTaskStatus(request *SearchWorkflowTaskStatusRequest) (response *SearchWorkflowTaskStatusResponse, err error) {
	response = CreateSearchWorkflowTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SearchWorkflowTaskStatusWithChan invokes the lrg.SearchWorkflowTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/searchworkflowtaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWorkflowTaskStatusWithChan(request *SearchWorkflowTaskStatusRequest) (<-chan *SearchWorkflowTaskStatusResponse, <-chan error) {
	responseChan := make(chan *SearchWorkflowTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchWorkflowTaskStatus(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SearchWorkflowTaskStatusWithCallback invokes the lrg.SearchWorkflowTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/lrg/searchworkflowtaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SearchWorkflowTaskStatusWithCallback(request *SearchWorkflowTaskStatusRequest, callback func(response *SearchWorkflowTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchWorkflowTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.SearchWorkflowTaskStatus(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SearchWorkflowTaskStatusRequest is the request struct for api SearchWorkflowTaskStatus
type SearchWorkflowTaskStatusRequest struct {
	*requests.RoaRequest
	Id requests.Integer `position:"Path" name:"id"`
}

// SearchWorkflowTaskStatusResponse is the response struct for api SearchWorkflowTaskStatus
type SearchWorkflowTaskStatusResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Data    map[string]interface{} `json:"data" xml:"data"`
	Message string                 `json:"message" xml:"message"`
	Success bool                   `json:"success" xml:"success"`
}

// CreateSearchWorkflowTaskStatusRequest creates a request to invoke SearchWorkflowTaskStatus API
func CreateSearchWorkflowTaskStatusRequest() (request *SearchWorkflowTaskStatusRequest) {
	request = &SearchWorkflowTaskStatusRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "SearchWorkflowTaskStatus", "/api/v2/tianji/process/[id]?action=queryTaskStatus", "", "")
	request.Method = requests.GET
	return
}

// CreateSearchWorkflowTaskStatusResponse creates a response to parse from SearchWorkflowTaskStatus response
func CreateSearchWorkflowTaskStatusResponse() (response *SearchWorkflowTaskStatusResponse) {
	response = &SearchWorkflowTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
