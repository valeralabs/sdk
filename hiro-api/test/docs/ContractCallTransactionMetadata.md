# ContractCallTransactionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** |  | 
**ContractCall** | [**ContractCallTransactionMetadataContractCall**](ContractCallTransactionMetadataContractCall.md) |  | 

## Methods

### NewContractCallTransactionMetadata

`func NewContractCallTransactionMetadata(txType string, contractCall ContractCallTransactionMetadataContractCall, ) *ContractCallTransactionMetadata`

NewContractCallTransactionMetadata instantiates a new ContractCallTransactionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallTransactionMetadataWithDefaults

`func NewContractCallTransactionMetadataWithDefaults() *ContractCallTransactionMetadata`

NewContractCallTransactionMetadataWithDefaults instantiates a new ContractCallTransactionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *ContractCallTransactionMetadata) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *ContractCallTransactionMetadata) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *ContractCallTransactionMetadata) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetContractCall

`func (o *ContractCallTransactionMetadata) GetContractCall() ContractCallTransactionMetadataContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *ContractCallTransactionMetadata) GetContractCallOk() (*ContractCallTransactionMetadataContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *ContractCallTransactionMetadata) SetContractCall(v ContractCallTransactionMetadataContractCall)`

SetContractCall sets ContractCall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


