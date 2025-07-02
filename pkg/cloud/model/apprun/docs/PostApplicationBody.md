# PostApplicationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | アプリケーション名 | 
**TimeoutSeconds** | **int32** | アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限 | [default to 60]
**Port** | **int32** | アプリケーションがリクエストを待ち受けるポート番号 | 
**MinScale** | **int32** | アプリケーション全体の最小スケール数 | [default to 0]
**MaxScale** | **int32** | アプリケーション全体の最大スケール数 | [default to 10]
**Components** | [**[]PostApplicationBodyComponent**](PostApplicationBodyComponent.md) | アプリケーションのコンポーネント情報 | 

## Methods

### NewPostApplicationBody

`func NewPostApplicationBody(name string, timeoutSeconds int32, port int32, minScale int32, maxScale int32, components []PostApplicationBodyComponent, ) *PostApplicationBody`

NewPostApplicationBody instantiates a new PostApplicationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApplicationBodyWithDefaults

`func NewPostApplicationBodyWithDefaults() *PostApplicationBody`

NewPostApplicationBodyWithDefaults instantiates a new PostApplicationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApplicationBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApplicationBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApplicationBody) SetName(v string)`

SetName sets Name field to given value.


### GetTimeoutSeconds

`func (o *PostApplicationBody) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *PostApplicationBody) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *PostApplicationBody) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.


### GetPort

`func (o *PostApplicationBody) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostApplicationBody) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostApplicationBody) SetPort(v int32)`

SetPort sets Port field to given value.


### GetMinScale

`func (o *PostApplicationBody) GetMinScale() int32`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *PostApplicationBody) GetMinScaleOk() (*int32, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *PostApplicationBody) SetMinScale(v int32)`

SetMinScale sets MinScale field to given value.


### GetMaxScale

`func (o *PostApplicationBody) GetMaxScale() int32`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *PostApplicationBody) GetMaxScaleOk() (*int32, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *PostApplicationBody) SetMaxScale(v int32)`

SetMaxScale sets MaxScale field to given value.


### GetComponents

`func (o *PostApplicationBody) GetComponents() []PostApplicationBodyComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PostApplicationBody) GetComponentsOk() (*[]PostApplicationBodyComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PostApplicationBody) SetComponents(v []PostApplicationBodyComponent)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


