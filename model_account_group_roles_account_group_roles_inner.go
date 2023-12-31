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

// checks if the AccountGroupRolesAccountGroupRolesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountGroupRolesAccountGroupRolesInner{}

// AccountGroupRolesAccountGroupRolesInner struct for AccountGroupRolesAccountGroupRolesInner
type AccountGroupRolesAccountGroupRolesInner struct {
	AccountGroup *AccountGroupUser `json:"accountGroup,omitempty"`
	Roles []Role `json:"roles,omitempty"`
}

// NewAccountGroupRolesAccountGroupRolesInner instantiates a new AccountGroupRolesAccountGroupRolesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGroupRolesAccountGroupRolesInner() *AccountGroupRolesAccountGroupRolesInner {
	this := AccountGroupRolesAccountGroupRolesInner{}
	return &this
}

// NewAccountGroupRolesAccountGroupRolesInnerWithDefaults instantiates a new AccountGroupRolesAccountGroupRolesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGroupRolesAccountGroupRolesInnerWithDefaults() *AccountGroupRolesAccountGroupRolesInner {
	this := AccountGroupRolesAccountGroupRolesInner{}
	return &this
}

// GetAccountGroup returns the AccountGroup field value if set, zero value otherwise.
func (o *AccountGroupRolesAccountGroupRolesInner) GetAccountGroup() AccountGroupUser {
	if o == nil || IsNil(o.AccountGroup) {
		var ret AccountGroupUser
		return ret
	}
	return *o.AccountGroup
}

// GetAccountGroupOk returns a tuple with the AccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupRolesAccountGroupRolesInner) GetAccountGroupOk() (*AccountGroupUser, bool) {
	if o == nil || IsNil(o.AccountGroup) {
		return nil, false
	}
	return o.AccountGroup, true
}

// HasAccountGroup returns a boolean if a field has been set.
func (o *AccountGroupRolesAccountGroupRolesInner) HasAccountGroup() bool {
	if o != nil && !IsNil(o.AccountGroup) {
		return true
	}

	return false
}

// SetAccountGroup gets a reference to the given AccountGroupUser and assigns it to the AccountGroup field.
func (o *AccountGroupRolesAccountGroupRolesInner) SetAccountGroup(v AccountGroupUser) {
	o.AccountGroup = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AccountGroupRolesAccountGroupRolesInner) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupRolesAccountGroupRolesInner) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AccountGroupRolesAccountGroupRolesInner) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *AccountGroupRolesAccountGroupRolesInner) SetRoles(v []Role) {
	o.Roles = v
}

func (o AccountGroupRolesAccountGroupRolesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGroupRolesAccountGroupRolesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountGroup) {
		toSerialize["accountGroup"] = o.AccountGroup
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableAccountGroupRolesAccountGroupRolesInner struct {
	value *AccountGroupRolesAccountGroupRolesInner
	isSet bool
}

func (v NullableAccountGroupRolesAccountGroupRolesInner) Get() *AccountGroupRolesAccountGroupRolesInner {
	return v.value
}

func (v *NullableAccountGroupRolesAccountGroupRolesInner) Set(val *AccountGroupRolesAccountGroupRolesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGroupRolesAccountGroupRolesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGroupRolesAccountGroupRolesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGroupRolesAccountGroupRolesInner(val *AccountGroupRolesAccountGroupRolesInner) *NullableAccountGroupRolesAccountGroupRolesInner {
	return &NullableAccountGroupRolesAccountGroupRolesInner{value: val, isSet: true}
}

func (v NullableAccountGroupRolesAccountGroupRolesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGroupRolesAccountGroupRolesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


