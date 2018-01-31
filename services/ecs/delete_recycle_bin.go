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

func (client *Client) DeleteRecycleBin(request *DeleteRecycleBinRequest) (response *DeleteRecycleBinResponse, err error) {
	response = CreateDeleteRecycleBinResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteRecycleBinWithChan(request *DeleteRecycleBinRequest) (<-chan *DeleteRecycleBinResponse, <-chan error) {
	responseChan := make(chan *DeleteRecycleBinResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRecycleBin(request)
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

func (client *Client) DeleteRecycleBinWithCallback(request *DeleteRecycleBinRequest, callback func(response *DeleteRecycleBinResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRecycleBinResponse
		var err error
		defer close(result)
		response, err = client.DeleteRecycleBin(request)
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

type DeleteRecycleBinRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceIds          string           `position:"Query" name:"resourceIds"`
}

type DeleteRecycleBinResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteRecycleBinRequest() (request *DeleteRecycleBinRequest) {
	request = &DeleteRecycleBinRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteRecycleBin", "ecs", "openAPI")
	return
}

func CreateDeleteRecycleBinResponse() (response *DeleteRecycleBinResponse) {
	response = &DeleteRecycleBinResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
