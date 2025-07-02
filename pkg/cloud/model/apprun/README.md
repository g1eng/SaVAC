# Go API client for apprun

---
「AppRun」が提供するAPIの利用方法とサンプルを公開しております。

# 基本的な使い方

## APIキーの発行

APIを利用するためには、認証のための「APIキー」が必要です。事前にキーを発行しておきます。  
APIキーは「ユーザーID」「パスワード」に相当する「トークン」と呼ばれる認証情報で構成されています。

|   項目名   | APIキー発行時の項目名        | このドキュメント内での例             |
|------------|------------------------------|--------------------------------------|
| ユーザーID | アクセストークン(UUID)       | 01234567-89ab-cdef-0123-456789abcdef |
| パスワード | アクセストークンシークレット | SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM |

<div class=\"warning\">
<b>操作マニュアル</b><br />
<ul><li><a href=\"https://manual.sakura.ad.jp/cloud/api/apikey.html\">APIキー | さくらのクラウド ドキュメント</a></li></ul>
</div>

## 入力パラメータ

APIの入力には送信先URLに対して、いくつかのヘッダーとAPIキーを送信します。

* 認証方式はHTTP Basic認証です。APIキーのアクセストークンをユーザーID、アクセストークンシークレットをパスワードとして指定します。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     'https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications'
```

## 出力結果と応答コード（HTTPステータスコード）

APIからの結果は、「応答コード（HTTPステータスコード）」と、「JSON形式(UTF-8)の結果」として出力されます。

応答コードは、リクエストが成功したのか、失敗したのか大まかな情報を判断することができるもので、例えば失敗したときには、なぜこのような結果になったのかなど、具体的な情報は応答コードと主に返された本文を見ることで把握することができます。

| 結果                                | 応答コード/status   |
|-------------------------------------|---------------------|
| 成功（要求を受け付けた）            | 2xx                 |
| 失敗（要求が受け付けられなかった）  | 4xx, 5xx            |

```
# 出力結果サンプル（レスポンスヘッダー）
HTTP/1.1 200 OK
Server: nginx
Date: Tue, 16 Nov 2021 12:39:48 GMT
Content-Type: application/json; charset=UTF-8
Content-Length: 443
Connection: keep-alive
Status: 200 OK
Pragma: no-cache
Cache-Control: no-cache
X-Sakura-Proxy-Microtime: 66245
X-Sakura-Proxy-Decode-Microtime: 62
X-Sakura-Content-Length: 443
X-Sakura-Serial: 86ab6c743f72aa5ea6f17e254fd5f803
X-Content-Type-Options: nosniff
X-XSS-Protection: 1; mode=block
X-Frame-Options: DENY
X-Sakura-Encode-Microtime: 260
Vary: Accept-Encoding
```

```
# 出力結果サンプル（レスポンスボディー）
{
  \"error\": {
    \"code\": 401,
    \"message\": \"Login Required\",
    \"errors\": [
      {
        \"domain\": \"global\",
        \"reason\": \"required\",
        \"message\": \"Login Required\",
        \"location_type\": \"header\",
        \"location\": \"Authorization\"
      }
    ]
  }
}
```

# 利用例

## 1.ユーザーの作成

AppRunの利用を開始するには**ユーザー**を作成します。

ユーザーとは、AppRunを利用するための独立したユーザーであり、ユーザー作成および削除による料金の発生はございません。  
なお、すでにユーザーを作成済みの場合は、再度ユーザーの作成は不要です。

ユーザーを作成するには以下のような入力を行います。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X POST \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/user
```

ユーザーの作成が完了すると、

* アプリケーションの作成、更新、削除
* バージョンの確認、削除
* トラフィック分散の確認、変更

などの操作が可能になります。

## 2.アプリケーションの作成、取得、更新、削除

ユーザーを作成後、**アプリケーション**の作成、更新、削除が可能になります。

アプリケーションを作成するには以下のような入力を行います。

```
# 入力サンプル
vi request_body.json
cat request_body.json
{
  \"name\": \"Application\",
  \"timeout_seconds\": 60,
  \"port\": 8080,
  \"min_scale\": 0,
  \"max_scale\": 1,
  \"components\": [
    {
      \"name\": \"Component01\",
      \"max_cpu\": \"0.1\",
      \"max_memory\": \"256Mi\",
      \"deploy_source\": {
        \"container_registry\": {
          \"image\": \"my-app.sakuracr.jp/my-app:latest\"
        }
      },
      \"env\": [
        {
          \"key\": \"TARGET\",
          \"value\": \"World\"
        }
      ],
      \"probe\": {
        \"http_get\": {
          \"path\": \"/\",
          \"port\": 8080,
          \"headers\": [
            {
              \"name\": \"Custom-Header\",
              \"value\": \"Awesome\"
            }
          ]
        }
      }
    }
  ]
}
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X POST \\
     -d '@request_body.json' \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications
```

