# AccountGroupDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupName** | Pointer to **string** | Unique name for the account group within an organization. | [optional] 
**Aid** | Pointer to **string** | Unique ID representing the account group. | [optional] 
**IsCurrentAccountGroup** | Pointer to **bool** | Indicates whether the requested aid is the context of the current account. | [optional] 
**IsDefaultAccountGroup** | Pointer to **bool** | Indicates whether the aid is the default one for the requesting user. | [optional] 
**OrganizationName** | Pointer to **string** | (Optional) Indicates whether the aid is the default one for the requesting user. | [optional] 
**Users** | Pointer to [**[]UserAccountGroup**](UserAccountGroup.md) |  | [optional] 
**Agents** | Pointer to [**[]EnterpriseAgent**](EnterpriseAgent.md) |  | [optional] 

## Methods

### NewAccountGroupDetail

`func NewAccountGroupDetail() *AccountGroupDetail`

NewAccountGroupDetail instantiates a new AccountGroupDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupDetailWithDefaults

`func NewAccountGroupDetailWithDefaults() *AccountGroupDetail`

NewAccountGroupDetailWithDefaults instantiates a new AccountGroupDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupName

`func (o *AccountGroupDetail) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *AccountGroupDetail) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *AccountGroupDetail) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *AccountGroupDetail) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAid

`func (o *AccountGroupDetail) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *AccountGroupDetail) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *AccountGroupDetail) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *AccountGroupDetail) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetIsCurrentAccountGroup

`func (o *AccountGroupDetail) GetIsCurrentAccountGroup() bool`

GetIsCurrentAccountGroup returns the IsCurrentAccountGroup field if non-nil, zero value otherwise.

### GetIsCurrentAccountGroupOk

`func (o *AccountGroupDetail) GetIsCurrentAccountGroupOk() (*bool, bool)`

GetIsCurrentAccountGroupOk returns a tuple with the IsCurrentAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentAccountGroup

`func (o *AccountGroupDetail) SetIsCurrentAccountGroup(v bool)`

SetIsCurrentAccountGroup sets IsCurrentAccountGroup field to given value.

### HasIsCurrentAccountGroup

`func (o *AccountGroupDetail) HasIsCurrentAccountGroup() bool`

HasIsCurrentAccountGroup returns a boolean if a field has been set.

### GetIsDefaultAccountGroup

`func (o *AccountGroupDetail) GetIsDefaultAccountGroup() bool`

GetIsDefaultAccountGroup returns the IsDefaultAccountGroup field if non-nil, zero value otherwise.

### GetIsDefaultAccountGroupOk

`func (o *AccountGroupDetail) GetIsDefaultAccountGroupOk() (*bool, bool)`

GetIsDefaultAccountGroupOk returns a tuple with the IsDefaultAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccountGroup

`func (o *AccountGroupDetail) SetIsDefaultAccountGroup(v bool)`

SetIsDefaultAccountGroup sets IsDefaultAccountGroup field to given value.

### HasIsDefaultAccountGroup

`func (o *AccountGroupDetail) HasIsDefaultAccountGroup() bool`

HasIsDefaultAccountGroup returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AccountGroupDetail) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AccountGroupDetail) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AccountGroupDetail) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AccountGroupDetail) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetUsers

`func (o *AccountGroupDetail) GetUsers() []UserAccountGroup`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *AccountGroupDetail) GetUsersOk() (*[]UserAccountGroup, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *AccountGroupDetail) SetUsers(v []UserAccountGroup)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *AccountGroupDetail) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAgents

`func (o *AccountGroupDetail) GetAgents() []EnterpriseAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AccountGroupDetail) GetAgentsOk() (*[]EnterpriseAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AccountGroupDetail) SetAgents(v []EnterpriseAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AccountGroupDetail) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


