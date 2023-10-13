# UserAccountGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. | [optional] 
**Uid** | Pointer to **string** | Unique ID representing the user. | [optional] 
**LastLogin** | Pointer to **time.Time** | User&#39;s UTC last login date (ISO date-time format). | [optional] 
**DateRegistered** | Pointer to **time.Time** | User&#39;s UTC registration date (ISO date-time format). | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 

## Methods

### NewUserAccountGroup

`func NewUserAccountGroup() *UserAccountGroup`

NewUserAccountGroup instantiates a new UserAccountGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountGroupWithDefaults

`func NewUserAccountGroupWithDefaults() *UserAccountGroup`

NewUserAccountGroupWithDefaults instantiates a new UserAccountGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserAccountGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAccountGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAccountGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAccountGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserAccountGroup) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserAccountGroup) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserAccountGroup) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserAccountGroup) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *UserAccountGroup) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserAccountGroup) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserAccountGroup) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserAccountGroup) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserAccountGroup) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserAccountGroup) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserAccountGroup) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserAccountGroup) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetDateRegistered

`func (o *UserAccountGroup) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *UserAccountGroup) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *UserAccountGroup) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *UserAccountGroup) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetRoles

`func (o *UserAccountGroup) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserAccountGroup) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserAccountGroup) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserAccountGroup) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


