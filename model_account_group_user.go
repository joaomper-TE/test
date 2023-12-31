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

// checks if the AccountGroupUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountGroupUser{}

// AccountGroupUser struct for AccountGroupUser
type AccountGroupUser struct {
	// The name of the account group
	AccountGroupName *string `json:"accountGroupName,omitempty"`
	// Unique ID of the account group.
	Aid *string `json:"aid,omitempty"`
}

// NewAccountGroupUser instantiates a new AccountGroupUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGroupUser() *AccountGroupUser {
	this := AccountGroupUser{}
	return &this
}

// NewAccountGroupUserWithDefaults instantiates a new AccountGroupUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGroupUserWithDefaults() *AccountGroupUser {
	this := AccountGroupUser{}
	return &this
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *AccountGroupUser) GetAccountGroupName() string {
	if o == nil || IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupUser) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *AccountGroupUser) HasAccountGroupName() bool {
	if o != nil && !IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *AccountGroupUser) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *AccountGroupUser) GetAid() string {
	if o == nil || IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupUser) GetAidOk() (*string, bool) {
	if o == nil || IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *AccountGroupUser) HasAid() bool {
	if o != nil && !IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *AccountGroupUser) SetAid(v string) {
	o.Aid = &v
}

func (o AccountGroupUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGroupUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	if !IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	return toSerialize, nil
}

type NullableAccountGroupUser struct {
	value *AccountGroupUser
	isSet bool
}

func (v NullableAccountGroupUser) Get() *AccountGroupUser {
	return v.value
}

func (v *NullableAccountGroupUser) Set(val *AccountGroupUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGroupUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGroupUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGroupUser(val *AccountGroupUser) *NullableAccountGroupUser {
	return &NullableAccountGroupUser{value: val, isSet: true}
}

func (v NullableAccountGroupUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGroupUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


