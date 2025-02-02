// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200324

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-03-24"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewApplyInstanceSnapshotRequest() (request *ApplyInstanceSnapshotRequest) {
    request = &ApplyInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ApplyInstanceSnapshot")
    return
}

func NewApplyInstanceSnapshotResponse() (response *ApplyInstanceSnapshotResponse) {
    response = &ApplyInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ApplyInstanceSnapshot）用于回滚指定实例的系统盘快照。
// <li>仅支持回滚到原系统盘。</li>
// <li>用于回滚的快照必须处于 NORMAL 状态。快照状态可以通 DescribeSnapshots 接口查询，见输出参数中 SnapshotState 字段解释。</li>
// <li>回滚快照时，实例的状态必须为 STOPPED 或 RUNNING，可通过 DescribeInstances 接口查询实例状态。处于 RUNNING 状态的实例会强制关机，然后回滚快照。</li>
func (c *Client) ApplyInstanceSnapshot(request *ApplyInstanceSnapshotRequest) (response *ApplyInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewApplyInstanceSnapshotRequest()
    }
    response = NewApplyInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewCreateBlueprintRequest() (request *CreateBlueprintRequest) {
    request = &CreateBlueprintRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateBlueprint")
    return
}

func NewCreateBlueprintResponse() (response *CreateBlueprintResponse) {
    response = &CreateBlueprintResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (CreateBlueprint) 用于创建镜像。
func (c *Client) CreateBlueprint(request *CreateBlueprintRequest) (response *CreateBlueprintResponse, err error) {
    if request == nil {
        request = NewCreateBlueprintRequest()
    }
    response = NewCreateBlueprintResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFirewallRulesRequest() (request *CreateFirewallRulesRequest) {
    request = &CreateFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateFirewallRules")
    return
}

func NewCreateFirewallRulesResponse() (response *CreateFirewallRulesResponse) {
    response = &CreateFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateFirewallRules）用于在实例上添加防火墙规则。
// 
// 
// * FirewallVersion 为防火墙版本号，用户每次更新防火墙规则版本会自动加1，防止您更新的规则已过期，不填不考虑冲突。
// 
// 在 FirewallRules 参数中：
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
// * Action 字段只允许输入 ACCEPT 或 DROP。
// * FirewallRuleDescription 字段长度不得超过 64。
func (c *Client) CreateFirewallRules(request *CreateFirewallRulesRequest) (response *CreateFirewallRulesResponse, err error) {
    if request == nil {
        request = NewCreateFirewallRulesRequest()
    }
    response = NewCreateFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateInstanceSnapshotRequest() (request *CreateInstanceSnapshotRequest) {
    request = &CreateInstanceSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "CreateInstanceSnapshot")
    return
}

func NewCreateInstanceSnapshotResponse() (response *CreateInstanceSnapshotResponse) {
    response = &CreateInstanceSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（CreateInstanceSnapshot）用于创建指定实例的系统盘快照。
func (c *Client) CreateInstanceSnapshot(request *CreateInstanceSnapshotRequest) (response *CreateInstanceSnapshotResponse, err error) {
    if request == nil {
        request = NewCreateInstanceSnapshotRequest()
    }
    response = NewCreateInstanceSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteBlueprintsRequest() (request *DeleteBlueprintsRequest) {
    request = &DeleteBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteBlueprints")
    return
}

func NewDeleteBlueprintsResponse() (response *DeleteBlueprintsResponse) {
    response = &DeleteBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DeleteBlueprints) 用于删除镜像。
func (c *Client) DeleteBlueprints(request *DeleteBlueprintsRequest) (response *DeleteBlueprintsResponse, err error) {
    if request == nil {
        request = NewDeleteBlueprintsRequest()
    }
    response = NewDeleteBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFirewallRulesRequest() (request *DeleteFirewallRulesRequest) {
    request = &DeleteFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteFirewallRules")
    return
}

func NewDeleteFirewallRulesResponse() (response *DeleteFirewallRulesResponse) {
    response = &DeleteFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteFirewallRules）用于删除实例的防火墙规则。
// 
// * FirewallVersion 用于指定要操作的防火墙的版本。传入 FirewallVersion 版本号若不等于当前防火墙的最新版本，将返回失败；若不传 FirewallVersion 则直接删除指定的规则。
// 
// 在 FirewallRules 参数中：
// * Protocol 字段支持输入 TCP，UDP，ICMP，ALL。
// * Port 字段允许输入 ALL，或者一个单独的端口号，或者用逗号分隔的离散端口号，或者用减号分隔的两个端口号代表的端口范围。当 Port 为范围时，减号分隔的第一个端口号小于第二个端口号。当 Protocol 字段不是 TCP 或 UDP 时，Port 字段只能为空或 ALL。Port 字段长度不得超过 64。
// * CidrBlock 字段允许输入符合 cidr 格式标准的任意字符串。租户之间网络隔离规则优先于防火墙中的内网规则。
// * Action 字段只允许输入 ACCEPT 或 DROP。
// * FirewallRuleDescription 字段长度不得超过 64。
func (c *Client) DeleteFirewallRules(request *DeleteFirewallRulesRequest) (response *DeleteFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDeleteFirewallRulesRequest()
    }
    response = NewDeleteFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteSnapshotsRequest() (request *DeleteSnapshotsRequest) {
    request = &DeleteSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DeleteSnapshots")
    return
}

func NewDeleteSnapshotsResponse() (response *DeleteSnapshotsResponse) {
    response = &DeleteSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DeleteSnapshots）用于删除快照。
// 快照必须处于 NORMAL 状态，快照状态可以通过 DescribeSnapshots 接口查询，见输出参数中 SnapshotState 字段解释。
func (c *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (response *DeleteSnapshotsResponse, err error) {
    if request == nil {
        request = NewDeleteSnapshotsRequest()
    }
    response = NewDeleteSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBlueprintsRequest() (request *DescribeBlueprintsRequest) {
    request = &DescribeBlueprintsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBlueprints")
    return
}

func NewDescribeBlueprintsResponse() (response *DescribeBlueprintsResponse) {
    response = &DescribeBlueprintsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBlueprints）用于查询镜像信息。
func (c *Client) DescribeBlueprints(request *DescribeBlueprintsRequest) (response *DescribeBlueprintsResponse, err error) {
    if request == nil {
        request = NewDescribeBlueprintsRequest()
    }
    response = NewDescribeBlueprintsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeBundlesRequest() (request *DescribeBundlesRequest) {
    request = &DescribeBundlesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeBundles")
    return
}

func NewDescribeBundlesResponse() (response *DescribeBundlesResponse) {
    response = &DescribeBundlesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeBundles）用于查询套餐信息。
func (c *Client) DescribeBundles(request *DescribeBundlesRequest) (response *DescribeBundlesResponse, err error) {
    if request == nil {
        request = NewDescribeBundlesRequest()
    }
    response = NewDescribeBundlesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFirewallRulesRequest() (request *DescribeFirewallRulesRequest) {
    request = &DescribeFirewallRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeFirewallRules")
    return
}

func NewDescribeFirewallRulesResponse() (response *DescribeFirewallRulesResponse) {
    response = &DescribeFirewallRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeFirewallRules）用于查询实例的防火墙规则。
func (c *Client) DescribeFirewallRules(request *DescribeFirewallRulesRequest) (response *DescribeFirewallRulesResponse, err error) {
    if request == nil {
        request = NewDescribeFirewallRulesRequest()
    }
    response = NewDescribeFirewallRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
    request = &DescribeInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstances")
    return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
    response = &DescribeInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstances）用于查询一个或多个实例的详细信息。
// 
// * 可以根据实例 ID、实例名称或者实例的内网 IP 查询实例的详细信息。
// * 过滤信息详细请见过滤器 [Filters](https://cloud.tencent.com/document/product/1207/47576#Filter) 。
// * 如果参数为空，返回当前用户一定数量（Limit 所指定的数量，默认为 20）的实例。
// * 支持查询实例的最新操作（LatestOperation）以及最新操作状态（LatestOperationState）。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesRequest()
    }
    response = NewDescribeInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeInstancesTrafficPackagesRequest() (request *DescribeInstancesTrafficPackagesRequest) {
    request = &DescribeInstancesTrafficPackagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeInstancesTrafficPackages")
    return
}

func NewDescribeInstancesTrafficPackagesResponse() (response *DescribeInstancesTrafficPackagesResponse) {
    response = &DescribeInstancesTrafficPackagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeInstancesTrafficPackages）用于查询一个或多个实例的流量包详情。
func (c *Client) DescribeInstancesTrafficPackages(request *DescribeInstancesTrafficPackagesRequest) (response *DescribeInstancesTrafficPackagesResponse, err error) {
    if request == nil {
        request = NewDescribeInstancesTrafficPackagesRequest()
    }
    response = NewDescribeInstancesTrafficPackagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeSnapshotsRequest() (request *DescribeSnapshotsRequest) {
    request = &DescribeSnapshotsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "DescribeSnapshots")
    return
}

func NewDescribeSnapshotsResponse() (response *DescribeSnapshotsResponse) {
    response = &DescribeSnapshotsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（DescribeSnapshots）用于查询快照的详细信息。
func (c *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (response *DescribeSnapshotsResponse, err error) {
    if request == nil {
        request = NewDescribeSnapshotsRequest()
    }
    response = NewDescribeSnapshotsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyBlueprintAttributeRequest() (request *ModifyBlueprintAttributeRequest) {
    request = &ModifyBlueprintAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifyBlueprintAttribute")
    return
}

func NewModifyBlueprintAttributeResponse() (response *ModifyBlueprintAttributeResponse) {
    response = &ModifyBlueprintAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (ModifyBlueprintAttribute) 用于修改镜像属性。
func (c *Client) ModifyBlueprintAttribute(request *ModifyBlueprintAttributeRequest) (response *ModifyBlueprintAttributeResponse, err error) {
    if request == nil {
        request = NewModifyBlueprintAttributeRequest()
    }
    response = NewModifyBlueprintAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewModifySnapshotAttributeRequest() (request *ModifySnapshotAttributeRequest) {
    request = &ModifySnapshotAttributeRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ModifySnapshotAttribute")
    return
}

func NewModifySnapshotAttributeResponse() (response *ModifySnapshotAttributeResponse) {
    response = &ModifySnapshotAttributeResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ModifySnapshotAttribute）用于修改指定快照的属性。
// <li>“快照名称”仅为方便用户自己管理之用，腾讯云并不以此名称作为提交工单或是进行快照管理操作的依据。</li>
func (c *Client) ModifySnapshotAttribute(request *ModifySnapshotAttributeRequest) (response *ModifySnapshotAttributeResponse, err error) {
    if request == nil {
        request = NewModifySnapshotAttributeRequest()
    }
    response = NewModifySnapshotAttributeResponse()
    err = c.Send(request, response)
    return
}

func NewRebootInstancesRequest() (request *RebootInstancesRequest) {
    request = &RebootInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "RebootInstances")
    return
}

func NewRebootInstancesResponse() (response *RebootInstancesResponse) {
    response = &RebootInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（RebootInstances）用于重启实例。
// 
// * 只有状态为 RUNNING 的实例才可以进行此操作。
// * 接口调用成功时，实例会进入 REBOOTING 状态；重启实例成功时，实例会进入 RUNNING 状态。
// * 支持批量操作，每次请求批量实例的上限为 100。
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
func (c *Client) RebootInstances(request *RebootInstancesRequest) (response *RebootInstancesResponse, err error) {
    if request == nil {
        request = NewRebootInstancesRequest()
    }
    response = NewRebootInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewResetInstanceRequest() (request *ResetInstanceRequest) {
    request = &ResetInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "ResetInstance")
    return
}

func NewResetInstanceResponse() (response *ResetInstanceResponse) {
    response = &ResetInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（ResetInstance）用于重装指定实例上的镜像。
// 
// * 如果指定了 BlueprintId 参数，则使用指定的镜像重装；否则按照当前实例使用的镜像进行重装。
// * 系统盘将会被格式化，并重置；请确保系统盘中无重要文件。
// * 目前不支持实例使用该接口实现 LINUX_UNIX 和 WINDOWS 操作系统切换。
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
func (c *Client) ResetInstance(request *ResetInstanceRequest) (response *ResetInstanceResponse, err error) {
    if request == nil {
        request = NewResetInstanceRequest()
    }
    response = NewResetInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewStartInstancesRequest() (request *StartInstancesRequest) {
    request = &StartInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "StartInstances")
    return
}

func NewStartInstancesResponse() (response *StartInstancesResponse) {
    response = &StartInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StartInstances）用于启动一个或多个实例。
// 
// * 只有状态为 STOPPED 的实例才可以进行此操作。
// * 接口调用成功时，实例会进入 STARTING 状态；启动实例成功时，实例会进入 RUNNING 状态。
// * 支持批量操作。每次请求批量实例的上限为 100。
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
func (c *Client) StartInstances(request *StartInstancesRequest) (response *StartInstancesResponse, err error) {
    if request == nil {
        request = NewStartInstancesRequest()
    }
    response = NewStartInstancesResponse()
    err = c.Send(request, response)
    return
}

func NewStopInstancesRequest() (request *StopInstancesRequest) {
    request = &StopInstancesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("lighthouse", APIVersion, "StopInstances")
    return
}

func NewStopInstancesResponse() (response *StopInstancesResponse) {
    response = &StopInstancesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口（StopInstances）用于关闭一个或多个实例。
// * 只有状态为 RUNNING 的实例才可以进行此操作。
// * 接口调用成功时，实例会进入 STOPPING 状态；关闭实例成功时，实例会进入 STOPPED 状态。
// * 支持批量操作。每次请求批量实例的上限为 100。
// * 本接口为异步接口，请求发送成功后会返回一个 RequestId，此时操作并未立即完成。实例操作结果可以通过调用 DescribeInstances 接口查询，如果实例的最新操作状态（LatestOperationState）为“SUCCESS”，则代表操作成功。
func (c *Client) StopInstances(request *StopInstancesRequest) (response *StopInstancesResponse, err error) {
    if request == nil {
        request = NewStopInstancesRequest()
    }
    response = NewStopInstancesResponse()
    err = c.Send(request, response)
    return
}
