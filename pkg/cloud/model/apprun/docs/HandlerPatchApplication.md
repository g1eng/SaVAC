# HandlerPatchApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | アプリケーションID | 
**Name** | **string** | アプリケーション名 | 
**TimeoutSeconds** | **int32** | アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限 | 
**Port** | **int32** | アプリケーションがリクエストを待ち受けるポート番号 | 
**MinScale** | **int32** | アプリケーション全体の最小スケール数 | 
**MaxScale** | **int32** | アプリケーション全体の最大スケール数 | 
**Components** | [**[]HandlerApplicationComponent**](HandlerApplicationComponent.md) | アプリケーションのコンポーネント情報 | 
**Status** | **string** | アプリケーションステータス | 
**PublicUrl** | **string** | 公開URL | 
**UpdatedAt** | **time.Time** | 更新日時 | 

## Methods

### NewHandlerPatchApplication

`func NewHandlerPatchApplication(id string, name string, timeoutSeconds int32, port int32, minScale int32, maxScale int32, components []HandlerApplicationComponent, status string, publicUrl string, updatedAt time.Time, ) *HandlerPatchApplication`

NewHandlerPatchApplication instantiates a new HandlerPatchApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerPatchApplicationWithDefaults

`func NewHandlerPatchApplicationWithDefaults() *HandlerPatchApplication`

NewHandlerPatchApplicationWithDefaults instantiates a new HandlerPatchApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HandlerPatchApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlerPatchApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlerPatchApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HandlerPatchApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlerPatchApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlerPatchApplication) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutSeconds

`func (o *HandlerPatchApplication) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *HandlerPatchApplication) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *HandlerPatchApplication) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetPort

`func (o *HandlerPatchApplication) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlerPatchApplication) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlerPatchApplication) SetPort(v int32)`

SetPort sets Port field to given value.


### GetMinScale

`func (o *HandlerPatchApplication) GetMinScale() int32`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *HandlerPatchApplication) GetMinScaleOk() (*int32, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *HandlerPatchApplication) SetMinScale(v int32)`

SetMinScale sets MinScale field to given value.


### GetMaxScale

`func (o *HandlerPatchApplication) GetMaxScale() int32`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *HandlerPatchApplication) GetMaxScaleOk() (*int32, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *HandlerPatchApplication) SetMaxScale(v int32)`

SetMaxScale sets MaxScale field to given value.


### GetComponents

`func (o *HandlerPatchApplication) GetComponents() []HandlerApplicationComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HandlerPatchApplication) GetComponentsOk() (*[]HandlerApplicationComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HandlerPatchApplication) SetComponents(v []HandlerApplicationComponent)`

SetComponents sets Components field to given value.


### GetStatus

`func (o *HandlerPatchApplication) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerPatchApplication) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerPatchApplication) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPublicUrl

`func (o *HandlerPatchApplication) GetPublicUrl() string`

GetPublicUrl returns the PublicUrl field if non-nil, zero value otherwise.

### GetPublicUrlOk

`func (o *HandlerPatchApplication) GetPublicUrlOk() (*string, bool)`

GetPublicUrlOk returns a tuple with the PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrl

`func (o *HandlerPatchApplication) SetPublicUrl(v string)`

SetPublicUrl sets PublicUrl field to given value.


### GetUpdatedAt

`func (o *HandlerPatchApplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HandlerPatchApplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HandlerPatchApplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


