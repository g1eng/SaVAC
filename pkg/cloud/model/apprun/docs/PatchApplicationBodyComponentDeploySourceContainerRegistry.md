# PatchApplicationBodyComponentDeploySourceContainerRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | コンテナイメージ名 | 
**Server** | Pointer to **NullableString** | コンテナレジストリのサーバー名 | [optional] 
**Username** | Pointer to **NullableString** | コンテナレジストリの認証情報 | [optional] 
**Password** | Pointer to **NullableString** | コンテナレジストリの認証情報 | [optional] 

## Methods

### NewPatchApplicationBodyComponentDeploySourceContainerRegistry

`func NewPatchApplicationBodyComponentDeploySourceContainerRegistry(image string, ) *PatchApplicationBodyComponentDeploySourceContainerRegistry`

NewPatchApplicationBodyComponentDeploySourceContainerRegistry instantiates a new PatchApplicationBodyComponentDeploySourceContainerRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchApplicationBodyComponentDeploySourceContainerRegistryWithDefaults

`func NewPatchApplicationBodyComponentDeploySourceContainerRegistryWithDefaults() *PatchApplicationBodyComponentDeploySourceContainerRegistry`

NewPatchApplicationBodyComponentDeploySourceContainerRegistryWithDefaults instantiates a new PatchApplicationBodyComponentDeploySourceContainerRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImage

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetImage(v string)`

SetImage sets Image field to given value.


### GetServer

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetUsername

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchApplicationBodyComponentDeploySourceContainerRegistry) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


