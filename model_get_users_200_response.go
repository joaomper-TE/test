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

// checks if the GetUsers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUsers200Response{}

// GetUsers200Response struct for GetUsers200Response
type GetUsers200Response struct {
	Users []ExtendedUser `json:"users,omitempty"`
	Links *SelfLinksLinks `json:"_links,omitempty"`
}

// NewGetUsers200Response instantiates a new GetUsers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsers200Response() *GetUsers200Response {
	this := GetUsers200Response{}
	return &this
}

// NewGetUsers200ResponseWithDefaults instantiates a new GetUsers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsers200ResponseWithDefaults() *GetUsers200Response {
	this := GetUsers200Response{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GetUsers200Response) GetUsers() []ExtendedUser {
	if o == nil || IsNil(o.Users) {
		var ret []ExtendedUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200Response) GetUsersOk() ([]ExtendedUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GetUsers200Response) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []ExtendedUser and assigns it to the Users field.
func (o *GetUsers200Response) SetUsers(v []ExtendedUser) {
	o.Users = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetUsers200Response) GetLinks() SelfLinksLinks {
	if o == nil || IsNil(o.Links) {
		var ret SelfLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUsers200Response) GetLinksOk() (*SelfLinksLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetUsers200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinksLinks and assigns it to the Links field.
func (o *GetUsers200Response) SetLinks(v SelfLinksLinks) {
	o.Links = &v
}

func (o GetUsers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUsers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetUsers200Response struct {
	value *GetUsers200Response
	isSet bool
}

func (v NullableGetUsers200Response) Get() *GetUsers200Response {
	return v.value
}

func (v *NullableGetUsers200Response) Set(val *GetUsers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsers200Response(val *GetUsers200Response) *NullableGetUsers200Response {
	return &NullableGetUsers200Response{value: val, isSet: true}
}

func (v NullableGetUsers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


