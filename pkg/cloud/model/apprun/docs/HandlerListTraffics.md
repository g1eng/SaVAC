# HandlerListTraffics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | [**[]Traffic**](Traffic.md) |  | 

## Methods

### NewHandlerListTraffics

`func NewHandlerListTraffics(data []Traffic, ) *HandlerListTraffics`

NewHandlerListTraffics instantiates a new HandlerListTraffics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerListTrafficsWithDefaults

`func NewHandlerListTrafficsWithDefaults() *HandlerListTraffics`

NewHandlerListTrafficsWithDefaults instantiates a new HandlerListTraffics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *HandlerListTraffics) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HandlerListTraffics) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HandlerListTraffics) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HandlerListTraffics) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *HandlerListTraffics) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *HandlerListTraffics) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetData

`func (o *HandlerListTraffics) GetData() []Traffic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HandlerListTraffics) GetDataOk() (*[]Traffic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HandlerListTraffics) SetData(v []Traffic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


