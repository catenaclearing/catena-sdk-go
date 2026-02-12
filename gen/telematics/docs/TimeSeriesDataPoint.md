# TimeSeriesDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Calendar date for this data point (YYYY-MM-DD). | 
**DailyCount** | **int32** | Number of new resources created on this specific date. | 
**CumulativeCount** | **int32** | Running total of resources up to and including this date. | 

## Methods

### NewTimeSeriesDataPoint

`func NewTimeSeriesDataPoint(date string, dailyCount int32, cumulativeCount int32, ) *TimeSeriesDataPoint`

NewTimeSeriesDataPoint instantiates a new TimeSeriesDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesDataPointWithDefaults

`func NewTimeSeriesDataPointWithDefaults() *TimeSeriesDataPoint`

NewTimeSeriesDataPointWithDefaults instantiates a new TimeSeriesDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TimeSeriesDataPoint) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimeSeriesDataPoint) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimeSeriesDataPoint) SetDate(v string)`

SetDate sets Date field to given value.


### GetDailyCount

`func (o *TimeSeriesDataPoint) GetDailyCount() int32`

GetDailyCount returns the DailyCount field if non-nil, zero value otherwise.

### GetDailyCountOk

`func (o *TimeSeriesDataPoint) GetDailyCountOk() (*int32, bool)`

GetDailyCountOk returns a tuple with the DailyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCount

`func (o *TimeSeriesDataPoint) SetDailyCount(v int32)`

SetDailyCount sets DailyCount field to given value.


### GetCumulativeCount

`func (o *TimeSeriesDataPoint) GetCumulativeCount() int32`

GetCumulativeCount returns the CumulativeCount field if non-nil, zero value otherwise.

### GetCumulativeCountOk

`func (o *TimeSeriesDataPoint) GetCumulativeCountOk() (*int32, bool)`

GetCumulativeCountOk returns a tuple with the CumulativeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeCount

`func (o *TimeSeriesDataPoint) SetCumulativeCount(v int32)`

SetCumulativeCount sets CumulativeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


