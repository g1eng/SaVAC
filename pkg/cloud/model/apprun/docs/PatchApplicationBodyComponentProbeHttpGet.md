# PatchApplicationBodyComponentProbeHttpGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | HTTPサーバーへアクセスしプローブをチェックする際のパス | 
**Port** | **int32** | HTTPサーバーへアクセスしプローブをチェックする際のポート番号 | 
**Headers** | Pointer to [**[]PatchApplicationBodyComponentProbeHttpGetHeader**](PatchApplicationBodyComponentProbeHttpGetHeader.md) |  | [optional] 

## Methods

### NewPatchApplicationBodyComponentProbeHttpGet

`func NewPatchApplicationBodyComponentProbeHttpGet(path string, port int32, ) *PatchApplicationBodyComponentProbeHttpGet`

NewPatchApplicationBodyComponentProbeHttpGet instantiates a new PatchApplicationBodyComponentProbeHttpGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationBodyComponentProbeHttpGetWithDefaults

`func NewPatchApplicationBodyComponentProbeHttpGetWithDefaults() *PatchApplicationBodyComponentProbeHttpGet`

NewPatchApplicationBodyComponentProbeHttpGetWithDefaults instantiates a new PatchApplicationBodyComponentProbeHttpGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PatchApplicationBodyComponentProbeHttpGet) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchApplicationBodyComponentProbeHttpGet) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchApplicationBodyComponentProbeHttpGet) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *PatchApplicationBodyComponentProbeHttpGet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchApplicationBodyComponentProbeHttpGet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchApplicationBodyComponentProbeHttpGet) SetPort(v int32)`

SetPort sets Port field to given value.


### GetHeaders

`func (o *PatchApplicationBodyComponentProbeHttpGet) GetHeaders() []PatchApplicationBodyComponentProbeHttpGetHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PatchApplicationBodyComponentProbeHttpGet) GetHeadersOk() (*[]PatchApplicationBodyComponentProbeHttpGetHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PatchApplicationBodyComponentProbeHttpGet) SetHeaders(v []PatchApplicationBodyComponentProbeHttpGetHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PatchApplicationBodyComponentProbeHttpGet) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


