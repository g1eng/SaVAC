# PatchApplicationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeoutSeconds** | Pointer to **int32** | アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限 | [optional] 
**Port** | Pointer to **int32** | アプリケーションがリクエストを待ち受けるポート番号 | [optional] 
**MinScale** | Pointer to **int32** | アプリケーション全体の最小スケール数 | [optional] 
**MaxScale** | Pointer to **int32** | アプリケーション全体の最大スケール数 | [optional] 
**Components** | Pointer to [**[]PatchApplicationBodyComponent**](PatchApplicationBodyComponent.md) | アプリケーションのコンポーネント情報 | [optional] 
**AllTrafficAvailable** | Pointer to **bool** | アプリケーションを最新のバージョンにすべてのトラフィックを割り当てるかどうか | [optional] 

## Methods

### NewPatchApplicationBody

`func NewPatchApplicationBody() *PatchApplicationBody`

NewPatchApplicationBody instantiates a new PatchApplicationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationBodyWithDefaults

`func NewPatchApplicationBodyWithDefaults() *PatchApplicationBody`

NewPatchApplicationBodyWithDefaults instantiates a new PatchApplicationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeoutSeconds

`func (o *PatchApplicationBody) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *PatchApplicationBody) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *PatchApplicationBody) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *PatchApplicationBody) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetPort

`func (o *PatchApplicationBody) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchApplicationBody) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchApplicationBody) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchApplicationBody) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetMinScale

`func (o *PatchApplicationBody) GetMinScale() int32`

GetMinScale returns the MinScale field if non-nil, zero value otherwise.

### GetMinScaleOk

`func (o *PatchApplicationBody) GetMinScaleOk() (*int32, bool)`

GetMinScaleOk returns a tuple with the MinScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScale

`func (o *PatchApplicationBody) SetMinScale(v int32)`

SetMinScale sets MinScale field to given value.

### HasMinScale

`func (o *PatchApplicationBody) HasMinScale() bool`

HasMinScale returns a boolean if a field has been set.

### GetMaxScale

`func (o *PatchApplicationBody) GetMaxScale() int32`

GetMaxScale returns the MaxScale field if non-nil, zero value otherwise.

### GetMaxScaleOk

`func (o *PatchApplicationBody) GetMaxScaleOk() (*int32, bool)`

GetMaxScaleOk returns a tuple with the MaxScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScale

`func (o *PatchApplicationBody) SetMaxScale(v int32)`

SetMaxScale sets MaxScale field to given value.

### HasMaxScale

`func (o *PatchApplicationBody) HasMaxScale() bool`

HasMaxScale returns a boolean if a field has been set.

### GetComponents

`func (o *PatchApplicationBody) GetComponents() []PatchApplicationBodyComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *PatchApplicationBody) GetComponentsOk() (*[]PatchApplicationBodyComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *PatchApplicationBody) SetComponents(v []PatchApplicationBodyComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *PatchApplicationBody) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetAllTrafficAvailable

`func (o *PatchApplicationBody) GetAllTrafficAvailable() bool`

GetAllTrafficAvailable returns the AllTrafficAvailable field if non-nil, zero value otherwise.

### GetAllTrafficAvailableOk

`func (o *PatchApplicationBody) GetAllTrafficAvailableOk() (*bool, bool)`

GetAllTrafficAvailableOk returns a tuple with the AllTrafficAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllTrafficAvailable

`func (o *PatchApplicationBody) SetAllTrafficAvailable(v bool)`

SetAllTrafficAvailable sets AllTrafficAvailable field to given value.

### HasAllTrafficAvailable

`func (o *PatchApplicationBody) HasAllTrafficAvailable() bool`

HasAllTrafficAvailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


