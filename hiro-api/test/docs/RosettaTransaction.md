# RosettaTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionIdentifier** | [**Object**](Object.md) |  | 
**Operations** | [**[]Object**](Object.md) | List of operations | 
**Metadata** | Pointer to [**RosettaTransactionMetadata**](RosettaTransactionMetadata.md) |  | [optional] 

## Methods

### NewRosettaTransaction

`func NewRosettaTransaction(transactionIdentifier Object, operations []Object, ) *RosettaTransaction`

NewRosettaTransaction instantiates a new RosettaTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaTransactionWithDefaults

`func NewRosettaTransactionWithDefaults() *RosettaTransaction`

NewRosettaTransactionWithDefaults instantiates a new RosettaTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionIdentifier

`func (o *RosettaTransaction) GetTransactionIdentifier() Object`

GetTransactionIdentifier returns the TransactionIdentifier field if non-nil, zero value otherwise.

### GetTransactionIdentifierOk

`func (o *RosettaTransaction) GetTransactionIdentifierOk() (*Object, bool)`

GetTransactionIdentifierOk returns a tuple with the TransactionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIdentifier

`func (o *RosettaTransaction) SetTransactionIdentifier(v Object)`

SetTransactionIdentifier sets TransactionIdentifier field to given value.


### GetOperations

`func (o *RosettaTransaction) GetOperations() []Object`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *RosettaTransaction) GetOperationsOk() (*[]Object, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *RosettaTransaction) SetOperations(v []Object)`

SetOperations sets Operations field to given value.


### GetMetadata

`func (o *RosettaTransaction) GetMetadata() RosettaTransactionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RosettaTransaction) GetMetadataOk() (*RosettaTransactionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RosettaTransaction) SetMetadata(v RosettaTransactionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RosettaTransaction) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


