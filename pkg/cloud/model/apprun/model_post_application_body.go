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

// checks if the PostApplicationBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostApplicationBody{}

// PostApplicationBody struct for PostApplicationBody
type PostApplicationBody struct {
	// アプリケーション名
	Name string `json:"name"`
	// アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds int32 `json:"timeout_seconds"`
	// アプリケーションがリクエストを待ち受けるポート番号
	Port int32 `json:"port"`
	// アプリケーション全体の最小スケール数
	MinScale int32 `json:"min_scale"`
	// アプリケーション全体の最大スケール数
	MaxScale int32 `json:"max_scale"`
	// アプリケーションのコンポーネント情報
	Components []PostApplicationBodyComponent `json:"components"`
}

type _PostApplicationBody PostApplicationBody

// NewPostApplicationBody instantiates a new PostApplicationBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostApplicationBody(name string, timeoutSeconds int32, port int32, minScale int32, maxScale int32, components []PostApplicationBodyComponent) *PostApplicationBody {
	this := PostApplicationBody{}
	this.Name = name
	this.TimeoutSeconds = timeoutSeconds
	this.Port = port
	this.MinScale = minScale
	this.MaxScale = maxScale
	this.Components = components
	return &this
}

// NewPostApplicationBodyWithDefaults instantiates a new PostApplicationBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostApplicationBodyWithDefaults() *PostApplicationBody {
	this := PostApplicationBody{}
	var timeoutSeconds int32 = 60
	this.TimeoutSeconds = timeoutSeconds
	var minScale int32 = 0
	this.MinScale = minScale
	var maxScale int32 = 10
	this.MaxScale = maxScale
	return &this
}

// GetName returns the Name field value
func (o *PostApplicationBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostApplicationBody) SetName(v string) {
	o.Name = v
}

// GetTimeoutSeconds returns the TimeoutSeconds field value
func (o *PostApplicationBody) GetTimeoutSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeoutSeconds
}

// GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBody) GetTimeoutSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutSeconds, true
}

// SetTimeoutSeconds sets field value
func (o *PostApplicationBody) SetTimeoutSeconds(v int32) {
	o.TimeoutSeconds = v
}

// GetPort returns the Port field value
func (o *PostApplicationBody) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBody) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *PostApplicationBody) SetPort(v int32) {
	o.Port = v
}

// GetMinScale returns the MinScale field value
func (o *PostApplicationBody) GetMinScale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinScale
}

// GetMinScaleOk returns a tuple with the MinScale field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBody) GetMinScaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinScale, true
}

// SetMinScale sets field value
func (o *PostApplicationBody) SetMinScale(v int32) {
	o.MinScale = v
}

// GetMaxScale returns the MaxScale field value
func (o *PostApplicationBody) GetMaxScale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxScale
}

// GetMaxScaleOk returns a tuple with the MaxScale field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBody) GetMaxScaleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxScale, true
}

// SetMaxScale sets field value
func (o *PostApplicationBody) SetMaxScale(v int32) {
	o.MaxScale = v
}

// GetComponents returns the Components field value
func (o *PostApplicationBody) GetComponents() []PostApplicationBodyComponent {
	if o == nil {
		var ret []PostApplicationBodyComponent
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBody) GetComponentsOk() ([]PostApplicationBodyComponent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *PostApplicationBody) SetComponents(v []PostApplicationBodyComponent) {
	o.Components = v
}

func (o PostApplicationBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostApplicationBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["timeout_seconds"] = o.TimeoutSeconds
	toSerialize["port"] = o.Port
	toSerialize["min_scale"] = o.MinScale
	toSerialize["max_scale"] = o.MaxScale
	toSerialize["components"] = o.Components
	return toSerialize, nil
}

func (o *PostApplicationBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"timeout_seconds",
		"port",
		"min_scale",
		"max_scale",
		"components",
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

	varPostApplicationBody := _PostApplicationBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostApplicationBody)

	if err != nil {
		return err
	}

	*o = PostApplicationBody(varPostApplicationBody)

	return err
}

type NullablePostApplicationBody struct {
	value *PostApplicationBody
	isSet bool
}

func (v NullablePostApplicationBody) Get() *PostApplicationBody {
	return v.value
}

func (v *NullablePostApplicationBody) Set(val *PostApplicationBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePostApplicationBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePostApplicationBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostApplicationBody(val *PostApplicationBody) *NullablePostApplicationBody {
	return &NullablePostApplicationBody{value: val, isSet: true}
}

func (v NullablePostApplicationBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostApplicationBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


