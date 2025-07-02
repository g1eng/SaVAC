# Traffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionName** | **string** | バージョン名 | 
**IsLatestVersion** | **bool** | 最新バージョンかどうか | 
**Percent** | **int32** | トラフィック分散の割合 | 

## Methods

### NewTraffic

`func NewTraffic(versionName string, isLatestVersion bool, percent int32, ) *Traffic`

NewTraffic instantiates a new Traffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficWithDefaults

`func NewTrafficWithDefaults() *Traffic`

NewTrafficWithDefaults instantiates a new Traffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionName

`func (o *Traffic) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *Traffic) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *Traffic) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.


### GetIsLatestVersion

`func (o *Traffic) GetIsLatestVersion() bool`

GetIsLatestVersion returns the IsLatestVersion field if non-nil, zero value otherwise.

### GetIsLatestVersionOk

`func (o *Traffic) GetIsLatestVersionOk() (*bool, bool)`

GetIsLatestVersionOk returns a tuple with the IsLatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLatestVersion

`func (o *Traffic) SetIsLatestVersion(v bool)`

SetIsLatestVersion sets IsLatestVersion field to given value.


### GetPercent

`func (o *Traffic) GetPercent() int32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *Traffic) GetPercentOk() (*int32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *Traffic) SetPercent(v int32)`

SetPercent sets Percent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


