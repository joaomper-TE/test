# Permission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | Label corresponding to the permission. | [optional] 
**PermissionId** | Pointer to **string** | Unique ID representing the permission. | [optional] 
**IsManagementPermission** | Pointer to **bool** | Flag indicating whether the permission is classified as a management permission. | [optional] 
**Permission** | Pointer to **string** | Permission name | [optional] 

## Methods

### NewPermission

`func NewPermission() *Permission`

NewPermission instantiates a new Permission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionWithDefaults

`func NewPermissionWithDefaults() *Permission`

NewPermissionWithDefaults instantiates a new Permission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *Permission) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Permission) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Permission) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Permission) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPermissionId

`func (o *Permission) GetPermissionId() string`

GetPermissionId returns the PermissionId field if non-nil, zero value otherwise.

### GetPermissionIdOk

`func (o *Permission) GetPermissionIdOk() (*string, bool)`

GetPermissionIdOk returns a tuple with the PermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionId

`func (o *Permission) SetPermissionId(v string)`

SetPermissionId sets PermissionId field to given value.

### HasPermissionId

`func (o *Permission) HasPermissionId() bool`

HasPermissionId returns a boolean if a field has been set.

### GetIsManagementPermission

`func (o *Permission) GetIsManagementPermission() bool`

GetIsManagementPermission returns the IsManagementPermission field if non-nil, zero value otherwise.

### GetIsManagementPermissionOk

`func (o *Permission) GetIsManagementPermissionOk() (*bool, bool)`

GetIsManagementPermissionOk returns a tuple with the IsManagementPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManagementPermission

`func (o *Permission) SetIsManagementPermission(v bool)`

SetIsManagementPermission sets IsManagementPermission field to given value.

### HasIsManagementPermission

`func (o *Permission) HasIsManagementPermission() bool`

HasIsManagementPermission returns a boolean if a field has been set.

### GetPermission

`func (o *Permission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *Permission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *Permission) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *Permission) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


