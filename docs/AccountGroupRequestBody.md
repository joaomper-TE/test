# AccountGroupRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupName** | **string** | The name of the account group | 
**Agents** | Pointer to **[]string** | To grant access to enterprise agents, specify the agent list. Note that this is not an additive list - the full list must be specified if changing access to agents. | [optional] 

## Methods

### NewAccountGroupRequestBody

`func NewAccountGroupRequestBody(accountGroupName string, ) *AccountGroupRequestBody`

NewAccountGroupRequestBody instantiates a new AccountGroupRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupRequestBodyWithDefaults

`func NewAccountGroupRequestBodyWithDefaults() *AccountGroupRequestBody`

NewAccountGroupRequestBodyWithDefaults instantiates a new AccountGroupRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupName

`func (o *AccountGroupRequestBody) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *AccountGroupRequestBody) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *AccountGroupRequestBody) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.


### GetAgents

`func (o *AccountGroupRequestBody) GetAgents() []string`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *AccountGroupRequestBody) GetAgentsOk() (*[]string, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *AccountGroupRequestBody) SetAgents(v []string)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *AccountGroupRequestBody) HasAgents() bool`

HasAgents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


