# RoleRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRoleRequestBody

`func NewRoleRequestBody() *RoleRequestBody`

NewRoleRequestBody instantiates a new RoleRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleRequestBodyWithDefaults

`func NewRoleRequestBodyWithDefaults() *RoleRequestBody`

NewRoleRequestBodyWithDefaults instantiates a new RoleRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleRequestBody) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleRequestBody) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleRequestBody) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleRequestBody) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


