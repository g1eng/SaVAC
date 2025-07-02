/*
さくらのオブジェクトストレージ APIドキュメント

 ---  「さくらのオブジェクトストレージ」が提供するAPIの利用方法とサンプルを公開しております。  JSON 形式の OpenAPI 仕様は、以下の URL からダウンロードしてください。   <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.json\">JSON形式でダウンロード</a></li></ul>  # 基本的な使い方  ## APIキーの発行  APIを利用するためには、認証のための「APIキー」が必要です。事前にキーを発行しておきます。 APIキーは「ユーザーID」「パスワード」に相当する「トークン」と呼ばれる認証情報で構成されています。  |   項目名   | APIキー発行時の項目名        | このドキュメント内での例             | |------------|------------------------------|--------------------------------------| | ユーザーID | アクセストークン(UUID)       | 01234567-89ab-cdef-0123-456789abcdef | | パスワード | アクセストークンシークレット | SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM |  <div class=\"warning\"> <b>操作マニュアル</b><br /> <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/api/apikey.html\">APIキー | さくらのクラウド ドキュメント</a></li></ul> </div>  ## 入力パラメータ  APIの入力には送信先URLに対して、いくつかのヘッダーとAPIキーを送信します。  * APIのURLは以下の2つが存在します。※ 各APIの使い分けは後述します。   * `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/(エンドポイント)`   * `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/(エンドポイント)` * 認証方式はHTTP Basic認証です。APIキーのアクセストークンをユーザーID、アクセストークンシークレットをパスワードとして指定します。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters' ```  ## 出力結果と応答コード（HTTPステータスコード）  APIからの結果は、「応答コード（HTTPステータスコード）」と、「JSON形式(UTF-8)の結果」として出力されます。  応答コードは、リクエストが成功したのか、失敗したのか大まかな情報を判断することができるもので、例えば失敗したときには、なぜこのような結果になったのかなど、具体的な情報は応答コードと主に返された本文を見ることで把握することができます。  | 結果                                | 応答コード/status   | |-------------------------------------|---------------------| | 成功（要求を受け付けた）             | 2xx                 | | 失敗（要求が受け付けられなかった）  | 4xx, 5xx            |  ``` # 出力結果サンプル（レスポンスヘッダ） HTTP/1.1 200 OK Server: nginx Date: Tue, 16 Nov 2021 12:39:48 GMT Content-Type: application/json; charset=UTF-8 Content-Length: 443 Connection: keep-alive Status: 200 OK Pragma: no-cache Cache-Control: no-cache X-Sakura-Proxy-Microtime: 66245 X-Sakura-Proxy-Decode-Microtime: 62 X-Sakura-Content-Length: 443 X-Sakura-Serial: 86ab6c743f72aa5ea6f17e254fd5f803 X-Content-Type-Options: nosniff X-XSS-Protection: 1; mode=block X-Frame-Options: DENY X-Sakura-Encode-Microtime: 260 Vary: Accept-Encoding ```  ``` # 出力結果サンプル（レスポンスボディー） {   \"error\": {     \"code\": 404,     \"errors\": [       {         \"domain\": \"fed.objectstorage.sacloud\",         \"location\": \"clusters\",         \"location_type\": \"path_parameter\",         \"message\": \"Cluster was not found\",         \"reason\": \"not_found\"       }     ],     \"message\": \"Cluster was not found\",     \"trace_id\": \"0f36837633984f3fc8871f515e8efa24\"   } } ```  # 利用例  ## 1.接続先サイト一覧の取得  さくらのオブジェクトストレージを利用するには、まずバケット作成先となる**サイト**を取得・選択します。  サイト一覧を取得するには、以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters' ```  実行結果として、サイトのリストが返却されます。  ``` # 出力結果サンプル {   \"data\": [     {       \"api_zone\": [],       \"control_panel_url\": \"https://secure.sakura.ad.jp/objectstorage/\",       \"display_name_en_us\": \"Ishikari Site #1\",       \"display_name_ja\": \"石狩第1サイト\",       \"display_name\": \"石狩第1サイト\",       \"display_order\": 1,       \"endpoint_base\": \"isk01.sakurastorage.jp\",       \"id\": \"isk01\",       \"region\": \"jp-north-1\",       \"s3_endpoint\": \"s3.isk01.sakurastorage.jp\",       \"s3_endpoint_for_control_panel\": \"s3.cp.isk01.sakurastorage.jp\",       \"storage_zone\": []     }   ] } ```  得られたサイトID（上記の`id`フィールド）を確認します。これは後続の利用例で使用します。  ## 2.サイトアカウントの作成  上記のサイトから利用したいサイトIDを選択し（ここではisk01を選択することにします）、**サイトアカウント**を作成します。  サイトアカウントとは、サイトを利用するための独立したアカウントであり、サイトアカウント作成・削除による料金の発生はございません。 なお、すでにサイトアカウントを作成済みの場合は、再度サイトアカウントの作成は不要です。  サイトアカウントを作成するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account ```  サイトアカウントの作成が完了すると、選択したサイトにて  * バケットの作成・削除 * アクセスキーの発行・削除 * パーミッションキーの発行・削除  などの操作が可能になります。  ## 3.バケットの作成・削除  選択したサイトにてサイトアカウントを作成後、**バケット**の作成・削除が可能です。  バケットを作成するには以下のような入力を行います。     この時、選択したサイト（ここではisk01とします）をリクエストボディーに入れ、作成したいバケット名をパスパラメータに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PUT \\      -d '{\"cluster_id\": \"isk01\"}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/buckets/sample ```  上記で作成したバケットを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      -d '{\"cluster_id\": \"isk01\"}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/buckets/sample ```  ## 4.アクセスキーの発行・削除  選択したサイトにてサイトアカウントを作成後、**アクセスキー**の発行・削除が可能です。  アクセスキーを発行するには以下のような入力を行います。      ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account/keys ```  コマンド結果には以下のフィールドが含まれます。  * `created_at` : 作成日時 * `id` : アクセスキーID * `secret` : シークレットアクセスキー  ``` # 出力結果サンプル {   \"data\": {     \"created_at\": \"2021-11-04T07:42:41.121418479Z\",     \"id\": \"XPJK4SC9883N91RHR253\",     \"secret\": \"jqRaUo5l+EiEYqP8wos9exbmFfq4/vG8CLPYI2XN\"   } } ```  上記で作成したアクセスキーを削除するには以下のような入力を行います。     この時、削除したいアクセスキーIDをパスパラメータに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account/keys/XPJK4SC9883N91RHR253 ```  ## 5.パーミッション及びパーミッションアクセスキーの発行・削除  選択したサイトにてサイトアカウントを作成後且つバケットが1つ以上ある場合、**パーミッション**の発行・削除が可能です。  パーミッションを作成するには以下のような入力を行います。 この時、パーミッション名、パーミッションで制御したいバケットとそれに対する操作をリクエストボディーに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      -d '{\"display_name\": \"sample_permission\", \"bucket_controls\": [{\"bucket_name\": \"sample\", \"can_read\": true, \"can_write\": true}]}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions ```  作成が完了すると、パーミッションIDが含まれたレスポンスを受け取ります。 ``` # 出力サンプル {   \"data\": {     \"bucket_controls\": [       {         \"bucket_name\":\"sample\",         \"can_read\":true,         \"can_write\":true,         \"created_at\":\"2021-11-11T13:36:08.767118492Z\"       }     ],     \"created_at\":\"2021-11-11T13:36:08.690384415Z\",     \"display_name\":\"sample_permission\",     \"id\":619   } } ```  このパーミッションのアクセスキーを発行するには以下のような入力を行います。 この時、パーミッション作成時に発行されたID（ここでは619とします）をパスパラメータに含める必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619/keys ```  コマンド結果には以下のフィールドが含まれます。  * `created_at` : 作成日時 * `id` : アクセスキーID * `secret` : シークレットアクセスキー  ``` # 出力結果サンプル {   \"data\": {     \"created_at\": \"2021-11-04T07:42:41.121418479Z\",     \"id\": \"XPJK4SC9883N91RHR253\",     \"secret\": \"jqRaUo5l+EiEYqP8wos9exbmFfq4/vG8CLPYI2XN\"   } } ```  パーミッションアクセスキーを削除するには以下のような入力を行います。 この時、パーミッションアクセスキー発行時に出力されたIDをパスパラメータに含める必要があります。 ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619/keys/XPJK4SC9883N91RHR253 ```  パーミッションを削除するには以下のような入力を行います。 この時、パーミッション作成時に発行されたID（ここでは619とします）をパスパラメータに含める必要があります。 ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619 ``` ----

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package object_storage

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateAccountRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiCreateAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.CreateAccountExecute(r)
}

/*
CreateAccount サイトアカウントの作成

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAccountRequest
*/
func (a *DefaultApiService) CreateAccount(ctx context.Context) ApiCreateAccountRequest {
	return ApiCreateAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Account
func (a *DefaultApiService) CreateAccountExecute(r ApiCreateAccountRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateAccountKeyRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiCreateAccountKeyRequest) Execute() (*AccountKey, *http.Response, error) {
	return r.ApiService.CreateAccountKeyExecute(r)
}

/*
CreateAccountKey サイトアカウントのアクセスキーの発行

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAccountKeyRequest
*/
func (a *DefaultApiService) CreateAccountKey(ctx context.Context) ApiCreateAccountKeyRequest {
	return ApiCreateAccountKeyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountKey
func (a *DefaultApiService) CreateAccountKeyExecute(r ApiCreateAccountKeyRequest) (*AccountKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccountKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateAccountKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateBucketRequest struct {
	ctx                     context.Context
	ApiService              *DefaultApiService
	name                    string
	handlerPutBucketReqBody *HandlerPutBucketReqBody
}

// 作成したいバケット情報
func (r ApiCreateBucketRequest) HandlerPutBucketReqBody(handlerPutBucketReqBody HandlerPutBucketReqBody) ApiCreateBucketRequest {
	r.handlerPutBucketReqBody = &handlerPutBucketReqBody
	return r
}

func (r ApiCreateBucketRequest) Execute() (*HandlerPutBucketRes, *http.Response, error) {
	return r.ApiService.CreateBucketExecute(r)
}

/*
CreateBucket バケットの作成

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/clusters` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters`
となります。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name バケット名
	@return ApiCreateBucketRequest
*/
func (a *DefaultApiService) CreateBucket(ctx context.Context, name string) ApiCreateBucketRequest {
	return ApiCreateBucketRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
//
//	@return HandlerPutBucketRes
func (a *DefaultApiService) CreateBucketExecute(r ApiCreateBucketRequest) (*HandlerPutBucketRes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HandlerPutBucketRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateBucket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.handlerPutBucketReqBody == nil {
		return localVarReturnValue, nil, reportError("handlerPutBucketReqBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.handlerPutBucketReqBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePermissionRequest struct {
	ctx                          context.Context
	ApiService                   *DefaultApiService
	permissionBucketControlsBody *PermissionBucketControlsBody
}

// どのバケットを制御するかを定義します。
func (r ApiCreatePermissionRequest) PermissionBucketControlsBody(permissionBucketControlsBody PermissionBucketControlsBody) ApiCreatePermissionRequest {
	r.permissionBucketControlsBody = &permissionBucketControlsBody
	return r
}

func (r ApiCreatePermissionRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.CreatePermissionExecute(r)
}

/*
CreatePermission パーミッションの作成

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreatePermissionRequest
*/
func (a *DefaultApiService) CreatePermission(ctx context.Context) ApiCreatePermissionRequest {
	return ApiCreatePermissionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Permission
func (a *DefaultApiService) CreatePermissionExecute(r ApiCreatePermissionRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreatePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.permissionBucketControlsBody == nil {
		return localVarReturnValue, nil, reportError("permissionBucketControlsBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.permissionBucketControlsBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePermissionKeyRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiCreatePermissionKeyRequest) Execute() (*PermissionKey, *http.Response, error) {
	return r.ApiService.CreatePermissionKeyExecute(r)
}

/*
CreatePermissionKey パーミッションのアクセスキーの発行

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@return ApiCreatePermissionKeyRequest
*/
func (a *DefaultApiService) CreatePermissionKey(ctx context.Context, id string) ApiCreatePermissionKeyRequest {
	return ApiCreatePermissionKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return PermissionKey
func (a *DefaultApiService) CreatePermissionKeyExecute(r ApiCreatePermissionKeyRequest) (*PermissionKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PermissionKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreatePermissionKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteAccountRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiDeleteAccountRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAccountExecute(r)
}

/*
DeleteAccount サイトアカウントの削除

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDeleteAccountRequest
*/
func (a *DefaultApiService) DeleteAccount(ctx context.Context) ApiDeleteAccountRequest {
	return ApiDeleteAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeleteAccountExecute(r ApiDeleteAccountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteAccount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteAccountKeyRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiDeleteAccountKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAccountKeyExecute(r)
}

/*
DeleteAccountKey サイトアカウントのアクセスキーの削除

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id アクセスキーID
	@return ApiDeleteAccountKeyRequest
*/
func (a *DefaultApiService) DeleteAccountKey(ctx context.Context, id string) ApiDeleteAccountKeyRequest {
	return ApiDeleteAccountKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeleteAccountKeyExecute(r ApiDeleteAccountKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteAccountKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteBucketRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	name       string
}

func (r ApiDeleteBucketRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteBucketExecute(r)
}

/*
DeleteBucket バケットの削除

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/clusters` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters`
となります。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name バケット名
	@return ApiDeleteBucketRequest
*/
func (a *DefaultApiService) DeleteBucket(ctx context.Context, name string) ApiDeleteBucketRequest {
	return ApiDeleteBucketRequest{
		ApiService: a,
		ctx:        ctx,
		name:       name,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeleteBucketExecute(r ApiDeleteBucketRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteBucket")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buckets/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeletePermissionRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiDeletePermissionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePermissionExecute(r)
}

/*
DeletePermission パーミッションの削除

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@return ApiDeletePermissionRequest
*/
func (a *DefaultApiService) DeletePermission(ctx context.Context, id string) ApiDeletePermissionRequest {
	return ApiDeletePermissionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeletePermissionExecute(r ApiDeletePermissionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeletePermission")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeletePermissionKeyRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
	keyId      string
}

func (r ApiDeletePermissionKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePermissionKeyExecute(r)
}

/*
DeletePermissionKey パーミッションが保有するアクセスキーの削除

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@param keyId パーミッションが保有するアクセスキーID
	@return ApiDeletePermissionKeyRequest
*/
func (a *DefaultApiService) DeletePermissionKey(ctx context.Context, id string, keyId string) ApiDeletePermissionKeyRequest {
	return ApiDeletePermissionKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		keyId:      keyId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DeletePermissionKeyExecute(r ApiDeletePermissionKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeletePermissionKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}/keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAccountRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiGetAccountRequest) Execute() (*Account, *http.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount サイトアカウントの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAccountRequest
*/
func (a *DefaultApiService) GetAccount(ctx context.Context) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Account
func (a *DefaultApiService) GetAccountExecute(r ApiGetAccountRequest) (*Account, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Account
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAccountKeyRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiGetAccountKeyRequest) Execute() (*AccountKey, *http.Response, error) {
	return r.ApiService.GetAccountKeyExecute(r)
}

/*
GetAccountKey サイトアカウントのアクセスキーの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id アクセスキーID
	@return ApiGetAccountKeyRequest
*/
func (a *DefaultApiService) GetAccountKey(ctx context.Context, id string) ApiGetAccountKeyRequest {
	return ApiGetAccountKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AccountKey
func (a *DefaultApiService) GetAccountKeyExecute(r ApiGetAccountKeyRequest) (*AccountKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccountKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetAccountKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/keys/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAccountKeysRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiGetAccountKeysRequest) Execute() (*AccountKeys, *http.Response, error) {
	return r.ApiService.GetAccountKeysExecute(r)
}

/*
GetAccountKeys サイトアカウントのアクセスキーの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAccountKeysRequest
*/
func (a *DefaultApiService) GetAccountKeys(ctx context.Context) ApiGetAccountKeysRequest {
	return ApiGetAccountKeysRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AccountKeys
func (a *DefaultApiService) GetAccountKeysExecute(r ApiGetAccountKeysRequest) (*AccountKeys, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AccountKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetAccountKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account/keys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClusterRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiGetClusterRequest) Execute() (*HandlerGetClusterRes, *http.Response, error) {
	return r.ApiService.GetClusterExecute(r)
}

/*
GetCluster サイトの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/clusters` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters`
となります。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id サイトID
	@return ApiGetClusterRequest
*/
func (a *DefaultApiService) GetCluster(ctx context.Context, id string) ApiGetClusterRequest {
	return ApiGetClusterRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return HandlerGetClusterRes
func (a *DefaultApiService) GetClusterExecute(r ApiGetClusterRequest) (*HandlerGetClusterRes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HandlerGetClusterRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clusters/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClustersRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiGetClustersRequest) Execute() (*HandlerListClustersRes, *http.Response, error) {
	return r.ApiService.GetClustersExecute(r)
}

/*
GetClusters サイト一覧の取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/clusters` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters`
となります。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetClustersRequest
*/
func (a *DefaultApiService) GetClusters(ctx context.Context) ApiGetClustersRequest {
	return ApiGetClustersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return HandlerListClustersRes
func (a *DefaultApiService) GetClustersExecute(r ApiGetClustersRequest) (*HandlerListClustersRes, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HandlerListClustersRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/clusters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPermissionRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiGetPermissionRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.GetPermissionExecute(r)
}

/*
GetPermission パーミッションの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@return ApiGetPermissionRequest
*/
func (a *DefaultApiService) GetPermission(ctx context.Context, id string) ApiGetPermissionRequest {
	return ApiGetPermissionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Permission
func (a *DefaultApiService) GetPermissionExecute(r ApiGetPermissionRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPermissionKeyRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
	keyId      string
}

func (r ApiGetPermissionKeyRequest) Execute() (*PermissionKey, *http.Response, error) {
	return r.ApiService.GetPermissionKeyExecute(r)
}

/*
GetPermissionKey パーミッションが保有するアクセスキーの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@param keyId パーミッションが保有するアクセスキーID
	@return ApiGetPermissionKeyRequest
*/
func (a *DefaultApiService) GetPermissionKey(ctx context.Context, id string, keyId string) ApiGetPermissionKeyRequest {
	return ApiGetPermissionKeyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
		keyId:      keyId,
	}
}

// Execute executes the request
//
//	@return PermissionKey
func (a *DefaultApiService) GetPermissionKeyExecute(r ApiGetPermissionKeyRequest) (*PermissionKey, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PermissionKey
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPermissionKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}/keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPermissionKeysRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
	id         string
}

func (r ApiGetPermissionKeysRequest) Execute() (*PermissionKeys, *http.Response, error) {
	return r.ApiService.GetPermissionKeysExecute(r)
}

/*
GetPermissionKeys パーミッションが保有するアクセスキー一覧の取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@return ApiGetPermissionKeysRequest
*/
func (a *DefaultApiService) GetPermissionKeys(ctx context.Context, id string) ApiGetPermissionKeysRequest {
	return ApiGetPermissionKeysRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return PermissionKeys
func (a *DefaultApiService) GetPermissionKeysExecute(r ApiGetPermissionKeysRequest) (*PermissionKeys, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PermissionKeys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPermissionKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPermissionsRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiGetPermissionsRequest) Execute() (*Permissions, *http.Response, error) {
	return r.ApiService.GetPermissionsExecute(r)
}

/*
GetPermissions パーミッション一覧の取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetPermissionsRequest
*/
func (a *DefaultApiService) GetPermissions(ctx context.Context) ApiGetPermissionsRequest {
	return ApiGetPermissionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Permissions
func (a *DefaultApiService) GetPermissionsExecute(r ApiGetPermissionsRequest) (*Permissions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetStatusRequest struct {
	ctx        context.Context
	ApiService *DefaultApiService
}

func (r ApiGetStatusRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.GetStatusExecute(r)
}

/*
GetStatus サイトのステータスの取得

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetStatusRequest
*/
func (a *DefaultApiService) GetStatus(ctx context.Context) ApiGetStatusRequest {
	return ApiGetStatusRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Status
func (a *DefaultApiService) GetStatusExecute(r ApiGetStatusRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePermissionRequest struct {
	ctx                          context.Context
	ApiService                   *DefaultApiService
	id                           string
	permissionBucketControlsBody *PermissionBucketControlsBody
}

// どのバケットを制御するかを定義します。
func (r ApiUpdatePermissionRequest) PermissionBucketControlsBody(permissionBucketControlsBody PermissionBucketControlsBody) ApiUpdatePermissionRequest {
	r.permissionBucketControlsBody = &permissionBucketControlsBody
	return r
}

func (r ApiUpdatePermissionRequest) Execute() (*Permission, *http.Response, error) {
	return r.ApiService.UpdatePermissionExecute(r)
}

/*
UpdatePermission パーミッションの更新

アクセス先ベースURL: `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/`

上記アクセス先ベースURLと右ペインのパスを組み合わせたものがアクセス先URLになります。
例えば、右ペインのパスが `/account` の場合、アクセス先URLは、
`https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/account`
となります。

なお、サイト名の取得方法は「利用例＞接続先サイト一覧の取得」をご参考下さい。

**注意**： 本APIは指定したパーミッションIDの内容を全て置き換えます。
そのため、一部のバケットの権限を変更したい場合であっても、本パーミッションIDが管理する全てのバケットに関する情報も合わせて送る必要があります。

例えば、パーミッションID`1`が管理するバケットが下記の内容であるとします（パーミッションID`1`をGETした際の内容を示しています）。
```

	{
	  "data": {
	    "bucket_controls": [
	      {
	        "bucket_name": "your-bucket-name-001",
	        "can_read": false,
	        "can_write": false,
	        "created_at": "2021-01-01T12:59:59.371074+09:00"
	      },
	      {
	        "bucket_name": "your-bucket-name-002",
	        "can_read": false,
	        "can_write": false,
	        "created_at": "2021-01-01T12:59:59.371074+09:00"
	      }
	    ]
	    "created_at": "2021-01-01T12:59:59.371074+09:00"
	    "display_name": "your-permission-001",
	    "id": 1
	  },
	}

```
この内、`your-bucket-name-002`の読み込み及び書き込み権限だけを許可したい場合は下記のようなリクエストボディーを送る必要があります。
```

	{
	  "id": 1,
	  "display_name": "your-permission-001",
	  "bucket_controls": [
	    {
	      "bucket_name": "your-bucket-name-001",
	      "can_read": false,
	      "can_write": false,
	    },
	    {
	      "bucket_name": "your-bucket-name-002",
	      "can_read": true,
	      "can_write": true,
	    }
	  ]
	}

```

実際にリクエストを送る際には上記の内容を`data.json`というファイルに保存し、以下のような入力を行います。
```
# 入力サンプル

	curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \
	     -X PUT \
	     -d @data.json \

https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/permissions/1
```

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id パーミッションID
	@return ApiUpdatePermissionRequest
*/
func (a *DefaultApiService) UpdatePermission(ctx context.Context, id string) ApiUpdatePermissionRequest {
	return ApiUpdatePermissionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Permission
func (a *DefaultApiService) UpdatePermissionExecute(r ApiUpdatePermissionRequest) (*Permission, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Permission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdatePermission")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/permissions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.permissionBucketControlsBody == nil {
		return localVarReturnValue, nil, reportError("permissionBucketControlsBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.permissionBucketControlsBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorDefault
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
