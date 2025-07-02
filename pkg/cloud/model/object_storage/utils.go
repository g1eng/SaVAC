/*
さくらのオブジェクトストレージ APIドキュメント

 ---  「さくらのオブジェクトストレージ」が提供するAPIの利用方法とサンプルを公開しております。  JSON 形式の OpenAPI 仕様は、以下の URL からダウンロードしてください。   <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.json\">JSON形式でダウンロード</a></li></ul>  # 基本的な使い方  ## APIキーの発行  APIを利用するためには、認証のための「APIキー」が必要です。事前にキーを発行しておきます。 APIキーは「ユーザーID」「パスワード」に相当する「トークン」と呼ばれる認証情報で構成されています。  |   項目名   | APIキー発行時の項目名        | このドキュメント内での例             | |------------|------------------------------|--------------------------------------| | ユーザーID | アクセストークン(UUID)       | 01234567-89ab-cdef-0123-456789abcdef | | パスワード | アクセストークンシークレット | SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM |  <div class=\"warning\"> <b>操作マニュアル</b><br /> <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/api/apikey.html\">APIキー | さくらのクラウド ドキュメント</a></li></ul> </div>  ## 入力パラメータ  APIの入力には送信先URLに対して、いくつかのヘッダーとAPIキーを送信します。  * APIのURLは以下の2つが存在します。※ 各APIの使い分けは後述します。   * `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/(エンドポイント)`   * `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/(エンドポイント)` * 認証方式はHTTP Basic認証です。APIキーのアクセストークンをユーザーID、アクセストークンシークレットをパスワードとして指定します。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters' ```  ## 出力結果と応答コード（HTTPステータスコード）  APIからの結果は、「応答コード（HTTPステータスコード）」と、「JSON形式(UTF-8)の結果」として出力されます。  応答コードは、リクエストが成功したのか、失敗したのか大まかな情報を判断することができるもので、例えば失敗したときには、なぜこのような結果になったのかなど、具体的な情報は応答コードと主に返された本文を見ることで把握することができます。  | 結果                                | 応答コード/status   | |-------------------------------------|---------------------| | 成功（要求を受け付けた）             | 2xx                 | | 失敗（要求が受け付けられなかった）  | 4xx, 5xx            |  ``` # 出力結果サンプル（レスポンスヘッダ） HTTP/1.1 200 OK Server: nginx Date: Tue, 16 Nov 2021 12:39:48 GMT Content-Type: application/json; charset=UTF-8 Content-Length: 443 Connection: keep-alive Status: 200 OK Pragma: no-cache Cache-Control: no-cache X-Sakura-Proxy-Microtime: 66245 X-Sakura-Proxy-Decode-Microtime: 62 X-Sakura-Content-Length: 443 X-Sakura-Serial: 86ab6c743f72aa5ea6f17e254fd5f803 X-Content-Type-Options: nosniff X-XSS-Protection: 1; mode=block X-Frame-Options: DENY X-Sakura-Encode-Microtime: 260 Vary: Accept-Encoding ```  ``` # 出力結果サンプル（レスポンスボディー） {   \"error\": {     \"code\": 404,     \"errors\": [       {         \"domain\": \"fed.objectstorage.sacloud\",         \"location\": \"clusters\",         \"location_type\": \"path_parameter\",         \"message\": \"Cluster was not found\",         \"reason\": \"not_found\"       }     ],     \"message\": \"Cluster was not found\",     \"trace_id\": \"0f36837633984f3fc8871f515e8efa24\"   } } ```  # 利用例  ## 1.接続先サイト一覧の取得  さくらのオブジェクトストレージを利用するには、まずバケット作成先となる**サイト**を取得・選択します。  サイト一覧を取得するには、以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters' ```  実行結果として、サイトのリストが返却されます。  ``` # 出力結果サンプル {   \"data\": [     {       \"api_zone\": [],       \"control_panel_url\": \"https://secure.sakura.ad.jp/objectstorage/\",       \"display_name_en_us\": \"Ishikari Site #1\",       \"display_name_ja\": \"石狩第1サイト\",       \"display_name\": \"石狩第1サイト\",       \"display_order\": 1,       \"endpoint_base\": \"isk01.sakurastorage.jp\",       \"id\": \"isk01\",       \"region\": \"jp-north-1\",       \"s3_endpoint\": \"s3.isk01.sakurastorage.jp\",       \"s3_endpoint_for_control_panel\": \"s3.cp.isk01.sakurastorage.jp\",       \"storage_zone\": []     }   ] } ```  得られたサイトID（上記の`id`フィールド）を確認します。これは後続の利用例で使用します。  ## 2.サイトアカウントの作成  上記のサイトから利用したいサイトIDを選択し（ここではisk01を選択することにします）、**サイトアカウント**を作成します。  サイトアカウントとは、サイトを利用するための独立したアカウントであり、サイトアカウント作成・削除による料金の発生はございません。 なお、すでにサイトアカウントを作成済みの場合は、再度サイトアカウントの作成は不要です。  サイトアカウントを作成するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account ```  サイトアカウントの作成が完了すると、選択したサイトにて  * バケットの作成・削除 * アクセスキーの発行・削除 * パーミッションキーの発行・削除  などの操作が可能になります。  ## 3.バケットの作成・削除  選択したサイトにてサイトアカウントを作成後、**バケット**の作成・削除が可能です。  バケットを作成するには以下のような入力を行います。     この時、選択したサイト（ここではisk01とします）をリクエストボディーに入れ、作成したいバケット名をパスパラメータに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PUT \\      -d '{\"cluster_id\": \"isk01\"}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/buckets/sample ```  上記で作成したバケットを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      -d '{\"cluster_id\": \"isk01\"}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/buckets/sample ```  ## 4.アクセスキーの発行・削除  選択したサイトにてサイトアカウントを作成後、**アクセスキー**の発行・削除が可能です。  アクセスキーを発行するには以下のような入力を行います。      ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account/keys ```  コマンド結果には以下のフィールドが含まれます。  * `created_at` : 作成日時 * `id` : アクセスキーID * `secret` : シークレットアクセスキー  ``` # 出力結果サンプル {   \"data\": {     \"created_at\": \"2021-11-04T07:42:41.121418479Z\",     \"id\": \"XPJK4SC9883N91RHR253\",     \"secret\": \"jqRaUo5l+EiEYqP8wos9exbmFfq4/vG8CLPYI2XN\"   } } ```  上記で作成したアクセスキーを削除するには以下のような入力を行います。     この時、削除したいアクセスキーIDをパスパラメータに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account/keys/XPJK4SC9883N91RHR253 ```  ## 5.パーミッション及びパーミッションアクセスキーの発行・削除  選択したサイトにてサイトアカウントを作成後且つバケットが1つ以上ある場合、**パーミッション**の発行・削除が可能です。  パーミッションを作成するには以下のような入力を行います。 この時、パーミッション名、パーミッションで制御したいバケットとそれに対する操作をリクエストボディーに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      -d '{\"display_name\": \"sample_permission\", \"bucket_controls\": [{\"bucket_name\": \"sample\", \"can_read\": true, \"can_write\": true}]}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions ```  作成が完了すると、パーミッションIDが含まれたレスポンスを受け取ります。 ``` # 出力サンプル {   \"data\": {     \"bucket_controls\": [       {         \"bucket_name\":\"sample\",         \"can_read\":true,         \"can_write\":true,         \"created_at\":\"2021-11-11T13:36:08.767118492Z\"       }     ],     \"created_at\":\"2021-11-11T13:36:08.690384415Z\",     \"display_name\":\"sample_permission\",     \"id\":619   } } ```  このパーミッションのアクセスキーを発行するには以下のような入力を行います。 この時、パーミッション作成時に発行されたID（ここでは619とします）をパスパラメータに含める必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619/keys ```  コマンド結果には以下のフィールドが含まれます。  * `created_at` : 作成日時 * `id` : アクセスキーID * `secret` : シークレットアクセスキー  ``` # 出力結果サンプル {   \"data\": {     \"created_at\": \"2021-11-04T07:42:41.121418479Z\",     \"id\": \"XPJK4SC9883N91RHR253\",     \"secret\": \"jqRaUo5l+EiEYqP8wos9exbmFfq4/vG8CLPYI2XN\"   } } ```  パーミッションアクセスキーを削除するには以下のような入力を行います。 この時、パーミッションアクセスキー発行時に出力されたIDをパスパラメータに含める必要があります。 ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619/keys/XPJK4SC9883N91RHR253 ```  パーミッションを削除するには以下のような入力を行います。 この時、パーミッション作成時に発行されたID（ここでは619とします）をパスパラメータに含める必要があります。 ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619 ``` ----

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package object_storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// PtrBool is a helper routine that returns a pointer to given boolean value.
func PtrBool(v bool) *bool { return &v }

// PtrInt is a helper routine that returns a pointer to given integer value.
func PtrInt(v int) *int { return &v }

// PtrInt32 is a helper routine that returns a pointer to given integer value.
func PtrInt32(v int32) *int32 { return &v }

// PtrInt64 is a helper routine that returns a pointer to given integer value.
func PtrInt64(v int64) *int64 { return &v }

// PtrFloat32 is a helper routine that returns a pointer to given float value.
func PtrFloat32(v float32) *float32 { return &v }

// PtrFloat64 is a helper routine that returns a pointer to given float value.
func PtrFloat64(v float64) *float64 { return &v }

// PtrString is a helper routine that returns a pointer to given string value.
func PtrString(v string) *string { return &v }

// PtrTime is helper routine that returns a pointer to given Time value.
func PtrTime(v time.Time) *time.Time { return &v }

type NullableBool struct {
	value *bool
	isSet bool
}

func (v NullableBool) Get() *bool {
	return v.value
}

func (v *NullableBool) Set(val *bool) {
	v.value = val
	v.isSet = true
}

func (v NullableBool) IsSet() bool {
	return v.isSet
}

func (v *NullableBool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBool(val *bool) *NullableBool {
	return &NullableBool{value: val, isSet: true}
}

func (v NullableBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt struct {
	value *int
	isSet bool
}

func (v NullableInt) Get() *int {
	return v.value
}

func (v *NullableInt) Set(val *int) {
	v.value = val
	v.isSet = true
}

func (v NullableInt) IsSet() bool {
	return v.isSet
}

func (v *NullableInt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt(val *int) *NullableInt {
	return &NullableInt{value: val, isSet: true}
}

func (v NullableInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt32 struct {
	value *int32
	isSet bool
}

func (v NullableInt32) Get() *int32 {
	return v.value
}

func (v *NullableInt32) Set(val *int32) {
	v.value = val
	v.isSet = true
}

func (v NullableInt32) IsSet() bool {
	return v.isSet
}

func (v *NullableInt32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt32(val *int32) *NullableInt32 {
	return &NullableInt32{value: val, isSet: true}
}

func (v NullableInt32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableInt64 struct {
	value *int64
	isSet bool
}

func (v NullableInt64) Get() *int64 {
	return v.value
}

func (v *NullableInt64) Set(val *int64) {
	v.value = val
	v.isSet = true
}

func (v NullableInt64) IsSet() bool {
	return v.isSet
}

func (v *NullableInt64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInt64(val *int64) *NullableInt64 {
	return &NullableInt64{value: val, isSet: true}
}

func (v NullableInt64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInt64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableFloat32 struct {
	value *float32
	isSet bool
}

func (v NullableFloat32) Get() *float32 {
	return v.value
}

func (v *NullableFloat32) Set(val *float32) {
	v.value = val
	v.isSet = true
}

func (v NullableFloat32) IsSet() bool {
	return v.isSet
}

func (v *NullableFloat32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloat32(val *float32) *NullableFloat32 {
	return &NullableFloat32{value: val, isSet: true}
}

func (v NullableFloat32) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloat32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableFloat64 struct {
	value *float64
	isSet bool
}

func (v NullableFloat64) Get() *float64 {
	return v.value
}

func (v *NullableFloat64) Set(val *float64) {
	v.value = val
	v.isSet = true
}

func (v NullableFloat64) IsSet() bool {
	return v.isSet
}

func (v *NullableFloat64) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFloat64(val *float64) *NullableFloat64 {
	return &NullableFloat64{value: val, isSet: true}
}

func (v NullableFloat64) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFloat64) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableString struct {
	value *string
	isSet bool
}

func (v NullableString) Get() *string {
	return v.value
}

func (v *NullableString) Set(val *string) {
	v.value = val
	v.isSet = true
}

func (v NullableString) IsSet() bool {
	return v.isSet
}

func (v *NullableString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableString(val *string) *NullableString {
	return &NullableString{value: val, isSet: true}
}

func (v NullableString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type NullableTime struct {
	value *time.Time
	isSet bool
}

func (v NullableTime) Get() *time.Time {
	return v.value
}

func (v *NullableTime) Set(val *time.Time) {
	v.value = val
	v.isSet = true
}

func (v NullableTime) IsSet() bool {
	return v.isSet
}

func (v *NullableTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTime(val *time.Time) *NullableTime {
	return &NullableTime{value: val, isSet: true}
}

func (v NullableTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

// IsNil checks if an input is nil
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	case reflect.Array:
		return reflect.ValueOf(i).IsZero()
	}
	return false
}

type MappedNullable interface {
	ToMap() (map[string]interface{}, error)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
