# UserEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupName** | Pointer to **string** | The name of the account group. | [optional] 
**Aid** | Pointer to **string** | Unique ID representing the account group. | [optional] 
**Date** | Pointer to **time.Time** | Date of the event in YYYY-mm-ddTHH:MM:SSZ format (UTC). | [optional] 
**Event** | Pointer to **string** | Event type. | [optional] 
**IpAddress** | Pointer to **string** | Source IP address of the user. | [optional] 
**Uid** | Pointer to **string** | Unique id representing the user. | [optional] 
**User** | Pointer to **string** | The name and email address of the user. | [optional] 
**Resources** | Pointer to [**[]UserEventResourcesInner**](UserEventResourcesInner.md) |  | [optional] 

## Methods

### NewUserEvent

`func NewUserEvent() *UserEvent`

NewUserEvent instantiates a new UserEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEventWithDefaults

`func NewUserEventWithDefaults() *UserEvent`

NewUserEventWithDefaults instantiates a new UserEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupName

`func (o *UserEvent) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *UserEvent) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *UserEvent) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *UserEvent) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAid

`func (o *UserEvent) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *UserEvent) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *UserEvent) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *UserEvent) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetDate

`func (o *UserEvent) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UserEvent) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UserEvent) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *UserEvent) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetEvent

`func (o *UserEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *UserEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *UserEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *UserEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetIpAddress

`func (o *UserEvent) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UserEvent) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UserEvent) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UserEvent) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUid

`func (o *UserEvent) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *UserEvent) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *UserEvent) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *UserEvent) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUser

`func (o *UserEvent) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserEvent) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserEvent) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UserEvent) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetResources

`func (o *UserEvent) GetResources() []UserEventResourcesInner`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UserEvent) GetResourcesOk() (*[]UserEventResourcesInner, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UserEvent) SetResources(v []UserEventResourcesInner)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UserEvent) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


