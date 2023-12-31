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

// checks if the GetRoles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoles200Response{}

// GetRoles200Response struct for GetRoles200Response
type GetRoles200Response struct {
	Roles []Role `json:"roles,omitempty"`
	Links *SelfLinksLinks `json:"_links,omitempty"`
}

// NewGetRoles200Response instantiates a new GetRoles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoles200Response() *GetRoles200Response {
	this := GetRoles200Response{}
	return &this
}

// NewGetRoles200ResponseWithDefaults instantiates a new GetRoles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoles200ResponseWithDefaults() *GetRoles200Response {
	this := GetRoles200Response{}
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *GetRoles200Response) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoles200Response) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GetRoles200Response) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *GetRoles200Response) SetRoles(v []Role) {
	o.Roles = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetRoles200Response) GetLinks() SelfLinksLinks {
	if o == nil || IsNil(o.Links) {
		var ret SelfLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoles200Response) GetLinksOk() (*SelfLinksLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetRoles200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinksLinks and assigns it to the Links field.
func (o *GetRoles200Response) SetLinks(v SelfLinksLinks) {
	o.Links = &v
}

func (o GetRoles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetRoles200Response struct {
	value *GetRoles200Response
	isSet bool
}

func (v NullableGetRoles200Response) Get() *GetRoles200Response {
	return v.value
}

func (v *NullableGetRoles200Response) Set(val *GetRoles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoles200Response(val *GetRoles200Response) *NullableGetRoles200Response {
	return &NullableGetRoles200Response{value: val, isSet: true}
}

func (v NullableGetRoles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


