# EnterpriseAgent

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
**ClusterMembers** | Pointer to [**[]ClusterMember**](ClusterMember.md) | If an enterprise agent is clustered, detailed information about each cluster member will be shown as array entries in the clusterMembers field. This field is not shown for Enterprise Agents in standalone mode, or for Cloud Agents. | [optional] [readonly] 
**Utilization** | Pointer to **int32** | Shows overall utilization percentage (online Enterprise Agents and Enterprise Clusters only). | [optional] [readonly] 
**AccountGroups** | Pointer to [**[]EnterpriseAgentAccountGroup**](EnterpriseAgentAccountGroup.md) | List of account groups. See /accounts-groups to pull a list of account IDs | [optional] 
**Prefix** | Pointer to **string** | Prefix containing agents public IP address. | [optional] [readonly] 
**Ipv6Policy** | Pointer to **string** | IP version policy, (Enterprise Agents and Enterprise Clusters only) | [optional] 
**ErrorDetails** | Pointer to [**[]ErrorDetail**](ErrorDetail.md) | If an enterprise agent or a cluster member presents at least one error, the errors will be shown as an array of entries in the errorDetails field (Enterprise Agents and Enterprise Cluster members only) | [optional] [readonly] 
**Hostname** | Pointer to **string** | Fully qualified domain name of the agent (Enterprise Agents only) | [optional] [readonly] 
**LastSeen** | Pointer to **time.Time** | UTC last seen date (ISO date-time format). | [optional] [readonly] 
**AgentState** | Pointer to **string** | State of the agent. | [optional] [readonly] 
**KeepBrowserCache** | Pointer to **bool** | Flag indicating if the agent retains cache. | [optional] 
**CreatedDate** | Pointer to **time.Time** | UTC Agent creation date (ISO date-time format). | [optional] [readonly] 
**TargetForTests** | Pointer to **string** | Test target IP address. | [optional] 
**LocalResolutionPrefixes** | Pointer to **[]string** | To perform rDNS lookups for public IP ranges, this field represents the public IP ranges. The range must be in CIDR notation; for example, 10.1.1.0/24. Maximum of 5 prefixes allowed (Enterprise Agents and Enterprise Agent clusters only). | [optional] 
**InterfaceIpMappings** | Pointer to [**[]InterfaceIpMapping**](InterfaceIpMapping.md) |  | [optional] [readonly] 

## Methods

### NewEnterpriseAgent

`func NewEnterpriseAgent() *EnterpriseAgent`

NewEnterpriseAgent instantiates a new EnterpriseAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnterpriseAgentWithDefaults

`func NewEnterpriseAgentWithDefaults() *EnterpriseAgent`

NewEnterpriseAgentWithDefaults instantiates a new EnterpriseAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddresses

`func (o *EnterpriseAgent) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *EnterpriseAgent) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *EnterpriseAgent) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *EnterpriseAgent) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetPublicIpAddresses

`func (o *EnterpriseAgent) GetPublicIpAddresses() []string`

GetPublicIpAddresses returns the PublicIpAddresses field if non-nil, zero value otherwise.

### GetPublicIpAddressesOk

`func (o *EnterpriseAgent) GetPublicIpAddressesOk() (*[]string, bool)`

GetPublicIpAddressesOk returns a tuple with the PublicIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpAddresses

`func (o *EnterpriseAgent) SetPublicIpAddresses(v []string)`

SetPublicIpAddresses sets PublicIpAddresses field to given value.

### HasPublicIpAddresses

`func (o *EnterpriseAgent) HasPublicIpAddresses() bool`

HasPublicIpAddresses returns a boolean if a field has been set.

### GetNetwork

`func (o *EnterpriseAgent) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *EnterpriseAgent) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *EnterpriseAgent) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *EnterpriseAgent) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetAgentId

