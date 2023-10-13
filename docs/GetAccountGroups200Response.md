# GetAccountGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroups** | Pointer to [**[]AccountGroup**](AccountGroup.md) |  | [optional] 
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewGetAccountGroups200Response

`func NewGetAccountGroups200Response() *GetAccountGroups200Response`

NewGetAccountGroups200Response instantiates a new GetAccountGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountGroups200ResponseWithDefaults

`func NewGetAccountGroups200ResponseWithDefaults() *GetAccountGroups200Response`

NewGetAccountGroups200ResponseWithDefaults instantiates a new GetAccountGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroups

`func (o *GetAccountGroups200Response) GetAccountGroups() []AccountGroup`

GetAccountGroups returns the AccountGroups field if non-nil, zero value otherwise.

### GetAccountGroupsOk

`func (o *GetAccountGroups200Response) GetAccountGroupsOk() (*[]AccountGroup, bool)`

GetAccountGroupsOk returns a tuple with the AccountGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroups

`func (o *GetAccountGroups200Response) SetAccountGroups(v []AccountGroup)`

SetAccountGroups sets AccountGroups field to given value.

### HasAccountGroups

`func (o *GetAccountGroups200Response) HasAccountGroups() bool`

HasAccountGroups returns a boolean if a field has been set.

### GetLinks

`func (o *GetAccountGroups200Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAccountGroups200Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAccountGroups200Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAccountGroups200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


