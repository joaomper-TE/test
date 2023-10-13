# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** | Its value is either a URI [RFC3986] or a URI template [RFC6570]. | 
**Templated** | Pointer to **bool** | Should be true when the link object&#39;s \&quot;href\&quot; property is a URI template. | [optional] 
**Type** | Pointer to **string** | Used as a hint to indicate the media type expected when dereferencing the target resource. | [optional] 
**Deprecation** | Pointer to **string** | Its presence indicates that the link is to be deprecated at a future date. Its value is a URL that should provide further information about the deprecation. | [optional] 
**Name** | Pointer to **string** | Its value may be used as a secondary key for selecting link objects that share the same relation type. | [optional] 
**Profile** | Pointer to **string** | A URI that hints about the profile of the target resource. | [optional] 
**Title** | Pointer to **string** | Intended for labelling the link with a human-readable identifier | [optional] 
**Hreflang** | Pointer to **string** | Indicates the language of the target resource | [optional] 

## Methods

### NewLink

`func NewLink(href string, ) *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Link) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Link) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Link) SetHref(v string)`

SetHref sets Href field to given value.


### GetTemplated

`func (o *Link) GetTemplated() bool`

GetTemplated returns the Templated field if non-nil, zero value otherwise.

### GetTemplatedOk

`func (o *Link) GetTemplatedOk() (*bool, bool)`

GetTemplatedOk returns a tuple with the Templated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplated

`func (o *Link) SetTemplated(v bool)`

SetTemplated sets Templated field to given value.

### HasTemplated

`func (o *Link) HasTemplated() bool`

HasTemplated returns a boolean if a field has been set.

### GetType

`func (o *Link) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Link) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Link) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Link) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDeprecation

`func (o *Link) GetDeprecation() string`

GetDeprecation returns the Deprecation field if non-nil, zero value otherwise.

### GetDeprecationOk

`func (o *Link) GetDeprecationOk() (*string, bool)`

GetDeprecationOk returns a tuple with the Deprecation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecation

`func (o *Link) SetDeprecation(v string)`

SetDeprecation sets Deprecation field to given value.

### HasDeprecation

`func (o *Link) HasDeprecation() bool`

HasDeprecation returns a boolean if a field has been set.

### GetName

`func (o *Link) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Link) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Link) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Link) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfile

`func (o *Link) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Link) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Link) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Link) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTitle

`func (o *Link) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Link) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Link) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Link) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHreflang

`func (o *Link) GetHreflang() string`

GetHreflang returns the Hreflang field if non-nil, zero value otherwise.

### GetHreflangOk

`func (o *Link) GetHreflangOk() (*string, bool)`

GetHreflangOk returns a tuple with the Hreflang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHreflang

`func (o *Link) SetHreflang(v string)`

SetHreflang sets Hreflang field to given value.

### HasHreflang

`func (o *Link) HasHreflang() bool`

HasHreflang returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


