# PostApplicationBodyComponentProbeHttpGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | HTTPサーバーへアクセスしプローブをチェックする際のパス | 
**Port** | **int32** | HTTPサーバーへアクセスしプローブをチェックする際のポート番号 | 
**Headers** | Pointer to [**[]PostApplicationBodyComponentProbeHttpGetHeader**](PostApplicationBodyComponentProbeHttpGetHeader.md) |  | [optional] 

## Methods

### NewPostApplicationBodyComponentProbeHttpGet

`func NewPostApplicationBodyComponentProbeHttpGet(path string, port int32, ) *PostApplicationBodyComponentProbeHttpGet`

NewPostApplicationBodyComponentProbeHttpGet instantiates a new PostApplicationBodyComponentProbeHttpGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostApplicationBodyComponentProbeHttpGetWithDefaults

`func NewPostApplicationBodyComponentProbeHttpGetWithDefaults() *PostApplicationBodyComponentProbeHttpGet`

NewPostApplicationBodyComponentProbeHttpGetWithDefaults instantiates a new PostApplicationBodyComponentProbeHttpGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PostApplicationBodyComponentProbeHttpGet) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PostApplicationBodyComponentProbeHttpGet) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PostApplicationBodyComponentProbeHttpGet) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *PostApplicationBodyComponentProbeHttpGet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostApplicationBodyComponentProbeHttpGet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostApplicationBodyComponentProbeHttpGet) SetPort(v int32)`

SetPort sets Port field to given value.


### GetHeaders

`func (o *PostApplicationBodyComponentProbeHttpGet) GetHeaders() []PostApplicationBodyComponentProbeHttpGetHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PostApplicationBodyComponentProbeHttpGet) GetHeadersOk() (*[]PostApplicationBodyComponentProbeHttpGetHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PostApplicationBodyComponentProbeHttpGet) SetHeaders(v []PostApplicationBodyComponentProbeHttpGetHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PostApplicationBodyComponentProbeHttpGet) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


