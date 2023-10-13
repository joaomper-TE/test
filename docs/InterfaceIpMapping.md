# InterfaceIpMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceName** | Pointer to **string** | Name of the mapping | [optional] [readonly] 
**IpAddresses** | Pointer to **[]string** | Array of ipAddress entries | [optional] [readonly] 

## Methods

### NewInterfaceIpMapping

`func NewInterfaceIpMapping() *InterfaceIpMapping`

NewInterfaceIpMapping instantiates a new InterfaceIpMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceIpMappingWithDefaults

`func NewInterfaceIpMappingWithDefaults() *InterfaceIpMapping`

NewInterfaceIpMappingWithDefaults instantiates a new InterfaceIpMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceName

`func (o *InterfaceIpMapping) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *InterfaceIpMapping) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *InterfaceIpMapping) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *InterfaceIpMapping) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpAddresses

`func (o *InterfaceIpMapping) GetIpAddresses() []string`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *InterfaceIpMapping) GetIpAddressesOk() (*[]string, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *InterfaceIpMapping) SetIpAddresses(v []string)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *InterfaceIpMapping) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


