# AccountGroupRolesRequestBodyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupId** | Pointer to **string** | Unique ID of the account group. | [optional] 
**RoleIds** | Pointer to **[]string** | Unique role IDs. | [optional] 

## Methods

### NewAccountGroupRolesRequestBodyInner

`func NewAccountGroupRolesRequestBodyInner() *AccountGroupRolesRequestBodyInner`

NewAccountGroupRolesRequestBodyInner instantiates a new AccountGroupRolesRequestBodyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupRolesRequestBodyInnerWithDefaults

`func NewAccountGroupRolesRequestBodyInnerWithDefaults() *AccountGroupRolesRequestBodyInner`

NewAccountGroupRolesRequestBodyInnerWithDefaults instantiates a new AccountGroupRolesRequestBodyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupId

`func (o *AccountGroupRolesRequestBodyInner) GetAccountGroupId() string`

GetAccountGroupId returns the AccountGroupId field if non-nil, zero value otherwise.

### GetAccountGroupIdOk

`func (o *AccountGroupRolesRequestBodyInner) GetAccountGroupIdOk() (*string, bool)`

GetAccountGroupIdOk returns a tuple with the AccountGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupId

`func (o *AccountGroupRolesRequestBodyInner) SetAccountGroupId(v string)`

SetAccountGroupId sets AccountGroupId field to given value.

### HasAccountGroupId

`func (o *AccountGroupRolesRequestBodyInner) HasAccountGroupId() bool`

HasAccountGroupId returns a boolean if a field has been set.

### GetRoleIds

`func (o *AccountGroupRolesRequestBodyInner) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *AccountGroupRolesRequestBodyInner) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *AccountGroupRolesRequestBodyInner) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *AccountGroupRolesRequestBodyInner) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


