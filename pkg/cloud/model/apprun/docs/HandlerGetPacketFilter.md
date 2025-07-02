# HandlerGetPacketFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | 有効フラグ | 
**Settings** | [**[]HandlerGetPacketFilterSettingsInner**](HandlerGetPacketFilterSettingsInner.md) |  | 

## Methods

### NewHandlerGetPacketFilter

`func NewHandlerGetPacketFilter(isEnabled bool, settings []HandlerGetPacketFilterSettingsInner, ) *HandlerGetPacketFilter`

NewHandlerGetPacketFilter instantiates a new HandlerGetPacketFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerGetPacketFilterWithDefaults

`func NewHandlerGetPacketFilterWithDefaults() *HandlerGetPacketFilter`

NewHandlerGetPacketFilterWithDefaults instantiates a new HandlerGetPacketFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *HandlerGetPacketFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *HandlerGetPacketFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *HandlerGetPacketFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetSettings

`func (o *HandlerGetPacketFilter) GetSettings() []HandlerGetPacketFilterSettingsInner`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *HandlerGetPacketFilter) GetSettingsOk() (*[]HandlerGetPacketFilterSettingsInner, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *HandlerGetPacketFilter) SetSettings(v []HandlerGetPacketFilterSettingsInner)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


