/*
AppRun β APIドキュメント

--- 「AppRun」が提供するAPIの利用方法とサンプルを公開しております。  # 基本的な使い方  ## APIキーの発行  APIを利用するためには、認証のための「APIキー」が必要です。事前にキーを発行しておきます。   APIキーは「ユーザーID」「パスワード」に相当する「トークン」と呼ばれる認証情報で構成されています。  |   項目名   | APIキー発行時の項目名        | このドキュメント内での例             | |------------|------------------------------|--------------------------------------| | ユーザーID | アクセストークン(UUID)       | 01234567-89ab-cdef-0123-456789abcdef | | パスワード | アクセストークンシークレット | SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM |  <div class=\"warning\"> <b>操作マニュアル</b><br /> <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/api/apikey.html\">APIキー | さくらのクラウド ドキュメント</a></li></ul> </div>  ## 入力パラメータ  APIの入力には送信先URLに対して、いくつかのヘッダーとAPIキーを送信します。  * 認証方式はHTTP Basic認証です。APIキーのアクセストークンをユーザーID、アクセストークンシークレットをパスワードとして指定します。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications' ```  ## 出力結果と応答コード（HTTPステータスコード）  APIからの結果は、「応答コード（HTTPステータスコード）」と、「JSON形式(UTF-8)の結果」として出力されます。  応答コードは、リクエストが成功したのか、失敗したのか大まかな情報を判断することができるもので、例えば失敗したときには、なぜこのような結果になったのかなど、具体的な情報は応答コードと主に返された本文を見ることで把握することができます。  | 結果                                | 応答コード/status   | |-------------------------------------|---------------------| | 成功（要求を受け付けた）            | 2xx                 | | 失敗（要求が受け付けられなかった）  | 4xx, 5xx            |  ``` # 出力結果サンプル（レスポンスヘッダー） HTTP/1.1 200 OK Server: nginx Date: Tue, 16 Nov 2021 12:39:48 GMT Content-Type: application/json; charset=UTF-8 Content-Length: 443 Connection: keep-alive Status: 200 OK Pragma: no-cache Cache-Control: no-cache X-Sakura-Proxy-Microtime: 66245 X-Sakura-Proxy-Decode-Microtime: 62 X-Sakura-Content-Length: 443 X-Sakura-Serial: 86ab6c743f72aa5ea6f17e254fd5f803 X-Content-Type-Options: nosniff X-XSS-Protection: 1; mode=block X-Frame-Options: DENY X-Sakura-Encode-Microtime: 260 Vary: Accept-Encoding ```  ``` # 出力結果サンプル（レスポンスボディー） {   \"error\": {     \"code\": 401,     \"message\": \"Login Required\",     \"errors\": [       {         \"domain\": \"global\",         \"reason\": \"required\",         \"message\": \"Login Required\",         \"location_type\": \"header\",         \"location\": \"Authorization\"       }     ]   } } ```  # 利用例  ## 1.ユーザーの作成  AppRunの利用を開始するには**ユーザー**を作成します。  ユーザーとは、AppRunを利用するための独立したユーザーであり、ユーザー作成および削除による料金の発生はございません。   なお、すでにユーザーを作成済みの場合は、再度ユーザーの作成は不要です。  ユーザーを作成するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/user ```  ユーザーの作成が完了すると、  * アプリケーションの作成、更新、削除 * バージョンの確認、削除 * トラフィック分散の確認、変更  などの操作が可能になります。  ## 2.アプリケーションの作成、取得、更新、削除  ユーザーを作成後、**アプリケーション**の作成、更新、削除が可能になります。  アプリケーションを作成するには以下のような入力を行います。  ``` # 入力サンプル vi request_body.json cat request_body.json {   \"name\": \"Application\",   \"timeout_seconds\": 60,   \"port\": 8080,   \"min_scale\": 0,   \"max_scale\": 1,   \"components\": [     {       \"name\": \"Component01\",       \"max_cpu\": \"0.1\",       \"max_memory\": \"256Mi\",       \"deploy_source\": {         \"container_registry\": {           \"image\": \"my-app.sakuracr.jp/my-app:latest\"         }       },       \"env\": [         {           \"key\": \"TARGET\",           \"value\": \"World\"         }       ],       \"probe\": {         \"http_get\": {           \"path\": \"/\",           \"port\": 8080,           \"headers\": [             {               \"name\": \"Custom-Header\",               \"value\": \"Awesome\"             }           ]         }       }     }   ] } curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      -d '@request_body.json' \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications ```  上記で作成したアプリケーションを取得するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X GET \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id} ```  上記で作成したアプリケーションを更新するには以下のような入力を行います。  ``` # 入力サンプル vi request_body.json cat request_body.json {   \"components\": [     {       \"name\": \"Component01 updated\",       \"max_cpu\": \"0.1\",       \"max_memory\": \"256Mi\",       \"deploy_source\": {         \"container_registry\": {           \"image\": \"my-app.sakuracr.jp/my-app-v2:latest\"         }       }     }   ],   \"all_traffic_available\": true }  curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PATCH \\      -d '@request_body.json' \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id} ```  上記で作成したアプリケーションを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id} ```  ## 3.バージョンの取得、削除  アプリケーションを作成、更新した際、その設定情報をバージョンとして保存します。  バージョンを取得するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X GET \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/versions/{version_id} ```  上記で作成したバージョンを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/versions/{version_id} ```  ## 4.トラフィック分散の確認、変更  アプリケーションは指定のバージョンへトラフィックを分散します。  トラフィック分散を確認するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X GET \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/traffics ```  トラフィック分散を変更するには以下のような入力を行います。  ``` # 入力サンプル vi request_body.json cat request_body.json [   {     \"is_latest_version\": true,     \"percent\": 50   },   {     \"version_name\": \"Application-861850d6-8240-7c31-9b69-80ea4466918d-1726726814\",     \"percent\": 50   } ] curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PUT \\      https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api/applications/{id}/traffics ``` ----

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apprun

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostApplicationBodyComponentDeploySourceContainerRegistry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostApplicationBodyComponentDeploySourceContainerRegistry{}

// PostApplicationBodyComponentDeploySourceContainerRegistry コンテナレジストリ
type PostApplicationBodyComponentDeploySourceContainerRegistry struct {
	// コンテナイメージ名
	Image string `json:"image"`
	// コンテナレジストリのサーバー名
	Server NullableString `json:"server,omitempty"`
	// コンテナレジストリの認証情報
	Username NullableString `json:"username,omitempty"`
	// コンテナレジストリの認証情報
	Password NullableString `json:"password,omitempty"`
}

type _PostApplicationBodyComponentDeploySourceContainerRegistry PostApplicationBodyComponentDeploySourceContainerRegistry

// NewPostApplicationBodyComponentDeploySourceContainerRegistry instantiates a new PostApplicationBodyComponentDeploySourceContainerRegistry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostApplicationBodyComponentDeploySourceContainerRegistry(image string) *PostApplicationBodyComponentDeploySourceContainerRegistry {
	this := PostApplicationBodyComponentDeploySourceContainerRegistry{}
	this.Image = image
	return &this
}

// NewPostApplicationBodyComponentDeploySourceContainerRegistryWithDefaults instantiates a new PostApplicationBodyComponentDeploySourceContainerRegistry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostApplicationBodyComponentDeploySourceContainerRegistryWithDefaults() *PostApplicationBodyComponentDeploySourceContainerRegistry {
	this := PostApplicationBodyComponentDeploySourceContainerRegistry{}
	return &this
}

// GetImage returns the Image field value
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetImage(v string) {
	o.Image = v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetServer() string {
	if o == nil || IsNil(o.Server.Get()) {
		var ret string
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableString and assigns it to the Server field.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetServer(v string) {
	o.Server.Set(&v)
}
// SetServerNil sets the value for Server to be an explicit nil
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) UnsetServer() {
	o.Server.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) UnsetPassword() {
	o.Password.Unset()
}

func (o PostApplicationBodyComponentDeploySourceContainerRegistry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostApplicationBodyComponentDeploySourceContainerRegistry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	if o.Server.IsSet() {
		toSerialize["server"] = o.Server.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	return toSerialize, nil
}

func (o *PostApplicationBodyComponentDeploySourceContainerRegistry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostApplicationBodyComponentDeploySourceContainerRegistry := _PostApplicationBodyComponentDeploySourceContainerRegistry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostApplicationBodyComponentDeploySourceContainerRegistry)

	if err != nil {
		return err
	}

	*o = PostApplicationBodyComponentDeploySourceContainerRegistry(varPostApplicationBodyComponentDeploySourceContainerRegistry)

	return err
}

type NullablePostApplicationBodyComponentDeploySourceContainerRegistry struct {
	value *PostApplicationBodyComponentDeploySourceContainerRegistry
	isSet bool
}

func (v NullablePostApplicationBodyComponentDeploySourceContainerRegistry) Get() *PostApplicationBodyComponentDeploySourceContainerRegistry {
	return v.value
}

func (v *NullablePostApplicationBodyComponentDeploySourceContainerRegistry) Set(val *PostApplicationBodyComponentDeploySourceContainerRegistry) {
	v.value = val
	v.isSet = true
}

func (v NullablePostApplicationBodyComponentDeploySourceContainerRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullablePostApplicationBodyComponentDeploySourceContainerRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostApplicationBodyComponentDeploySourceContainerRegistry(val *PostApplicationBodyComponentDeploySourceContainerRegistry) *NullablePostApplicationBodyComponentDeploySourceContainerRegistry {
	return &NullablePostApplicationBodyComponentDeploySourceContainerRegistry{value: val, isSet: true}
}

func (v NullablePostApplicationBodyComponentDeploySourceContainerRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostApplicationBodyComponentDeploySourceContainerRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


