package amqp_open

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

// InstanceVO is a nested struct in amqp_open response
type InstanceVO struct {
	InstanceId        string `json:"InstanceId" xml:"InstanceId"`
	InstanceName      string `json:"InstanceName" xml:"InstanceName"`
	InstanceType      string `json:"InstanceType" xml:"InstanceType"`
	Status            string `json:"Status" xml:"Status"`
	OrderType         string `json:"OrderType" xml:"OrderType"`
	OrderCreateTime   int64  `json:"OrderCreateTime" xml:"OrderCreateTime"`
	ExpireTime        int64  `json:"ExpireTime" xml:"ExpireTime"`
	AutoRenewInstance bool   `json:"AutoRenewInstance" xml:"AutoRenewInstance"`
	SupportEIP        bool   `json:"SupportEIP" xml:"SupportEIP"`
	PrivateEndpoint   string `json:"PrivateEndpoint" xml:"PrivateEndpoint"`
	PublicEndpoint    string `json:"PublicEndpoint" xml:"PublicEndpoint"`
}