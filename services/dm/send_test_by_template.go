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

func (client *Client) SendTestByTemplate(request *SendTestByTemplateRequest) (response *SendTestByTemplateResponse, err error) {
	response = CreateSendTestByTemplateResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SendTestByTemplateWithChan(request *SendTestByTemplateRequest) (<-chan *SendTestByTemplateResponse, <-chan error) {
	responseChan := make(chan *SendTestByTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SendTestByTemplate(request)
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

func (client *Client) SendTestByTemplateWithCallback(request *SendTestByTemplateRequest, callback func(response *SendTestByTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendTestByTemplateResponse
		var err error
		defer close(result)
		response, err = client.SendTestByTemplate(request)
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

type SendTestByTemplateRequest struct {
	*requests.RpcRequest
	Birthday             string           `position:"Query" name:"Birthday"`
	AccountName          string           `position:"Query" name:"AccountName"`
	NickName             string           `position:"Query" name:"NickName"`
	TemplateId           requests.Integer `position:"Query" name:"TemplateId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UserName             string           `position:"Query" name:"UserName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Email                string           `position:"Query" name:"Email"`
	Gender               string           `position:"Query" name:"Gender"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Mobile               string           `position:"Query" name:"Mobile"`
}

type SendTestByTemplateResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateSendTestByTemplateRequest() (request *SendTestByTemplateRequest) {
	request = &SendTestByTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "SendTestByTemplate", "", "")
	return
}

func CreateSendTestByTemplateResponse() (response *SendTestByTemplateResponse) {
	response = &SendTestByTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
