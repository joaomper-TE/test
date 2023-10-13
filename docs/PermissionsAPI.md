# \PermissionsAPI

All URIs are relative to *https://api.stg.thousandeyes.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissions**](PermissionsAPI.md#GetPermissions) | **Get** /v7/permissions | List assignable permissions



## GetPermissions

> GetPermissions200Response GetPermissions(ctx).Aid(aid).Execute()

List assignable permissions



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
    resp, r, err := apiClient.PermissionsAPI.GetPermissions(context.Background()).Aid(aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissions`: GetPermissions200Response
    fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 

### Return type

[**GetPermissions200Response**](GetPermissions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

