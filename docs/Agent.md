# Agent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddresses** | Pointer to **[]string** | Array of private IP addresses. | [optional] [readonly] 
**PublicIpAddresses** | Pointer to **[]string** | Array of public IP addresses. | [optional] [readonly] 
**Network** | Pointer to **string** | Network (including ASN) of agent’s public IP. | [optional] [readonly] 
**AgentId** | Pointer to **string** | Unique ID of the agent. | [optional] [readonly] 
**AgentName** | Pointer to **string** | Name of the agent. | [optional] 
**AgentType** | Pointer to **string** | Type of the agent. | [optional] [readonly] 
**Location** | Pointer to **string** | Location of the agent. | [optional] [readonly] 
**CountryId** | Pointer to **string** | 2-digit ISO country code | [optional] [readonly] 
**Enabled** | Pointer to **bool** | Flag indicating if the agent is enabled. | [optional] 
**VerifySslCertificates** | Pointer to **bool** | Flag indicating if has normal SSL operations or  if instead it&#39;s set to ignore SSL errors on browserbot-based tests. | [optional] [readonly] 

## Methods

### NewAgent

`func NewAgent() *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *Agent) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Agent) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Agent) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *Agent) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetPublicIpAddresses

`func (o *Agent) GetPublicIpAddresses() []string`

GetPublicIpAddresses returns the PublicIpAddresses field if non-nil, zero value otherwise.

### GetPublicIpAddressesOk

`func (o *Agent) GetPublicIpAddressesOk() (*[]string, bool)`

GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddresses

`func (o *Agent) SetPublicIpAddresses(v []string)`

SetPublicIpAddresses sets PublicIpAddresses field to given value.

### HasPublicIpAddresses

`func (o *Agent) HasPublicIpAddresses() bool`

HasPublicIpAddresses returns a boolean if a field has been set.

### GetNetwork

`func (o *Agent) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Agent) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Agent) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Agent) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAgentId

`func (o *Agent) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *Agent) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *Agent) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *Agent) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentName

`func (o *Agent) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *Agent) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *Agent) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *Agent) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetAgentType

`func (o *Agent) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *Agent) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *Agent) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *Agent) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetLocation

`func (o *Agent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Agent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Agent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Agent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountryId

`func (o *Agent) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *Agent) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *Agent) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *Agent) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetEnabled

`func (o *Agent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Agent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Agent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Agent) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVerifySslCertificates

`func (o *Agent) GetVerifySslCertificates() bool`

GetVerifySslCertificates returns the VerifySslCertificates field if non-nil, zero value otherwise.

### GetVerifySslCertificatesOk

`func (o *Agent) GetVerifySslCertificatesOk() (*bool, bool)`

GetVerifySslCertificatesOk returns a tuple with the VerifySslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySslCertificates

`func (o *Agent) SetVerifySslCertificates(v bool)`

SetVerifySslCertificates sets VerifySslCertificates field to given value.

### HasVerifySslCertificates

`func (o *Agent) HasVerifySslCertificates() bool`

HasVerifySslCertificates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


