# PartnerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of your organization | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 
**Categories** | Pointer to [**[]PartnerCategory**](PartnerCategory.md) | A list of categories for the services offered by your organization | [optional] 
**IsSandbox** | Pointer to **bool** | Indicates whether the account is a sandbox account for testing purposes. | [optional] [default to false]
**Properties** | Pointer to [**[]PartnerPropertyCreate**](PartnerPropertyCreate.md) | Optional custom properties for the partner organization | [optional] [default to []]

## Methods

### NewPartnerCreate

`func NewPartnerCreate(name string, ) *PartnerCreate`

NewPartnerCreate instantiates a new PartnerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerCreateWithDefaults

`func NewPartnerCreateWithDefaults() *PartnerCreate`

NewPartnerCreateWithDefaults instantiates a new PartnerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PartnerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PartnerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PartnerCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PartnerCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PartnerCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PartnerCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PartnerCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PartnerCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PartnerCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWebsites

`func (o *PartnerCreate) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *PartnerCreate) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *PartnerCreate) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *PartnerCreate) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsitesNil

`func (o *PartnerCreate) SetWebsitesNil(b bool)`

 SetWebsitesNil sets the value for Websites to be an explicit nil

### UnsetWebsites
`func (o *PartnerCreate) UnsetWebsites()`

UnsetWebsites ensures that no value is present for Websites, not even an explicit nil
### GetCategories

`func (o *PartnerCreate) GetCategories() []PartnerCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PartnerCreate) GetCategoriesOk() (*[]PartnerCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PartnerCreate) SetCategories(v []PartnerCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PartnerCreate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetIsSandbox

`func (o *PartnerCreate) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *PartnerCreate) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *PartnerCreate) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *PartnerCreate) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.

### GetProperties

`func (o *PartnerCreate) GetProperties() []PartnerPropertyCreate`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PartnerCreate) GetPropertiesOk() (*[]PartnerPropertyCreate, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PartnerCreate) SetProperties(v []PartnerPropertyCreate)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PartnerCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


