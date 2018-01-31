package cms

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

func (client *Client) ListProductOfActiveAlert(request *ListProductOfActiveAlertRequest) (response *ListProductOfActiveAlertResponse, err error) {
	response = CreateListProductOfActiveAlertResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListProductOfActiveAlertWithChan(request *ListProductOfActiveAlertRequest) (<-chan *ListProductOfActiveAlertResponse, <-chan error) {
	responseChan := make(chan *ListProductOfActiveAlertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProductOfActiveAlert(request)
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

func (client *Client) ListProductOfActiveAlertWithCallback(request *ListProductOfActiveAlertRequest, callback func(response *ListProductOfActiveAlertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProductOfActiveAlertResponse
		var err error
		defer close(result)
		response, err = client.ListProductOfActiveAlert(request)
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

type ListProductOfActiveAlertRequest struct {
	*requests.RpcRequest
	UserId string `position:"Query" name:"UserId"`
}

type ListProductOfActiveAlertResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Success    bool   `json:"Success" xml:"Success"`
	Code       int    `json:"Code" xml:"Code"`
	Message    string `json:"Message" xml:"Message"`
	Datapoints string `json:"Datapoints" xml:"Datapoints"`
}

func CreateListProductOfActiveAlertRequest() (request *ListProductOfActiveAlertRequest) {
	request = &ListProductOfActiveAlertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "ListProductOfActiveAlert", "cms", "openAPI")
	return
}

func CreateListProductOfActiveAlertResponse() (response *ListProductOfActiveAlertResponse) {
	response = &ListProductOfActiveAlertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