上記で作成したアプリケーションを取得するには以下のような入力を行います。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X GET \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}
```

上記で作成したアプリケーションを更新するには以下のような入力を行います。

```
# 入力サンプル
vi request_body.json
cat request_body.json
{
  \"components\": [
    {
      \"name\": \"Component01 updated\",
      \"max_cpu\": \"0.1\",
      \"max_memory\": \"256Mi\",
      \"deploy_source\": {
        \"container_registry\": {
          \"image\": \"my-app.sakuracr.jp/my-app-v2:latest\"
        }
      }
    }
  ],
  \"all_traffic_available\": true
}

curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X PATCH \\
     -d '@request_body.json' \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}
```

上記で作成したアプリケーションを削除するには以下のような入力を行います。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X DELETE \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}
```

## 3.バージョンの取得、削除

アプリケーションを作成、更新した際、その設定情報をバージョンとして保存します。

バージョンを取得するには以下のような入力を行います。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X GET \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/versions/{version_id}
```

上記で作成したバージョンを削除するには以下のような入力を行います。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X DELETE \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/versions/{version_id}
```

## 4.トラフィック分散の確認、変更

アプリケーションは指定のバージョンへトラフィックを分散します。

トラフィック分散を確認するには以下のような入力を行います。

```
# 入力サンプル
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X GET \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/traffics
```

トラフィック分散を変更するには以下のような入力を行います。

