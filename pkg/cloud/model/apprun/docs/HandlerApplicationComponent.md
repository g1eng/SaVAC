# HandlerApplicationComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | コンポーネント名 | 
**MaxCpu** | **string** | コンポーネントの最大CPU数 | 
**MaxMemory** | **string** | コンポーネントの最大メモリ | 
**DeploySource** | [**HandlerApplicationComponentDeploySource**](HandlerApplicationComponentDeploySource.md) |  | 
**Env** | Pointer to [**[]HandlerApplicationComponentEnv**](HandlerApplicationComponentEnv.md) | コンポーネントに渡す環境変数 | [optional] 
**Probe** | Pointer to [**NullableHandlerApplicationComponentProbe**](HandlerApplicationComponentProbe.md) |  | [optional] 

## Methods

### NewHandlerApplicationComponent

`func NewHandlerApplicationComponent(name string, maxCpu string, maxMemory string, deploySource HandlerApplicationComponentDeploySource, ) *HandlerApplicationComponent`

NewHandlerApplicationComponent instantiates a new HandlerApplicationComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerApplicationComponentWithDefaults

`func NewHandlerApplicationComponentWithDefaults() *HandlerApplicationComponent`

NewHandlerApplicationComponentWithDefaults instantiates a new HandlerApplicationComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HandlerApplicationComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlerApplicationComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlerApplicationComponent) SetName(v string)`

SetName sets Name field to given value.


### GetMaxCpu

`func (o *HandlerApplicationComponent) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *HandlerApplicationComponent) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *HandlerApplicationComponent) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.


### GetMaxMemory

`func (o *HandlerApplicationComponent) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *HandlerApplicationComponent) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *HandlerApplicationComponent) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.


### GetDeploySource

`func (o *HandlerApplicationComponent) GetDeploySource() HandlerApplicationComponentDeploySource`

GetDeploySource returns the DeploySource field if non-nil, zero value otherwise.

### GetDeploySourceOk

`func (o *HandlerApplicationComponent) GetDeploySourceOk() (*HandlerApplicationComponentDeploySource, bool)`

GetDeploySourceOk returns a tuple with the DeploySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySource

`func (o *HandlerApplicationComponent) SetDeploySource(v HandlerApplicationComponentDeploySource)`

SetDeploySource sets DeploySource field to given value.


### GetEnv

`func (o *HandlerApplicationComponent) GetEnv() []HandlerApplicationComponentEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *HandlerApplicationComponent) GetEnvOk() (*[]HandlerApplicationComponentEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *HandlerApplicationComponent) SetEnv(v []HandlerApplicationComponentEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *HandlerApplicationComponent) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetProbe

`func (o *HandlerApplicationComponent) GetProbe() HandlerApplicationComponentProbe`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *HandlerApplicationComponent) GetProbeOk() (*HandlerApplicationComponentProbe, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *HandlerApplicationComponent) SetProbe(v HandlerApplicationComponentProbe)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *HandlerApplicationComponent) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### SetProbeNil

`func (o *HandlerApplicationComponent) SetProbeNil(b bool)`

 SetProbeNil sets the value for Probe to be an explicit nil

### UnsetProbe
`func (o *HandlerApplicationComponent) UnsetProbe()`

UnsetProbe ensures that no value is present for Probe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


