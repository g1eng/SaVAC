/*
さくらのオブジェクトストレージ APIドキュメント

 ---  「さくらのオブジェクトストレージ」が提供するAPIの利用方法とサンプルを公開しております。  JSON 形式の OpenAPI 仕様は、以下の URL からダウンロードしてください。   <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/objectstorage/api/api-json.json\">JSON形式でダウンロード</a></li></ul>  # 基本的な使い方  ## APIキーの発行  APIを利用するためには、認証のための「APIキー」が必要です。事前にキーを発行しておきます。 APIキーは「ユーザーID」「パスワード」に相当する「トークン」と呼ばれる認証情報で構成されています。  |   項目名   | APIキー発行時の項目名        | このドキュメント内での例             | |------------|------------------------------|--------------------------------------| | ユーザーID | アクセストークン(UUID)       | 01234567-89ab-cdef-0123-456789abcdef | | パスワード | アクセストークンシークレット | SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM |  <div class=\"warning\"> <b>操作マニュアル</b><br /> <ul><li><a href=\"https://manual.sakura.ad.jp/cloud/api/apikey.html\">APIキー | さくらのクラウド ドキュメント</a></li></ul> </div>  ## 入力パラメータ  APIの入力には送信先URLに対して、いくつかのヘッダーとAPIキーを送信します。  * APIのURLは以下の2つが存在します。※ 各APIの使い分けは後述します。   * `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/(エンドポイント)`   * `https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/（サイト名）/v2/(エンドポイント)` * 認証方式はHTTP Basic認証です。APIキーのアクセストークンをユーザーID、アクセストークンシークレットをパスワードとして指定します。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters' ```  ## 出力結果と応答コード（HTTPステータスコード）  APIからの結果は、「応答コード（HTTPステータスコード）」と、「JSON形式(UTF-8)の結果」として出力されます。  応答コードは、リクエストが成功したのか、失敗したのか大まかな情報を判断することができるもので、例えば失敗したときには、なぜこのような結果になったのかなど、具体的な情報は応答コードと主に返された本文を見ることで把握することができます。  | 結果                                | 応答コード/status   | |-------------------------------------|---------------------| | 成功（要求を受け付けた）             | 2xx                 | | 失敗（要求が受け付けられなかった）  | 4xx, 5xx            |  ``` # 出力結果サンプル（レスポンスヘッダ） HTTP/1.1 200 OK Server: nginx Date: Tue, 16 Nov 2021 12:39:48 GMT Content-Type: application/json; charset=UTF-8 Content-Length: 443 Connection: keep-alive Status: 200 OK Pragma: no-cache Cache-Control: no-cache X-Sakura-Proxy-Microtime: 66245 X-Sakura-Proxy-Decode-Microtime: 62 X-Sakura-Content-Length: 443 X-Sakura-Serial: 86ab6c743f72aa5ea6f17e254fd5f803 X-Content-Type-Options: nosniff X-XSS-Protection: 1; mode=block X-Frame-Options: DENY X-Sakura-Encode-Microtime: 260 Vary: Accept-Encoding ```  ``` # 出力結果サンプル（レスポンスボディー） {   \"error\": {     \"code\": 404,     \"errors\": [       {         \"domain\": \"fed.objectstorage.sacloud\",         \"location\": \"clusters\",         \"location_type\": \"path_parameter\",         \"message\": \"Cluster was not found\",         \"reason\": \"not_found\"       }     ],     \"message\": \"Cluster was not found\",     \"trace_id\": \"0f36837633984f3fc8871f515e8efa24\"   } } ```  # 利用例  ## 1.接続先サイト一覧の取得  さくらのオブジェクトストレージを利用するには、まずバケット作成先となる**サイト**を取得・選択します。  サイト一覧を取得するには、以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      'https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/clusters' ```  実行結果として、サイトのリストが返却されます。  ``` # 出力結果サンプル {   \"data\": [     {       \"api_zone\": [],       \"control_panel_url\": \"https://secure.sakura.ad.jp/objectstorage/\",       \"display_name_en_us\": \"Ishikari Site #1\",       \"display_name_ja\": \"石狩第1サイト\",       \"display_name\": \"石狩第1サイト\",       \"display_order\": 1,       \"endpoint_base\": \"isk01.sakurastorage.jp\",       \"id\": \"isk01\",       \"region\": \"jp-north-1\",       \"s3_endpoint\": \"s3.isk01.sakurastorage.jp\",       \"s3_endpoint_for_control_panel\": \"s3.cp.isk01.sakurastorage.jp\",       \"storage_zone\": []     }   ] } ```  得られたサイトID（上記の`id`フィールド）を確認します。これは後続の利用例で使用します。  ## 2.サイトアカウントの作成  上記のサイトから利用したいサイトIDを選択し（ここではisk01を選択することにします）、**サイトアカウント**を作成します。  サイトアカウントとは、サイトを利用するための独立したアカウントであり、サイトアカウント作成・削除による料金の発生はございません。 なお、すでにサイトアカウントを作成済みの場合は、再度サイトアカウントの作成は不要です。  サイトアカウントを作成するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account ```  サイトアカウントの作成が完了すると、選択したサイトにて  * バケットの作成・削除 * アクセスキーの発行・削除 * パーミッションキーの発行・削除  などの操作が可能になります。  ## 3.バケットの作成・削除  選択したサイトにてサイトアカウントを作成後、**バケット**の作成・削除が可能です。  バケットを作成するには以下のような入力を行います。     この時、選択したサイト（ここではisk01とします）をリクエストボディーに入れ、作成したいバケット名をパスパラメータに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X PUT \\      -d '{\"cluster_id\": \"isk01\"}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/buckets/sample ```  上記で作成したバケットを削除するには以下のような入力を行います。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      -d '{\"cluster_id\": \"isk01\"}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/fed/v1/buckets/sample ```  ## 4.アクセスキーの発行・削除  選択したサイトにてサイトアカウントを作成後、**アクセスキー**の発行・削除が可能です。  アクセスキーを発行するには以下のような入力を行います。      ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account/keys ```  コマンド結果には以下のフィールドが含まれます。  * `created_at` : 作成日時 * `id` : アクセスキーID * `secret` : シークレットアクセスキー  ``` # 出力結果サンプル {   \"data\": {     \"created_at\": \"2021-11-04T07:42:41.121418479Z\",     \"id\": \"XPJK4SC9883N91RHR253\",     \"secret\": \"jqRaUo5l+EiEYqP8wos9exbmFfq4/vG8CLPYI2XN\"   } } ```  上記で作成したアクセスキーを削除するには以下のような入力を行います。     この時、削除したいアクセスキーIDをパスパラメータに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/account/keys/XPJK4SC9883N91RHR253 ```  ## 5.パーミッション及びパーミッションアクセスキーの発行・削除  選択したサイトにてサイトアカウントを作成後且つバケットが1つ以上ある場合、**パーミッション**の発行・削除が可能です。  パーミッションを作成するには以下のような入力を行います。 この時、パーミッション名、パーミッションで制御したいバケットとそれに対する操作をリクエストボディーに入れる必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      -d '{\"display_name\": \"sample_permission\", \"bucket_controls\": [{\"bucket_name\": \"sample\", \"can_read\": true, \"can_write\": true}]}' \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions ```  作成が完了すると、パーミッションIDが含まれたレスポンスを受け取ります。 ``` # 出力サンプル {   \"data\": {     \"bucket_controls\": [       {         \"bucket_name\":\"sample\",         \"can_read\":true,         \"can_write\":true,         \"created_at\":\"2021-11-11T13:36:08.767118492Z\"       }     ],     \"created_at\":\"2021-11-11T13:36:08.690384415Z\",     \"display_name\":\"sample_permission\",     \"id\":619   } } ```  このパーミッションのアクセスキーを発行するには以下のような入力を行います。 この時、パーミッション作成時に発行されたID（ここでは619とします）をパスパラメータに含める必要があります。  ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X POST \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619/keys ```  コマンド結果には以下のフィールドが含まれます。  * `created_at` : 作成日時 * `id` : アクセスキーID * `secret` : シークレットアクセスキー  ``` # 出力結果サンプル {   \"data\": {     \"created_at\": \"2021-11-04T07:42:41.121418479Z\",     \"id\": \"XPJK4SC9883N91RHR253\",     \"secret\": \"jqRaUo5l+EiEYqP8wos9exbmFfq4/vG8CLPYI2XN\"   } } ```  パーミッションアクセスキーを削除するには以下のような入力を行います。 この時、パーミッションアクセスキー発行時に出力されたIDをパスパラメータに含める必要があります。 ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619/keys/XPJK4SC9883N91RHR253 ```  パーミッションを削除するには以下のような入力を行います。 この時、パーミッション作成時に発行されたID（ここでは619とします）をパスパラメータに含める必要があります。 ``` # 入力サンプル curl -u '01234567-89ab-cdef-0123-456789abcdef:SAMPLETOKENSAMPLETOKENSAMPLETOKENSAM' \\      -X DELETE \\      https://secure.sakura.ad.jp/cloud/zone/is1a/api/objectstorage/1.0/isk01/v2/permissions/619 ``` ----

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package object_storage

import (
	"encoding/json"
)

// checks if the ModelCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCluster{}

// ModelCluster struct for ModelCluster
type ModelCluster struct {
	// API Servers Zones
	ApiZone []string `json:"api_zone,omitempty"`
	// URL of Control Panel
	ControlPanelUrl *string `json:"control_panel_url,omitempty"`
	// Display Name (en-us)
	DisplayNameEnUs *string `json:"display_name_en_us,omitempty"`
	// Display Name (ja)
	DisplayNameJa *string `json:"display_name_ja,omitempty"`
	// Display Name (Depending on Accept-Language)
	DisplayName *string `json:"display_name,omitempty"`
	// Display Order (Can be ignored)
	DisplayOrder *int32 `json:"display_order,omitempty"`
	// Endpoint Base of Cluster
	EndpointBase *string `json:"endpoint_base,omitempty"`
	// URL of IAM-compat API
	IamEndpoint *string `json:"iam_endpoint,omitempty"`
	// URL of IAM-compat API (w/ resigning)
	IamEndpointForControlPanel *string `json:"iam_endpoint_for_control_panel,omitempty"`
	Id                         *string `json:"id,omitempty"`
	Region                     *string `json:"region,omitempty"`
	// URL of S3-compat API
	S3Endpoint *string `json:"s3_endpoint,omitempty"`
	// URL of S3-compat API (w/ resigning)
	S3EndpointForControlPanel *string `json:"s3_endpoint_for_control_panel,omitempty"`
	// Storage Servers Zones
	StorageZone []string `json:"storage_zone,omitempty"`
}

// NewModelCluster instantiates a new ModelCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCluster() *ModelCluster {
	this := ModelCluster{}
	return &this
}

// NewModelClusterWithDefaults instantiates a new ModelCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelClusterWithDefaults() *ModelCluster {
	this := ModelCluster{}
	return &this
}

// GetApiZone returns the ApiZone field value if set, zero value otherwise.
func (o *ModelCluster) GetApiZone() []string {
	if o == nil || IsNil(o.ApiZone) {
		var ret []string
		return ret
	}
	return o.ApiZone
}

// GetApiZoneOk returns a tuple with the ApiZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetApiZoneOk() ([]string, bool) {
	if o == nil || IsNil(o.ApiZone) {
		return nil, false
	}
	return o.ApiZone, true
}

// HasApiZone returns a boolean if a field has been set.
func (o *ModelCluster) HasApiZone() bool {
	if o != nil && !IsNil(o.ApiZone) {
		return true
	}

	return false
}

// SetApiZone gets a reference to the given []string and assigns it to the ApiZone field.
func (o *ModelCluster) SetApiZone(v []string) {
	o.ApiZone = v
}

// GetControlPanelUrl returns the ControlPanelUrl field value if set, zero value otherwise.
func (o *ModelCluster) GetControlPanelUrl() string {
	if o == nil || IsNil(o.ControlPanelUrl) {
		var ret string
		return ret
	}
	return *o.ControlPanelUrl
}

// GetControlPanelUrlOk returns a tuple with the ControlPanelUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetControlPanelUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ControlPanelUrl) {
		return nil, false
	}
	return o.ControlPanelUrl, true
}

// HasControlPanelUrl returns a boolean if a field has been set.
func (o *ModelCluster) HasControlPanelUrl() bool {
	if o != nil && !IsNil(o.ControlPanelUrl) {
		return true
	}

	return false
}

// SetControlPanelUrl gets a reference to the given string and assigns it to the ControlPanelUrl field.
func (o *ModelCluster) SetControlPanelUrl(v string) {
	o.ControlPanelUrl = &v
}

// GetDisplayNameEnUs returns the DisplayNameEnUs field value if set, zero value otherwise.
func (o *ModelCluster) GetDisplayNameEnUs() string {
	if o == nil || IsNil(o.DisplayNameEnUs) {
		var ret string
		return ret
	}
	return *o.DisplayNameEnUs
}

// GetDisplayNameEnUsOk returns a tuple with the DisplayNameEnUs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetDisplayNameEnUsOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayNameEnUs) {
		return nil, false
	}
	return o.DisplayNameEnUs, true
}

// HasDisplayNameEnUs returns a boolean if a field has been set.
func (o *ModelCluster) HasDisplayNameEnUs() bool {
	if o != nil && !IsNil(o.DisplayNameEnUs) {
		return true
	}

	return false
}

// SetDisplayNameEnUs gets a reference to the given string and assigns it to the DisplayNameEnUs field.
func (o *ModelCluster) SetDisplayNameEnUs(v string) {
	o.DisplayNameEnUs = &v
}

// GetDisplayNameJa returns the DisplayNameJa field value if set, zero value otherwise.
func (o *ModelCluster) GetDisplayNameJa() string {
	if o == nil || IsNil(o.DisplayNameJa) {
		var ret string
		return ret
	}
	return *o.DisplayNameJa
}

// GetDisplayNameJaOk returns a tuple with the DisplayNameJa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetDisplayNameJaOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayNameJa) {
		return nil, false
	}
	return o.DisplayNameJa, true
}

// HasDisplayNameJa returns a boolean if a field has been set.
func (o *ModelCluster) HasDisplayNameJa() bool {
	if o != nil && !IsNil(o.DisplayNameJa) {
		return true
	}

	return false
}

// SetDisplayNameJa gets a reference to the given string and assigns it to the DisplayNameJa field.
func (o *ModelCluster) SetDisplayNameJa(v string) {
	o.DisplayNameJa = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ModelCluster) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ModelCluster) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ModelCluster) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *ModelCluster) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *ModelCluster) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *ModelCluster) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetEndpointBase returns the EndpointBase field value if set, zero value otherwise.
func (o *ModelCluster) GetEndpointBase() string {
	if o == nil || IsNil(o.EndpointBase) {
		var ret string
		return ret
	}
	return *o.EndpointBase
}

// GetEndpointBaseOk returns a tuple with the EndpointBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetEndpointBaseOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointBase) {
		return nil, false
	}
	return o.EndpointBase, true
}

// HasEndpointBase returns a boolean if a field has been set.
func (o *ModelCluster) HasEndpointBase() bool {
	if o != nil && !IsNil(o.EndpointBase) {
		return true
	}

	return false
}

// SetEndpointBase gets a reference to the given string and assigns it to the EndpointBase field.
func (o *ModelCluster) SetEndpointBase(v string) {
	o.EndpointBase = &v
}

// GetIamEndpoint returns the IamEndpoint field value if set, zero value otherwise.
func (o *ModelCluster) GetIamEndpoint() string {
	if o == nil || IsNil(o.IamEndpoint) {
		var ret string
		return ret
	}
	return *o.IamEndpoint
}

// GetIamEndpointOk returns a tuple with the IamEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetIamEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.IamEndpoint) {
		return nil, false
	}
	return o.IamEndpoint, true
}

// HasIamEndpoint returns a boolean if a field has been set.
func (o *ModelCluster) HasIamEndpoint() bool {
	if o != nil && !IsNil(o.IamEndpoint) {
		return true
	}

	return false
}

// SetIamEndpoint gets a reference to the given string and assigns it to the IamEndpoint field.
func (o *ModelCluster) SetIamEndpoint(v string) {
	o.IamEndpoint = &v
}

// GetIamEndpointForControlPanel returns the IamEndpointForControlPanel field value if set, zero value otherwise.
func (o *ModelCluster) GetIamEndpointForControlPanel() string {
	if o == nil || IsNil(o.IamEndpointForControlPanel) {
		var ret string
		return ret
	}
	return *o.IamEndpointForControlPanel
}

// GetIamEndpointForControlPanelOk returns a tuple with the IamEndpointForControlPanel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetIamEndpointForControlPanelOk() (*string, bool) {
	if o == nil || IsNil(o.IamEndpointForControlPanel) {
		return nil, false
	}
	return o.IamEndpointForControlPanel, true
}

// HasIamEndpointForControlPanel returns a boolean if a field has been set.
func (o *ModelCluster) HasIamEndpointForControlPanel() bool {
	if o != nil && !IsNil(o.IamEndpointForControlPanel) {
		return true
	}

	return false
}

// SetIamEndpointForControlPanel gets a reference to the given string and assigns it to the IamEndpointForControlPanel field.
func (o *ModelCluster) SetIamEndpointForControlPanel(v string) {
	o.IamEndpointForControlPanel = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelCluster) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelCluster) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelCluster) SetId(v string) {
	o.Id = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ModelCluster) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ModelCluster) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ModelCluster) SetRegion(v string) {
	o.Region = &v
}

// GetS3Endpoint returns the S3Endpoint field value if set, zero value otherwise.
func (o *ModelCluster) GetS3Endpoint() string {
	if o == nil || IsNil(o.S3Endpoint) {
		var ret string
		return ret
	}
	return *o.S3Endpoint
}

// GetS3EndpointOk returns a tuple with the S3Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetS3EndpointOk() (*string, bool) {
	if o == nil || IsNil(o.S3Endpoint) {
		return nil, false
	}
	return o.S3Endpoint, true
}

// HasS3Endpoint returns a boolean if a field has been set.
func (o *ModelCluster) HasS3Endpoint() bool {
	if o != nil && !IsNil(o.S3Endpoint) {
		return true
	}

	return false
}

// SetS3Endpoint gets a reference to the given string and assigns it to the S3Endpoint field.
func (o *ModelCluster) SetS3Endpoint(v string) {
	o.S3Endpoint = &v
}

// GetS3EndpointForControlPanel returns the S3EndpointForControlPanel field value if set, zero value otherwise.
func (o *ModelCluster) GetS3EndpointForControlPanel() string {
	if o == nil || IsNil(o.S3EndpointForControlPanel) {
		var ret string
		return ret
	}
	return *o.S3EndpointForControlPanel
}

// GetS3EndpointForControlPanelOk returns a tuple with the S3EndpointForControlPanel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetS3EndpointForControlPanelOk() (*string, bool) {
	if o == nil || IsNil(o.S3EndpointForControlPanel) {
		return nil, false
	}
	return o.S3EndpointForControlPanel, true
}

// HasS3EndpointForControlPanel returns a boolean if a field has been set.
func (o *ModelCluster) HasS3EndpointForControlPanel() bool {
	if o != nil && !IsNil(o.S3EndpointForControlPanel) {
		return true
	}

	return false
}

// SetS3EndpointForControlPanel gets a reference to the given string and assigns it to the S3EndpointForControlPanel field.
func (o *ModelCluster) SetS3EndpointForControlPanel(v string) {
	o.S3EndpointForControlPanel = &v
}

// GetStorageZone returns the StorageZone field value if set, zero value otherwise.
func (o *ModelCluster) GetStorageZone() []string {
	if o == nil || IsNil(o.StorageZone) {
		var ret []string
		return ret
	}
	return o.StorageZone
}

// GetStorageZoneOk returns a tuple with the StorageZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCluster) GetStorageZoneOk() ([]string, bool) {
	if o == nil || IsNil(o.StorageZone) {
		return nil, false
	}
	return o.StorageZone, true
}

// HasStorageZone returns a boolean if a field has been set.
func (o *ModelCluster) HasStorageZone() bool {
	if o != nil && !IsNil(o.StorageZone) {
		return true
	}

	return false
}

// SetStorageZone gets a reference to the given []string and assigns it to the StorageZone field.
func (o *ModelCluster) SetStorageZone(v []string) {
	o.StorageZone = v
}

func (o ModelCluster) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiZone) {
		toSerialize["api_zone"] = o.ApiZone
	}
	if !IsNil(o.ControlPanelUrl) {
		toSerialize["control_panel_url"] = o.ControlPanelUrl
	}
	if !IsNil(o.DisplayNameEnUs) {
		toSerialize["display_name_en_us"] = o.DisplayNameEnUs
	}
	if !IsNil(o.DisplayNameJa) {
		toSerialize["display_name_ja"] = o.DisplayNameJa
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["display_order"] = o.DisplayOrder
	}
	if !IsNil(o.EndpointBase) {
		toSerialize["endpoint_base"] = o.EndpointBase
	}
	if !IsNil(o.IamEndpoint) {
		toSerialize["iam_endpoint"] = o.IamEndpoint
	}
	if !IsNil(o.IamEndpointForControlPanel) {
		toSerialize["iam_endpoint_for_control_panel"] = o.IamEndpointForControlPanel
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.S3Endpoint) {
		toSerialize["s3_endpoint"] = o.S3Endpoint
	}
	if !IsNil(o.S3EndpointForControlPanel) {
		toSerialize["s3_endpoint_for_control_panel"] = o.S3EndpointForControlPanel
	}
	if !IsNil(o.StorageZone) {
		toSerialize["storage_zone"] = o.StorageZone
	}
	return toSerialize, nil
}

type NullableModelCluster struct {
	value *ModelCluster
	isSet bool
}

func (v NullableModelCluster) Get() *ModelCluster {
	return v.value
}

func (v *NullableModelCluster) Set(val *ModelCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCluster(val *ModelCluster) *NullableModelCluster {
	return &NullableModelCluster{value: val, isSet: true}
}

func (v NullableModelCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
