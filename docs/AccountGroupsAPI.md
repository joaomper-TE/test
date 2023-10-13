# \AccountGroupsAPI

All URIs are relative to *https://api.stg.thousandeyes.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountGroup**](AccountGroupsAPI.md#CreateAccountGroup) | **Post** /v7/account-groups | Create account group
[**DeleteAccountGroup**](AccountGroupsAPI.md#DeleteAccountGroup) | **Delete** /v7/account-groups/{id} | Delete account group
[**GetAccountGroup**](AccountGroupsAPI.md#GetAccountGroup) | **Get** /v7/account-groups/{id} | Retrieve account group
[**GetAccountGroups**](AccountGroupsAPI.md#GetAccountGroups) | **Get** /v7/account-groups | List account groups
[**UpdateAccountGroup**](AccountGroupsAPI.md#UpdateAccountGroup) | **Put** /v7/account-groups/{id} | Update account group



## CreateAccountGroup

> CreateAccountGroup201Response CreateAccountGroup(ctx).AccountGroupRequestBody(accountGroupRequestBody).Expand(expand).Execute()

Create account group



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
    accountGroupRequestBody := *openapiclient.NewAccountGroupRequestBody("My testing account group") // AccountGroupRequestBody | 
    expand := []string{"Expand_example"} // []string | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the \"User\" resource, pass the `?expand=user` query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountGroupsAPI.CreateAccountGroup(context.Background()).AccountGroupRequestBody(accountGroupRequestBody).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.CreateAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountGroup`: CreateAccountGroup201Response
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.CreateAccountGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountGroupRequestBody** | [**AccountGroupRequestBody**](AccountGroupRequestBody.md) |  | 
 **expand** | **[]string** | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the \&quot;User\&quot; resource, pass the &#x60;?expand&#x3D;user&#x60; query. | 

### Return type

[**CreateAccountGroup201Response**](CreateAccountGroup201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountGroup

> DeleteAccountGroup(ctx, id).Execute()

Delete account group



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
    id := "2067" // string | The ID of the desired account group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccountGroupsAPI.DeleteAccountGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.DeleteAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the desired account group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAccountGroup

> GetAccountGroup200Response GetAccountGroup(ctx, id).Expand(expand).Execute()

Retrieve account group



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
    id := "2067" // string | The ID of the desired account group.
    expand := []string{"Expand_example"} // []string | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the \"User\" resource, pass the `?expand=user` query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountGroupsAPI.GetAccountGroup(context.Background(), id).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.GetAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountGroup`: GetAccountGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.GetAccountGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the desired account group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **[]string** | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the \&quot;User\&quot; resource, pass the &#x60;?expand&#x3D;user&#x60; query. | 

### Return type

[**GetAccountGroup200Response**](GetAccountGroup200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountGroups

> GetAccountGroups200Response GetAccountGroups(ctx).Aid(aid).Execute()

List account groups



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
    resp, r, err := apiClient.AccountGroupsAPI.GetAccountGroups(context.Background()).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.GetAccountGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountGroups`: GetAccountGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.GetAccountGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

[**GetAccountGroups200Response**](GetAccountGroups200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountGroup

> GetAccountGroup200Response UpdateAccountGroup(ctx, id).AccountGroupRequestBody(accountGroupRequestBody).Expand(expand).Execute()

Update account group



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
    id := "2067" // string | The ID of the desired account group.
    accountGroupRequestBody := *openapiclient.NewAccountGroupRequestBody("My testing account group") // AccountGroupRequestBody | 
    expand := []string{"Expand_example"} // []string | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the \"User\" resource, pass the `?expand=user` query. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountGroupsAPI.UpdateAccountGroup(context.Background(), id).AccountGroupRequestBody(accountGroupRequestBody).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountGroupsAPI.UpdateAccountGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountGroup`: GetAccountGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `AccountGroupsAPI.UpdateAccountGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the desired account group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountGroupRequestBody** | [**AccountGroupRequestBody**](AccountGroupRequestBody.md) |  | 
 **expand** | **[]string** | Optional parameter that specifies whether or not account group related resources should be expanded. By default, no expansion takes place if the query parameter is not passed. For example, to expand the \&quot;User\&quot; resource, pass the &#x60;?expand&#x3D;user&#x60; query. | 

### Return type

[**GetAccountGroup200Response**](GetAccountGroup200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

