# AccountGroupUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupName** | Pointer to **string** | The name of the account group | [optional] 
**Aid** | Pointer to **string** | Unique ID of the account group. | [optional] 

## Methods

### NewAccountGroupUser

`func NewAccountGroupUser() *AccountGroupUser`

NewAccountGroupUser instantiates a new AccountGroupUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupUserWithDefaults

`func NewAccountGroupUserWithDefaults() *AccountGroupUser`

NewAccountGroupUserWithDefaults instantiates a new AccountGroupUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupName

`func (o *AccountGroupUser) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *AccountGroupUser) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *AccountGroupUser) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *AccountGroupUser) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAid

`func (o *AccountGroupUser) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *AccountGroupUser) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *AccountGroupUser) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *AccountGroupUser) HasAid() bool`

HasAid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