`func (o *EnterpriseAgent) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *EnterpriseAgent) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *EnterpriseAgent) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *EnterpriseAgent) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentName

`func (o *EnterpriseAgent) GetAgentName() string`

GetAgentName returns the AgentName field if non-nil, zero value otherwise.

### GetAgentNameOk

`func (o *EnterpriseAgent) GetAgentNameOk() (*string, bool)`

GetAgentNameOk returns a tuple with the AgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentName

`func (o *EnterpriseAgent) SetAgentName(v string)`

SetAgentName sets AgentName field to given value.

### HasAgentName

`func (o *EnterpriseAgent) HasAgentName() bool`

HasAgentName returns a boolean if a field has been set.

### GetAgentType

`func (o *EnterpriseAgent) GetAgentType() string`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *EnterpriseAgent) GetAgentTypeOk() (*string, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *EnterpriseAgent) SetAgentType(v string)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *EnterpriseAgent) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetLocation

`func (o *EnterpriseAgent) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EnterpriseAgent) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EnterpriseAgent) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *EnterpriseAgent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCountryId

`func (o *EnterpriseAgent) GetCountryId() string`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *EnterpriseAgent) GetCountryIdOk() (*string, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *EnterpriseAgent) SetCountryId(v string)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *EnterpriseAgent) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetEnabled

`func (o *EnterpriseAgent) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EnterpriseAgent) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EnterpriseAgent) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EnterpriseAgent) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVerifySslCertificates

`func (o *EnterpriseAgent) GetVerifySslCertificates() bool`

GetVerifySslCertificates returns the VerifySslCertificates field if non-nil, zero value otherwise.

### GetVerifySslCertificatesOk

`func (o *EnterpriseAgent) GetVerifySslCertificatesOk() (*bool, bool)`

GetVerifySslCertificatesOk returns a tuple with the VerifySslCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySslCertificates

`func (o *EnterpriseAgent) SetVerifySslCertificates(v bool)`

SetVerifySslCertificates sets VerifySslCertificates field to given value.

### HasVerifySslCertificates

`func (o *EnterpriseAgent) HasVerifySslCertificates() bool`

HasVerifySslCertificates returns a boolean if a field has been set.

### GetClusterMembers

`func (o *EnterpriseAgent) GetClusterMembers() []ClusterMember`

GetClusterMembers returns the ClusterMembers field if non-nil, zero value otherwise.

### GetClusterMembersOk

`func (o *EnterpriseAgent) GetClusterMembersOk() (*[]ClusterMember, bool)`

GetClusterMembersOk returns a tuple with the ClusterMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMembers

`func (o *EnterpriseAgent) SetClusterMembers(v []ClusterMember)`

SetClusterMembers sets ClusterMembers field to given value.

### HasClusterMembers

`func (o *EnterpriseAgent) HasClusterMembers() bool`

HasClusterMembers returns a boolean if a field has been set.

### GetUtilization

`func (o *EnterpriseAgent) GetUtilization() int32`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *EnterpriseAgent) GetUtilizationOk() (*int32, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *EnterpriseAgent) SetUtilization(v int32)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *EnterpriseAgent) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetAccountGroups

`func (o *EnterpriseAgent) GetAccountGroups() []EnterpriseAgentAccountGroup`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *EnterpriseAgent) GetAccountGroupsOk() (*[]EnterpriseAgentAccountGroup, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *EnterpriseAgent) SetAccountGroups(v []EnterpriseAgentAccountGroup)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *EnterpriseAgent) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetPrefix

`func (o *EnterpriseAgent) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EnterpriseAgent) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EnterpriseAgent) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *EnterpriseAgent) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetIpv6Policy

`func (o *EnterpriseAgent) GetIpv6Policy() string`

GetIpv6Policy returns the Ipv6Policy field if non-nil, zero value otherwise.

### GetIpv6PolicyOk

`func (o *EnterpriseAgent) GetIpv6PolicyOk() (*string, bool)`

GetIpv6PolicyOk returns a tuple with the Ipv6Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Policy

`func (o *EnterpriseAgent) SetIpv6Policy(v string)`

SetIpv6Policy sets Ipv6Policy field to given value.

### HasIpv6Policy

`func (o *EnterpriseAgent) HasIpv6Policy() bool`

HasIpv6Policy returns a boolean if a field has been set.

### GetErrorDetails

`func (o *EnterpriseAgent) GetErrorDetails() []ErrorDetail`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *EnterpriseAgent) GetErrorDetailsOk() (*[]ErrorDetail, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *EnterpriseAgent) SetErrorDetails(v []ErrorDetail)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *EnterpriseAgent) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetHostname

