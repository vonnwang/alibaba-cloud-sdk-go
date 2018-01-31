package cdn

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

func (client *Client) DescribeLiveAppRecordConfig(request *DescribeLiveAppRecordConfigRequest) (response *DescribeLiveAppRecordConfigResponse, err error) {
	response = CreateDescribeLiveAppRecordConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveAppRecordConfigWithChan(request *DescribeLiveAppRecordConfigRequest) (<-chan *DescribeLiveAppRecordConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveAppRecordConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveAppRecordConfig(request)
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

func (client *Client) DescribeLiveAppRecordConfigWithCallback(request *DescribeLiveAppRecordConfigRequest, callback func(response *DescribeLiveAppRecordConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveAppRecordConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveAppRecordConfig(request)
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

type DescribeLiveAppRecordConfigRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveAppRecordConfigResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	LiveAppRecord struct {
		DomainName      string `json:"DomainName" xml:"DomainName"`
		AppName         string `json:"AppName" xml:"AppName"`
		OssEndpoint     string `json:"OssEndpoint" xml:"OssEndpoint"`
		OssBucket       string `json:"OssBucket" xml:"OssBucket"`
		OssObjectPrefix string `json:"OssObjectPrefix" xml:"OssObjectPrefix"`
		CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	} `json:"LiveAppRecord" xml:"LiveAppRecord"`
}

func CreateDescribeLiveAppRecordConfigRequest() (request *DescribeLiveAppRecordConfigRequest) {
	request = &DescribeLiveAppRecordConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveAppRecordConfig", "", "")
	return
}

func CreateDescribeLiveAppRecordConfigResponse() (response *DescribeLiveAppRecordConfigResponse) {
	response = &DescribeLiveAppRecordConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
