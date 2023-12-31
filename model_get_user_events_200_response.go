/*
Administrative API

## Overview Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API.  This API provides the following endpoints that define the operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These endpoints can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thousandeyes

import (
	"encoding/json"
	"time"
)

// checks if the GetUserEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserEvents200Response{}

// GetUserEvents200Response struct for GetUserEvents200Response
type GetUserEvents200Response struct {
	AuditEvents []UserEvent `json:"auditEvents,omitempty"`
	// (Optional) When passing `window`, the client will also receive the `startDate` and `endDate` fields.
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window`, the client will also receive the `startDate` and `endDate` fields.
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *SelfLinksLinks `json:"_links,omitempty"`
}

// NewGetUserEvents200Response instantiates a new GetUserEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserEvents200Response() *GetUserEvents200Response {
	this := GetUserEvents200Response{}
	return &this
}

// NewGetUserEvents200ResponseWithDefaults instantiates a new GetUserEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserEvents200ResponseWithDefaults() *GetUserEvents200Response {
	this := GetUserEvents200Response{}
	return &this
}

// GetAuditEvents returns the AuditEvents field value if set, zero value otherwise.
func (o *GetUserEvents200Response) GetAuditEvents() []UserEvent {
	if o == nil || IsNil(o.AuditEvents) {
		var ret []UserEvent
		return ret
	}
	return o.AuditEvents
}

// GetAuditEventsOk returns a tuple with the AuditEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserEvents200Response) GetAuditEventsOk() ([]UserEvent, bool) {
	if o == nil || IsNil(o.AuditEvents) {
		return nil, false
	}
	return o.AuditEvents, true
}

// HasAuditEvents returns a boolean if a field has been set.
func (o *GetUserEvents200Response) HasAuditEvents() bool {
	if o != nil && !IsNil(o.AuditEvents) {
		return true
	}

	return false
}

// SetAuditEvents gets a reference to the given []UserEvent and assigns it to the AuditEvents field.
func (o *GetUserEvents200Response) SetAuditEvents(v []UserEvent) {
	o.AuditEvents = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *GetUserEvents200Response) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserEvents200Response) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *GetUserEvents200Response) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *GetUserEvents200Response) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *GetUserEvents200Response) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserEvents200Response) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *GetUserEvents200Response) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *GetUserEvents200Response) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetUserEvents200Response) GetLinks() SelfLinksLinks {
	if o == nil || IsNil(o.Links) {
		var ret SelfLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserEvents200Response) GetLinksOk() (*SelfLinksLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetUserEvents200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinksLinks and assigns it to the Links field.
func (o *GetUserEvents200Response) SetLinks(v SelfLinksLinks) {
	o.Links = &v
}

func (o GetUserEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuditEvents) {
		toSerialize["auditEvents"] = o.AuditEvents
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGetUserEvents200Response struct {
	value *GetUserEvents200Response
	isSet bool
}

func (v NullableGetUserEvents200Response) Get() *GetUserEvents200Response {
	return v.value
}

func (v *NullableGetUserEvents200Response) Set(val *GetUserEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserEvents200Response(val *GetUserEvents200Response) *NullableGetUserEvents200Response {
	return &NullableGetUserEvents200Response{value: val, isSet: true}
}

func (v NullableGetUserEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


