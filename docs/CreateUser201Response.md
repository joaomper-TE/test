# CreateUser201Response

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
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewCreateUser201Response

`func NewCreateUser201Response() *CreateUser201Response`

NewCreateUser201Response instantiates a new CreateUser201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUser201ResponseWithDefaults

`func NewCreateUser201ResponseWithDefaults() *CreateUser201Response`

NewCreateUser201ResponseWithDefaults instantiates a new CreateUser201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUser201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUser201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUser201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUser201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUser201Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUser201Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUser201Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUser201Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *CreateUser201Response) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *CreateUser201Response) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *CreateUser201Response) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *CreateUser201Response) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDateRegistered

`func (o *CreateUser201Response) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *CreateUser201Response) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *CreateUser201Response) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *CreateUser201Response) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetLoginAccountGroup

`func (o *CreateUser201Response) GetLoginAccountGroup() AccountGroupUser`

GetLoginAccountGroup returns the LoginAccountGroup field if non-nil, zero value otherwise.

### GetLoginAccountGroupOk

`func (o *CreateUser201Response) GetLoginAccountGroupOk() (*AccountGroupUser, bool)`

GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroup

`func (o *CreateUser201Response) SetLoginAccountGroup(v AccountGroupUser)`

SetLoginAccountGroup sets LoginAccountGroup field to given value.

### HasLoginAccountGroup

`func (o *CreateUser201Response) HasLoginAccountGroup() bool`

HasLoginAccountGroup returns a boolean if a field has been set.

### GetAccountGroupRoles

`func (o *CreateUser201Response) GetAccountGroupRoles() []AccountGroupRolesAccountGroupRolesInner`

GetAccountGroupRoles returns the AccountGroupRoles field if non-nil, zero value otherwise.

### GetAccountGroupRolesOk

`func (o *CreateUser201Response) GetAccountGroupRolesOk() (*[]AccountGroupRolesAccountGroupRolesInner, bool)`

GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupRoles

`func (o *CreateUser201Response) SetAccountGroupRoles(v []AccountGroupRolesAccountGroupRolesInner)`

SetAccountGroupRoles sets AccountGroupRoles field to given value.

### HasAccountGroupRoles

`func (o *CreateUser201Response) HasAccountGroupRoles() bool`

HasAccountGroupRoles returns a boolean if a field has been set.

### GetAllAccountGroupRoles

`func (o *CreateUser201Response) GetAllAccountGroupRoles() []Role`

GetAllAccountGroupRoles returns the AllAccountGroupRoles field if non-nil, zero value otherwise.

### GetAllAccountGroupRolesOk

`func (o *CreateUser201Response) GetAllAccountGroupRolesOk() (*[]Role, bool)`

GetAllAccountGroupRolesOk returns a tuple with the AllAccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAccountGroupRoles

`func (o *CreateUser201Response) SetAllAccountGroupRoles(v []Role)`

SetAllAccountGroupRoles sets AllAccountGroupRoles field to given value.

### HasAllAccountGroupRoles

`func (o *CreateUser201Response) HasAllAccountGroupRoles() bool`

HasAllAccountGroupRoles returns a boolean if a field has been set.

### GetLinks

`func (o *CreateUser201Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateUser201Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateUser201Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateUser201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


