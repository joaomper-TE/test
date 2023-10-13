# CreateRole201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**RoleId** | Pointer to **string** | Unique ID representing the role. | [optional] 
**IsBuiltin** | Pointer to **bool** | Flag indicating if the role is built-in (Account Admin, Organization Admin, Regular User). | [optional] 
**Permissions** | Pointer to [**[]Permission**](Permission.md) |  | [optional] 
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewCreateRole201Response

`func NewCreateRole201Response() *CreateRole201Response`

NewCreateRole201Response instantiates a new CreateRole201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRole201ResponseWithDefaults

`func NewCreateRole201ResponseWithDefaults() *CreateRole201Response`

NewCreateRole201ResponseWithDefaults instantiates a new CreateRole201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRole201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRole201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRole201Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateRole201Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoleId

`func (o *CreateRole201Response) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateRole201Response) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateRole201Response) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *CreateRole201Response) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetIsBuiltin

`func (o *CreateRole201Response) GetIsBuiltin() bool`

GetIsBuiltin returns the IsBuiltin field if non-nil, zero value otherwise.

### GetIsBuiltinOk

`func (o *CreateRole201Response) GetIsBuiltinOk() (*bool, bool)`

GetIsBuiltinOk returns a tuple with the IsBuiltin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuiltin

`func (o *CreateRole201Response) SetIsBuiltin(v bool)`

SetIsBuiltin sets IsBuiltin field to given value.

### HasIsBuiltin

`func (o *CreateRole201Response) HasIsBuiltin() bool`

HasIsBuiltin returns a boolean if a field has been set.

### GetPermissions

`func (o *CreateRole201Response) GetPermissions() []Permission`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRole201Response) GetPermissionsOk() (*[]Permission, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRole201Response) SetPermissions(v []Permission)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRole201Response) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetLinks

`func (o *CreateRole201Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateRole201Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateRole201Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateRole201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


