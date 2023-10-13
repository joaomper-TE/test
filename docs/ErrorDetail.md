# ErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code for the agent error. | [optional] [readonly] 
**Description** | Pointer to **string** | Description for the agent error. | [optional] [readonly] 

## Methods

### NewErrorDetail

`func NewErrorDetail() *ErrorDetail`

NewErrorDetail instantiates a new ErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailWithDefaults

`func NewErrorDetailWithDefaults() *ErrorDetail`

NewErrorDetailWithDefaults instantiates a new ErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


