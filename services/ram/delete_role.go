package ram

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

func (client *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
	response = CreateDeleteRoleResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteRoleWithChan(request *DeleteRoleRequest) (<-chan *DeleteRoleResponse, <-chan error) {
	responseChan := make(chan *DeleteRoleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRole(request)
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

func (client *Client) DeleteRoleWithCallback(request *DeleteRoleRequest, callback func(response *DeleteRoleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRoleResponse
		var err error
		defer close(result)
		response, err = client.DeleteRole(request)
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

type DeleteRoleRequest struct {
	*requests.RpcRequest
	RoleName string `position:"Query" name:"RoleName"`
}

type DeleteRoleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteRoleRequest() (request *DeleteRoleRequest) {
	request = &DeleteRoleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "DeleteRole", "", "")
	return
}

func CreateDeleteRoleResponse() (response *DeleteRoleResponse) {
	response = &DeleteRoleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
