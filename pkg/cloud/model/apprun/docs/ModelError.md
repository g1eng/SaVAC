# ModelError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **NullableString** |  | 
**Reason** | **NullableString** |  | 
**Message** | **NullableString** |  | 
**LocationType** | **NullableString** |  | 
**Location** | **NullableString** |  | 

## Methods

### NewModelError

`func NewModelError(domain NullableString, reason NullableString, message NullableString, locationType NullableString, location NullableString, ) *ModelError`

NewModelError instantiates a new ModelError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelErrorWithDefaults

`func NewModelErrorWithDefaults() *ModelError`

NewModelErrorWithDefaults instantiates a new ModelError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ModelError) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ModelError) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ModelError) SetDomain(v string)`

SetDomain sets Domain field to given value.


### SetDomainNil

`func (o *ModelError) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *ModelError) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetReason

`func (o *ModelError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelError) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *ModelError) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ModelError) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetMessage

`func (o *ModelError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelError) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ModelError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ModelError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetLocationType

`func (o *ModelError) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *ModelError) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *ModelError) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.


### SetLocationTypeNil

`func (o *ModelError) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *ModelError) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetLocation

`func (o *ModelError) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ModelError) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ModelError) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *ModelError) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ModelError) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


