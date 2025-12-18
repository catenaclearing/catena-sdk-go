# RefHosRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionCode** | **string** | Canonical region code (ISO-3166-2), e.g., &#39;US-CA&#39;, &#39;CA-ON&#39;. Primary key. | 
**CountryCode** | **string** | Country code (ISO-3166-1 alpha-2), e.g., &#39;US&#39;, &#39;CA&#39;. | 
**CountryName** | **string** | Country name (e.g., &#39;United States&#39;, &#39;Canada&#39;). | 
**RegionCodeIso2** | **string** | Subdivision code component of ISO-3166-2 (the part after the hyphen), e.g., &#39;CA&#39; for US-CA, &#39;ON&#39; for CA-ON. | 
**RegionName** | **string** | Region/subdivision name, e.g., &#39;California&#39;, &#39;Ontario&#39;. | 

## Methods

### NewRefHosRegion

`func NewRefHosRegion(regionCode string, countryCode string, countryName string, regionCodeIso2 string, regionName string, ) *RefHosRegion`

NewRefHosRegion instantiates a new RefHosRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosRegionWithDefaults

`func NewRefHosRegionWithDefaults() *RefHosRegion`

NewRefHosRegionWithDefaults instantiates a new RefHosRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionCode

`func (o *RefHosRegion) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *RefHosRegion) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *RefHosRegion) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetCountryCode

`func (o *RefHosRegion) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *RefHosRegion) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *RefHosRegion) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.


### GetCountryName

`func (o *RefHosRegion) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *RefHosRegion) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *RefHosRegion) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetRegionCodeIso2

`func (o *RefHosRegion) GetRegionCodeIso2() string`

GetRegionCodeIso2 returns the RegionCodeIso2 field if non-nil, zero value otherwise.

### GetRegionCodeIso2Ok

`func (o *RefHosRegion) GetRegionCodeIso2Ok() (*string, bool)`

GetRegionCodeIso2Ok returns a tuple with the RegionCodeIso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCodeIso2

`func (o *RefHosRegion) SetRegionCodeIso2(v string)`

SetRegionCodeIso2 sets RegionCodeIso2 field to given value.


### GetRegionName

`func (o *RefHosRegion) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *RefHosRegion) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *RefHosRegion) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


