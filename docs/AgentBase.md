# AgentBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | Pointer to **[]string** | Array of private IP addresses. | [optional] [readonly] 
**PublicIpAddresses** | Pointer to **[]string** | Array of public IP addresses. | [optional] [readonly] 
**Network** | Pointer to **string** | Network (including ASN) of agent’s public IP. | [optional] [readonly] 

## Methods

### NewAgentBase

`func NewAgentBase() *AgentBase`

NewAgentBase instantiates a new AgentBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBaseWithDefaults

`func NewAgentBaseWithDefaults() *AgentBase`

NewAgentBaseWithDefaults instantiates a new AgentBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *AgentBase) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *AgentBase) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *AgentBase) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *AgentBase) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetPublicIpAddresses

`func (o *AgentBase) GetPublicIpAddresses() []string`

GetPublicIpAddresses returns the PublicIpAddresses field if non-nil, zero value otherwise.

### GetPublicIpAddressesOk

`func (o *AgentBase) GetPublicIpAddressesOk() (*[]string, bool)`

GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddresses

`func (o *AgentBase) SetPublicIpAddresses(v []string)`

SetPublicIpAddresses sets PublicIpAddresses field to given value.

### HasPublicIpAddresses

`func (o *AgentBase) HasPublicIpAddresses() bool`

HasPublicIpAddresses returns a boolean if a field has been set.

### GetNetwork

`func (o *AgentBase) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *AgentBase) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *AgentBase) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *AgentBase) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


