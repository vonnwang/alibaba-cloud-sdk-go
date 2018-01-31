package domain_intl

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

func (client *Client) QueryTaskDetailHistory(request *QueryTaskDetailHistoryRequest) (response *QueryTaskDetailHistoryResponse, err error) {
	response = CreateQueryTaskDetailHistoryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTaskDetailHistoryWithChan(request *QueryTaskDetailHistoryRequest) (<-chan *QueryTaskDetailHistoryResponse, <-chan error) {
	responseChan := make(chan *QueryTaskDetailHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskDetailHistory(request)
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

func (client *Client) QueryTaskDetailHistoryWithCallback(request *QueryTaskDetailHistoryRequest, callback func(response *QueryTaskDetailHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskDetailHistoryResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskDetailHistory(request)
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

type QueryTaskDetailHistoryRequest struct {
	*requests.RpcRequest
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
	TaskNo             string           `position:"Query" name:"TaskNo"`
	DomainName         string           `position:"Query" name:"DomainName"`
	TaskStatus         requests.Integer `position:"Query" name:"TaskStatus"`
	DomainNameCursor   string           `position:"Query" name:"DomainNameCursor"`
	UserClientIp       string           `position:"Query" name:"UserClientIp"`
	Lang               string           `position:"Query" name:"Lang"`
	TaskDetailNoCursor string           `position:"Query" name:"TaskDetailNoCursor"`
}

type QueryTaskDetailHistoryResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	PageSize          int    `json:"PageSize" xml:"PageSize"`
	CurrentPageCursor struct {
		TaskNo       string `json:"TaskNo" xml:"TaskNo"`
		TaskDetailNo string `json:"TaskDetailNo" xml:"TaskDetailNo"`
		TaskType     string `json:"TaskType" xml:"TaskType"`
		InstanceId   string `json:"InstanceId" xml:"InstanceId"`
		DomainName   string `json:"DomainName" xml:"DomainName"`
		TaskStatus   string `json:"TaskStatus" xml:"TaskStatus"`
		UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
		CreateTime   string `json:"CreateTime" xml:"CreateTime"`
		TryCount     int    `json:"TryCount" xml:"TryCount"`
		ErrorMsg     string `json:"ErrorMsg" xml:"ErrorMsg"`
	} `json:"CurrentPageCursor" xml:"CurrentPageCursor"`
	NextPageCursor struct {
		TaskNo       string `json:"TaskNo" xml:"TaskNo"`
		TaskDetailNo string `json:"TaskDetailNo" xml:"TaskDetailNo"`
		TaskType     string `json:"TaskType" xml:"TaskType"`
		InstanceId   string `json:"InstanceId" xml:"InstanceId"`
		DomainName   string `json:"DomainName" xml:"DomainName"`
		TaskStatus   string `json:"TaskStatus" xml:"TaskStatus"`
		UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
		CreateTime   string `json:"CreateTime" xml:"CreateTime"`
		TryCount     int    `json:"TryCount" xml:"TryCount"`
		ErrorMsg     string `json:"ErrorMsg" xml:"ErrorMsg"`
	} `json:"NextPageCursor" xml:"NextPageCursor"`
	PrePageCursor struct {
		TaskNo       string `json:"TaskNo" xml:"TaskNo"`
		TaskDetailNo string `json:"TaskDetailNo" xml:"TaskDetailNo"`
		TaskType     string `json:"TaskType" xml:"TaskType"`
		InstanceId   string `json:"InstanceId" xml:"InstanceId"`
		DomainName   string `json:"DomainName" xml:"DomainName"`
		TaskStatus   string `json:"TaskStatus" xml:"TaskStatus"`
		UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
		CreateTime   string `json:"CreateTime" xml:"CreateTime"`
		TryCount     int    `json:"TryCount" xml:"TryCount"`
		ErrorMsg     string `json:"ErrorMsg" xml:"ErrorMsg"`
	} `json:"PrePageCursor" xml:"PrePageCursor"`
	Objects []struct {
		TaskNo       string `json:"TaskNo" xml:"TaskNo"`
		TaskDetailNo string `json:"TaskDetailNo" xml:"TaskDetailNo"`
		TaskType     string `json:"TaskType" xml:"TaskType"`
		InstanceId   string `json:"InstanceId" xml:"InstanceId"`
		DomainName   string `json:"DomainName" xml:"DomainName"`
		TaskStatus   string `json:"TaskStatus" xml:"TaskStatus"`
		UpdateTime   string `json:"UpdateTime" xml:"UpdateTime"`
		CreateTime   string `json:"CreateTime" xml:"CreateTime"`
		TryCount     int    `json:"TryCount" xml:"TryCount"`
		ErrorMsg     string `json:"ErrorMsg" xml:"ErrorMsg"`
	} `json:"Objects" xml:"Objects"`
}

func CreateQueryTaskDetailHistoryRequest() (request *QueryTaskDetailHistoryRequest) {
	request = &QueryTaskDetailHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTaskDetailHistory", "", "")
	return
}

func CreateQueryTaskDetailHistoryResponse() (response *QueryTaskDetailHistoryResponse) {
	response = &QueryTaskDetailHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
