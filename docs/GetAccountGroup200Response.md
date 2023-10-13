# GetAccountGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountGroupName** | Pointer to **string** | Unique name for the account group within an organization. | [optional] 
**Aid** | Pointer to **string** | Unique ID representing the account group. | [optional] 
**IsCurrentAccountGroup** | Pointer to **bool** | Indicates whether the requested aid is the context of the current account. | [optional] 
**IsDefaultAccountGroup** | Pointer to **bool** | Indicates whether the aid is the default one for the requesting user. | [optional] 
**OrganizationName** | Pointer to **string** | (Optional) Indicates whether the aid is the default one for the requesting user. | [optional] 
**Users** | Pointer to [**[]UserAccountGroup**](UserAccountGroup.md) |  | [optional] 
**Agents** | Pointer to [**[]EnterpriseAgent**](EnterpriseAgent.md) |  | [optional] 
**Links** | Pointer to [**SelfLinksLinks**](SelfLinksLinks.md) |  | [optional] 

## Methods

### NewGetAccountGroup200Response

`func NewGetAccountGroup200Response() *GetAccountGroup200Response`

NewGetAccountGroup200Response instantiates a new GetAccountGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountGroup200ResponseWithDefaults

`func NewGetAccountGroup200ResponseWithDefaults() *GetAccountGroup200Response`

NewGetAccountGroup200ResponseWithDefaults instantiates a new GetAccountGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountGroupName

`func (o *GetAccountGroup200Response) GetAccountGroupName() string`

GetAccountGroupName returns the AccountGroupName field if non-nil, zero value otherwise.

### GetAccountGroupNameOk

`func (o *GetAccountGroup200Response) GetAccountGroupNameOk() (*string, bool)`

GetAccountGroupNameOk returns a tuple with the AccountGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountGroupName

`func (o *GetAccountGroup200Response) SetAccountGroupName(v string)`

SetAccountGroupName sets AccountGroupName field to given value.

### HasAccountGroupName

`func (o *GetAccountGroup200Response) HasAccountGroupName() bool`

HasAccountGroupName returns a boolean if a field has been set.

### GetAid

`func (o *GetAccountGroup200Response) GetAid() string`

GetAid returns the Aid field if non-nil, zero value otherwise.

### GetAidOk

`func (o *GetAccountGroup200Response) GetAidOk() (*string, bool)`

GetAidOk returns a tuple with the Aid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAid

`func (o *GetAccountGroup200Response) SetAid(v string)`

SetAid sets Aid field to given value.

### HasAid

`func (o *GetAccountGroup200Response) HasAid() bool`

HasAid returns a boolean if a field has been set.

### GetIsCurrentAccountGroup

`func (o *GetAccountGroup200Response) GetIsCurrentAccountGroup() bool`

GetIsCurrentAccountGroup returns the IsCurrentAccountGroup field if non-nil, zero value otherwise.

### GetIsCurrentAccountGroupOk

`func (o *GetAccountGroup200Response) GetIsCurrentAccountGroupOk() (*bool, bool)`

GetIsCurrentAccountGroupOk returns a tuple with the IsCurrentAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentAccountGroup

`func (o *GetAccountGroup200Response) SetIsCurrentAccountGroup(v bool)`

SetIsCurrentAccountGroup sets IsCurrentAccountGroup field to given value.

### HasIsCurrentAccountGroup

`func (o *GetAccountGroup200Response) HasIsCurrentAccountGroup() bool`

HasIsCurrentAccountGroup returns a boolean if a field has been set.

### GetIsDefaultAccountGroup

`func (o *GetAccountGroup200Response) GetIsDefaultAccountGroup() bool`

GetIsDefaultAccountGroup returns the IsDefaultAccountGroup field if non-nil, zero value otherwise.

### GetIsDefaultAccountGroupOk

`func (o *GetAccountGroup200Response) GetIsDefaultAccountGroupOk() (*bool, bool)`

GetIsDefaultAccountGroupOk returns a tuple with the IsDefaultAccountGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAccountGroup

`func (o *GetAccountGroup200Response) SetIsDefaultAccountGroup(v bool)`

SetIsDefaultAccountGroup sets IsDefaultAccountGroup field to given value.

### HasIsDefaultAccountGroup

`func (o *GetAccountGroup200Response) HasIsDefaultAccountGroup() bool`

HasIsDefaultAccountGroup returns a boolean if a field has been set.

### GetOrganizationName

`func (o *GetAccountGroup200Response) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *GetAccountGroup200Response) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *GetAccountGroup200Response) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *GetAccountGroup200Response) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetUsers

`func (o *GetAccountGroup200Response) GetUsers() []UserAccountGroup`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetAccountGroup200Response) GetUsersOk() (*[]UserAccountGroup, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetAccountGroup200Response) SetUsers(v []UserAccountGroup)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetAccountGroup200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetAgents

`func (o *GetAccountGroup200Response) GetAgents() []EnterpriseAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *GetAccountGroup200Response) GetAgentsOk() (*[]EnterpriseAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *GetAccountGroup200Response) SetAgents(v []EnterpriseAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *GetAccountGroup200Response) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetLinks

`func (o *GetAccountGroup200Response) GetLinks() SelfLinksLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetAccountGroup200Response) GetLinksOk() (*SelfLinksLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetAccountGroup200Response) SetLinks(v SelfLinksLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetAccountGroup200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


