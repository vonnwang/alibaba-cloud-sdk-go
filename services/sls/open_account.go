package sls

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

func (client *Client) OpenAccount(request *OpenAccountRequest) (response *OpenAccountResponse, err error) {
	response = CreateOpenAccountResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) OpenAccountWithChan(request *OpenAccountRequest) (<-chan *OpenAccountResponse, <-chan error) {
	responseChan := make(chan *OpenAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenAccount(request)
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

func (client *Client) OpenAccountWithCallback(request *OpenAccountRequest, callback func(response *OpenAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenAccountResponse
		var err error
		defer close(result)
		response, err = client.OpenAccount(request)
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

type OpenAccountRequest struct {
	*requests.RpcRequest
	AliUid requests.Integer `position:"Query" name:"AliUid"`
	Bid    string           `position:"Query" name:"Bid"`
}

type OpenAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateOpenAccountRequest() (request *OpenAccountRequest) {
	request = &OpenAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2016-08-01", "OpenAccount", "", "")
	return
}

func CreateOpenAccountResponse() (response *OpenAccountResponse) {
	response = &OpenAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
