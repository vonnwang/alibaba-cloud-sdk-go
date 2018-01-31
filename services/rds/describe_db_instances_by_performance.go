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

func (client *Client) DescribeDBInstancesByPerformance(request *DescribeDBInstancesByPerformanceRequest) (response *DescribeDBInstancesByPerformanceResponse, err error) {
	response = CreateDescribeDBInstancesByPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDBInstancesByPerformanceWithChan(request *DescribeDBInstancesByPerformanceRequest) (<-chan *DescribeDBInstancesByPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesByPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstancesByPerformance(request)
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

func (client *Client) DescribeDBInstancesByPerformanceWithCallback(request *DescribeDBInstancesByPerformanceRequest, callback func(response *DescribeDBInstancesByPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesByPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstancesByPerformance(request)
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

type DescribeDBInstancesByPerformanceRequest struct {
	*requests.RpcRequest
	Tags                 string           `position:"Query" name:"Tags"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tag5Key              string           `position:"Query" name:"Tag.5.key"`
	Tag5Value            string           `position:"Query" name:"Tag.5.value"`
	Tag3Key              string           `position:"Query" name:"Tag.3.key"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Tag2Key              string           `position:"Query" name:"Tag.2.key"`
	Tag1Key              string           `position:"Query" name:"Tag.1.key"`
	SortKey              string           `position:"Query" name:"SortKey"`
	Tag1Value            string           `position:"Query" name:"Tag.1.value"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Tag4Value            string           `position:"Query" name:"Tag.4.value"`
	Tag3Value            string           `position:"Query" name:"Tag.3.value"`
	SortMethod           string           `position:"Query" name:"SortMethod"`
	Tag2Value            string           `position:"Query" name:"Tag.2.value"`
	Tag4Key              string           `position:"Query" name:"Tag.4.key"`
}

type DescribeDBInstancesByPerformanceResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	PageNumber       int    `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int    `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int    `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            struct {
		DBInstancePerformance []struct {
			CPUUsage              string `json:"CPUUsage" xml:"CPUUsage"`
			IOPSUsage             string `json:"IOPSUsage" xml:"IOPSUsage"`
			DiskUsage             string `json:"DiskUsage" xml:"DiskUsage"`
			SessionUsage          string `json:"SessionUsage" xml:"SessionUsage"`
			DBInstanceId          string `json:"DBInstanceId" xml:"DBInstanceId"`
			DBInstanceDescription string `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
		} `json:"DBInstancePerformance" xml:"DBInstancePerformance"`
	} `json:"Items" xml:"Items"`
}

func CreateDescribeDBInstancesByPerformanceRequest() (request *DescribeDBInstancesByPerformanceRequest) {
	request = &DescribeDBInstancesByPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancesByPerformance", "rds", "openAPI")
	return
}

func CreateDescribeDBInstancesByPerformanceResponse() (response *DescribeDBInstancesByPerformanceResponse) {
	response = &DescribeDBInstancesByPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
