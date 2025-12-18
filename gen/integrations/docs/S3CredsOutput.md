# S3CredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **string** |  | 
**SecretKey** | **interface{}** |  | 
**BucketName** | **string** |  | 
**Region** | **string** |  | 

## Methods

### NewS3CredsOutput

`func NewS3CredsOutput(accessKey string, secretKey interface{}, bucketName string, region string, ) *S3CredsOutput`

NewS3CredsOutput instantiates a new S3CredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CredsOutputWithDefaults

`func NewS3CredsOutputWithDefaults() *S3CredsOutput`

NewS3CredsOutputWithDefaults instantiates a new S3CredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *S3CredsOutput) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *S3CredsOutput) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *S3CredsOutput) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *S3CredsOutput) GetSecretKey() interface{}`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *S3CredsOutput) GetSecretKeyOk() (*interface{}, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *S3CredsOutput) SetSecretKey(v interface{})`

SetSecretKey sets SecretKey field to given value.


### SetSecretKeyNil

`func (o *S3CredsOutput) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *S3CredsOutput) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetBucketName

`func (o *S3CredsOutput) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3CredsOutput) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3CredsOutput) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRegion

`func (o *S3CredsOutput) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3CredsOutput) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3CredsOutput) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


