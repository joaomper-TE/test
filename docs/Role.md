# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**RoleId** | Pointer to **string** | Unique ID representing the role. | [optional] 
**IsBuiltin** | Pointer to **bool** | Flag indicating if the role is built-in (Account Admin, Organization Admin, Regular User). | [optional] 
**HasManagementPermissions** | Pointer to **bool** | Flag indicating whether the user has management permissions. | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Role) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *Role) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *Role) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *Role) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *Role) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetIsBuiltin

`func (o *Role) GetIsBuiltin() bool`

GetIsBuiltin returns the IsBuiltin field if non-nil, zero value otherwise.

### GetIsBuiltinOk

`func (o *Role) GetIsBuiltinOk() (*bool, bool)`

GetIsBuiltinOk returns a tuple with the IsBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltin

`func (o *Role) SetIsBuiltin(v bool)`

SetIsBuiltin sets IsBuiltin field to given value.

### HasIsBuiltin

`func (o *Role) HasIsBuiltin() bool`

HasIsBuiltin returns a boolean if a field has been set.

### GetHasManagementPermissions

`func (o *Role) GetHasManagementPermissions() bool`

GetHasManagementPermissions returns the HasManagementPermissions field if non-nil, zero value otherwise.

### GetHasManagementPermissionsOk

`func (o *Role) GetHasManagementPermissionsOk() (*bool, bool)`

GetHasManagementPermissionsOk returns a tuple with the HasManagementPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasManagementPermissions

`func (o *Role) SetHasManagementPermissions(v bool)`

SetHasManagementPermissions sets HasManagementPermissions field to given value.

### HasHasManagementPermissions

`func (o *Role) HasHasManagementPermissions() bool`

HasHasManagementPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


