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

func (client *Client) DescribeInstanceAttribute(request *DescribeInstanceAttributeRequest) (response *DescribeInstanceAttributeResponse, err error) {
	response = CreateDescribeInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeInstanceAttributeWithChan(request *DescribeInstanceAttributeRequest) (<-chan *DescribeInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceAttribute(request)
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

func (client *Client) DescribeInstanceAttributeWithCallback(request *DescribeInstanceAttributeRequest, callback func(response *DescribeInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceAttribute(request)
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

type DescribeInstanceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

type DescribeInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	InstanceId              string `json:"InstanceId" xml:"InstanceId"`
	InstanceName            string `json:"InstanceName" xml:"InstanceName"`
	ImageId                 string `json:"ImageId" xml:"ImageId"`
	RegionId                string `json:"RegionId" xml:"RegionId"`
	ZoneId                  string `json:"ZoneId" xml:"ZoneId"`
	ClusterId               string `json:"ClusterId" xml:"ClusterId"`
	InstanceType            string `json:"InstanceType" xml:"InstanceType"`
	Cpu                     int    `json:"Cpu" xml:"Cpu"`
	Memory                  int    `json:"Memory" xml:"Memory"`
	HostName                string `json:"HostName" xml:"HostName"`
	Status                  string `json:"Status" xml:"Status"`
	InternetChargeType      string `json:"InternetChargeType" xml:"InternetChargeType"`
	InternetMaxBandwidthIn  int    `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	InternetMaxBandwidthOut int    `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	VlanId                  string `json:"VlanId" xml:"VlanId"`
	SerialNumber            string `json:"SerialNumber" xml:"SerialNumber"`
	CreationTime            string `json:"CreationTime" xml:"CreationTime"`
	Description             string `json:"Description" xml:"Description"`
	InstanceNetworkType     string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	IoOptimized             string `json:"IoOptimized" xml:"IoOptimized"`
	InstanceChargeType      string `json:"InstanceChargeType" xml:"InstanceChargeType"`
	ExpiredTime             string `json:"ExpiredTime" xml:"ExpiredTime"`
	StoppedMode             string `json:"StoppedMode" xml:"StoppedMode"`
	SecurityGroupIds        struct {
		SecurityGroupId []string `json:"SecurityGroupId" xml:"SecurityGroupId"`
	} `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	PublicIpAddress struct {
		IpAddress []string `json:"IpAddress" xml:"IpAddress"`
	} `json:"PublicIpAddress" xml:"PublicIpAddress"`
	InnerIpAddress struct {
		IpAddress []string `json:"IpAddress" xml:"IpAddress"`
	} `json:"InnerIpAddress" xml:"InnerIpAddress"`
	VpcAttributes struct {
		VpcId            string `json:"VpcId" xml:"VpcId"`
		VSwitchId        string `json:"VSwitchId" xml:"VSwitchId"`
		NatIpAddress     string `json:"NatIpAddress" xml:"NatIpAddress"`
		PrivateIpAddress struct {
			IpAddress []string `json:"IpAddress" xml:"IpAddress"`
		} `json:"PrivateIpAddress" xml:"PrivateIpAddress"`
	} `json:"VpcAttributes" xml:"VpcAttributes"`
	EipAddress struct {
		AllocationId       string `json:"AllocationId" xml:"AllocationId"`
		IpAddress          string `json:"IpAddress" xml:"IpAddress"`
		Bandwidth          int    `json:"Bandwidth" xml:"Bandwidth"`
		InternetChargeType string `json:"InternetChargeType" xml:"InternetChargeType"`
	} `json:"EipAddress" xml:"EipAddress"`
	OperationLocks struct {
		LockReason []struct {
			LockReason string `json:"LockReason" xml:"LockReason"`
		} `json:"LockReason" xml:"LockReason"`
	} `json:"OperationLocks" xml:"OperationLocks"`
}

func CreateDescribeInstanceAttributeRequest() (request *DescribeInstanceAttributeRequest) {
	request = &DescribeInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceAttribute", "ecs", "openAPI")
	return
}

func CreateDescribeInstanceAttributeResponse() (response *DescribeInstanceAttributeResponse) {
	response = &DescribeInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
