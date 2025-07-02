# HandlerApplicationComponentDeploySourceContainerRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | コンテナイメージ名 | 
**Server** | Pointer to **string** | コンテナレジストリのサーバー名 | [optional] 
**Username** | Pointer to **string** | コンテナレジストリの認証情報 | [optional] 

## Methods

### NewHandlerApplicationComponentDeploySourceContainerRegistry

`func NewHandlerApplicationComponentDeploySourceContainerRegistry(image string, ) *HandlerApplicationComponentDeploySourceContainerRegistry`

NewHandlerApplicationComponentDeploySourceContainerRegistry instantiates a new HandlerApplicationComponentDeploySourceContainerRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerApplicationComponentDeploySourceContainerRegistryWithDefaults

`func NewHandlerApplicationComponentDeploySourceContainerRegistryWithDefaults() *HandlerApplicationComponentDeploySourceContainerRegistry`

NewHandlerApplicationComponentDeploySourceContainerRegistryWithDefaults instantiates a new HandlerApplicationComponentDeploySourceContainerRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) SetImage(v string)`

SetImage sets Image field to given value.


### GetServer

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetUsername

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HandlerApplicationComponentDeploySourceContainerRegistry) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


