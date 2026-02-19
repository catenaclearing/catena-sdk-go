# TspRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | TSP enum identifier for API references | 
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) | The status of the TSP | [optional] 
**ConnType** | [**ConnectionTypeEnum**](ConnectionTypeEnum.md) | The type of connection usedt to authenticate with the TSP | 
**IsSandbox** | Pointer to **bool** | Indicates whether the TSP is a sandbox integrations for testing purposes. | [optional] [default to false]
**LogoUrl** | Pointer to **NullableString** |  | [optional] 
**LogoDarkUrl** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique Catena TSP identifier | 
**Slug** | **string** | URL-friendly TSP identifier | 
**CreatedAt** | **time.Time** | When the TSP was added to Catena | 
**UpdatedAt** | **time.Time** | Last modification timestamp | 

## Methods

### NewTspRead

`func NewTspRead(name string, sourceName TspEnum, connType ConnectionTypeEnum, id string, slug string, createdAt time.Time, updatedAt time.Time, ) *TspRead`

NewTspRead instantiates a new TspRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTspReadWithDefaults

`func NewTspReadWithDefaults() *TspRead`

NewTspReadWithDefaults instantiates a new TspRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TspRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TspRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TspRead) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TspRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TspRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TspRead) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TspRead) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TspRead) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TspRead) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsites

`func (o *TspRead) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *TspRead) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *TspRead) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *TspRead) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsitesNil

`func (o *TspRead) SetWebsitesNil(b bool)`

 SetWebsitesNil sets the value for Websites to be an explicit nil

### UnsetWebsites
`func (o *TspRead) UnsetWebsites()`

UnsetWebsites ensures that no value is present for Websites, not even an explicit nil
### GetSourceName

`func (o *TspRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TspRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TspRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetStatus

`func (o *TspRead) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TspRead) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TspRead) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TspRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnType

`func (o *TspRead) GetConnType() ConnectionTypeEnum`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *TspRead) GetConnTypeOk() (*ConnectionTypeEnum, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *TspRead) SetConnType(v ConnectionTypeEnum)`

SetConnType sets ConnType field to given value.


### GetIsSandbox

`func (o *TspRead) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *TspRead) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *TspRead) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *TspRead) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.

### GetLogoUrl

`func (o *TspRead) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *TspRead) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *TspRead) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *TspRead) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *TspRead) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *TspRead) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetLogoDarkUrl

`func (o *TspRead) GetLogoDarkUrl() string`

GetLogoDarkUrl returns the LogoDarkUrl field if non-nil, zero value otherwise.

### GetLogoDarkUrlOk

`func (o *TspRead) GetLogoDarkUrlOk() (*string, bool)`

GetLogoDarkUrlOk returns a tuple with the LogoDarkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDarkUrl

`func (o *TspRead) SetLogoDarkUrl(v string)`

SetLogoDarkUrl sets LogoDarkUrl field to given value.

### HasLogoDarkUrl

`func (o *TspRead) HasLogoDarkUrl() bool`

HasLogoDarkUrl returns a boolean if a field has been set.

### SetLogoDarkUrlNil

`func (o *TspRead) SetLogoDarkUrlNil(b bool)`

 SetLogoDarkUrlNil sets the value for LogoDarkUrl to be an explicit nil

### UnsetLogoDarkUrl
`func (o *TspRead) UnsetLogoDarkUrl()`

UnsetLogoDarkUrl ensures that no value is present for LogoDarkUrl, not even an explicit nil
### GetId

`func (o *TspRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TspRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TspRead) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *TspRead) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TspRead) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TspRead) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetCreatedAt

`func (o *TspRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TspRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TspRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TspRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TspRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TspRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


