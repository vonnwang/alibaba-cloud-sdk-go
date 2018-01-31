package cdn

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

func (client *Client) SetUserDomainBlackList(request *SetUserDomainBlackListRequest) (response *SetUserDomainBlackListResponse, err error) {
	response = CreateSetUserDomainBlackListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetUserDomainBlackListWithChan(request *SetUserDomainBlackListRequest) (<-chan *SetUserDomainBlackListResponse, <-chan error) {
	responseChan := make(chan *SetUserDomainBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetUserDomainBlackList(request)
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

func (client *Client) SetUserDomainBlackListWithCallback(request *SetUserDomainBlackListRequest, callback func(response *SetUserDomainBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetUserDomainBlackListResponse
		var err error
		defer close(result)
		response, err = client.SetUserDomainBlackList(request)
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

type SetUserDomainBlackListRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type SetUserDomainBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetUserDomainBlackListRequest() (request *SetUserDomainBlackListRequest) {
	request = &SetUserDomainBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetUserDomainBlackList", "", "")
	return
}

func CreateSetUserDomainBlackListResponse() (response *SetUserDomainBlackListResponse) {
	response = &SetUserDomainBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
