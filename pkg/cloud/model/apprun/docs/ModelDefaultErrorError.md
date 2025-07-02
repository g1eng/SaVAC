# ModelDefaultErrorError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** |  | 
**Message** | **string** |  | 
**Errors** | [**[]ModelError**](ModelError.md) |  | 

## Methods

### NewModelDefaultErrorError

`func NewModelDefaultErrorError(code int32, message string, errors []ModelError, ) *ModelDefaultErrorError`

NewModelDefaultErrorError instantiates a new ModelDefaultErrorError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDefaultErrorErrorWithDefaults

`func NewModelDefaultErrorErrorWithDefaults() *ModelDefaultErrorError`

NewModelDefaultErrorErrorWithDefaults instantiates a new ModelDefaultErrorError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModelDefaultErrorError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModelDefaultErrorError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModelDefaultErrorError) SetCode(v int32)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ModelDefaultErrorError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelDefaultErrorError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelDefaultErrorError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrors

`func (o *ModelDefaultErrorError) GetErrors() []ModelError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModelDefaultErrorError) GetErrorsOk() (*[]ModelError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModelDefaultErrorError) SetErrors(v []ModelError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


