# HandlerPatchPacketFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | 有効フラグ | 
**Settings** | [**[]HandlerPatchPacketFilterSettingsInner**](HandlerPatchPacketFilterSettingsInner.md) |  | 

## Methods

### NewHandlerPatchPacketFilter

`func NewHandlerPatchPacketFilter(isEnabled bool, settings []HandlerPatchPacketFilterSettingsInner, ) *HandlerPatchPacketFilter`

NewHandlerPatchPacketFilter instantiates a new HandlerPatchPacketFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerPatchPacketFilterWithDefaults

`func NewHandlerPatchPacketFilterWithDefaults() *HandlerPatchPacketFilter`

NewHandlerPatchPacketFilterWithDefaults instantiates a new HandlerPatchPacketFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *HandlerPatchPacketFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HandlerPatchPacketFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HandlerPatchPacketFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetSettings

`func (o *HandlerPatchPacketFilter) GetSettings() []HandlerPatchPacketFilterSettingsInner`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *HandlerPatchPacketFilter) GetSettingsOk() (*[]HandlerPatchPacketFilterSettingsInner, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *HandlerPatchPacketFilter) SetSettings(v []HandlerPatchPacketFilterSettingsInner)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


