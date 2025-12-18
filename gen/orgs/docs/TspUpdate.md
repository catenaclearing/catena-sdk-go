# TspUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**SourceName** | Pointer to [**NullableTspEnum**](TspEnum.md) |  | [optional] 
**ConnType** | Pointer to [**NullableConnectionTypeEnum**](ConnectionTypeEnum.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTspUpdate

`func NewTspUpdate() *TspUpdate`

NewTspUpdate instantiates a new TspUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTspUpdateWithDefaults

`func NewTspUpdateWithDefaults() *TspUpdate`

NewTspUpdateWithDefaults instantiates a new TspUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TspUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TspUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TspUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TspUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TspUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TspUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSourceName

`func (o *TspUpdate) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TspUpdate) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TspUpdate) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *TspUpdate) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *TspUpdate) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *TspUpdate) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetConnType

`func (o *TspUpdate) GetConnType() ConnectionTypeEnum`

GetConnType returns the ConnType field if non-nil, zero value otherwise.

### GetConnTypeOk

`func (o *TspUpdate) GetConnTypeOk() (*ConnectionTypeEnum, bool)`

GetConnTypeOk returns a tuple with the ConnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnType

`func (o *TspUpdate) SetConnType(v ConnectionTypeEnum)`

SetConnType sets ConnType field to given value.

### HasConnType

`func (o *TspUpdate) HasConnType() bool`

HasConnType returns a boolean if a field has been set.

### SetConnTypeNil

`func (o *TspUpdate) SetConnTypeNil(b bool)`

 SetConnTypeNil sets the value for ConnType to be an explicit nil

### UnsetConnType
`func (o *TspUpdate) UnsetConnType()`

UnsetConnType ensures that no value is present for ConnType, not even an explicit nil
### GetDescription

`func (o *TspUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TspUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TspUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TspUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TspUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TspUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


