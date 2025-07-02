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

// checks if the PostApplicationBodyComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostApplicationBodyComponent{}

// PostApplicationBodyComponent struct for PostApplicationBodyComponent
type PostApplicationBodyComponent struct {
	// コンポーネント名
	Name string `json:"name"`
	// コンポーネントの最大CPU数
	MaxCpu string `json:"max_cpu"`
	// コンポーネントの最大メモリ
	MaxMemory string `json:"max_memory"`
	DeploySource PostApplicationBodyComponentDeploySource `json:"deploy_source"`
	// コンポーネントに渡す環境変数
	Env []PostApplicationBodyComponentEnv `json:"env,omitempty"`
	Probe NullablePostApplicationBodyComponentProbe `json:"probe,omitempty"`
}

type _PostApplicationBodyComponent PostApplicationBodyComponent

// NewPostApplicationBodyComponent instantiates a new PostApplicationBodyComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostApplicationBodyComponent(name string, maxCpu string, maxMemory string, deploySource PostApplicationBodyComponentDeploySource) *PostApplicationBodyComponent {
	this := PostApplicationBodyComponent{}
	this.Name = name
	this.MaxCpu = maxCpu
	this.MaxMemory = maxMemory
	this.DeploySource = deploySource
	return &this
}

// NewPostApplicationBodyComponentWithDefaults instantiates a new PostApplicationBodyComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostApplicationBodyComponentWithDefaults() *PostApplicationBodyComponent {
	this := PostApplicationBodyComponent{}
	return &this
}

// GetName returns the Name field value
func (o *PostApplicationBodyComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBodyComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PostApplicationBodyComponent) SetName(v string) {
	o.Name = v
}

// GetMaxCpu returns the MaxCpu field value
func (o *PostApplicationBodyComponent) GetMaxCpu() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxCpu
}

// GetMaxCpuOk returns a tuple with the MaxCpu field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBodyComponent) GetMaxCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCpu, true
}

// SetMaxCpu sets field value
func (o *PostApplicationBodyComponent) SetMaxCpu(v string) {
	o.MaxCpu = v
}

// GetMaxMemory returns the MaxMemory field value
func (o *PostApplicationBodyComponent) GetMaxMemory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxMemory
}

// GetMaxMemoryOk returns a tuple with the MaxMemory field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBodyComponent) GetMaxMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMemory, true
}

// SetMaxMemory sets field value
func (o *PostApplicationBodyComponent) SetMaxMemory(v string) {
	o.MaxMemory = v
}

// GetDeploySource returns the DeploySource field value
func (o *PostApplicationBodyComponent) GetDeploySource() PostApplicationBodyComponentDeploySource {
	if o == nil {
		var ret PostApplicationBodyComponentDeploySource
		return ret
	}

	return o.DeploySource
}

// GetDeploySourceOk returns a tuple with the DeploySource field value
// and a boolean to check if the value has been set.
func (o *PostApplicationBodyComponent) GetDeploySourceOk() (*PostApplicationBodyComponentDeploySource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploySource, true
}

// SetDeploySource sets field value
func (o *PostApplicationBodyComponent) SetDeploySource(v PostApplicationBodyComponentDeploySource) {
	o.DeploySource = v
}

// GetEnv returns the Env field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostApplicationBodyComponent) GetEnv() []PostApplicationBodyComponentEnv {
	if o == nil {
		var ret []PostApplicationBodyComponentEnv
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostApplicationBodyComponent) GetEnvOk() ([]PostApplicationBodyComponentEnv, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *PostApplicationBodyComponent) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []PostApplicationBodyComponentEnv and assigns it to the Env field.
func (o *PostApplicationBodyComponent) SetEnv(v []PostApplicationBodyComponentEnv) {
	o.Env = v
}

// GetProbe returns the Probe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostApplicationBodyComponent) GetProbe() PostApplicationBodyComponentProbe {
	if o == nil || IsNil(o.Probe.Get()) {
		var ret PostApplicationBodyComponentProbe
		return ret
	}
	return *o.Probe.Get()
}

// GetProbeOk returns a tuple with the Probe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostApplicationBodyComponent) GetProbeOk() (*PostApplicationBodyComponentProbe, bool) {
	if o == nil {
		return nil, false
	}
	return o.Probe.Get(), o.Probe.IsSet()
}

// HasProbe returns a boolean if a field has been set.
func (o *PostApplicationBodyComponent) HasProbe() bool {
	if o != nil && o.Probe.IsSet() {
		return true
	}

	return false
}

// SetProbe gets a reference to the given NullablePostApplicationBodyComponentProbe and assigns it to the Probe field.
func (o *PostApplicationBodyComponent) SetProbe(v PostApplicationBodyComponentProbe) {
	o.Probe.Set(&v)
}
// SetProbeNil sets the value for Probe to be an explicit nil
func (o *PostApplicationBodyComponent) SetProbeNil() {
	o.Probe.Set(nil)
}

// UnsetProbe ensures that no value is present for Probe, not even an explicit nil
func (o *PostApplicationBodyComponent) UnsetProbe() {
	o.Probe.Unset()
}

func (o PostApplicationBodyComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostApplicationBodyComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["max_cpu"] = o.MaxCpu
	toSerialize["max_memory"] = o.MaxMemory
	toSerialize["deploy_source"] = o.DeploySource
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Probe.IsSet() {
		toSerialize["probe"] = o.Probe.Get()
	}
	return toSerialize, nil
}

func (o *PostApplicationBodyComponent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"max_cpu",
		"max_memory",
		"deploy_source",
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

	varPostApplicationBodyComponent := _PostApplicationBodyComponent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostApplicationBodyComponent)

	if err != nil {
		return err
	}

	*o = PostApplicationBodyComponent(varPostApplicationBodyComponent)

	return err
}

type NullablePostApplicationBodyComponent struct {
	value *PostApplicationBodyComponent
	isSet bool
}

func (v NullablePostApplicationBodyComponent) Get() *PostApplicationBodyComponent {
	return v.value
}

func (v *NullablePostApplicationBodyComponent) Set(val *PostApplicationBodyComponent) {
	v.value = val
	v.isSet = true
}

func (v NullablePostApplicationBodyComponent) IsSet() bool {
	return v.isSet
}

func (v *NullablePostApplicationBodyComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostApplicationBodyComponent(val *PostApplicationBodyComponent) *NullablePostApplicationBodyComponent {
	return &NullablePostApplicationBodyComponent{value: val, isSet: true}
}

func (v NullablePostApplicationBodyComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostApplicationBodyComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


