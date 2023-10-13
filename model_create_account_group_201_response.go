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

// checks if the CreateAccountGroup201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccountGroup201Response{}

// CreateAccountGroup201Response struct for CreateAccountGroup201Response
type CreateAccountGroup201Response struct {
	// Unique name for the account group within an organization.
	AccountGroupName *string `json:"accountGroupName,omitempty"`
	// Unique ID representing the account group.
	Aid *string `json:"aid,omitempty"`
	// Indicates whether the requested aid is the context of the current account.
	IsCurrentAccountGroup *bool `json:"isCurrentAccountGroup,omitempty"`
	// Indicates whether the aid is the default one for the requesting user.
	IsDefaultAccountGroup *bool `json:"isDefaultAccountGroup,omitempty"`
	// (Optional) Indicates whether the aid is the default one for the requesting user.
	OrganizationName *string `json:"organizationName,omitempty"`
	Users []UserAccountGroup `json:"users,omitempty"`
	Links *SelfLinksLinks `json:"_links,omitempty"`
}

// NewCreateAccountGroup201Response instantiates a new CreateAccountGroup201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountGroup201Response() *CreateAccountGroup201Response {
	this := CreateAccountGroup201Response{}
	return &this
}

// NewCreateAccountGroup201ResponseWithDefaults instantiates a new CreateAccountGroup201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountGroup201ResponseWithDefaults() *CreateAccountGroup201Response {
	this := CreateAccountGroup201Response{}
	return &this
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetAccountGroupName() string {
	if o == nil || IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasAccountGroupName() bool {
	if o != nil && !IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *CreateAccountGroup201Response) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetAid() string {
	if o == nil || IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetAidOk() (*string, bool) {
	if o == nil || IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasAid() bool {
	if o != nil && !IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *CreateAccountGroup201Response) SetAid(v string) {
	o.Aid = &v
}

// GetIsCurrentAccountGroup returns the IsCurrentAccountGroup field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetIsCurrentAccountGroup() bool {
	if o == nil || IsNil(o.IsCurrentAccountGroup) {
		var ret bool
		return ret
	}
	return *o.IsCurrentAccountGroup
}

// GetIsCurrentAccountGroupOk returns a tuple with the IsCurrentAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetIsCurrentAccountGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCurrentAccountGroup) {
		return nil, false
	}
	return o.IsCurrentAccountGroup, true
}

// HasIsCurrentAccountGroup returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasIsCurrentAccountGroup() bool {
	if o != nil && !IsNil(o.IsCurrentAccountGroup) {
		return true
	}

	return false
}

// SetIsCurrentAccountGroup gets a reference to the given bool and assigns it to the IsCurrentAccountGroup field.
func (o *CreateAccountGroup201Response) SetIsCurrentAccountGroup(v bool) {
	o.IsCurrentAccountGroup = &v
}

// GetIsDefaultAccountGroup returns the IsDefaultAccountGroup field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetIsDefaultAccountGroup() bool {
	if o == nil || IsNil(o.IsDefaultAccountGroup) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccountGroup
}

// GetIsDefaultAccountGroupOk returns a tuple with the IsDefaultAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetIsDefaultAccountGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccountGroup) {
		return nil, false
	}
	return o.IsDefaultAccountGroup, true
}

// HasIsDefaultAccountGroup returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasIsDefaultAccountGroup() bool {
	if o != nil && !IsNil(o.IsDefaultAccountGroup) {
		return true
	}

	return false
}

// SetIsDefaultAccountGroup gets a reference to the given bool and assigns it to the IsDefaultAccountGroup field.
func (o *CreateAccountGroup201Response) SetIsDefaultAccountGroup(v bool) {
	o.IsDefaultAccountGroup = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *CreateAccountGroup201Response) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetUsers() []UserAccountGroup {
	if o == nil || IsNil(o.Users) {
		var ret []UserAccountGroup
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetUsersOk() ([]UserAccountGroup, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserAccountGroup and assigns it to the Users field.
func (o *CreateAccountGroup201Response) SetUsers(v []UserAccountGroup) {
	o.Users = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CreateAccountGroup201Response) GetLinks() SelfLinksLinks {
	if o == nil || IsNil(o.Links) {
		var ret SelfLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountGroup201Response) GetLinksOk() (*SelfLinksLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CreateAccountGroup201Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinksLinks and assigns it to the Links field.
func (o *CreateAccountGroup201Response) SetLinks(v SelfLinksLinks) {
	o.Links = &v
}

func (o CreateAccountGroup201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccountGroup201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	if !IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !IsNil(o.IsCurrentAccountGroup) {
		toSerialize["isCurrentAccountGroup"] = o.IsCurrentAccountGroup
	}
	if !IsNil(o.IsDefaultAccountGroup) {
		toSerialize["isDefaultAccountGroup"] = o.IsDefaultAccountGroup
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCreateAccountGroup201Response struct {
	value *CreateAccountGroup201Response
	isSet bool
}

func (v NullableCreateAccountGroup201Response) Get() *CreateAccountGroup201Response {
	return v.value
}

func (v *NullableCreateAccountGroup201Response) Set(val *CreateAccountGroup201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountGroup201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountGroup201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountGroup201Response(val *CreateAccountGroup201Response) *NullableCreateAccountGroup201Response {
	return &NullableCreateAccountGroup201Response{value: val, isSet: true}
}

func (v NullableCreateAccountGroup201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountGroup201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

