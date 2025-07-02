# HandlerListApplicationsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | アプリケーションID | 
**Name** | **string** | アプリケーション名 | 
**Status** | **string** | アプリケーションステータス | 
**PublicUrl** | **string** | 公開URL | 
**CreatedAt** | **time.Time** | 作成日時 | 

## Methods

### NewHandlerListApplicationsData

`func NewHandlerListApplicationsData(id string, name string, status string, publicUrl string, createdAt time.Time, ) *HandlerListApplicationsData`

NewHandlerListApplicationsData instantiates a new HandlerListApplicationsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerListApplicationsDataWithDefaults

`func NewHandlerListApplicationsDataWithDefaults() *HandlerListApplicationsData`

NewHandlerListApplicationsDataWithDefaults instantiates a new HandlerListApplicationsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HandlerListApplicationsData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HandlerListApplicationsData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HandlerListApplicationsData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HandlerListApplicationsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandlerListApplicationsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandlerListApplicationsData) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *HandlerListApplicationsData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerListApplicationsData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerListApplicationsData) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPublicUrl

`func (o *HandlerListApplicationsData) GetPublicUrl() string`

GetPublicUrl returns the PublicUrl field if non-nil, zero value otherwise.

### GetPublicUrlOk

`func (o *HandlerListApplicationsData) GetPublicUrlOk() (*string, bool)`

GetPublicUrlOk returns a tuple with the PublicUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUrl

`func (o *HandlerListApplicationsData) SetPublicUrl(v string)`

SetPublicUrl sets PublicUrl field to given value.


### GetCreatedAt

`func (o *HandlerListApplicationsData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HandlerListApplicationsData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HandlerListApplicationsData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