```
# 入力サンプル
vi request_body.json
cat request_body.json
[
  {
    \"is_latest_version\": true,
    \"percent\": 50
  },
  {
    \"version_name\": \"Application-861850d6-8240-7c31-9b69-80ea4466918d-1726726814\",
    \"percent\": 50
  }
]
curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\
     -X PUT \\
     https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/traffics
```
----

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Generator version: 7.13.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import apprun "github.com/g1eng/savac/pkg/cloud/model/apprun"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `apprun.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), apprun.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `apprun.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), apprun.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `apprun.ContextOperationServerIndices` and `apprun.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), apprun.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), apprun.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**DeleteApplication**](docs/DefaultApi.md#deleteapplication) | **Delete** /applications/{id} | アプリケーションを削除します。
*DefaultApi* | [**DeleteApplicationVersion**](docs/DefaultApi.md#deleteapplicationversion) | **Delete** /applications/{id}/versions/{version_id} | アプリケーションバージョンを削除します。
*DefaultApi* | [**GetApplication**](docs/DefaultApi.md#getapplication) | **Get** /applications/{id} | アプリケーション詳細を取得します。
*DefaultApi* | [**GetApplicationStatus**](docs/DefaultApi.md#getapplicationstatus) | **Get** /applications/{id}/status | アプリケーションステータスを取得します。
*DefaultApi* | [**GetApplicationVersion**](docs/DefaultApi.md#getapplicationversion) | **Get** /applications/{id}/versions/{version_id} | アプリケーションバージョン詳細を取得します。
*DefaultApi* | [**GetPacketFilter**](docs/DefaultApi.md#getpacketfilter) | **Get** /applications/{id}/packet_filter | パケットフィルタを取得します。
*DefaultApi* | [**GetUser**](docs/DefaultApi.md#getuser) | **Get** /user | 
*DefaultApi* | [**ListApplicationTraffics**](docs/DefaultApi.md#listapplicationtraffics) | **Get** /applications/{id}/traffics | アプリケーショントラフィック分散を取得します。
*DefaultApi* | [**ListApplicationVersions**](docs/DefaultApi.md#listapplicationversions) | **Get** /applications/{id}/versions | アプリケーションバージョン一覧を取得します。
*DefaultApi* | [**ListApplications**](docs/DefaultApi.md#listapplications) | **Get** /applications | アプリケーション一覧を取得します。
*DefaultApi* | [**PatchApplication**](docs/DefaultApi.md#patchapplication) | **Patch** /applications/{id} | アプリケーションを部分的に変更します。
*DefaultApi* | [**PatchPacketFilter**](docs/DefaultApi.md#patchpacketfilter) | **Patch** /applications/{id}/packet_filter | パケットフィルタを部分的に変更します。
*DefaultApi* | [**PostApplication**](docs/DefaultApi.md#postapplication) | **Post** /applications | アプリケーションを作成します。
*DefaultApi* | [**PostUser**](docs/DefaultApi.md#postuser) | **Post** /user | 
*DefaultApi* | [**PutApplicationTraffic**](docs/DefaultApi.md#putapplicationtraffic) | **Put** /applications/{id}/traffics | アプリケーショントラフィック分散を変更します。


## Documentation For Models

 - [Application](docs/Application.md)
 - [HandlerApplicationComponent](docs/HandlerApplicationComponent.md)
 - [HandlerApplicationComponentDeploySource](docs/HandlerApplicationComponentDeploySource.md)
 - [HandlerApplicationComponentDeploySourceContainerRegistry](docs/HandlerApplicationComponentDeploySourceContainerRegistry.md)
 - [HandlerApplicationComponentEnv](docs/HandlerApplicationComponentEnv.md)
 - [HandlerApplicationComponentProbe](docs/HandlerApplicationComponentProbe.md)
 - [HandlerApplicationComponentProbeHttpGet](docs/HandlerApplicationComponentProbeHttpGet.md)
 - [HandlerApplicationComponentProbeHttpGetHeader](docs/HandlerApplicationComponentProbeHttpGetHeader.md)
 - [HandlerGetApplicationStatus](docs/HandlerGetApplicationStatus.md)
 - [HandlerGetPacketFilter](docs/HandlerGetPacketFilter.md)
 - [HandlerGetPacketFilterSettingsInner](docs/HandlerGetPacketFilterSettingsInner.md)
 - [HandlerGetVersion](docs/HandlerGetVersion.md)
 - [HandlerListApplications](docs/HandlerListApplications.md)
 - [HandlerListApplicationsData](docs/HandlerListApplicationsData.md)
 - [HandlerListApplicationsMeta](docs/HandlerListApplicationsMeta.md)
 - [HandlerListTraffics](docs/HandlerListTraffics.md)
 - [HandlerListVersions](docs/HandlerListVersions.md)
 - [HandlerListVersionsMeta](docs/HandlerListVersionsMeta.md)
 - [HandlerPatchApplication](docs/HandlerPatchApplication.md)
 - [HandlerPatchPacketFilter](docs/HandlerPatchPacketFilter.md)
 - [HandlerPatchPacketFilterSettingsInner](docs/HandlerPatchPacketFilterSettingsInner.md)
 - [HandlerPutTraffics](docs/HandlerPutTraffics.md)
 - [ModelDefaultError](docs/ModelDefaultError.md)
 - [ModelDefaultErrorError](docs/ModelDefaultErrorError.md)
 - [ModelError](docs/ModelError.md)
 - [PatchApplicationBody](docs/PatchApplicationBody.md)
 - [PatchApplicationBodyComponent](docs/PatchApplicationBodyComponent.md)
 - [PatchApplicationBodyComponentDeploySource](docs/PatchApplicationBodyComponentDeploySource.md)
 - [PatchApplicationBodyComponentDeploySourceContainerRegistry](docs/PatchApplicationBodyComponentDeploySourceContainerRegistry.md)
 - [PatchApplicationBodyComponentEnv](docs/PatchApplicationBodyComponentEnv.md)
 - [PatchApplicationBodyComponentProbe](docs/PatchApplicationBodyComponentProbe.md)
 - [PatchApplicationBodyComponentProbeHttpGet](docs/PatchApplicationBodyComponentProbeHttpGet.md)
 - [PatchApplicationBodyComponentProbeHttpGetHeader](docs/PatchApplicationBodyComponentProbeHttpGetHeader.md)
 - [PatchPacketFilter](docs/PatchPacketFilter.md)
 - [PostApplicationBody](docs/PostApplicationBody.md)
 - [PostApplicationBodyComponent](docs/PostApplicationBodyComponent.md)
 - [PostApplicationBodyComponentDeploySource](docs/PostApplicationBodyComponentDeploySource.md)
 - [PostApplicationBodyComponentDeploySourceContainerRegistry](docs/PostApplicationBodyComponentDeploySourceContainerRegistry.md)
 - [PostApplicationBodyComponentEnv](docs/PostApplicationBodyComponentEnv.md)
 - [PostApplicationBodyComponentProbe](docs/PostApplicationBodyComponentProbe.md)
 - [PostApplicationBodyComponentProbeHttpGet](docs/PostApplicationBodyComponentProbeHttpGet.md)
 - [PostApplicationBodyComponentProbeHttpGetHeader](docs/PostApplicationBodyComponentProbeHttpGetHeader.md)
 - [Traffic](docs/Traffic.md)
 - [Version](docs/Version.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



