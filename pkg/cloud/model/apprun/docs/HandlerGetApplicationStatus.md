# HandlerGetApplicationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | アプリケーションステータス | 
**Message** | **string** | ステータス失敗時のメッセージ | 

## Methods

### NewHandlerGetApplicationStatus

`func NewHandlerGetApplicationStatus(status string, message string, ) *HandlerGetApplicationStatus`

NewHandlerGetApplicationStatus instantiates a new HandlerGetApplicationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerGetApplicationStatusWithDefaults

`func NewHandlerGetApplicationStatusWithDefaults() *HandlerGetApplicationStatus`

NewHandlerGetApplicationStatusWithDefaults instantiates a new HandlerGetApplicationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HandlerGetApplicationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerGetApplicationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerGetApplicationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *HandlerGetApplicationStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HandlerGetApplicationStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HandlerGetApplicationStatus) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


