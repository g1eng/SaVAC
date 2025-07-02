# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | アプリケーションID | 
**Name** | **string** | アプリケーション名 | 
**TimeoutSeconds** | **int32** | アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限 | [default to 60]
**Port** | **int32** | アプリケーションがリクエストを待ち受けるポート番号 | 
**MinScale** | **int32** | アプリケーション全体の最小スケール数 | [default to 0]
**MaxScale** | **int32** | アプリケーション全体の最大スケール数 | [default to 10]
**Components** | [**[]HandlerApplicationComponent**](HandlerApplicationComponent.md) | アプリケーションのコンポーネント情報 | 
**Status** | **string** | アプリケーションステータス | 
**PublicUrl** | **string** | 公開URL | 
**CreatedAt** | **time.Time** | 作成日時 | 

## Methods

### NewApplication

`func NewApplication(id string, name string, timeoutSeconds int32, port int32, minScale int32, maxScale int32, components []HandlerApplicationComponent, status string, publicUrl string, createdAt time.Time, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutSeconds

`func (o *Application) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *Application) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *Application) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetPort

`func (o *Application) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Application) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Application) SetPort(v int32)`

SetPort sets Port field to given value.


### GetMinScale

`func (o *Application) GetMinScale() int32`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *Application) GetMinScaleOk() (*int32, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *Application) SetMinScale(v int32)`

SetMinScale sets MinScale field to given value.


### GetMaxScale

`func (o *Application) GetMaxScale() int32`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *Application) GetMaxScaleOk() (*int32, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *Application) SetMaxScale(v int32)`

SetMaxScale sets MaxScale field to given value.


### GetComponents

`func (o *Application) GetComponents() []HandlerApplicationComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Application) GetComponentsOk() (*[]HandlerApplicationComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Application) SetComponents(v []HandlerApplicationComponent)`

SetComponents sets Components field to given value.


### GetStatus

`func (o *Application) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Application) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Application) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPublicUrl

`func (o *Application) GetPublicUrl() string`

GetPublicUrl returns the PublicUrl field if non-nil, zero value otherwise.

### GetPublicUrlOk

`func (o *Application) GetPublicUrlOk() (*string, bool)`

GetPublicUrlOk returns a tuple with the PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrl

`func (o *Application) SetPublicUrl(v string)`

SetPublicUrl sets PublicUrl field to given value.


### GetCreatedAt

`func (o *Application) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Application) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Application) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


