package push

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

func (client *Client) ListPushRecords(request *ListPushRecordsRequest) (response *ListPushRecordsResponse, err error) {
	response = CreateListPushRecordsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListPushRecordsWithChan(request *ListPushRecordsRequest) (<-chan *ListPushRecordsResponse, <-chan error) {
	responseChan := make(chan *ListPushRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListPushRecords(request)
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

func (client *Client) ListPushRecordsWithCallback(request *ListPushRecordsRequest, callback func(response *ListPushRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListPushRecordsResponse
		var err error
		defer close(result)
		response, err = client.ListPushRecords(request)
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

type ListPushRecordsRequest struct {
	*requests.RpcRequest
	EndTime   string           `position:"Query" name:"EndTime"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	StartTime string           `position:"Query" name:"StartTime"`
	Page      requests.Integer `position:"Query" name:"Page"`
	AppKey    requests.Integer `position:"Query" name:"AppKey"`
	PushType  string           `position:"Query" name:"PushType"`
}

type ListPushRecordsResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	Total            int    `json:"Total" xml:"Total"`
	Page             int    `json:"Page" xml:"Page"`
	PageSize         int    `json:"PageSize" xml:"PageSize"`
	PushMessageInfos struct {
		PushMessageInfo []struct {
			AppKey     int    `json:"AppKey" xml:"AppKey"`
			AppName    string `json:"AppName" xml:"AppName"`
			MessageId  string `json:"MessageId" xml:"MessageId"`
			Type       string `json:"Type" xml:"Type"`
			DeviceType string `json:"DeviceType" xml:"DeviceType"`
			PushTime   string `json:"PushTime" xml:"PushTime"`
			Title      string `json:"Title" xml:"Title"`
			Body       string `json:"Body" xml:"Body"`
		} `json:"PushMessageInfo" xml:"PushMessageInfo"`
	} `json:"PushMessageInfos" xml:"PushMessageInfos"`
}

func CreateListPushRecordsRequest() (request *ListPushRecordsRequest) {
	request = &ListPushRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "ListPushRecords", "", "")
	return
}

func CreateListPushRecordsResponse() (response *ListPushRecordsResponse) {
	response = &ListPushRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