`func (o *EnterpriseAgent) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EnterpriseAgent) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EnterpriseAgent) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EnterpriseAgent) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLastSeen

`func (o *EnterpriseAgent) GetLastSeen() time.Time`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *EnterpriseAgent) GetLastSeenOk() (*time.Time, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *EnterpriseAgent) SetLastSeen(v time.Time)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *EnterpriseAgent) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetAgentState

`func (o *EnterpriseAgent) GetAgentState() string`

GetAgentState returns the AgentState field if non-nil, zero value otherwise.

### GetAgentStateOk

`func (o *EnterpriseAgent) GetAgentStateOk() (*string, bool)`

GetAgentStateOk returns a tuple with the AgentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentState

`func (o *EnterpriseAgent) SetAgentState(v string)`

SetAgentState sets AgentState field to given value.

### HasAgentState

`func (o *EnterpriseAgent) HasAgentState() bool`

HasAgentState returns a boolean if a field has been set.

### GetKeepBrowserCache

`func (o *EnterpriseAgent) GetKeepBrowserCache() bool`

GetKeepBrowserCache returns the KeepBrowserCache field if non-nil, zero value otherwise.

### GetKeepBrowserCacheOk

`func (o *EnterpriseAgent) GetKeepBrowserCacheOk() (*bool, bool)`

GetKeepBrowserCacheOk returns a tuple with the KeepBrowserCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepBrowserCache

`func (o *EnterpriseAgent) SetKeepBrowserCache(v bool)`

SetKeepBrowserCache sets KeepBrowserCache field to given value.

### HasKeepBrowserCache

`func (o *EnterpriseAgent) HasKeepBrowserCache() bool`

HasKeepBrowserCache returns a boolean if a field has been set.

### GetCreatedDate

`func (o *EnterpriseAgent) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *EnterpriseAgent) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *EnterpriseAgent) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *EnterpriseAgent) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetTargetForTests

`func (o *EnterpriseAgent) GetTargetForTests() string`

GetTargetForTests returns the TargetForTests field if non-nil, zero value otherwise.

### GetTargetForTestsOk

`func (o *EnterpriseAgent) GetTargetForTestsOk() (*string, bool)`

GetTargetForTestsOk returns a tuple with the TargetForTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetForTests

`func (o *EnterpriseAgent) SetTargetForTests(v string)`

SetTargetForTests sets TargetForTests field to given value.

### HasTargetForTests

`func (o *EnterpriseAgent) HasTargetForTests() bool`

HasTargetForTests returns a boolean if a field has been set.

### GetLocalResolutionPrefixes

`func (o *EnterpriseAgent) GetLocalResolutionPrefixes() []string`

GetLocalResolutionPrefixes returns the LocalResolutionPrefixes field if non-nil, zero value otherwise.

### GetLocalResolutionPrefixesOk

`func (o *EnterpriseAgent) GetLocalResolutionPrefixesOk() (*[]string, bool)`

GetLocalResolutionPrefixesOk returns a tuple with the LocalResolutionPrefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalResolutionPrefixes

`func (o *EnterpriseAgent) SetLocalResolutionPrefixes(v []string)`

SetLocalResolutionPrefixes sets LocalResolutionPrefixes field to given value.

### HasLocalResolutionPrefixes

`func (o *EnterpriseAgent) HasLocalResolutionPrefixes() bool`

HasLocalResolutionPrefixes returns a boolean if a field has been set.

### GetInterfaceIpMappings

`func (o *EnterpriseAgent) GetInterfaceIpMappings() []InterfaceIpMapping`

GetInterfaceIpMappings returns the InterfaceIpMappings field if non-nil, zero value otherwise.

### GetInterfaceIpMappingsOk

`func (o *EnterpriseAgent) GetInterfaceIpMappingsOk() (*[]InterfaceIpMapping, bool)`

GetInterfaceIpMappingsOk returns a tuple with the InterfaceIpMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIpMappings

`func (o *EnterpriseAgent) SetInterfaceIpMappings(v []InterfaceIpMapping)`

SetInterfaceIpMappings sets InterfaceIpMappings field to given value.

### HasInterfaceIpMappings

`func (o *EnterpriseAgent) HasInterfaceIpMappings() bool`

HasInterfaceIpMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


