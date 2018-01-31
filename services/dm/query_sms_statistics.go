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

func (client *Client) QuerySmsStatistics(request *QuerySmsStatisticsRequest) (response *QuerySmsStatisticsResponse, err error) {
	response = CreateQuerySmsStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QuerySmsStatisticsWithChan(request *QuerySmsStatisticsRequest) (<-chan *QuerySmsStatisticsResponse, <-chan error) {
	responseChan := make(chan *QuerySmsStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySmsStatistics(request)
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

func (client *Client) QuerySmsStatisticsWithCallback(request *QuerySmsStatisticsRequest, callback func(response *QuerySmsStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySmsStatisticsResponse
		var err error
		defer close(result)
		response, err = client.QuerySmsStatistics(request)
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

type QuerySmsStatisticsRequest struct {
	*requests.RpcRequest
	EndTime              string           `position:"Query" name:"EndTime"`
	StartTime            string           `position:"Query" name:"StartTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
}

type QuerySmsStatisticsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Data       struct {
		Stat []struct {
			CreateTime   string `json:"CreateTime" xml:"CreateTime"`
			RequestCount string `json:"requestCount" xml:"requestCount"`
			SuccessCount string `json:"successCount" xml:"successCount"`
			FaildCount   string `json:"faildCount" xml:"faildCount"`
		} `json:"stat" xml:"stat"`
	} `json:"data" xml:"data"`
}

func CreateQuerySmsStatisticsRequest() (request *QuerySmsStatisticsRequest) {
	request = &QuerySmsStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QuerySmsStatistics", "", "")
	return
}

func CreateQuerySmsStatisticsResponse() (response *QuerySmsStatisticsResponse) {
	response = &QuerySmsStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
