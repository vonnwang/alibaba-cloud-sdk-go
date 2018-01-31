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

func (client *Client) AddCdnDomain(request *AddCdnDomainRequest) (response *AddCdnDomainResponse, err error) {
	response = CreateAddCdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddCdnDomainWithChan(request *AddCdnDomainRequest) (<-chan *AddCdnDomainResponse, <-chan error) {
	responseChan := make(chan *AddCdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCdnDomain(request)
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

func (client *Client) AddCdnDomainWithCallback(request *AddCdnDomainRequest, callback func(response *AddCdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCdnDomainResponse
		var err error
		defer close(result)
		response, err = client.AddCdnDomain(request)
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

type AddCdnDomainRequest struct {
	*requests.RpcRequest
	Region          string           `position:"Query" name:"Region"`
	CheckUrl        string           `position:"Query" name:"CheckUrl"`
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	Scope           string           `position:"Query" name:"Scope"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	Sources         string           `position:"Query" name:"Sources"`
	SourcePort      requests.Integer `position:"Query" name:"SourcePort"`
	Priorities      string           `position:"Query" name:"Priorities"`
	DomainName      string           `position:"Query" name:"DomainName"`
	LiveType        string           `position:"Query" name:"LiveType"`
	CdnType         string           `position:"Query" name:"CdnType"`
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
	SourceType      string           `position:"Query" name:"SourceType"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
}

type AddCdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateAddCdnDomainRequest() (request *AddCdnDomainRequest) {
	request = &AddCdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "AddCdnDomain", "", "")
	return
}

func CreateAddCdnDomainResponse() (response *AddCdnDomainResponse) {
	response = &AddCdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
