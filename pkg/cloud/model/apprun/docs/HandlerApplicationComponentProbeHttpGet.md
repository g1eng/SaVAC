# HandlerApplicationComponentProbeHttpGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | HTTPサーバーへアクセスしプローブをチェックする際のパス | 
**Port** | **int32** | HTTPサーバーへアクセスしプローブをチェックする際のポート番号 | 
**Headers** | Pointer to [**[]HandlerApplicationComponentProbeHttpGetHeader**](HandlerApplicationComponentProbeHttpGetHeader.md) |  | [optional] 

## Methods

### NewHandlerApplicationComponentProbeHttpGet

`func NewHandlerApplicationComponentProbeHttpGet(path string, port int32, ) *HandlerApplicationComponentProbeHttpGet`

NewHandlerApplicationComponentProbeHttpGet instantiates a new HandlerApplicationComponentProbeHttpGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerApplicationComponentProbeHttpGetWithDefaults

`func NewHandlerApplicationComponentProbeHttpGetWithDefaults() *HandlerApplicationComponentProbeHttpGet`

NewHandlerApplicationComponentProbeHttpGetWithDefaults instantiates a new HandlerApplicationComponentProbeHttpGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *HandlerApplicationComponentProbeHttpGet) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HandlerApplicationComponentProbeHttpGet) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HandlerApplicationComponentProbeHttpGet) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *HandlerApplicationComponentProbeHttpGet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HandlerApplicationComponentProbeHttpGet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HandlerApplicationComponentProbeHttpGet) SetPort(v int32)`

SetPort sets Port field to given value.


### GetHeaders

`func (o *HandlerApplicationComponentProbeHttpGet) GetHeaders() []HandlerApplicationComponentProbeHttpGetHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HandlerApplicationComponentProbeHttpGet) GetHeadersOk() (*[]HandlerApplicationComponentProbeHttpGetHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HandlerApplicationComponentProbeHttpGet) SetHeaders(v []HandlerApplicationComponentProbeHttpGetHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HandlerApplicationComponentProbeHttpGet) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


