# GetRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewGetRoles200Response

`func NewGetRoles200Response() *GetRoles200Response`

NewGetRoles200Response instantiates a new GetRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRoles200ResponseWithDefaults

`func NewGetRoles200ResponseWithDefaults() *GetRoles200Response`

NewGetRoles200ResponseWithDefaults instantiates a new GetRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GetRoles200Response) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetRoles200Response) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetRoles200Response) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetRoles200Response) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetLinks

`func (o *GetRoles200Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetRoles200Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetRoles200Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetRoles200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


