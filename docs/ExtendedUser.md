# ExtendedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. | [optional] 
**Uid** | Pointer to **string** | Unique ID of the user. | [optional] 
**DateRegistered** | Pointer to **time.Time** | The date the user registered their account (UTC). | [optional] 
**LoginAccountGroup** | Pointer to [**AccountGroupUser**](AccountGroupUser.md) |  | [optional] 
**LastLogin** | Pointer to **time.Time** | The last login of the user (UTC). | [optional] 

## Methods

### NewExtendedUser

`func NewExtendedUser() *ExtendedUser`

NewExtendedUser instantiates a new ExtendedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedUserWithDefaults

`func NewExtendedUserWithDefaults() *ExtendedUser`

NewExtendedUserWithDefaults instantiates a new ExtendedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtendedUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExtendedUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ExtendedUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExtendedUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExtendedUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ExtendedUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *ExtendedUser) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ExtendedUser) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ExtendedUser) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *ExtendedUser) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDateRegistered

`func (o *ExtendedUser) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *ExtendedUser) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *ExtendedUser) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *ExtendedUser) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetLoginAccountGroup

`func (o *ExtendedUser) GetLoginAccountGroup() AccountGroupUser`

GetLoginAccountGroup returns the LoginAccountGroup field if non-nil, zero value otherwise.

### GetLoginAccountGroupOk

`func (o *ExtendedUser) GetLoginAccountGroupOk() (*AccountGroupUser, bool)`

GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroup

`func (o *ExtendedUser) SetLoginAccountGroup(v AccountGroupUser)`

SetLoginAccountGroup sets LoginAccountGroup field to given value.

### HasLoginAccountGroup

`func (o *ExtendedUser) HasLoginAccountGroup() bool`

HasLoginAccountGroup returns a boolean if a field has been set.

### GetLastLogin

`func (o *ExtendedUser) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *ExtendedUser) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *ExtendedUser) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *ExtendedUser) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


