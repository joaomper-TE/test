# UserEventResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of resource affected. Can be “testName”, “reportTitle”, “userDisplayName”, “alertRuleName”, etc. | [optional] 
**Name** | Pointer to **string** | Name of the affected resource. | [optional] 

## Methods

### NewUserEventResourcesInner

`func NewUserEventResourcesInner() *UserEventResourcesInner`

NewUserEventResourcesInner instantiates a new UserEventResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEventResourcesInnerWithDefaults

`func NewUserEventResourcesInnerWithDefaults() *UserEventResourcesInner`

NewUserEventResourcesInnerWithDefaults instantiates a new UserEventResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserEventResourcesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserEventResourcesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserEventResourcesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserEventResourcesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UserEventResourcesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserEventResourcesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserEventResourcesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserEventResourcesInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


