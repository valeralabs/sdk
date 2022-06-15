# RosettaBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockIdentifier** | [**Object**](Object.md) |  | 
**ParentBlockIdentifier** | [**RosettaParentBlockIdentifier**](RosettaParentBlockIdentifier.md) |  | 
**Timestamp** | **int32** | The timestamp of the block in milliseconds since the Unix Epoch. The timestamp is stored in milliseconds because some blockchains produce blocks more often than once a second. | 
**Transactions** | [**[]Object**](Object.md) | All the transactions in the block | 
**Metadata** | Pointer to [**RosettaBlockMetadata**](RosettaBlockMetadata.md) |  | [optional] 

## Methods

### NewRosettaBlock

`func NewRosettaBlock(blockIdentifier Object, parentBlockIdentifier RosettaParentBlockIdentifier, timestamp int32, transactions []Object, ) *RosettaBlock`

NewRosettaBlock instantiates a new RosettaBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaBlockWithDefaults

`func NewRosettaBlockWithDefaults() *RosettaBlock`

NewRosettaBlockWithDefaults instantiates a new RosettaBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockIdentifier

`func (o *RosettaBlock) GetBlockIdentifier() Object`

GetBlockIdentifier returns the BlockIdentifier field if non-nil, zero value otherwise.

### GetBlockIdentifierOk

`func (o *RosettaBlock) GetBlockIdentifierOk() (*Object, bool)`

GetBlockIdentifierOk returns a tuple with the BlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIdentifier

`func (o *RosettaBlock) SetBlockIdentifier(v Object)`

SetBlockIdentifier sets BlockIdentifier field to given value.


### GetParentBlockIdentifier

`func (o *RosettaBlock) GetParentBlockIdentifier() RosettaParentBlockIdentifier`

GetParentBlockIdentifier returns the ParentBlockIdentifier field if non-nil, zero value otherwise.

### GetParentBlockIdentifierOk

`func (o *RosettaBlock) GetParentBlockIdentifierOk() (*RosettaParentBlockIdentifier, bool)`

GetParentBlockIdentifierOk returns a tuple with the ParentBlockIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlockIdentifier

`func (o *RosettaBlock) SetParentBlockIdentifier(v RosettaParentBlockIdentifier)`

SetParentBlockIdentifier sets ParentBlockIdentifier field to given value.


### GetTimestamp

`func (o *RosettaBlock) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RosettaBlock) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RosettaBlock) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetTransactions

`func (o *RosettaBlock) GetTransactions() []Object`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *RosettaBlock) GetTransactionsOk() (*[]Object, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *RosettaBlock) SetTransactions(v []Object)`

SetTransactions sets Transactions field to given value.


### GetMetadata

`func (o *RosettaBlock) GetMetadata() RosettaBlockMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaBlock) GetMetadataOk() (*RosettaBlockMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaBlock) SetMetadata(v RosettaBlockMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaBlock) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


