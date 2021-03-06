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

// QueryNodeDeployProcesses invokes the lrg.QueryNodeDeployProcesses API synchronously
// api document: https://help.aliyun.com/api/lrg/querynodedeployprocesses.html
func (client *Client) QueryNodeDeployProcesses(request *QueryNodeDeployProcessesRequest) (response *QueryNodeDeployProcessesResponse, err error) {
	response = CreateQueryNodeDeployProcessesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryNodeDeployProcessesWithChan invokes the lrg.QueryNodeDeployProcesses API asynchronously
// api document: https://help.aliyun.com/api/lrg/querynodedeployprocesses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNodeDeployProcessesWithChan(request *QueryNodeDeployProcessesRequest) (<-chan *QueryNodeDeployProcessesResponse, <-chan error) {
	responseChan := make(chan *QueryNodeDeployProcessesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryNodeDeployProcesses(request)
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

// QueryNodeDeployProcessesWithCallback invokes the lrg.QueryNodeDeployProcesses API asynchronously
// api document: https://help.aliyun.com/api/lrg/querynodedeployprocesses.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryNodeDeployProcessesWithCallback(request *QueryNodeDeployProcessesRequest, callback func(response *QueryNodeDeployProcessesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryNodeDeployProcessesResponse
		var err error
		defer close(result)
		response, err = client.QueryNodeDeployProcesses(request)
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

// QueryNodeDeployProcessesRequest is the request struct for api QueryNodeDeployProcesses
type QueryNodeDeployProcessesRequest struct {
	*requests.RoaRequest
	Id requests.Integer `position:"Path" name:"id"`
}

// QueryNodeDeployProcessesResponse is the response struct for api QueryNodeDeployProcesses
type QueryNodeDeployProcessesResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Data    map[string]interface{} `json:"data" xml:"data"`
	Message string                 `json:"message" xml:"message"`
	Success bool                   `json:"success" xml:"success"`
}

// CreateQueryNodeDeployProcessesRequest creates a request to invoke QueryNodeDeployProcesses API
func CreateQueryNodeDeployProcessesRequest() (request *QueryNodeDeployProcessesRequest) {
	request = &QueryNodeDeployProcessesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "QueryNodeDeployProcesses", "/api/v2/tianji/node/[id]/process", "", "")
	request.Method = requests.GET
	return
}

// CreateQueryNodeDeployProcessesResponse creates a response to parse from QueryNodeDeployProcesses response
func CreateQueryNodeDeployProcessesResponse() (response *QueryNodeDeployProcessesResponse) {
	response = &QueryNodeDeployProcessesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
