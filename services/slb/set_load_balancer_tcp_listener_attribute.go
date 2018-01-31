package slb

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

func (client *Client) SetLoadBalancerTCPListenerAttribute(request *SetLoadBalancerTCPListenerAttributeRequest) (response *SetLoadBalancerTCPListenerAttributeResponse, err error) {
	response = CreateSetLoadBalancerTCPListenerAttributeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SetLoadBalancerTCPListenerAttributeWithChan(request *SetLoadBalancerTCPListenerAttributeRequest) (<-chan *SetLoadBalancerTCPListenerAttributeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerTCPListenerAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerTCPListenerAttribute(request)
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

func (client *Client) SetLoadBalancerTCPListenerAttributeWithCallback(request *SetLoadBalancerTCPListenerAttributeRequest, callback func(response *SetLoadBalancerTCPListenerAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerTCPListenerAttributeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerTCPListenerAttribute(request)
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

type SetLoadBalancerTCPListenerAttributeRequest struct {
	*requests.RpcRequest
	VServerGroup              string           `position:"Query" name:"VServerGroup"`
	SynProxy                  string           `position:"Query" name:"SynProxy"`
	UnhealthyThreshold        requests.Integer `position:"Query" name:"UnhealthyThreshold"`
	Bandwidth                 requests.Integer `position:"Query" name:"Bandwidth"`
	HealthCheckType           string           `position:"Query" name:"HealthCheckType"`
	EstablishedTimeout        requests.Integer `position:"Query" name:"EstablishedTimeout"`
	HealthCheckDomain         string           `position:"Query" name:"HealthCheckDomain"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	MasterSlaveServerGroupId  string           `position:"Query" name:"MasterSlaveServerGroupId"`
	PersistenceTimeout        requests.Integer `position:"Query" name:"PersistenceTimeout"`
	Tags                      string           `position:"Query" name:"Tags"`
	HealthCheckHttpCode       string           `position:"Query" name:"HealthCheckHttpCode"`
	Scheduler                 string           `position:"Query" name:"Scheduler"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	VServerGroupId            string           `position:"Query" name:"VServerGroupId"`
	HealthCheckInterval       requests.Integer `position:"Query" name:"HealthCheckInterval"`
	ListenerPort              requests.Integer `position:"Query" name:"ListenerPort"`
	HealthCheckURI            string           `position:"Query" name:"HealthCheckURI"`
	MaxConnection             requests.Integer `position:"Query" name:"MaxConnection"`
	AccessKeyId               string           `position:"Query" name:"access_key_id"`
	HealthCheckConnectPort    requests.Integer `position:"Query" name:"HealthCheckConnectPort"`
	LoadBalancerId            string           `position:"Query" name:"LoadBalancerId"`
	MasterSlaveServerGroup    string           `position:"Query" name:"MasterSlaveServerGroup"`
	HealthyThreshold          requests.Integer `position:"Query" name:"HealthyThreshold"`
	HealthCheckConnectTimeout requests.Integer `position:"Query" name:"HealthCheckConnectTimeout"`
}

type SetLoadBalancerTCPListenerAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSetLoadBalancerTCPListenerAttributeRequest() (request *SetLoadBalancerTCPListenerAttributeRequest) {
	request = &SetLoadBalancerTCPListenerAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerTCPListenerAttribute", "slb", "openAPI")
	return
}

func CreateSetLoadBalancerTCPListenerAttributeResponse() (response *SetLoadBalancerTCPListenerAttributeResponse) {
	response = &SetLoadBalancerTCPListenerAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
