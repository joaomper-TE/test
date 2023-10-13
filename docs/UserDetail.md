# UserDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. | [optional] 
**Uid** | Pointer to **string** | Unique ID of the user. | [optional] 
**DateRegistered** | Pointer to **time.Time** | The date the user registered their account (UTC). | [optional] 
**LoginAccountGroup** | Pointer to [**AccountGroupUser**](AccountGroupUser.md) |  | [optional] 
**LastLogin** | Pointer to **time.Time** | The last login of the user (UTC). | [optional] 
**AccountGroupRoles** | Pointer to [**[]AccountGroupRolesAccountGroupRolesInner**](AccountGroupRolesAccountGroupRolesInner.md) |  | [optional] 
**AllAccountGroupRoles** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewUserDetail

`func NewUserDetail() *UserDetail`

NewUserDetail instantiates a new UserDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailWithDefaults

`func NewUserDetailWithDefaults() *UserDetail`

NewUserDetailWithDefaults instantiates a new UserDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserDetail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDetail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDetail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserDetail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *UserDetail) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserDetail) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserDetail) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserDetail) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDateRegistered

`func (o *UserDetail) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *UserDetail) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *UserDetail) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *UserDetail) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetLoginAccountGroup

`func (o *UserDetail) GetLoginAccountGroup() AccountGroupUser`

GetLoginAccountGroup returns the LoginAccountGroup field if non-nil, zero value otherwise.

### GetLoginAccountGroupOk

`func (o *UserDetail) GetLoginAccountGroupOk() (*AccountGroupUser, bool)`

GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroup

`func (o *UserDetail) SetLoginAccountGroup(v AccountGroupUser)`

SetLoginAccountGroup sets LoginAccountGroup field to given value.

### HasLoginAccountGroup

`func (o *UserDetail) HasLoginAccountGroup() bool`

HasLoginAccountGroup returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserDetail) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserDetail) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserDetail) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserDetail) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetAccountGroupRoles

`func (o *UserDetail) GetAccountGroupRoles() []AccountGroupRolesAccountGroupRolesInner`

GetAccountGroupRoles returns the AccountGroupRoles field if non-nil, zero value otherwise.

### GetAccountGroupRolesOk

`func (o *UserDetail) GetAccountGroupRolesOk() (*[]AccountGroupRolesAccountGroupRolesInner, bool)`

GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupRoles

`func (o *UserDetail) SetAccountGroupRoles(v []AccountGroupRolesAccountGroupRolesInner)`

SetAccountGroupRoles sets AccountGroupRoles field to given value.

### HasAccountGroupRoles

`func (o *UserDetail) HasAccountGroupRoles() bool`

HasAccountGroupRoles returns a boolean if a field has been set.

### GetAllAccountGroupRoles

`func (o *UserDetail) GetAllAccountGroupRoles() []Role`

GetAllAccountGroupRoles returns the AllAccountGroupRoles field if non-nil, zero value otherwise.

### GetAllAccountGroupRolesOk

`func (o *UserDetail) GetAllAccountGroupRolesOk() (*[]Role, bool)`

GetAllAccountGroupRolesOk returns a tuple with the AllAccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAccountGroupRoles

`func (o *UserDetail) SetAllAccountGroupRoles(v []Role)`

SetAllAccountGroupRoles sets AllAccountGroupRoles field to given value.

### HasAllAccountGroupRoles

`func (o *UserDetail) HasAllAccountGroupRoles() bool`

HasAllAccountGroupRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


