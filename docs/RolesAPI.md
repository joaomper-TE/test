# \RolesAPI

All URIs are relative to *https://api.stg.thousandeyes.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRole**](RolesAPI.md#CreateRole) | **Post** /v7/roles | Create role
[**DeleteRole**](RolesAPI.md#DeleteRole) | **Delete** /v7/roles/{id} | Delete role
[**GetRole**](RolesAPI.md#GetRole) | **Get** /v7/roles/{id} | Retrieve role
[**GetRoles**](RolesAPI.md#GetRoles) | **Get** /v7/roles | List roles
[**UpdateRole**](RolesAPI.md#UpdateRole) | **Put** /v7/roles/{id} | Update role



## CreateRole

> CreateRole201Response CreateRole(ctx).RoleRequestBody(roleRequestBody).Aid(aid).Execute()

Create role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/joaomper/test/v1"
)

func main() {
    roleRequestBody := *openapiclient.NewRoleRequestBody() // RoleRequestBody | 
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.CreateRole(context.Background()).RoleRequestBody(roleRequestBody).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.CreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRole`: CreateRole201Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.CreateRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleRequestBody** | [**RoleRequestBody**](RoleRequestBody.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

[**CreateRole201Response**](CreateRole201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRole

> DeleteRole(ctx, id).Aid(aid).Execute()

Delete role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/joaomper/test/v1"
)

func main() {
    id := "23" // string | The ID of the desired role.
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RolesAPI.DeleteRole(context.Background(), id).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.DeleteRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the desired role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRole

> CreateRole201Response GetRole(ctx, id).Aid(aid).Execute()

Retrieve role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/joaomper/test/v1"
)

func main() {
    id := "23" // string | The ID of the desired role.
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.GetRole(context.Background(), id).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRole`: CreateRole201Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the desired role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

[**CreateRole201Response**](CreateRole201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoles

> GetRoles200Response GetRoles(ctx).Aid(aid).Execute()

List roles



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/joaomper/test/v1"
)

func main() {
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.GetRoles(context.Background()).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.GetRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoles`: GetRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.GetRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

[**GetRoles200Response**](GetRoles200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRole

> CreateRole201Response UpdateRole(ctx, id).RoleRequestBody(roleRequestBody).Aid(aid).Execute()

Update role



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/joaomper/test/v1"
)

func main() {
    id := "23" // string | The ID of the desired role.
    roleRequestBody := *openapiclient.NewRoleRequestBody() // RoleRequestBody | 
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesAPI.UpdateRole(context.Background(), id).RoleRequestBody(roleRequestBody).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesAPI.UpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRole`: CreateRole201Response
    fmt.Fprintf(os.Stdout, "Response from `RolesAPI.UpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the desired role. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleRequestBody** | [**RoleRequestBody**](RoleRequestBody.md) |  | 
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

[**CreateRole201Response**](CreateRole201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

