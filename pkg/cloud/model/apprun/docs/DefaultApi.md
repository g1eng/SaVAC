# \DefaultApi

All URIs are relative to *https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplication**](DefaultApi.md#DeleteApplication) | **Delete** /applications/{id} | アプリケーションを削除します。
[**DeleteApplicationVersion**](DefaultApi.md#DeleteApplicationVersion) | **Delete** /applications/{id}/versions/{version_id} | アプリケーションバージョンを削除します。
[**GetApplication**](DefaultApi.md#GetApplication) | **Get** /applications/{id} | アプリケーション詳細を取得します。
[**GetApplicationStatus**](DefaultApi.md#GetApplicationStatus) | **Get** /applications/{id}/status | アプリケーションステータスを取得します。
[**GetApplicationVersion**](DefaultApi.md#GetApplicationVersion) | **Get** /applications/{id}/versions/{version_id} | アプリケーションバージョン詳細を取得します。
[**GetPacketFilter**](DefaultApi.md#GetPacketFilter) | **Get** /applications/{id}/packet_filter | パケットフィルタを取得します。
[**GetUser**](DefaultApi.md#GetUser) | **Get** /user | 
[**ListApplicationTraffics**](DefaultApi.md#ListApplicationTraffics) | **Get** /applications/{id}/traffics | アプリケーショントラフィック分散を取得します。
[**ListApplicationVersions**](DefaultApi.md#ListApplicationVersions) | **Get** /applications/{id}/versions | アプリケーションバージョン一覧を取得します。
[**ListApplications**](DefaultApi.md#ListApplications) | **Get** /applications | アプリケーション一覧を取得します。
[**PatchApplication**](DefaultApi.md#PatchApplication) | **Patch** /applications/{id} | アプリケーションを部分的に変更します。
[**PatchPacketFilter**](DefaultApi.md#PatchPacketFilter) | **Patch** /applications/{id}/packet_filter | パケットフィルタを部分的に変更します。
[**PostApplication**](DefaultApi.md#PostApplication) | **Post** /applications | アプリケーションを作成します。
[**PostUser**](DefaultApi.md#PostUser) | **Post** /user | 
[**PutApplicationTraffic**](DefaultApi.md#PutApplicationTraffic) | **Put** /applications/{id}/traffics | アプリケーショントラフィック分散を変更します。



## DeleteApplication

> DeleteApplication(ctx, id).Execute()

アプリケーションを削除します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationVersion

> DeleteApplicationVersion(ctx, id, versionId).Execute()

アプリケーションバージョンを削除します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 
	versionId := "versionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.DeleteApplicationVersion(context.Background(), id, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, id).Execute()

アプリケーション詳細を取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: Application
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationStatus

> HandlerGetApplicationStatus GetApplicationStatus(ctx, id).Execute()

アプリケーションステータスを取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetApplicationStatus(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationStatus`: HandlerGetApplicationStatus
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlerGetApplicationStatus**](HandlerGetApplicationStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationVersion

> HandlerGetVersion GetApplicationVersion(ctx, id, versionId).Execute()

アプリケーションバージョン詳細を取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 
	versionId := "versionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetApplicationVersion(context.Background(), id, versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApplicationVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationVersion`: HandlerGetVersion
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApplicationVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HandlerGetVersion**](HandlerGetVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPacketFilter

> HandlerGetPacketFilter GetPacketFilter(ctx, id).Execute()

パケットフィルタを取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.GetPacketFilter(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPacketFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPacketFilter`: HandlerGetPacketFilter
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPacketFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPacketFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlerGetPacketFilter**](HandlerGetPacketFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUser(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.GetUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationTraffics

> HandlerListTraffics ListApplicationTraffics(ctx, id).Execute()

アプリケーショントラフィック分散を取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ListApplicationTraffics(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplicationTraffics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationTraffics`: HandlerListTraffics
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplicationTraffics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationTrafficsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlerListTraffics**](HandlerListTraffics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationVersions

> HandlerListVersions ListApplicationVersions(ctx, id).PageNum(pageNum).PageSize(pageSize).SortField(sortField).SortOrder(sortOrder).Execute()

アプリケーションバージョン一覧を取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 
	pageNum := int32(20) // int32 | 表示したいページ番号 (optional) (default to 1)
	pageSize := int32(10) // int32 | 表示したい1ページあたりのサイズ (optional) (default to 50)
	sortField := "created_at" // string | ソートしたいフィールド名 (optional) (default to "created_at")
	sortOrder := "asc" // string | ソート順（昇順、降順） (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ListApplicationVersions(context.Background(), id).PageNum(pageNum).PageSize(pageSize).SortField(sortField).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplicationVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicationVersions`: HandlerListVersions
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplicationVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNum** | **int32** | 表示したいページ番号 | [default to 1]
 **pageSize** | **int32** | 表示したい1ページあたりのサイズ | [default to 50]
 **sortField** | **string** | ソートしたいフィールド名 | [default to &quot;created_at&quot;]
 **sortOrder** | **string** | ソート順（昇順、降順） | [default to &quot;desc&quot;]

### Return type

[**HandlerListVersions**](HandlerListVersions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplications

> HandlerListApplications ListApplications(ctx).PageNum(pageNum).PageSize(pageSize).SortField(sortField).SortOrder(sortOrder).Execute()

アプリケーション一覧を取得します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	pageNum := int32(20) // int32 | 表示したいページ番号 (optional) (default to 1)
	pageSize := int32(10) // int32 | 表示したい1ページあたりのサイズ (optional) (default to 50)
	sortField := "created_at" // string | ソートしたいフィールド名 (optional) (default to "created_at")
	sortOrder := "asc" // string | ソート順（昇順、降順） (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.ListApplications(context.Background()).PageNum(pageNum).PageSize(pageSize).SortField(sortField).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplications`: HandlerListApplications
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNum** | **int32** | 表示したいページ番号 | [default to 1]
 **pageSize** | **int32** | 表示したい1ページあたりのサイズ | [default to 50]
 **sortField** | **string** | ソートしたいフィールド名 | [default to &quot;created_at&quot;]
 **sortOrder** | **string** | ソート順（昇順、降順） | [default to &quot;desc&quot;]

### Return type

[**HandlerListApplications**](HandlerListApplications.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchApplication

> HandlerPatchApplication PatchApplication(ctx, id).PatchApplicationBody(patchApplicationBody).Execute()

アプリケーションを部分的に変更します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 
	patchApplicationBody := *openapiclient.NewPatchApplicationBody() // PatchApplicationBody | 部分的に変更するアプリケーションの情報を入力します。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PatchApplication(context.Background(), id).PatchApplicationBody(patchApplicationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchApplication`: HandlerPatchApplication
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchApplicationBody** | [**PatchApplicationBody**](PatchApplicationBody.md) | 部分的に変更するアプリケーションの情報を入力します。 | 

### Return type

[**HandlerPatchApplication**](HandlerPatchApplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPacketFilter

> HandlerPatchPacketFilter PatchPacketFilter(ctx, id).PatchPacketFilter(patchPacketFilter).Execute()

パケットフィルタを部分的に変更します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	patchPacketFilter := *openapiclient.NewPatchPacketFilter() // PatchPacketFilter | 部分的に変更するパケットフィルタの情報を入力します。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PatchPacketFilter(context.Background(), id).PatchPacketFilter(patchPacketFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PatchPacketFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchPacketFilter`: HandlerPatchPacketFilter
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PatchPacketFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPacketFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchPacketFilter** | [**PatchPacketFilter**](PatchPacketFilter.md) | 部分的に変更するパケットフィルタの情報を入力します。 | 

### Return type

[**HandlerPatchPacketFilter**](HandlerPatchPacketFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApplication

> Application PostApplication(ctx).PostApplicationBody(postApplicationBody).Execute()

アプリケーションを作成します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	postApplicationBody := *openapiclient.NewPostApplicationBody("アプリケーション1", int32(60), int32(8080), int32(0), int32(2), []openapiclient.PostApplicationBodyComponent{*openapiclient.NewPostApplicationBodyComponent("コンポーネント1", "0.1", "256Mi", *openapiclient.NewPostApplicationBodyComponentDeploySource())}) // PostApplicationBody | 作成するアプリケーションの情報を入力します。

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PostApplication(context.Background()).PostApplicationBody(postApplicationBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApplication`: Application
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postApplicationBody** | [**PostApplicationBody**](PostApplicationBody.md) | 作成するアプリケーションの情報を入力します。 | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUser

> PostUser(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultApi.PostUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostUserRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutApplicationTraffic

> HandlerPutTraffics PutApplicationTraffic(ctx, id).Traffic(traffic).Execute()

アプリケーショントラフィック分散を変更します。



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/g1eng/savac/pkg/cloud/model/apprun"
)

func main() {
	id := "id_example" // string | 
	traffic := []openapiclient.Traffic{*openapiclient.NewTraffic("Version1", false, int32(100))} // []Traffic | トラフィック分散の割合を入力します。 version_nameまたはis_latest_versionのどちらかを指定して何％の割合でトラフィックを転送するかを入力します。 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultApi.PutApplicationTraffic(context.Background(), id).Traffic(traffic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PutApplicationTraffic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutApplicationTraffic`: HandlerPutTraffics
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PutApplicationTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutApplicationTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **traffic** | [**[]Traffic**](Traffic.md) | トラフィック分散の割合を入力します。 version_nameまたはis_latest_versionのどちらかを指定して何％の割合でトラフィックを転送するかを入力します。  | 

### Return type

[**HandlerPutTraffics**](HandlerPutTraffics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

