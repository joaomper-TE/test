# UserRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | User&#39;s display name. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. | [optional] 
**LoginAccountGroupId** | Pointer to **string** | Unique ID of the login account group. | [optional] 
**AccountGroupRoles** | Pointer to [**[]AccountGroupRolesRequestBodyInner**](AccountGroupRolesRequestBodyInner.md) |  | [optional] 
**AllAccountGroupRoleIds** | Pointer to **[]string** | Unique IDs representing the roles. | [optional] 

## Methods

### NewUserRequestBody

`func NewUserRequestBody() *UserRequestBody`

NewUserRequestBody instantiates a new UserRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestBodyWithDefaults

`func NewUserRequestBodyWithDefaults() *UserRequestBody`

NewUserRequestBodyWithDefaults instantiates a new UserRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UserRequestBody) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRequestBody) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRequestBody) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRequestBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLoginAccountGroupId

`func (o *UserRequestBody) GetLoginAccountGroupId() string`

GetLoginAccountGroupId returns the LoginAccountGroupId field if non-nil, zero value otherwise.

### GetLoginAccountGroupIdOk

`func (o *UserRequestBody) GetLoginAccountGroupIdOk() (*string, bool)`

GetLoginAccountGroupIdOk returns a tuple with the LoginAccountGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAccountGroupId

`func (o *UserRequestBody) SetLoginAccountGroupId(v string)`

SetLoginAccountGroupId sets LoginAccountGroupId field to given value.

### HasLoginAccountGroupId

`func (o *UserRequestBody) HasLoginAccountGroupId() bool`

HasLoginAccountGroupId returns a boolean if a field has been set.

### GetAccountGroupRoles

`func (o *UserRequestBody) GetAccountGroupRoles() []AccountGroupRolesRequestBodyInner`

GetAccountGroupRoles returns the AccountGroupRoles field if non-nil, zero value otherwise.

### GetAccountGroupRolesOk

`func (o *UserRequestBody) GetAccountGroupRolesOk() (*[]AccountGroupRolesRequestBodyInner, bool)`

GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupRoles

`func (o *UserRequestBody) SetAccountGroupRoles(v []AccountGroupRolesRequestBodyInner)`

SetAccountGroupRoles sets AccountGroupRoles field to given value.

### HasAccountGroupRoles

`func (o *UserRequestBody) HasAccountGroupRoles() bool`

HasAccountGroupRoles returns a boolean if a field has been set.

### GetAllAccountGroupRoleIds

`func (o *UserRequestBody) GetAllAccountGroupRoleIds() []string`

GetAllAccountGroupRoleIds returns the AllAccountGroupRoleIds field if non-nil, zero value otherwise.

### GetAllAccountGroupRoleIdsOk

`func (o *UserRequestBody) GetAllAccountGroupRoleIdsOk() (*[]string, bool)`

GetAllAccountGroupRoleIdsOk returns a tuple with the AllAccountGroupRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAccountGroupRoleIds

`func (o *UserRequestBody) SetAllAccountGroupRoleIds(v []string)`

SetAllAccountGroupRoleIds sets AllAccountGroupRoleIds field to given value.

### HasAllAccountGroupRoleIds

`func (o *UserRequestBody) HasAllAccountGroupRoleIds() bool`

HasAllAccountGroupRoleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


