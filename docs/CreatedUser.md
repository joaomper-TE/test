# CreatedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. | [optional] 
**Uid** | Pointer to **string** | Unique ID of the user. | [optional] 
**DateRegistered** | Pointer to **time.Time** | The date the user registered their account (UTC). | [optional] 
**LoginAccountGroup** | Pointer to [**AccountGroupUser**](AccountGroupUser.md) |  | [optional] 
**AccountGroupRoles** | Pointer to [**[]AccountGroupRolesAccountGroupRolesInner**](AccountGroupRolesAccountGroupRolesInner.md) |  | [optional] 
**AllAccountGroupRoles** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewCreatedUser

`func NewCreatedUser() *CreatedUser`

NewCreatedUser instantiates a new CreatedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedUserWithDefaults

`func NewCreatedUserWithDefaults() *CreatedUser`

NewCreatedUserWithDefaults instantiates a new CreatedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatedUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatedUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *CreatedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreatedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreatedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreatedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *CreatedUser) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreatedUser) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreatedUser) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CreatedUser) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDateRegistered

`func (o *CreatedUser) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *CreatedUser) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *CreatedUser) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *CreatedUser) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetLoginAccountGroup

`func (o *CreatedUser) GetLoginAccountGroup() AccountGroupUser`

GetLoginAccountGroup returns the LoginAccountGroup field if non-nil, zero value otherwise.

### GetLoginAccountGroupOk

`func (o *CreatedUser) GetLoginAccountGroupOk() (*AccountGroupUser, bool)`

GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroup

`func (o *CreatedUser) SetLoginAccountGroup(v AccountGroupUser)`

SetLoginAccountGroup sets LoginAccountGroup field to given value.

### HasLoginAccountGroup

`func (o *CreatedUser) HasLoginAccountGroup() bool`

HasLoginAccountGroup returns a boolean if a field has been set.

### GetAccountGroupRoles

`func (o *CreatedUser) GetAccountGroupRoles() []AccountGroupRolesAccountGroupRolesInner`

GetAccountGroupRoles returns the AccountGroupRoles field if non-nil, zero value otherwise.

### GetAccountGroupRolesOk

`func (o *CreatedUser) GetAccountGroupRolesOk() (*[]AccountGroupRolesAccountGroupRolesInner, bool)`

GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupRoles

`func (o *CreatedUser) SetAccountGroupRoles(v []AccountGroupRolesAccountGroupRolesInner)`

SetAccountGroupRoles sets AccountGroupRoles field to given value.

### HasAccountGroupRoles

`func (o *CreatedUser) HasAccountGroupRoles() bool`

HasAccountGroupRoles returns a boolean if a field has been set.

### GetAllAccountGroupRoles

`func (o *CreatedUser) GetAllAccountGroupRoles() []Role`

GetAllAccountGroupRoles returns the AllAccountGroupRoles field if non-nil, zero value otherwise.

### GetAllAccountGroupRolesOk

`func (o *CreatedUser) GetAllAccountGroupRolesOk() (*[]Role, bool)`

GetAllAccountGroupRolesOk returns a tuple with the AllAccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAccountGroupRoles

`func (o *CreatedUser) SetAllAccountGroupRoles(v []Role)`

SetAllAccountGroupRoles sets AllAccountGroupRoles field to given value.

### HasAllAccountGroupRoles

`func (o *CreatedUser) HasAllAccountGroupRoles() bool`

HasAllAccountGroupRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


