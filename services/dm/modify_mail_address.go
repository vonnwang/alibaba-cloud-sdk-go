package dm

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

func (client *Client) ModifyMailAddress(request *ModifyMailAddressRequest) (response *ModifyMailAddressResponse, err error) {
	response = CreateModifyMailAddressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyMailAddressWithChan(request *ModifyMailAddressRequest) (<-chan *ModifyMailAddressResponse, <-chan error) {
	responseChan := make(chan *ModifyMailAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyMailAddress(request)
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

func (client *Client) ModifyMailAddressWithCallback(request *ModifyMailAddressRequest, callback func(response *ModifyMailAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyMailAddressResponse
		var err error
		defer close(result)
		response, err = client.ModifyMailAddress(request)
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

type ModifyMailAddressRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	MailAddressId        requests.Integer `position:"Query" name:"MailAddressId"`
	ReplyAddress         string           `position:"Query" name:"ReplyAddress"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Password             string           `position:"Query" name:"Password"`
}

type ModifyMailAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyMailAddressRequest() (request *ModifyMailAddressRequest) {
	request = &ModifyMailAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "ModifyMailAddress", "", "")
	return
}

func CreateModifyMailAddressResponse() (response *ModifyMailAddressResponse) {
	response = &ModifyMailAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
