package rds

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

func (client *Client) DescribeSQLCollectorPolicy(request *DescribeSQLCollectorPolicyRequest) (response *DescribeSQLCollectorPolicyResponse, err error) {
	response = CreateDescribeSQLCollectorPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeSQLCollectorPolicyWithChan(request *DescribeSQLCollectorPolicyRequest) (<-chan *DescribeSQLCollectorPolicyResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLCollectorPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLCollectorPolicy(request)
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

func (client *Client) DescribeSQLCollectorPolicyWithCallback(request *DescribeSQLCollectorPolicyRequest, callback func(response *DescribeSQLCollectorPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLCollectorPolicyResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLCollectorPolicy(request)
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

type DescribeSQLCollectorPolicyRequest struct {
	*requests.RpcRequest
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

type DescribeSQLCollectorPolicyResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	SQLCollectorStatus string `json:"SQLCollectorStatus" xml:"SQLCollectorStatus"`
	StoragePeriod      int    `json:"StoragePeriod" xml:"StoragePeriod"`
}

func CreateDescribeSQLCollectorPolicyRequest() (request *DescribeSQLCollectorPolicyRequest) {
	request = &DescribeSQLCollectorPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLCollectorPolicy", "rds", "openAPI")
	return
}

func CreateDescribeSQLCollectorPolicyResponse() (response *DescribeSQLCollectorPolicyResponse) {
	response = &DescribeSQLCollectorPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
