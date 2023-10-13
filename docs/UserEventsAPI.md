# \UserEventsAPI

All URIs are relative to *https://api.stg.thousandeyes.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserEvents**](UserEventsAPI.md#GetUserEvents) | **Get** /v7/audit-user-events/ | List activity log events



## GetUserEvents

> GetUserEvents200Response GetUserEvents(ctx).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Execute()

List activity log events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/joaomper/test/v1"
)

func main() {
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)
    window := "12h" // string | The `window` parameter helps you get data from a specific time in the past up to the present request time. You specify an amount of time back, like a number followed by an optional time interval type: `s` for seconds, `m` for minutes, `h` for hours, `d` for days, and `w` for weeks. If no type is given, it's assumed to be seconds. If you need a precise range, use startDate and endDate instead. (optional)
    startDate := time.Now() // time.Time | To focus on specific data, use the `startDate` and `endDate` parameters. Dates should follow the ISO 8601 date-time format, with hyphens for dates and colons for times. You need to include the complete time (hours, minutes, and seconds). The date and time can be separated by either a space or the letter T. For instance: `2012-01-01T00:00:00Z`. The date range includes both start and end dates. The time zone is UTC. Please note that this parameter can't be used alongside `window`. (optional)
    endDate := time.Now() // time.Time | The endDate parameter is optional; use alongside the `startDate` parameter. If not provided, the current time (at the request moment) is assumed. This parameter is incompatible with `window`. (optional)
    cursor := "cursor_example" // string | Optional opaque cursor that is used for pagination. Clients  must not use this parameter directly and should instead use the `_links` provided by the api. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserEventsAPI.GetUserEvents(context.Background()).Aid(aid).Window(window).StartDate(startDate).EndDate(endDate).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserEventsAPI.GetUserEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserEvents`: GetUserEvents200Response
    fmt.Fprintf(os.Stdout, "Response from `UserEventsAPI.GetUserEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aid** | **string** | A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. | 
 **window** | **string** | The &#x60;window&#x60; parameter helps you get data from a specific time in the past up to the present request time. You specify an amount of time back, like a number followed by an optional time interval type: &#x60;s&#x60; for seconds, &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. If no type is given, it&#39;s assumed to be seconds. If you need a precise range, use startDate and endDate instead. | 
 **startDate** | **time.Time** | To focus on specific data, use the &#x60;startDate&#x60; and &#x60;endDate&#x60; parameters. Dates should follow the ISO 8601 date-time format, with hyphens for dates and colons for times. You need to include the complete time (hours, minutes, and seconds). The date and time can be separated by either a space or the letter T. For instance: &#x60;2012-01-01T00:00:00Z&#x60;. The date range includes both start and end dates. The time zone is UTC. Please note that this parameter can&#39;t be used alongside &#x60;window&#x60;. | 
 **endDate** | **time.Time** | The endDate parameter is optional; use alongside the &#x60;startDate&#x60; parameter. If not provided, the current time (at the request moment) is assumed. This parameter is incompatible with &#x60;window&#x60;. | 
 **cursor** | **string** | Optional opaque cursor that is used for pagination. Clients  must not use this parameter directly and should instead use the &#x60;_links&#x60; provided by the api. | 

### Return type

[**GetUserEvents200Response**](GetUserEvents200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

