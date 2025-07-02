# PostApplicationBodyComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | コンポーネント名 | 
**MaxCpu** | **string** | コンポーネントの最大CPU数 | 
**MaxMemory** | **string** | コンポーネントの最大メモリ | 
**DeploySource** | [**PostApplicationBodyComponentDeploySource**](PostApplicationBodyComponentDeploySource.md) |  | 
**Env** | Pointer to [**[]PostApplicationBodyComponentEnv**](PostApplicationBodyComponentEnv.md) | コンポーネントに渡す環境変数 | [optional] 
**Probe** | Pointer to [**NullablePostApplicationBodyComponentProbe**](PostApplicationBodyComponentProbe.md) |  | [optional] 

## Methods

### NewPostApplicationBodyComponent

`func NewPostApplicationBodyComponent(name string, maxCpu string, maxMemory string, deploySource PostApplicationBodyComponentDeploySource, ) *PostApplicationBodyComponent`

NewPostApplicationBodyComponent instantiates a new PostApplicationBodyComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApplicationBodyComponentWithDefaults

`func NewPostApplicationBodyComponentWithDefaults() *PostApplicationBodyComponent`

NewPostApplicationBodyComponentWithDefaults instantiates a new PostApplicationBodyComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostApplicationBodyComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostApplicationBodyComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostApplicationBodyComponent) SetName(v string)`

SetName sets Name field to given value.


### GetMaxCpu

`func (o *PostApplicationBodyComponent) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *PostApplicationBodyComponent) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *PostApplicationBodyComponent) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.


### GetMaxMemory

`func (o *PostApplicationBodyComponent) GetMaxMemory() string`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *PostApplicationBodyComponent) GetMaxMemoryOk() (*string, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *PostApplicationBodyComponent) SetMaxMemory(v string)`

SetMaxMemory sets MaxMemory field to given value.


### GetDeploySource

`func (o *PostApplicationBodyComponent) GetDeploySource() PostApplicationBodyComponentDeploySource`

GetDeploySource returns the DeploySource field if non-nil, zero value otherwise.

### GetDeploySourceOk

`func (o *PostApplicationBodyComponent) GetDeploySourceOk() (*PostApplicationBodyComponentDeploySource, bool)`

GetDeploySourceOk returns a tuple with the DeploySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploySource

`func (o *PostApplicationBodyComponent) SetDeploySource(v PostApplicationBodyComponentDeploySource)`

SetDeploySource sets DeploySource field to given value.


### GetEnv

`func (o *PostApplicationBodyComponent) GetEnv() []PostApplicationBodyComponentEnv`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *PostApplicationBodyComponent) GetEnvOk() (*[]PostApplicationBodyComponentEnv, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *PostApplicationBodyComponent) SetEnv(v []PostApplicationBodyComponentEnv)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *PostApplicationBodyComponent) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnvNil

`func (o *PostApplicationBodyComponent) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *PostApplicationBodyComponent) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetProbe

`func (o *PostApplicationBodyComponent) GetProbe() PostApplicationBodyComponentProbe`

GetProbe returns the Probe field if non-nil, zero value otherwise.

### GetProbeOk

`func (o *PostApplicationBodyComponent) GetProbeOk() (*PostApplicationBodyComponentProbe, bool)`

GetProbeOk returns a tuple with the Probe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbe

`func (o *PostApplicationBodyComponent) SetProbe(v PostApplicationBodyComponentProbe)`

SetProbe sets Probe field to given value.

### HasProbe

`func (o *PostApplicationBodyComponent) HasProbe() bool`

HasProbe returns a boolean if a field has been set.

### SetProbeNil

`func (o *PostApplicationBodyComponent) SetProbeNil(b bool)`

 SetProbeNil sets the value for Probe to be an explicit nil

### UnsetProbe
`func (o *PostApplicationBodyComponent) UnsetProbe()`

UnsetProbe ensures that no value is present for Probe, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


