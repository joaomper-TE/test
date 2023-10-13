/*
Administrative API

## Overview Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API.  This API provides the following endpoints that define the operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These endpoints can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thousandeyes

import (
	"encoding/json"
)

// checks if the RoleDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleDetail{}

// RoleDetail struct for RoleDetail
type RoleDetail struct {
	// Name of the role.
	Name *string `json:"name,omitempty"`
	// Unique ID representing the role.
	RoleId *string `json:"roleId,omitempty"`
	// Flag indicating if the role is built-in (Account Admin, Organization Admin, Regular User).
	IsBuiltin *bool `json:"isBuiltin,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

// NewRoleDetail instantiates a new RoleDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleDetail() *RoleDetail {
	this := RoleDetail{}
	return &this
}

// NewRoleDetailWithDefaults instantiates a new RoleDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleDetailWithDefaults() *RoleDetail {
	this := RoleDetail{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleDetail) SetName(v string) {
	o.Name = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *RoleDetail) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDetail) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *RoleDetail) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *RoleDetail) SetRoleId(v string) {
	o.RoleId = &v
}

// GetIsBuiltin returns the IsBuiltin field value if set, zero value otherwise.
func (o *RoleDetail) GetIsBuiltin() bool {
	if o == nil || IsNil(o.IsBuiltin) {
		var ret bool
		return ret
	}
	return *o.IsBuiltin
}

// GetIsBuiltinOk returns a tuple with the IsBuiltin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDetail) GetIsBuiltinOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBuiltin) {
		return nil, false
	}
	return o.IsBuiltin, true
}

// HasIsBuiltin returns a boolean if a field has been set.
func (o *RoleDetail) HasIsBuiltin() bool {
	if o != nil && !IsNil(o.IsBuiltin) {
		return true
	}

	return false
}

// SetIsBuiltin gets a reference to the given bool and assigns it to the IsBuiltin field.
func (o *RoleDetail) SetIsBuiltin(v bool) {
	o.IsBuiltin = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *RoleDetail) GetPermissions() []Permission {
	if o == nil || IsNil(o.Permissions) {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleDetail) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *RoleDetail) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *RoleDetail) SetPermissions(v []Permission) {
	o.Permissions = v
}

func (o RoleDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !IsNil(o.IsBuiltin) {
		toSerialize["isBuiltin"] = o.IsBuiltin
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableRoleDetail struct {
	value *RoleDetail
	isSet bool
}

func (v NullableRoleDetail) Get() *RoleDetail {
	return v.value
}

func (v *NullableRoleDetail) Set(val *RoleDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleDetail(val *RoleDetail) *NullableRoleDetail {
	return &NullableRoleDetail{value: val, isSet: true}
}

func (v NullableRoleDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

