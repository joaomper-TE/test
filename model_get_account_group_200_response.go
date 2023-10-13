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

// checks if the GetAccountGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountGroup200Response{}

// GetAccountGroup200Response struct for GetAccountGroup200Response
type GetAccountGroup200Response struct {
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
	Agents []EnterpriseAgent `json:"agents,omitempty"`
	Links *SelfLinksLinks `json:"_links,omitempty"`
}

// NewGetAccountGroup200Response instantiates a new GetAccountGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountGroup200Response() *GetAccountGroup200Response {
	this := GetAccountGroup200Response{}
	return &this
}

// NewGetAccountGroup200ResponseWithDefaults instantiates a new GetAccountGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountGroup200ResponseWithDefaults() *GetAccountGroup200Response {
	this := GetAccountGroup200Response{}
	return &this
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetAccountGroupName() string {
	if o == nil || IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasAccountGroupName() bool {
	if o != nil && !IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *GetAccountGroup200Response) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetAid() string {
	if o == nil || IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetAidOk() (*string, bool) {
	if o == nil || IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasAid() bool {
	if o != nil && !IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *GetAccountGroup200Response) SetAid(v string) {
	o.Aid = &v
}

// GetIsCurrentAccountGroup returns the IsCurrentAccountGroup field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetIsCurrentAccountGroup() bool {
	if o == nil || IsNil(o.IsCurrentAccountGroup) {
		var ret bool
		return ret
	}
	return *o.IsCurrentAccountGroup
}

// GetIsCurrentAccountGroupOk returns a tuple with the IsCurrentAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetIsCurrentAccountGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCurrentAccountGroup) {
		return nil, false
	}
	return o.IsCurrentAccountGroup, true
}

// HasIsCurrentAccountGroup returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasIsCurrentAccountGroup() bool {
	if o != nil && !IsNil(o.IsCurrentAccountGroup) {
		return true
	}

	return false
}

// SetIsCurrentAccountGroup gets a reference to the given bool and assigns it to the IsCurrentAccountGroup field.
func (o *GetAccountGroup200Response) SetIsCurrentAccountGroup(v bool) {
	o.IsCurrentAccountGroup = &v
}

// GetIsDefaultAccountGroup returns the IsDefaultAccountGroup field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetIsDefaultAccountGroup() bool {
	if o == nil || IsNil(o.IsDefaultAccountGroup) {
		var ret bool
		return ret
	}
	return *o.IsDefaultAccountGroup
}

// GetIsDefaultAccountGroupOk returns a tuple with the IsDefaultAccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetIsDefaultAccountGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefaultAccountGroup) {
		return nil, false
	}
	return o.IsDefaultAccountGroup, true
}

// HasIsDefaultAccountGroup returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasIsDefaultAccountGroup() bool {
	if o != nil && !IsNil(o.IsDefaultAccountGroup) {
		return true
	}

	return false
}

// SetIsDefaultAccountGroup gets a reference to the given bool and assigns it to the IsDefaultAccountGroup field.
func (o *GetAccountGroup200Response) SetIsDefaultAccountGroup(v bool) {
	o.IsDefaultAccountGroup = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *GetAccountGroup200Response) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetUsers() []UserAccountGroup {
	if o == nil || IsNil(o.Users) {
		var ret []UserAccountGroup
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetUsersOk() ([]UserAccountGroup, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserAccountGroup and assigns it to the Users field.
func (o *GetAccountGroup200Response) SetUsers(v []UserAccountGroup) {
	o.Users = v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetAgents() []EnterpriseAgent {
	if o == nil || IsNil(o.Agents) {
		var ret []EnterpriseAgent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetAgentsOk() ([]EnterpriseAgent, bool) {
	if o == nil || IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasAgents() bool {
	if o != nil && !IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []EnterpriseAgent and assigns it to the Agents field.
func (o *GetAccountGroup200Response) SetAgents(v []EnterpriseAgent) {
	o.Agents = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetAccountGroup200Response) GetLinks() SelfLinksLinks {
	if o == nil || IsNil(o.Links) {
		var ret SelfLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccountGroup200Response) GetLinksOk() (*SelfLinksLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetAccountGroup200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinksLinks and assigns it to the Links field.
func (o *GetAccountGroup200Response) SetLinks(v SelfLinksLinks) {
	o.Links = &v
}

func (o GetAccountGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountGroup200Response) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetAccountGroup200Response struct {
	value *GetAccountGroup200Response
	isSet bool
}

func (v NullableGetAccountGroup200Response) Get() *GetAccountGroup200Response {
	return v.value
}

func (v *NullableGetAccountGroup200Response) Set(val *GetAccountGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountGroup200Response(val *GetAccountGroup200Response) *NullableGetAccountGroup200Response {
	return &NullableGetAccountGroup200Response{value: val, isSet: true}
}

func (v NullableGetAccountGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

