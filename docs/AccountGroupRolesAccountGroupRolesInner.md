# AccountGroupRolesAccountGroupRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroup** | Pointer to [**AccountGroupUser**](AccountGroupUser.md) |  | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewAccountGroupRolesAccountGroupRolesInner

`func NewAccountGroupRolesAccountGroupRolesInner() *AccountGroupRolesAccountGroupRolesInner`

NewAccountGroupRolesAccountGroupRolesInner instantiates a new AccountGroupRolesAccountGroupRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupRolesAccountGroupRolesInnerWithDefaults

`func NewAccountGroupRolesAccountGroupRolesInnerWithDefaults() *AccountGroupRolesAccountGroupRolesInner`

NewAccountGroupRolesAccountGroupRolesInnerWithDefaults instantiates a new AccountGroupRolesAccountGroupRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroup

`func (o *AccountGroupRolesAccountGroupRolesInner) GetAccountGroup() AccountGroupUser`

GetAccountGroup returns the AccountGroup field if non-nil, zero value otherwise.

### GetAccountGroupOk

`func (o *AccountGroupRolesAccountGroupRolesInner) GetAccountGroupOk() (*AccountGroupUser, bool)`

GetAccountGroupOk returns a tuple with the AccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroup

`func (o *AccountGroupRolesAccountGroupRolesInner) SetAccountGroup(v AccountGroupUser)`

SetAccountGroup sets AccountGroup field to given value.

### HasAccountGroup

`func (o *AccountGroupRolesAccountGroupRolesInner) HasAccountGroup() bool`

HasAccountGroup returns a boolean if a field has been set.

### GetRoles

`func (o *AccountGroupRolesAccountGroupRolesInner) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountGroupRolesAccountGroupRolesInner) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountGroupRolesAccountGroupRolesInner) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccountGroupRolesAccountGroupRolesInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


