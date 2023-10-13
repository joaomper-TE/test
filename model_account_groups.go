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

// checks if the AccountGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountGroups{}

// AccountGroups struct for AccountGroups
type AccountGroups struct {
	AccountGroups []AccountGroup `json:"accountGroups,omitempty"`
}

// NewAccountGroups instantiates a new AccountGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGroups() *AccountGroups {
	this := AccountGroups{}
	return &this
}

// NewAccountGroupsWithDefaults instantiates a new AccountGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGroupsWithDefaults() *AccountGroups {
	this := AccountGroups{}
	return &this
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *AccountGroups) GetAccountGroups() []AccountGroup {
	if o == nil || IsNil(o.AccountGroups) {
		var ret []AccountGroup
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroups) GetAccountGroupsOk() ([]AccountGroup, bool) {
	if o == nil || IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *AccountGroups) HasAccountGroups() bool {
	if o != nil && !IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []AccountGroup and assigns it to the AccountGroups field.
func (o *AccountGroups) SetAccountGroups(v []AccountGroup) {
	o.AccountGroups = v
}

func (o AccountGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	return toSerialize, nil
}

type NullableAccountGroups struct {
	value *AccountGroups
	isSet bool
}

func (v NullableAccountGroups) Get() *AccountGroups {
	return v.value
}

func (v *NullableAccountGroups) Set(val *AccountGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGroups(val *AccountGroups) *NullableAccountGroups {
	return &NullableAccountGroups{value: val, isSet: true}
}

func (v NullableAccountGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

