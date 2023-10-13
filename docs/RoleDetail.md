# RoleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**RoleId** | Pointer to **string** | Unique ID representing the role. | [optional] 
**IsBuiltin** | Pointer to **bool** | Flag indicating if the role is built-in (Account Admin, Organization Admin, Regular User). | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) |  | [optional] 

## Methods

### NewRoleDetail

`func NewRoleDetail() *RoleDetail`

NewRoleDetail instantiates a new RoleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleDetailWithDefaults

`func NewRoleDetailWithDefaults() *RoleDetail`

NewRoleDetailWithDefaults instantiates a new RoleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RoleDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *RoleDetail) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *RoleDetail) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *RoleDetail) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *RoleDetail) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetIsBuiltin

`func (o *RoleDetail) GetIsBuiltin() bool`

GetIsBuiltin returns the IsBuiltin field if non-nil, zero value otherwise.

### GetIsBuiltinOk

`func (o *RoleDetail) GetIsBuiltinOk() (*bool, bool)`

GetIsBuiltinOk returns a tuple with the IsBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltin

`func (o *RoleDetail) SetIsBuiltin(v bool)`

SetIsBuiltin sets IsBuiltin field to given value.

### HasIsBuiltin

`func (o *RoleDetail) HasIsBuiltin() bool`

HasIsBuiltin returns a boolean if a field has been set.

### GetPermissions

`func (o *RoleDetail) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RoleDetail) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RoleDetail) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RoleDetail) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


