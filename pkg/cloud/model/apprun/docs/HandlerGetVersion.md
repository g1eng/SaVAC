# HandlerGetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | バージョンID | 
**Name** | **string** | バージョン名 | 
**Status** | **string** | バージョンステータス | 
**TimeoutSeconds** | **int32** | アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限 | [default to 60]
**Port** | **int32** | アプリケーションがリクエストを待ち受けるポート番号 | 
**MinScale** | **int32** | バージョンの最小スケール数 | [default to 0]
**MaxScale** | **int32** | バージョンの最大スケール数 | [default to 10]
**Components** | [**[]HandlerApplicationComponent**](HandlerApplicationComponent.md) | バージョンのコンポーネント情報 | 
**CreatedAt** | **time.Time** | 作成日時 | 

## Methods

### NewHandlerGetVersion

`func NewHandlerGetVersion(id string, name string, status string, timeoutSeconds int32, port int32, minScale int32, maxScale int32, components []HandlerApplicationComponent, createdAt time.Time, ) *HandlerGetVersion`

NewHandlerGetVersion instantiates a new HandlerGetVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerGetVersionWithDefaults

`func NewHandlerGetVersionWithDefaults() *HandlerGetVersion`

NewHandlerGetVersionWithDefaults instantiates a new HandlerGetVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HandlerGetVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlerGetVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlerGetVersion) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HandlerGetVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlerGetVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlerGetVersion) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *HandlerGetVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerGetVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerGetVersion) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimeoutSeconds

`func (o *HandlerGetVersion) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *HandlerGetVersion) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *HandlerGetVersion) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetPort

`func (o *HandlerGetVersion) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlerGetVersion) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlerGetVersion) SetPort(v int32)`

SetPort sets Port field to given value.


### GetMinScale

`func (o *HandlerGetVersion) GetMinScale() int32`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *HandlerGetVersion) GetMinScaleOk() (*int32, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *HandlerGetVersion) SetMinScale(v int32)`

SetMinScale sets MinScale field to given value.


### GetMaxScale

`func (o *HandlerGetVersion) GetMaxScale() int32`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *HandlerGetVersion) GetMaxScaleOk() (*int32, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *HandlerGetVersion) SetMaxScale(v int32)`

SetMaxScale sets MaxScale field to given value.


### GetComponents

`func (o *HandlerGetVersion) GetComponents() []HandlerApplicationComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HandlerGetVersion) GetComponentsOk() (*[]HandlerApplicationComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HandlerGetVersion) SetComponents(v []HandlerApplicationComponent)`

SetComponents sets Components field to given value.


### GetCreatedAt

`func (o *HandlerGetVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlerGetVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlerGetVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


