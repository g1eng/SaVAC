# HandlerListVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**HandlerListVersionsMeta**](HandlerListVersionsMeta.md) |  | 
**Data** | [**[]Version**](Version.md) |  | 

## Methods

### NewHandlerListVersions

`func NewHandlerListVersions(meta HandlerListVersionsMeta, data []Version, ) *HandlerListVersions`

NewHandlerListVersions instantiates a new HandlerListVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerListVersionsWithDefaults

`func NewHandlerListVersionsWithDefaults() *HandlerListVersions`

NewHandlerListVersionsWithDefaults instantiates a new HandlerListVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *HandlerListVersions) GetMeta() HandlerListVersionsMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HandlerListVersions) GetMetaOk() (*HandlerListVersionsMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HandlerListVersions) SetMeta(v HandlerListVersionsMeta)`

SetMeta sets Meta field to given value.


### GetData

`func (o *HandlerListVersions) GetData() []Version`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HandlerListVersions) GetDataOk() (*[]Version, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HandlerListVersions) SetData(v []Version)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


