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

// checks if the UserAccountGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAccountGroups{}

// UserAccountGroups struct for UserAccountGroups
type UserAccountGroups struct {
	Users []UserAccountGroup `json:"users,omitempty"`
}

// NewUserAccountGroups instantiates a new UserAccountGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccountGroups() *UserAccountGroups {
	this := UserAccountGroups{}
	return &this
}

// NewUserAccountGroupsWithDefaults instantiates a new UserAccountGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccountGroupsWithDefaults() *UserAccountGroups {
	this := UserAccountGroups{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UserAccountGroups) GetUsers() []UserAccountGroup {
	if o == nil || IsNil(o.Users) {
		var ret []UserAccountGroup
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccountGroups) GetUsersOk() ([]UserAccountGroup, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserAccountGroups) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserAccountGroup and assigns it to the Users field.
func (o *UserAccountGroups) SetUsers(v []UserAccountGroup) {
	o.Users = v
}

func (o UserAccountGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAccountGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUserAccountGroups struct {
	value *UserAccountGroups
	isSet bool
}

func (v NullableUserAccountGroups) Get() *UserAccountGroups {
	return v.value
}

func (v *NullableUserAccountGroups) Set(val *UserAccountGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccountGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccountGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccountGroups(val *UserAccountGroups) *NullableUserAccountGroups {
	return &NullableUserAccountGroups{value: val, isSet: true}
}

func (v NullableUserAccountGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccountGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


