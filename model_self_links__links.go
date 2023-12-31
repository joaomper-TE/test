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

// checks if the SelfLinksLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelfLinksLinks{}

// SelfLinksLinks A links object containing the self link.
type SelfLinksLinks struct {
	Self *Link `json:"self,omitempty"`
}

// NewSelfLinksLinks instantiates a new SelfLinksLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfLinksLinks() *SelfLinksLinks {
	this := SelfLinksLinks{}
	return &this
}

// NewSelfLinksLinksWithDefaults instantiates a new SelfLinksLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfLinksLinksWithDefaults() *SelfLinksLinks {
	this := SelfLinksLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SelfLinksLinks) GetSelf() Link {
	if o == nil || IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfLinksLinks) GetSelfOk() (*Link, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SelfLinksLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *SelfLinksLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o SelfLinksLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfLinksLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableSelfLinksLinks struct {
	value *SelfLinksLinks
	isSet bool
}

func (v NullableSelfLinksLinks) Get() *SelfLinksLinks {
	return v.value
}

func (v *NullableSelfLinksLinks) Set(val *SelfLinksLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfLinksLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfLinksLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfLinksLinks(val *SelfLinksLinks) *NullableSelfLinksLinks {
	return &NullableSelfLinksLinks{value: val, isSet: true}
}

func (v NullableSelfLinksLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfLinksLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


