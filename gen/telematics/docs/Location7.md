# Location7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bbox** | Pointer to [**NullableBbox**](Bbox.md) |  | [optional] 
**Type** | **string** |  | 
**Coordinates** | [**Coordinates**](Coordinates.md) |  | 

## Methods

### NewLocation7

`func NewLocation7(type_ string, coordinates Coordinates, ) *Location7`

NewLocation7 instantiates a new Location7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocation7WithDefaults

`func NewLocation7WithDefaults() *Location7`

NewLocation7WithDefaults instantiates a new Location7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBbox

`func (o *Location7) GetBbox() Bbox`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *Location7) GetBboxOk() (*Bbox, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *Location7) SetBbox(v Bbox)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *Location7) HasBbox() bool`

HasBbox returns a boolean if a field has been set.

### SetBboxNil

`func (o *Location7) SetBboxNil(b bool)`

 SetBboxNil sets the value for Bbox to be an explicit nil

### UnsetBbox
`func (o *Location7) UnsetBbox()`

UnsetBbox ensures that no value is present for Bbox, not even an explicit nil
### GetType

`func (o *Location7) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Location7) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Location7) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *Location7) GetCoordinates() Coordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *Location7) GetCoordinatesOk() (*Coordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *Location7) SetCoordinates(v Coordinates)`

SetCoordinates sets Coordinates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


