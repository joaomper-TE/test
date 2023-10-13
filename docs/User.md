# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. | [optional] 
**Uid** | Pointer to **string** | Unique ID of the user. | [optional] 
**DateRegistered** | Pointer to **time.Time** | The date the user registered their account (UTC). | [optional] 
**LoginAccountGroup** | Pointer to [**AccountGroupUser**](AccountGroupUser.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *User) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *User) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *User) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *User) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDateRegistered

`func (o *User) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *User) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *User) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *User) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetLoginAccountGroup

`func (o *User) GetLoginAccountGroup() AccountGroupUser`

GetLoginAccountGroup returns the LoginAccountGroup field if non-nil, zero value otherwise.

### GetLoginAccountGroupOk

`func (o *User) GetLoginAccountGroupOk() (*AccountGroupUser, bool)`

GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroup

`func (o *User) SetLoginAccountGroup(v AccountGroupUser)`

SetLoginAccountGroup sets LoginAccountGroup field to given value.

### HasLoginAccountGroup

`func (o *User) HasLoginAccountGroup() bool`

HasLoginAccountGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


