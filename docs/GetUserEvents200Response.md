# GetUserEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEvents** | Pointer to [**[]UserEvent**](UserEvent.md) |  | [optional] 
**StartDate** | Pointer to **time.Time** | (Optional) When passing &#x60;window&#x60;, the client will also receive the &#x60;startDate&#x60; and &#x60;endDate&#x60; fields. | [optional] 
**EndDate** | Pointer to **time.Time** | (Optional) When passing &#x60;window&#x60;, the client will also receive the &#x60;startDate&#x60; and &#x60;endDate&#x60; fields. | [optional] 
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewGetUserEvents200Response

`func NewGetUserEvents200Response() *GetUserEvents200Response`

NewGetUserEvents200Response instantiates a new GetUserEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserEvents200ResponseWithDefaults

`func NewGetUserEvents200ResponseWithDefaults() *GetUserEvents200Response`

NewGetUserEvents200ResponseWithDefaults instantiates a new GetUserEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEvents

`func (o *GetUserEvents200Response) GetAuditEvents() []UserEvent`

GetAuditEvents returns the AuditEvents field if non-nil, zero value otherwise.

### GetAuditEventsOk

`func (o *GetUserEvents200Response) GetAuditEventsOk() (*[]UserEvent, bool)`

GetAuditEventsOk returns a tuple with the AuditEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEvents

`func (o *GetUserEvents200Response) SetAuditEvents(v []UserEvent)`

SetAuditEvents sets AuditEvents field to given value.

### HasAuditEvents

`func (o *GetUserEvents200Response) HasAuditEvents() bool`

HasAuditEvents returns a boolean if a field has been set.

### GetStartDate

`func (o *GetUserEvents200Response) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetUserEvents200Response) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetUserEvents200Response) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetUserEvents200Response) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetUserEvents200Response) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetUserEvents200Response) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetUserEvents200Response) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetUserEvents200Response) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLinks

`func (o *GetUserEvents200Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUserEvents200Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUserEvents200Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUserEvents200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


