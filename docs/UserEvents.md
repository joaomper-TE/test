# UserEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEvents** | Pointer to [**[]UserEvent**](UserEvent.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | (Optional) When passing &#x60;window&#x60;, the client will also receive the &#x60;startDate&#x60; and &#x60;endDate&#x60; fields. | [optional] 
**EndDate** | Pointer to **time.Time** | (Optional) When passing &#x60;window&#x60;, the client will also receive the &#x60;startDate&#x60; and &#x60;endDate&#x60; fields. | [optional] 

## Methods

### NewUserEvents

`func NewUserEvents() *UserEvents`

NewUserEvents instantiates a new UserEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEventsWithDefaults

`func NewUserEventsWithDefaults() *UserEvents`

NewUserEventsWithDefaults instantiates a new UserEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEvents

`func (o *UserEvents) GetAuditEvents() []UserEvent`

GetAuditEvents returns the AuditEvents field if non-nil, zero value otherwise.

### GetAuditEventsOk

`func (o *UserEvents) GetAuditEventsOk() (*[]UserEvent, bool)`

GetAuditEventsOk returns a tuple with the AuditEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEvents

`func (o *UserEvents) SetAuditEvents(v []UserEvent)`

SetAuditEvents sets AuditEvents field to given value.

### HasAuditEvents

`func (o *UserEvents) HasAuditEvents() bool`

HasAuditEvents returns a boolean if a field has been set.

### GetStartDate

`func (o *UserEvents) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UserEvents) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UserEvents) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UserEvents) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UserEvents) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UserEvents) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UserEvents) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UserEvents) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


