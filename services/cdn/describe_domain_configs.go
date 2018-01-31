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

func (client *Client) DescribeDomainConfigs(request *DescribeDomainConfigsRequest) (response *DescribeDomainConfigsResponse, err error) {
	response = CreateDescribeDomainConfigsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainConfigsWithChan(request *DescribeDomainConfigsRequest) (<-chan *DescribeDomainConfigsResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainConfigs(request)
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

func (client *Client) DescribeDomainConfigsWithCallback(request *DescribeDomainConfigsRequest, callback func(response *DescribeDomainConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainConfigsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainConfigs(request)
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

type DescribeDomainConfigsRequest struct {
	*requests.RpcRequest
	ConfigList    string           `position:"Query" name:"ConfigList"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeDomainConfigsResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	DomainConfigs struct {
		CcConfig struct {
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Enable   string `json:"Enable" xml:"Enable"`
			AllowIps string `json:"AllowIps" xml:"AllowIps"`
			BlockIps string `json:"BlockIps" xml:"BlockIps"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"CcConfig" xml:"CcConfig"`
		ErrorPageConfig struct {
			ConfigId      string `json:"ConfigId" xml:"ConfigId"`
			ErrorCode     string `json:"ErrorCode" xml:"ErrorCode"`
			PageType      string `json:"PageType" xml:"PageType"`
			CustomPageUrl string `json:"CustomPageUrl" xml:"CustomPageUrl"`
			Status        string `json:"Status" xml:"Status"`
		} `json:"ErrorPageConfig" xml:"ErrorPageConfig"`
		OptimizeConfig struct {
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Enable   string `json:"Enable" xml:"Enable"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"OptimizeConfig" xml:"OptimizeConfig"`
		PageCompressConfig struct {
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Enable   string `json:"Enable" xml:"Enable"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"PageCompressConfig" xml:"PageCompressConfig"`
		IgnoreQueryStringConfig struct {
			ConfigId    string `json:"ConfigId" xml:"ConfigId"`
			HashKeyArgs string `json:"HashKeyArgs" xml:"HashKeyArgs"`
			Enable      string `json:"Enable" xml:"Enable"`
			Status      string `json:"Status" xml:"Status"`
		} `json:"IgnoreQueryStringConfig" xml:"IgnoreQueryStringConfig"`
		RangeConfig struct {
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Enable   string `json:"Enable" xml:"Enable"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"RangeConfig" xml:"RangeConfig"`
		RefererConfig struct {
			ConfigId   string `json:"ConfigId" xml:"ConfigId"`
			ReferType  string `json:"ReferType" xml:"ReferType"`
			ReferList  string `json:"ReferList" xml:"ReferList"`
			AllowEmpty string `json:"AllowEmpty" xml:"AllowEmpty"`
			DisableAst string `json:"DisableAst" xml:"DisableAst"`
			Status     string `json:"Status" xml:"Status"`
		} `json:"RefererConfig" xml:"RefererConfig"`
		ReqAuthConfig struct {
			ConfigId         string `json:"ConfigId" xml:"ConfigId"`
			AuthType         string `json:"AuthType" xml:"AuthType"`
			Key1             string `json:"Key1" xml:"Key1"`
			Key2             string `json:"Key2" xml:"Key2"`
			Status           string `json:"Status" xml:"Status"`
			AliAuthWhiteList string `json:"AliAuthWhiteList" xml:"AliAuthWhiteList"`
			AuthM3u8         string `json:"AuthM3u8" xml:"AuthM3u8"`
			AuthAddr         string `json:"AuthAddr" xml:"AuthAddr"`
			AuthRemoteDesc   string `json:"AuthRemoteDesc" xml:"AuthRemoteDesc"`
			TimeOut          string `json:"TimeOut" xml:"TimeOut"`
		} `json:"ReqAuthConfig" xml:"ReqAuthConfig"`
		SrcHostConfig struct {
			ConfigId   string `json:"ConfigId" xml:"ConfigId"`
			DomainName string `json:"DomainName" xml:"DomainName"`
			Status     string `json:"Status" xml:"Status"`
		} `json:"SrcHostConfig" xml:"SrcHostConfig"`
		VideoSeekConfig struct {
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Enable   string `json:"Enable" xml:"Enable"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"VideoSeekConfig" xml:"VideoSeekConfig"`
		WafConfig struct {
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Enable   string `json:"Enable" xml:"Enable"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"WafConfig" xml:"WafConfig"`
		NotifyUrlConfig struct {
			Enable    string `json:"Enable" xml:"Enable"`
			NotifyUrl string `json:"NotifyUrl" xml:"NotifyUrl"`
		} `json:"NotifyUrlConfig" xml:"NotifyUrlConfig"`
		RedirectTypeConfig struct {
			RedirectType string `json:"RedirectType" xml:"RedirectType"`
		} `json:"RedirectTypeConfig" xml:"RedirectTypeConfig"`
		ForwardSchemeConfig struct {
			ConfigId         string `json:"ConfigId" xml:"ConfigId"`
			Enable           string `json:"Enable" xml:"Enable"`
			SchemeOrigin     string `json:"SchemeOrigin" xml:"SchemeOrigin"`
			SchemeOriginPort string `json:"SchemeOriginPort" xml:"SchemeOriginPort"`
			Status           string `json:"Status" xml:"Status"`
		} `json:"ForwardSchemeConfig" xml:"ForwardSchemeConfig"`
		RemoveQueryStringConfig struct {
			AliRemoveArgs string `json:"AliRemoveArgs" xml:"AliRemoveArgs"`
			ConfigId      string `json:"ConfigId" xml:"ConfigId"`
			Status        string `json:"Status" xml:"Status"`
		} `json:"RemoveQueryStringConfig" xml:"RemoveQueryStringConfig"`
		L2OssKeyConfig struct {
			PrivateOssAuth string `json:"PrivateOssAuth" xml:"PrivateOssAuth"`
			ConfigId       string `json:"ConfigId" xml:"ConfigId"`
			Status         string `json:"Status" xml:"Status"`
		} `json:"L2OssKeyConfig" xml:"L2OssKeyConfig"`
		MacServiceConfig struct {
			AppList       string `json:"AppList" xml:"AppList"`
			Enabled       string `json:"Enabled" xml:"Enabled"`
			ProcessResult string `json:"ProcessResult" xml:"ProcessResult"`
			ConfigId      string `json:"ConfigId" xml:"ConfigId"`
			Status        string `json:"Status" xml:"Status"`
		} `json:"MacServiceConfig" xml:"MacServiceConfig"`
		GreenManagerConfig struct {
			Enabled  string `json:"Enabled" xml:"Enabled"`
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"GreenManagerConfig" xml:"GreenManagerConfig"`
		HttpsOptionConfig struct {
			Http2    string `json:"Http2" xml:"Http2"`
			ConfigId string `json:"ConfigId" xml:"ConfigId"`
			Status   string `json:"Status" xml:"Status"`
		} `json:"HttpsOptionConfig" xml:"HttpsOptionConfig"`
		AliBusinessConfig struct {
			AliBusinessTable string `json:"AliBusinessTable" xml:"AliBusinessTable"`
			AliBusinessType  string `json:"AliBusinessType" xml:"AliBusinessType"`
			ConfigId         string `json:"ConfigId" xml:"ConfigId"`
			Status           string `json:"Status" xml:"Status"`
		} `json:"AliBusinessConfig" xml:"AliBusinessConfig"`
		IpAllowListConfig struct {
			ConfigId  string `json:"ConfigId" xml:"ConfigId"`
			IpList    string `json:"IpList" xml:"IpList"`
			IpAclXfwd string `json:"IpAclXfwd" xml:"IpAclXfwd"`
			Status    string `json:"Status" xml:"Status"`
		} `json:"IpAllowListConfig" xml:"IpAllowListConfig"`
		CacheExpiredConfigs struct {
			CacheExpiredConfig []struct {
				ConfigId     string `json:"ConfigId" xml:"ConfigId"`
				CacheType    string `json:"CacheType" xml:"CacheType"`
				CacheContent string `json:"CacheContent" xml:"CacheContent"`
				TTL          string `json:"TTL" xml:"TTL"`
				Weight       string `json:"Weight" xml:"Weight"`
				Status       string `json:"Status" xml:"Status"`
			} `json:"CacheExpiredConfig" xml:"CacheExpiredConfig"`
		} `json:"CacheExpiredConfigs" xml:"CacheExpiredConfigs"`
		HttpErrorPageConfigs struct {
			HttpErrorPageConfig []struct {
				ConfigId  string `json:"ConfigId" xml:"ConfigId"`
				ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
				PageUrl   string `json:"PageUrl" xml:"PageUrl"`
				Status    string `json:"Status" xml:"Status"`
			} `json:"HttpErrorPageConfig" xml:"HttpErrorPageConfig"`
		} `json:"HttpErrorPageConfigs" xml:"HttpErrorPageConfigs"`
		HttpHeaderConfigs struct {
			HttpHeaderConfig []struct {
				ConfigId    string `json:"ConfigId" xml:"ConfigId"`
				HeaderKey   string `json:"HeaderKey" xml:"HeaderKey"`
				HeaderValue string `json:"HeaderValue" xml:"HeaderValue"`
				Status      string `json:"Status" xml:"Status"`
			} `json:"HttpHeaderConfig" xml:"HttpHeaderConfig"`
		} `json:"HttpHeaderConfigs" xml:"HttpHeaderConfigs"`
		DynamicConfigs struct {
			DynamicConfig []struct {
				ConfigId            string `json:"ConfigId" xml:"ConfigId"`
				DynamicOrigin       string `json:"DynamicOrigin" xml:"DynamicOrigin"`
				StaticType          string `json:"StaticType" xml:"StaticType"`
				StaticUri           string `json:"StaticUri" xml:"StaticUri"`
				StaticPath          string `json:"StaticPath" xml:"StaticPath"`
				DynamicCacheControl string `json:"DynamicCacheControl" xml:"DynamicCacheControl"`
				Status              string `json:"Status" xml:"Status"`
			} `json:"DynamicConfig" xml:"DynamicConfig"`
		} `json:"DynamicConfigs" xml:"DynamicConfigs"`
		ReqHeaderConfigs struct {
			ReqHeaderConfig []struct {
				ConfigId string `json:"ConfigId" xml:"ConfigId"`
				Key      string `json:"Key" xml:"Key"`
				Value    string `json:"Value" xml:"Value"`
				Status   string `json:"Status" xml:"Status"`
			} `json:"ReqHeaderConfig" xml:"ReqHeaderConfig"`
		} `json:"ReqHeaderConfigs" xml:"ReqHeaderConfigs"`
		SetVarsConfigs struct {
			SetVarsConfig []struct {
				ConfigId string `json:"ConfigId" xml:"ConfigId"`
				VarName  string `json:"VarName" xml:"VarName"`
				VarValue string `json:"VarValue" xml:"VarValue"`
				Status   string `json:"Status" xml:"Status"`
			} `json:"SetVarsConfig" xml:"SetVarsConfig"`
		} `json:"SetVarsConfigs" xml:"SetVarsConfigs"`
	} `json:"DomainConfigs" xml:"DomainConfigs"`
}

func CreateDescribeDomainConfigsRequest() (request *DescribeDomainConfigsRequest) {
	request = &DescribeDomainConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainConfigs", "", "")
	return
}

func CreateDescribeDomainConfigsResponse() (response *DescribeDomainConfigsResponse) {
	response = &DescribeDomainConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
