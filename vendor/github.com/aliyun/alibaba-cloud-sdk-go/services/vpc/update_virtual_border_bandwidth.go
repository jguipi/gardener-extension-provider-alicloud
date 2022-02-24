package vpc

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

// UpdateVirtualBorderBandwidth invokes the vpc.UpdateVirtualBorderBandwidth API synchronously
func (client *Client) UpdateVirtualBorderBandwidth(request *UpdateVirtualBorderBandwidthRequest) (response *UpdateVirtualBorderBandwidthResponse, err error) {
	response = CreateUpdateVirtualBorderBandwidthResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateVirtualBorderBandwidthWithChan invokes the vpc.UpdateVirtualBorderBandwidth API asynchronously
func (client *Client) UpdateVirtualBorderBandwidthWithChan(request *UpdateVirtualBorderBandwidthRequest) (<-chan *UpdateVirtualBorderBandwidthResponse, <-chan error) {
	responseChan := make(chan *UpdateVirtualBorderBandwidthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateVirtualBorderBandwidth(request)
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

// UpdateVirtualBorderBandwidthWithCallback invokes the vpc.UpdateVirtualBorderBandwidth API asynchronously
func (client *Client) UpdateVirtualBorderBandwidthWithCallback(request *UpdateVirtualBorderBandwidthRequest, callback func(response *UpdateVirtualBorderBandwidthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateVirtualBorderBandwidthResponse
		var err error
		defer close(result)
		response, err = client.UpdateVirtualBorderBandwidth(request)
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

// UpdateVirtualBorderBandwidthRequest is the request struct for api UpdateVirtualBorderBandwidth
type UpdateVirtualBorderBandwidthRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	VirtualBorderRouterId string           `position:"Query" name:"VirtualBorderRouterId"`
	Bandwidth             requests.Integer `position:"Query" name:"Bandwidth"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateVirtualBorderBandwidthResponse is the response struct for api UpdateVirtualBorderBandwidth
type UpdateVirtualBorderBandwidthResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode string `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUpdateVirtualBorderBandwidthRequest creates a request to invoke UpdateVirtualBorderBandwidth API
func CreateUpdateVirtualBorderBandwidthRequest() (request *UpdateVirtualBorderBandwidthRequest) {
	request = &UpdateVirtualBorderBandwidthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UpdateVirtualBorderBandwidth", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateVirtualBorderBandwidthResponse creates a response to parse from UpdateVirtualBorderBandwidth response
func CreateUpdateVirtualBorderBandwidthResponse() (response *UpdateVirtualBorderBandwidthResponse) {
	response = &UpdateVirtualBorderBandwidthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
