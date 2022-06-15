# ContractCallTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** |  | 
**ContractCall** | [**ContractCallTransactionMetadataContractCall**](ContractCallTransactionMetadataContractCall.md) |  | 

## Methods

### NewContractCallTransaction

`func NewContractCallTransaction(txType string, contractCall ContractCallTransactionMetadataContractCall, ) *ContractCallTransaction`

NewContractCallTransaction instantiates a new ContractCallTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallTransactionWithDefaults

`func NewContractCallTransactionWithDefaults() *ContractCallTransaction`

NewContractCallTransactionWithDefaults instantiates a new ContractCallTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *ContractCallTransaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *ContractCallTransaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *ContractCallTransaction) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetContractCall

`func (o *ContractCallTransaction) GetContractCall() ContractCallTransactionMetadataContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *ContractCallTransaction) GetContractCallOk() (*ContractCallTransactionMetadataContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *ContractCallTransaction) SetContractCall(v ContractCallTransactionMetadataContractCall)`

SetContractCall sets ContractCall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


