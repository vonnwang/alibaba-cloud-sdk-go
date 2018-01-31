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

func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (response *ModifyDBInstanceNetworkTypeResponse, err error) {
	response = CreateModifyDBInstanceNetworkTypeResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyDBInstanceNetworkTypeWithChan(request *ModifyDBInstanceNetworkTypeRequest) (<-chan *ModifyDBInstanceNetworkTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceNetworkTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceNetworkType(request)
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

func (client *Client) ModifyDBInstanceNetworkTypeWithCallback(request *ModifyDBInstanceNetworkTypeRequest, callback func(response *ModifyDBInstanceNetworkTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceNetworkTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceNetworkType(request)
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

type ModifyDBInstanceNetworkTypeRequest struct {
	*requests.RpcRequest
	ReadWriteSplittingPrivateIpAddress   string           `position:"Query" name:"ReadWriteSplittingPrivateIpAddress"`
	VSwitchId                            string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress                     string           `position:"Query" name:"PrivateIpAddress"`
	DBInstanceId                         string           `position:"Query" name:"DBInstanceId"`
	ReadWriteSplittingClassicExpiredDays requests.Integer `position:"Query" name:"ReadWriteSplittingClassicExpiredDays"`
	OwnerId                              requests.Integer `position:"Query" name:"OwnerId"`
	RetainClassic                        string           `position:"Query" name:"RetainClassic"`
	ResourceOwnerAccount                 string           `position:"Query" name:"ResourceOwnerAccount"`
	ClassicExpiredDays                   string           `position:"Query" name:"ClassicExpiredDays"`
	ResourceOwnerId                      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                         string           `position:"Query" name:"OwnerAccount"`
	VPCId                                string           `position:"Query" name:"VPCId"`
	InstanceNetworkType                  string           `position:"Query" name:"InstanceNetworkType"`
}

type ModifyDBInstanceNetworkTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

func CreateModifyDBInstanceNetworkTypeRequest() (request *ModifyDBInstanceNetworkTypeRequest) {
	request = &ModifyDBInstanceNetworkTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDBInstanceNetworkType", "rds", "openAPI")
	return
}

func CreateModifyDBInstanceNetworkTypeResponse() (response *ModifyDBInstanceNetworkTypeResponse) {
	response = &ModifyDBInstanceNetworkTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
