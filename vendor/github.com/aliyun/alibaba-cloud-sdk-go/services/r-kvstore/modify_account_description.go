package r_kvstore

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

// ModifyAccountDescription invokes the r_kvstore.ModifyAccountDescription API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyaccountdescription.html
func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (response *ModifyAccountDescriptionResponse, err error) {
	response = CreateModifyAccountDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAccountDescriptionWithChan invokes the r_kvstore.ModifyAccountDescription API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyaccountdescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccountDescriptionWithChan(request *ModifyAccountDescriptionRequest) (<-chan *ModifyAccountDescriptionResponse, <-chan error) {
	responseChan := make(chan *ModifyAccountDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAccountDescription(request)
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

// ModifyAccountDescriptionWithCallback invokes the r_kvstore.ModifyAccountDescription API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/modifyaccountdescription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyAccountDescriptionWithCallback(request *ModifyAccountDescriptionRequest, callback func(response *ModifyAccountDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAccountDescriptionResponse
		var err error
		defer close(result)
		response, err = client.ModifyAccountDescription(request)
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

// ModifyAccountDescriptionRequest is the request struct for api ModifyAccountDescription
type ModifyAccountDescriptionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountDescription   string           `position:"Query" name:"AccountDescription"`
	AccountName          string           `position:"Query" name:"AccountName"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// ModifyAccountDescriptionResponse is the response struct for api ModifyAccountDescription
type ModifyAccountDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAccountDescriptionRequest creates a request to invoke ModifyAccountDescription API
func CreateModifyAccountDescriptionRequest() (request *ModifyAccountDescriptionRequest) {
	request = &ModifyAccountDescriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyAccountDescription", "", "")
	return
}

// CreateModifyAccountDescriptionResponse creates a response to parse from ModifyAccountDescription response
func CreateModifyAccountDescriptionResponse() (response *ModifyAccountDescriptionResponse) {
	response = &ModifyAccountDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
