# ClusterMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | Pointer to **[]string** | Array of private IP addresses. | [optional] [readonly] 
**PublicIpAddresses** | Pointer to **[]string** | Array of public IP addresses. | [optional] [readonly] 
**Network** | Pointer to **string** | Network (including ASN) of agent’s public IP. | [optional] [readonly] 
**MemberId** | Pointer to **string** | Unique ID of the cluster member | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the cluster member | [optional] [readonly] 
**ErrorDetails** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | If an enterprise agent or a cluster member presents at least one error, the errors will be shown as an array of entries in the errorDetails field (Enterprise Agents and Enterprise Cluster members only) | [optional] [readonly] 
**LastSeen** | Pointer to **time.Time** | UTC last seen date (ISO date-time format). | [optional] [readonly] 
**AgentState** | Pointer to **string** | State of the agent. | [optional] [readonly] 
**TargetForTests** | Pointer to **string** | Test target IP address. | [optional] 
**Utilization** | Pointer to **int32** | Shows overall utilization percentage (online Enterprise Agents and Enterprise Clusters only). | [optional] [readonly] 

## Methods

### NewClusterMember

`func NewClusterMember() *ClusterMember`

NewClusterMember instantiates a new ClusterMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMemberWithDefaults

`func NewClusterMemberWithDefaults() *ClusterMember`

NewClusterMemberWithDefaults instantiates a new ClusterMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *ClusterMember) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *ClusterMember) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *ClusterMember) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *ClusterMember) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetPublicIpAddresses

`func (o *ClusterMember) GetPublicIpAddresses() []string`

GetPublicIpAddresses returns the PublicIpAddresses field if non-nil, zero value otherwise.

### GetPublicIpAddressesOk

`func (o *ClusterMember) GetPublicIpAddressesOk() (*[]string, bool)`

GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddresses

`func (o *ClusterMember) SetPublicIpAddresses(v []string)`

SetPublicIpAddresses sets PublicIpAddresses field to given value.

### HasPublicIpAddresses

`func (o *ClusterMember) HasPublicIpAddresses() bool`

HasPublicIpAddresses returns a boolean if a field has been set.

### GetNetwork

`func (o *ClusterMember) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterMember) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterMember) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClusterMember) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetMemberId

`func (o *ClusterMember) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *ClusterMember) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *ClusterMember) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *ClusterMember) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetName

`func (o *ClusterMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetErrorDetails

`func (o *ClusterMember) GetErrorDetails() []ErrorDetail`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ClusterMember) GetErrorDetailsOk() (*[]ErrorDetail, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ClusterMember) SetErrorDetails(v []ErrorDetail)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ClusterMember) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetLastSeen

`func (o *ClusterMember) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClusterMember) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClusterMember) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ClusterMember) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetAgentState

`func (o *ClusterMember) GetAgentState() string`

GetAgentState returns the AgentState field if non-nil, zero value otherwise.

### GetAgentStateOk

`func (o *ClusterMember) GetAgentStateOk() (*string, bool)`

GetAgentStateOk returns a tuple with the AgentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentState

`func (o *ClusterMember) SetAgentState(v string)`

SetAgentState sets AgentState field to given value.

### HasAgentState

`func (o *ClusterMember) HasAgentState() bool`

HasAgentState returns a boolean if a field has been set.

### GetTargetForTests

`func (o *ClusterMember) GetTargetForTests() string`

GetTargetForTests returns the TargetForTests field if non-nil, zero value otherwise.

### GetTargetForTestsOk

`func (o *ClusterMember) GetTargetForTestsOk() (*string, bool)`

GetTargetForTestsOk returns a tuple with the TargetForTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetForTests

`func (o *ClusterMember) SetTargetForTests(v string)`

SetTargetForTests sets TargetForTests field to given value.

### HasTargetForTests

`func (o *ClusterMember) HasTargetForTests() bool`

HasTargetForTests returns a boolean if a field has been set.

### GetUtilization

`func (o *ClusterMember) GetUtilization() int32`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *ClusterMember) GetUtilizationOk() (*int32, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *ClusterMember) SetUtilization(v int32)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *ClusterMember) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


