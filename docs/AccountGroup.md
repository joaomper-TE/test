# AccountGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupName** | Pointer to **string** | Unique name for the account group within an organization. | [optional] 
**Aid** | Pointer to **string** | Unique ID representing the account group. | [optional] 
**IsCurrentAccountGroup** | Pointer to **bool** | Indicates whether the requested aid is the context of the current account. | [optional] 
**IsDefaultAccountGroup** | Pointer to **bool** | Indicates whether the aid is the default one for the requesting user. | [optional] 
**OrganizationName** | Pointer to **string** | (Optional) Indicates whether the aid is the default one for the requesting user. | [optional] 

## Methods

### NewAccountGroup

`func NewAccountGroup() *AccountGroup`

NewAccountGroup instantiates a new AccountGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupWithDefaults

`func NewAccountGroupWithDefaults() *AccountGroup`

NewAccountGroupWithDefaults instantiates a new AccountGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupName

`func (o *AccountGroup) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *AccountGroup) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *AccountGroup) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *AccountGroup) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAid

`func (o *AccountGroup) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *AccountGroup) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *AccountGroup) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *AccountGroup) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetIsCurrentAccountGroup

`func (o *AccountGroup) GetIsCurrentAccountGroup() bool`

GetIsCurrentAccountGroup returns the IsCurrentAccountGroup field if non-nil, zero value otherwise.

### GetIsCurrentAccountGroupOk

`func (o *AccountGroup) GetIsCurrentAccountGroupOk() (*bool, bool)`

GetIsCurrentAccountGroupOk returns a tuple with the IsCurrentAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentAccountGroup

`func (o *AccountGroup) SetIsCurrentAccountGroup(v bool)`

SetIsCurrentAccountGroup sets IsCurrentAccountGroup field to given value.

### HasIsCurrentAccountGroup

`func (o *AccountGroup) HasIsCurrentAccountGroup() bool`

HasIsCurrentAccountGroup returns a boolean if a field has been set.

### GetIsDefaultAccountGroup

`func (o *AccountGroup) GetIsDefaultAccountGroup() bool`

GetIsDefaultAccountGroup returns the IsDefaultAccountGroup field if non-nil, zero value otherwise.

### GetIsDefaultAccountGroupOk

`func (o *AccountGroup) GetIsDefaultAccountGroupOk() (*bool, bool)`

GetIsDefaultAccountGroupOk returns a tuple with the IsDefaultAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccountGroup

`func (o *AccountGroup) SetIsDefaultAccountGroup(v bool)`

SetIsDefaultAccountGroup sets IsDefaultAccountGroup field to given value.

### HasIsDefaultAccountGroup

`func (o *AccountGroup) HasIsDefaultAccountGroup() bool`

HasIsDefaultAccountGroup returns a boolean if a field has been set.

### GetOrganizationName

`func (o *AccountGroup) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *AccountGroup) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *AccountGroup) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *AccountGroup) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


