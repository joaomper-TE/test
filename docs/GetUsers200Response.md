# GetUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]ExtendedUser**](ExtendedUser.md) |  | [optional] 
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewGetUsers200Response

`func NewGetUsers200Response() *GetUsers200Response`

NewGetUsers200Response instantiates a new GetUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200ResponseWithDefaults

`func NewGetUsers200ResponseWithDefaults() *GetUsers200Response`

NewGetUsers200ResponseWithDefaults instantiates a new GetUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetUsers200Response) GetUsers() []ExtendedUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetUsers200Response) GetUsersOk() (*[]ExtendedUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetUsers200Response) SetUsers(v []ExtendedUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetLinks

`func (o *GetUsers200Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetUsers200Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetUsers200Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetUsers200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


