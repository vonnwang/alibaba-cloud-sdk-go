package ecs

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

func (client *Client) DescribeHaVips(request *DescribeHaVipsRequest) (response *DescribeHaVipsResponse, err error) {
	response = CreateDescribeHaVipsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeHaVipsWithChan(request *DescribeHaVipsRequest) (<-chan *DescribeHaVipsResponse, <-chan error) {
	responseChan := make(chan *DescribeHaVipsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHaVips(request)
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

func (client *Client) DescribeHaVipsWithCallback(request *DescribeHaVipsRequest, callback func(response *DescribeHaVipsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHaVipsResponse
		var err error
		defer close(result)
		response, err = client.DescribeHaVips(request)
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

type DescribeHaVipsRequest struct {
	*requests.RpcRequest
	PageSize             requests.Integer        `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string                  `position:"Query" name:"ResourceOwnerAccount"`
	PageNumber           requests.Integer        `position:"Query" name:"PageNumber"`
	ResourceOwnerId      requests.Integer        `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string                  `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer        `position:"Query" name:"OwnerId"`
	Filter               *[]DescribeHaVipsFilter `position:"Query" name:"Filter"  type:"Repeated"`
}

type DescribeHaVipsFilter struct {
	Key   string    `name:"Key"`
	Value *[]string `name:"Value" type:"Repeated"`
}

type DescribeHaVipsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	HaVips     struct {
		HaVip []struct {
			HaVipId             string `json:"HaVipId" xml:"HaVipId"`
			RegionId            string `json:"RegionId" xml:"RegionId"`
			VpcId               string `json:"VpcId" xml:"VpcId"`
			VSwitchId           string `json:"VSwitchId" xml:"VSwitchId"`
			IpAddress           string `json:"IpAddress" xml:"IpAddress"`
			Status              string `json:"Status" xml:"Status"`
			MasterInstanceId    string `json:"MasterInstanceId" xml:"MasterInstanceId"`
			Description         string `json:"Description" xml:"Description"`
			CreateTime          string `json:"CreateTime" xml:"CreateTime"`
			AssociatedInstances struct {
				AssociatedInstance []string `json:"associatedInstance" xml:"associatedInstance"`
			} `json:"AssociatedInstances" xml:"AssociatedInstances"`
			AssociatedEipAddresses struct {
				AssociatedEipAddresse []string `json:"associatedEipAddresse" xml:"associatedEipAddresse"`
			} `json:"AssociatedEipAddresses" xml:"AssociatedEipAddresses"`
		} `json:"HaVip" xml:"HaVip"`
	} `json:"HaVips" xml:"HaVips"`
}

func CreateDescribeHaVipsRequest() (request *DescribeHaVipsRequest) {
	request = &DescribeHaVipsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeHaVips", "ecs", "openAPI")
	return
}

func CreateDescribeHaVipsResponse() (response *DescribeHaVipsResponse) {
	response = &DescribeHaVipsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
