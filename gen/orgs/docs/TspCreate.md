# TspCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | TSP display name | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | Unique TSP identifier enum value | 
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) | The status of the TSP | [optional] 
**ConnType** | [**ConnectionTypeEnum**](ConnectionTypeEnum.md) | Authentication method (OAuth, API Key, etc.) | 
**IsSandbox** | Pointer to **bool** | Indicates whether the TSP is a sandbox integrations for testing purposes. | [optional] [default to false]
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**LogoDarkUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTspCreate

`func NewTspCreate(name string, sourceName TspEnum, connType ConnectionTypeEnum, ) *TspCreate`

NewTspCreate instantiates a new TspCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTspCreateWithDefaults

`func NewTspCreateWithDefaults() *TspCreate`

NewTspCreateWithDefaults instantiates a new TspCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TspCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TspCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TspCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TspCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TspCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TspCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TspCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TspCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TspCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsites

`func (o *TspCreate) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *TspCreate) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *TspCreate) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *TspCreate) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsitesNil

`func (o *TspCreate) SetWebsitesNil(b bool)`

 SetWebsitesNil sets the value for Websites to be an explicit nil

### UnsetWebsites
`func (o *TspCreate) UnsetWebsites()`

UnsetWebsites ensures that no value is present for Websites, not even an explicit nil
### GetSourceName

`func (o *TspCreate) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TspCreate) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TspCreate) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetStatus

`func (o *TspCreate) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TspCreate) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TspCreate) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TspCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnType

`func (o *TspCreate) GetConnType() ConnectionTypeEnum`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *TspCreate) GetConnTypeOk() (*ConnectionTypeEnum, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *TspCreate) SetConnType(v ConnectionTypeEnum)`

SetConnType sets ConnType field to given value.


### GetIsSandbox

`func (o *TspCreate) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *TspCreate) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *TspCreate) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *TspCreate) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.

### GetLogoUrl

`func (o *TspCreate) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *TspCreate) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *TspCreate) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *TspCreate) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *TspCreate) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *TspCreate) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetLogoDarkUrl

`func (o *TspCreate) GetLogoDarkUrl() string`

GetLogoDarkUrl returns the LogoDarkUrl field if non-nil, zero value otherwise.

### GetLogoDarkUrlOk

`func (o *TspCreate) GetLogoDarkUrlOk() (*string, bool)`

GetLogoDarkUrlOk returns a tuple with the LogoDarkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDarkUrl

`func (o *TspCreate) SetLogoDarkUrl(v string)`

SetLogoDarkUrl sets LogoDarkUrl field to given value.

### HasLogoDarkUrl

`func (o *TspCreate) HasLogoDarkUrl() bool`

HasLogoDarkUrl returns a boolean if a field has been set.

### SetLogoDarkUrlNil

`func (o *TspCreate) SetLogoDarkUrlNil(b bool)`

 SetLogoDarkUrlNil sets the value for LogoDarkUrl to be an explicit nil

### UnsetLogoDarkUrl
`func (o *TspCreate) UnsetLogoDarkUrl()`

UnsetLogoDarkUrl ensures that no value is present for LogoDarkUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


