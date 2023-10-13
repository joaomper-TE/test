# GetUser200Response

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
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewGetUser200Response

`func NewGetUser200Response() *GetUser200Response`

NewGetUser200Response instantiates a new GetUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUser200ResponseWithDefaults

`func NewGetUser200ResponseWithDefaults() *GetUser200Response`

NewGetUser200ResponseWithDefaults instantiates a new GetUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetUser200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUser200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUser200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUser200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *GetUser200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUser200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUser200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUser200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUid

`func (o *GetUser200Response) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *GetUser200Response) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *GetUser200Response) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *GetUser200Response) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetDateRegistered

`func (o *GetUser200Response) GetDateRegistered() time.Time`

GetDateRegistered returns the DateRegistered field if non-nil, zero value otherwise.

### GetDateRegisteredOk

`func (o *GetUser200Response) GetDateRegisteredOk() (*time.Time, bool)`

GetDateRegisteredOk returns a tuple with the DateRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRegistered

`func (o *GetUser200Response) SetDateRegistered(v time.Time)`

SetDateRegistered sets DateRegistered field to given value.

### HasDateRegistered

`func (o *GetUser200Response) HasDateRegistered() bool`

HasDateRegistered returns a boolean if a field has been set.

### GetLoginAccountGroup

`func (o *GetUser200Response) GetLoginAccountGroup() AccountGroupUser`

GetLoginAccountGroup returns the LoginAccountGroup field if non-nil, zero value otherwise.

### GetLoginAccountGroupOk

`func (o *GetUser200Response) GetLoginAccountGroupOk() (*AccountGroupUser, bool)`

GetLoginAccountGroupOk returns a tuple with the LoginAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroup

`func (o *GetUser200Response) SetLoginAccountGroup(v AccountGroupUser)`

SetLoginAccountGroup sets LoginAccountGroup field to given value.

### HasLoginAccountGroup

`func (o *GetUser200Response) HasLoginAccountGroup() bool`

HasLoginAccountGroup returns a boolean if a field has been set.

### GetLastLogin

`func (o *GetUser200Response) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *GetUser200Response) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *GetUser200Response) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *GetUser200Response) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetAccountGroupRoles

`func (o *GetUser200Response) GetAccountGroupRoles() []AccountGroupRolesAccountGroupRolesInner`

GetAccountGroupRoles returns the AccountGroupRoles field if non-nil, zero value otherwise.

### GetAccountGroupRolesOk

`func (o *GetUser200Response) GetAccountGroupRolesOk() (*[]AccountGroupRolesAccountGroupRolesInner, bool)`

GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupRoles

`func (o *GetUser200Response) SetAccountGroupRoles(v []AccountGroupRolesAccountGroupRolesInner)`

SetAccountGroupRoles sets AccountGroupRoles field to given value.

### HasAccountGroupRoles

`func (o *GetUser200Response) HasAccountGroupRoles() bool`

HasAccountGroupRoles returns a boolean if a field has been set.

### GetAllAccountGroupRoles

`func (o *GetUser200Response) GetAllAccountGroupRoles() []Role`

GetAllAccountGroupRoles returns the AllAccountGroupRoles field if non-nil, zero value otherwise.

### GetAllAccountGroupRolesOk

`func (o *GetUser200Response) GetAllAccountGroupRolesOk() (*[]Role, bool)`

GetAllAccountGroupRolesOk returns a tuple with the AllAccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAccountGroupRoles

`func (o *GetUser200Response) SetAllAccountGroupRoles(v []Role)`

SetAllAccountGroupRoles sets AllAccountGroupRoles field to given value.

### HasAllAccountGroupRoles

`func (o *GetUser200Response) HasAllAccountGroupRoles() bool`

HasAllAccountGroupRoles returns a boolean if a field has been set.

### GetLinks

`func (o *GetUser200Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUser200Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUser200Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUser200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


