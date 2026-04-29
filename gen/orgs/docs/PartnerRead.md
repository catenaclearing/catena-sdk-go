# PartnerRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Partner organization name | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to [**[]PartnerCategory**](PartnerCategory.md) | A list of categories for the services offered by your organization | [optional] 
**IsSandbox** | Pointer to **bool** | Indicates whether the account is a sandbox account for testing purposes. | [optional] [default to false]
**ParentPartnerId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique Catena partner identifier | 
**Slug** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | When the partner was onboarded to Catena | 
**UpdatedAt** | **time.Time** | Last modification timestamp | 

## Methods

### NewPartnerRead

`func NewPartnerRead(name string, id string, createdAt time.Time, updatedAt time.Time, ) *PartnerRead`

NewPartnerRead instantiates a new PartnerRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerReadWithDefaults

`func NewPartnerReadWithDefaults() *PartnerRead`

NewPartnerReadWithDefaults instantiates a new PartnerRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PartnerRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerRead) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PartnerRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartnerRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartnerRead) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartnerRead) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PartnerRead) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PartnerRead) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsites

`func (o *PartnerRead) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *PartnerRead) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *PartnerRead) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *PartnerRead) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsitesNil

`func (o *PartnerRead) SetWebsitesNil(b bool)`

 SetWebsitesNil sets the value for Websites to be an explicit nil

### UnsetWebsites
`func (o *PartnerRead) UnsetWebsites()`

UnsetWebsites ensures that no value is present for Websites, not even an explicit nil
### GetCategories

`func (o *PartnerRead) GetCategories() []PartnerCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PartnerRead) GetCategoriesOk() (*[]PartnerCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PartnerRead) SetCategories(v []PartnerCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PartnerRead) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetIsSandbox

`func (o *PartnerRead) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *PartnerRead) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *PartnerRead) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *PartnerRead) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.

### GetParentPartnerId

`func (o *PartnerRead) GetParentPartnerId() string`

GetParentPartnerId returns the ParentPartnerId field if non-nil, zero value otherwise.

### GetParentPartnerIdOk

`func (o *PartnerRead) GetParentPartnerIdOk() (*string, bool)`

GetParentPartnerIdOk returns a tuple with the ParentPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPartnerId

`func (o *PartnerRead) SetParentPartnerId(v string)`

SetParentPartnerId sets ParentPartnerId field to given value.

### HasParentPartnerId

`func (o *PartnerRead) HasParentPartnerId() bool`

HasParentPartnerId returns a boolean if a field has been set.

### SetParentPartnerIdNil

`func (o *PartnerRead) SetParentPartnerIdNil(b bool)`

 SetParentPartnerIdNil sets the value for ParentPartnerId to be an explicit nil

### UnsetParentPartnerId
`func (o *PartnerRead) UnsetParentPartnerId()`

UnsetParentPartnerId ensures that no value is present for ParentPartnerId, not even an explicit nil
### GetId

`func (o *PartnerRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerRead) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *PartnerRead) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PartnerRead) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PartnerRead) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PartnerRead) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *PartnerRead) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *PartnerRead) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetCreatedAt

`func (o *PartnerRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartnerRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartnerRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PartnerRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PartnerRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PartnerRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


