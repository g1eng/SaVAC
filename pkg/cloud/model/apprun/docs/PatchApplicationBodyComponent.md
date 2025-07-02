# PatchApplicationBodyComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | コンポーネント名 | 
**MaxCpu** | **string** | コンポーネントの最大CPU数 | 
**MaxMemory** | **string** | コンポーネントの最大メモリ | 
**DeploySource** | [**PatchApplicationBodyComponentDeploySource**](PatchApplicationBodyComponentDeploySource.md) |  | 
**Env** | Pointer to [**[]PatchApplicationBodyComponentEnv**](PatchApplicationBodyComponentEnv.md) | コンポーネントに渡す環境変数 | [optional] 
**Probe** | Pointer to [**NullablePatchApplicationBodyComponentProbe**](PatchApplicationBodyComponentProbe.md) |  | [optional] 

## Methods

### NewPatchApplicationBodyComponent

`func NewPatchApplicationBodyComponent(name string, maxCpu string, maxMemory string, deploySource PatchApplicationBodyComponentDeploySource, ) *PatchApplicationBodyComponent`

NewPatchApplicationBodyComponent instantiates a new PatchApplicationBodyComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationBodyComponentWithDefaults

`func NewPatchApplicationBodyComponentWithDefaults() *PatchApplicationBodyComponent`

NewPatchApplicationBodyComponentWithDefaults instantiates a new PatchApplicationBodyComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchApplicationBodyComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchApplicationBodyComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchApplicationBodyComponent) SetName(v string)`

SetName sets Name field to given value.


### GetMaxCpu

`func (o *PatchApplicationBodyComponent) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *PatchApplicationBodyComponent) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *PatchApplicationBodyComponent) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.


### GetMaxMemory

`func (o *PatchApplicationBodyComponent) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *PatchApplicationBodyComponent) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *PatchApplicationBodyComponent) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.


### GetDeploySource

`func (o *PatchApplicationBodyComponent) GetDeploySource() PatchApplicationBodyComponentDeploySource`

GetDeploySource returns the DeploySource field if non-nil, zero value otherwise.

### GetDeploySourceOk

`func (o *PatchApplicationBodyComponent) GetDeploySourceOk() (*PatchApplicationBodyComponentDeploySource, bool)`

GetDeploySourceOk returns a tuple with the DeploySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySource

`func (o *PatchApplicationBodyComponent) SetDeploySource(v PatchApplicationBodyComponentDeploySource)`

SetDeploySource sets DeploySource field to given value.


### GetEnv

`func (o *PatchApplicationBodyComponent) GetEnv() []PatchApplicationBodyComponentEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PatchApplicationBodyComponent) GetEnvOk() (*[]PatchApplicationBodyComponentEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PatchApplicationBodyComponent) SetEnv(v []PatchApplicationBodyComponentEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PatchApplicationBodyComponent) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnvNil

`func (o *PatchApplicationBodyComponent) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *PatchApplicationBodyComponent) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetProbe

`func (o *PatchApplicationBodyComponent) GetProbe() PatchApplicationBodyComponentProbe`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *PatchApplicationBodyComponent) GetProbeOk() (*PatchApplicationBodyComponentProbe, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *PatchApplicationBodyComponent) SetProbe(v PatchApplicationBodyComponentProbe)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *PatchApplicationBodyComponent) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### SetProbeNil

`func (o *PatchApplicationBodyComponent) SetProbeNil(b bool)`

 SetProbeNil sets the value for Probe to be an explicit nil

### UnsetProbe
`func (o *PatchApplicationBodyComponent) UnsetProbe()`

UnsetProbe ensures that no value is present for Probe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


