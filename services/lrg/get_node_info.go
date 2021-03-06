package lrg

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

// GetNodeInfo invokes the lrg.GetNodeInfo API synchronously
// api document: https://help.aliyun.com/api/lrg/getnodeinfo.html
func (client *Client) GetNodeInfo(request *GetNodeInfoRequest) (response *GetNodeInfoResponse, err error) {
	response = CreateGetNodeInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetNodeInfoWithChan invokes the lrg.GetNodeInfo API asynchronously
// api document: https://help.aliyun.com/api/lrg/getnodeinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNodeInfoWithChan(request *GetNodeInfoRequest) (<-chan *GetNodeInfoResponse, <-chan error) {
	responseChan := make(chan *GetNodeInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNodeInfo(request)
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

// GetNodeInfoWithCallback invokes the lrg.GetNodeInfo API asynchronously
// api document: https://help.aliyun.com/api/lrg/getnodeinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetNodeInfoWithCallback(request *GetNodeInfoRequest, callback func(response *GetNodeInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNodeInfoResponse
		var err error
		defer close(result)
		response, err = client.GetNodeInfo(request)
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

// GetNodeInfoRequest is the request struct for api GetNodeInfo
type GetNodeInfoRequest struct {
	*requests.RoaRequest
}

// GetNodeInfoResponse is the response struct for api GetNodeInfo
type GetNodeInfoResponse struct {
	*responses.BaseResponse
	Code    int                    `json:"code" xml:"code"`
	Success bool                   `json:"success" xml:"success"`
	Message string                 `json:"message" xml:"message"`
	Data    map[string]interface{} `json:"data" xml:"data"`
}

// CreateGetNodeInfoRequest creates a request to invoke GetNodeInfo API
func CreateGetNodeInfoRequest() (request *GetNodeInfoRequest) {
	request = &GetNodeInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("LRG", "2019-10-10", "GetNodeInfo", "/api/v2/tianji/node", "", "")
	request.Method = requests.GET
	return
}

// CreateGetNodeInfoResponse creates a response to parse from GetNodeInfo response
func CreateGetNodeInfoResponse() (response *GetNodeInfoResponse) {
	response = &GetNodeInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
