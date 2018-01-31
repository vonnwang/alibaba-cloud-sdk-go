package ecs

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

func (client *Client) ModifyCommand(request *ModifyCommandRequest) (response *ModifyCommandResponse, err error) {
	response = CreateModifyCommandResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyCommandWithChan(request *ModifyCommandRequest) (<-chan *ModifyCommandResponse, <-chan error) {
	responseChan := make(chan *ModifyCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCommand(request)
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

func (client *Client) ModifyCommandWithCallback(request *ModifyCommandRequest, callback func(response *ModifyCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCommandResponse
		var err error
		defer close(result)
		response, err = client.ModifyCommand(request)
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

type ModifyCommandRequest struct {
	*requests.RpcRequest
	WorkingDir           string           `position:"Query" name:"WorkingDir"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Description          string           `position:"Query" name:"Description"`
	Name                 string           `position:"Query" name:"Name"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	CommandId            string           `position:"Query" name:"CommandId"`
	CommandContent       string           `position:"Query" name:"CommandContent"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Timeout              requests.Integer `position:"Query" name:"Timeout"`
}

type ModifyCommandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyCommandRequest() (request *ModifyCommandRequest) {
	request = &ModifyCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyCommand", "ecs", "openAPI")
	return
}

func CreateModifyCommandResponse() (response *ModifyCommandResponse) {
	response = &ModifyCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
