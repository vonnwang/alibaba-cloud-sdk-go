package cs

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

func (client *Client) ModifyClusterName(request *ModifyClusterNameRequest) (response *ModifyClusterNameResponse, err error) {
	response = CreateModifyClusterNameResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyClusterNameWithChan(request *ModifyClusterNameRequest) (<-chan *ModifyClusterNameResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterName(request)
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

func (client *Client) ModifyClusterNameWithCallback(request *ModifyClusterNameRequest, callback func(response *ModifyClusterNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterNameResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterName(request)
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

type ModifyClusterNameRequest struct {
	*requests.RoaRequest
}

type ModifyClusterNameResponse struct {
	*responses.BaseResponse
}

func CreateModifyClusterNameRequest() (request *ModifyClusterNameRequest) {
	request = &ModifyClusterNameRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ModifyClusterName", "/clusters/[ClusterId]/name/[ClusterName]", "", "")
	request.Method = requests.POST
	return
}

func CreateModifyClusterNameResponse() (response *ModifyClusterNameResponse) {
	response = &ModifyClusterNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
