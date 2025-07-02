# PatchPacketFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | Pointer to **bool** | 有効フラグ | [optional] 
**Settings** | Pointer to [**[]HandlerPatchPacketFilterSettingsInner**](HandlerPatchPacketFilterSettingsInner.md) |  | [optional] 

## Methods

### NewPatchPacketFilter

`func NewPatchPacketFilter() *PatchPacketFilter`

NewPatchPacketFilter instantiates a new PatchPacketFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPacketFilterWithDefaults

`func NewPatchPacketFilterWithDefaults() *PatchPacketFilter`

NewPatchPacketFilterWithDefaults instantiates a new PatchPacketFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *PatchPacketFilter) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PatchPacketFilter) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PatchPacketFilter) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *PatchPacketFilter) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetSettings

`func (o *PatchPacketFilter) GetSettings() []HandlerPatchPacketFilterSettingsInner`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PatchPacketFilter) GetSettingsOk() (*[]HandlerPatchPacketFilterSettingsInner, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PatchPacketFilter) SetSettings(v []HandlerPatchPacketFilterSettingsInner)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PatchPacketFilter) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


