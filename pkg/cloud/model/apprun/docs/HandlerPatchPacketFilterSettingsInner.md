# HandlerPatchPacketFilterSettingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromIp** | **string** | 送信元IPv4アドレス | 
**FromIpPrefixLength** | **int32** | IPv4アドレスprefix長 | 

## Methods

### NewHandlerPatchPacketFilterSettingsInner

`func NewHandlerPatchPacketFilterSettingsInner(fromIp string, fromIpPrefixLength int32, ) *HandlerPatchPacketFilterSettingsInner`

NewHandlerPatchPacketFilterSettingsInner instantiates a new HandlerPatchPacketFilterSettingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerPatchPacketFilterSettingsInnerWithDefaults

`func NewHandlerPatchPacketFilterSettingsInnerWithDefaults() *HandlerPatchPacketFilterSettingsInner`

NewHandlerPatchPacketFilterSettingsInnerWithDefaults instantiates a new HandlerPatchPacketFilterSettingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromIp

`func (o *HandlerPatchPacketFilterSettingsInner) GetFromIp() string`

GetFromIp returns the FromIp field if non-nil, zero value otherwise.

### GetFromIpOk

`func (o *HandlerPatchPacketFilterSettingsInner) GetFromIpOk() (*string, bool)`

GetFromIpOk returns a tuple with the FromIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIp

`func (o *HandlerPatchPacketFilterSettingsInner) SetFromIp(v string)`

SetFromIp sets FromIp field to given value.


### GetFromIpPrefixLength

`func (o *HandlerPatchPacketFilterSettingsInner) GetFromIpPrefixLength() int32`

GetFromIpPrefixLength returns the FromIpPrefixLength field if non-nil, zero value otherwise.

### GetFromIpPrefixLengthOk

`func (o *HandlerPatchPacketFilterSettingsInner) GetFromIpPrefixLengthOk() (*int32, bool)`

GetFromIpPrefixLengthOk returns a tuple with the FromIpPrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromIpPrefixLength

`func (o *HandlerPatchPacketFilterSettingsInner) SetFromIpPrefixLength(v int32)`

SetFromIpPrefixLength sets FromIpPrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


