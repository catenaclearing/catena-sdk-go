# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bbox** | Pointer to [**NullableBbox**](Bbox.md) |  | [optional] 
**Type** | **string** |  | 
**Coordinates** | [**Coordinates**](Coordinates.md) |  | 

## Methods

### NewLocation

`func NewLocation(type_ string, coordinates Coordinates, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBbox

`func (o *Location) GetBbox() Bbox`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *Location) GetBboxOk() (*Bbox, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *Location) SetBbox(v Bbox)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *Location) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *Location) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *Location) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil
### GetType

`func (o *Location) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Location) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Location) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *Location) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Location) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Location) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


