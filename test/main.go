package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "github.com/joaomper/test/v1"
)

func main() {
    aid := "2067" // string | A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response. (optional)

	// auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
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
