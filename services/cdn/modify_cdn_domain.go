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

func (client *Client) ModifyCdnDomain(request *ModifyCdnDomainRequest) (response *ModifyCdnDomainResponse, err error) {
	response = CreateModifyCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyCdnDomainWithChan(request *ModifyCdnDomainRequest) (<-chan *ModifyCdnDomainResponse, <-chan error) {
	responseChan := make(chan *ModifyCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCdnDomain(request)
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

func (client *Client) ModifyCdnDomainWithCallback(request *ModifyCdnDomainRequest, callback func(response *ModifyCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.ModifyCdnDomain(request)
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

type ModifyCdnDomainRequest struct {
	*requests.RpcRequest
	Priorities      string           `position:"Query" name:"Priorities"`
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	DomainName      string           `position:"Query" name:"DomainName"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SourceType      string           `position:"Query" name:"SourceType"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	Sources         string           `position:"Query" name:"Sources"`
	SourcePort      requests.Integer `position:"Query" name:"SourcePort"`
}

type ModifyCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateModifyCdnDomainRequest() (request *ModifyCdnDomainRequest) {
	request = &ModifyCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ModifyCdnDomain", "", "")
	return
}

func CreateModifyCdnDomainResponse() (response *ModifyCdnDomainResponse) {
	response = &ModifyCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
