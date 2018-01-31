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

func (client *Client) ListPoliciesForGroup(request *ListPoliciesForGroupRequest) (response *ListPoliciesForGroupResponse, err error) {
	response = CreateListPoliciesForGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListPoliciesForGroupWithChan(request *ListPoliciesForGroupRequest) (<-chan *ListPoliciesForGroupResponse, <-chan error) {
	responseChan := make(chan *ListPoliciesForGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPoliciesForGroup(request)
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

func (client *Client) ListPoliciesForGroupWithCallback(request *ListPoliciesForGroupRequest, callback func(response *ListPoliciesForGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPoliciesForGroupResponse
		var err error
		defer close(result)
		response, err = client.ListPoliciesForGroup(request)
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

type ListPoliciesForGroupRequest struct {
	*requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

type ListPoliciesForGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Policies  struct {
		Policy []struct {
			PolicyName     string `json:"PolicyName" xml:"PolicyName"`
			PolicyType     string `json:"PolicyType" xml:"PolicyType"`
			Description    string `json:"Description" xml:"Description"`
			DefaultVersion string `json:"DefaultVersion" xml:"DefaultVersion"`
			AttachDate     string `json:"AttachDate" xml:"AttachDate"`
		} `json:"Policy" xml:"Policy"`
	} `json:"Policies" xml:"Policies"`
}

func CreateListPoliciesForGroupRequest() (request *ListPoliciesForGroupRequest) {
	request = &ListPoliciesForGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListPoliciesForGroup", "", "")
	return
}

func CreateListPoliciesForGroupResponse() (response *ListPoliciesForGroupResponse) {
	response = &ListPoliciesForGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
