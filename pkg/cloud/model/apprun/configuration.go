/*
AppRun β APIドキュメント

--- 「AppRun」が提供するAPIの利用方法とサンプルを公開しております。  # 基本的な使い方  ## APIキーの発行  APIを利用するためには、認証のための「APIキー」が必要です。事前にキーを発行しておきます。   APIキーは「ユーザーID」「パスワード」に相当する「トークン」と呼ばれる認証情報で構成されています。  |   項目名   | APIキー発行時の項目名        | このドキュメント内での例             | |------------|------------------------------|--------------------------------------| | ユーザーID | アクセストークン(UUID)       | 01234567-89ab-cdef-0123-456789abcdef | | パスワード | アクセストークンシークレット | SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM |  <div class=\"warning\"> <b>操作マニュアル</b><br /> <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/api/apikey.html\">APIキー | さくらのクラウド ドキュメント</a></li></ul> </div>  ## 入力パラメータ  APIの入力には送信先URLに対して、いくつかのヘッダーとAPIキーを送信します。  * 認証方式はHTTP Basic認証です。APIキーのアクセストークンをユーザーID、アクセストークンシークレットをパスワードとして指定します。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications' ```  ## 出力結果と応答コード（HTTPステータスコード）  APIからの結果は、「応答コード（HTTPステータスコード）」と、「JSON形式(UTF-8)の結果」として出力されます。  応答コードは、リクエストが成功したのか、失敗したのか大まかな情報を判断することができるもので、例えば失敗したときには、なぜこのような結果になったのかなど、具体的な情報は応答コードと主に返された本文を見ることで把握することができます。  | 結果                                | 応答コード/status   | |-------------------------------------|---------------------| | 成功（要求を受け付けた）            | 2xx                 | | 失敗（要求が受け付けられなかった）  | 4xx, 5xx            |  ``` # 出力結果サンプル（レスポンスヘッダー） HTTP/1.1 200 OK Server: nginx Date: Tue, 16 Nov 2021 12:39:48 GMT Content-Type: application/json; charset=UTF-8 Content-Length: 443 Connection: keep-alive Status: 200 OK Pragma: no-cache Cache-Control: no-cache X-Sakura-Proxy-Microtime: 66245 X-Sakura-Proxy-Decode-Microtime: 62 X-Sakura-Content-Length: 443 X-Sakura-Serial: 86ab6c743f72aa5ea6f17e254fd5f803 X-Content-Type-Options: nosniff X-XSS-Protection: 1; mode=block X-Frame-Options: DENY X-Sakura-Encode-Microtime: 260 Vary: Accept-Encoding ```  ``` # 出力結果サンプル（レスポンスボディー） {   \"error\": {     \"code\": 401,     \"message\": \"Login Required\",     \"errors\": [       {         \"domain\": \"global\",         \"reason\": \"required\",         \"message\": \"Login Required\",         \"location_type\": \"header\",         \"location\": \"Authorization\"       }     ]   } } ```  # 利用例  ## 1.ユーザーの作成  AppRunの利用を開始するには**ユーザー**を作成します。  ユーザーとは、AppRunを利用するための独立したユーザーであり、ユーザー作成および削除による料金の発生はございません。   なお、すでにユーザーを作成済みの場合は、再度ユーザーの作成は不要です。  ユーザーを作成するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/user ```  ユーザーの作成が完了すると、  * アプリケーションの作成、更新、削除 * バージョンの確認、削除 * トラフィック分散の確認、変更  などの操作が可能になります。  ## 2.アプリケーションの作成、取得、更新、削除  ユーザーを作成後、**アプリケーション**の作成、更新、削除が可能になります。  アプリケーションを作成するには以下のような入力を行います。  ``` # 入力サンプル vi request_body.json cat request_body.json {   \"name\": \"Application\",   \"timeout_seconds\": 60,   \"port\": 8080,   \"min_scale\": 0,   \"max_scale\": 1,   \"components\": [     {       \"name\": \"Component01\",       \"max_cpu\": \"0.1\",       \"max_memory\": \"256Mi\",       \"deploy_source\": {         \"container_registry\": {           \"image\": \"my-app.sakuracr.jp/my-app:latest\"         }       },       \"env\": [         {           \"key\": \"TARGET\",           \"value\": \"World\"         }       ],       \"probe\": {         \"http_get\": {           \"path\": \"/\",           \"port\": 8080,           \"headers\": [             {               \"name\": \"Custom-Header\",               \"value\": \"Awesome\"             }           ]         }       }     }   ] } curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      -d '@request_body.json' \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications ```  上記で作成したアプリケーションを取得するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X GET \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id} ```  上記で作成したアプリケーションを更新するには以下のような入力を行います。  ``` # 入力サンプル vi request_body.json cat request_body.json {   \"components\": [     {       \"name\": \"Component01 updated\",       \"max_cpu\": \"0.1\",       \"max_memory\": \"256Mi\",       \"deploy_source\": {         \"container_registry\": {           \"image\": \"my-app.sakuracr.jp/my-app-v2:latest\"         }       }     }   ],   \"all_traffic_available\": true }  curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PATCH \\      -d '@request_body.json' \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id} ```  上記で作成したアプリケーションを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id} ```  ## 3.バージョンの取得、削除  アプリケーションを作成、更新した際、その設定情報をバージョンとして保存します。  バージョンを取得するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X GET \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/versions/{version_id} ```  上記で作成したバージョンを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/versions/{version_id} ```  ## 4.トラフィック分散の確認、変更  アプリケーションは指定のバージョンへトラフィックを分散します。  トラフィック分散を確認するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X GET \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/traffics ```  トラフィック分散を変更するには以下のような入力を行います。  ``` # 入力サンプル vi request_body.json cat request_body.json [   {     \"is_latest_version\": true,     \"percent\": 50   },   {     \"version_name\": \"Application-861850d6-8240-7c31-9b69-80ea4466918d-1726726814\",     \"percent\": 50   } ] curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PUT \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/traffics ``` ----

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apprun

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
