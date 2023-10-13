# UnauthorizedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewUnauthorizedError

`func NewUnauthorizedError() *UnauthorizedError`

NewUnauthorizedError instantiates a new UnauthorizedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorWithDefaults

`func NewUnauthorizedErrorWithDefaults() *UnauthorizedError`

NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *UnauthorizedError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UnauthorizedError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UnauthorizedError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *UnauthorizedError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorDescription

`func (o *UnauthorizedError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *UnauthorizedError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *UnauthorizedError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *UnauthorizedError) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


