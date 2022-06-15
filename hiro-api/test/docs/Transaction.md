# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** |  | 
**TokenTransfer** | [**TokenTransferTransactionMetadataTokenTransfer**](TokenTransferTransactionMetadataTokenTransfer.md) |  | 
**ContractCall** | [**ContractCallTransactionMetadataContractCall**](ContractCallTransactionMetadataContractCall.md) |  | 
**PoisonMicroblock** | [**PoisonMicroblockTransactionMetadataPoisonMicroblock**](PoisonMicroblockTransactionMetadataPoisonMicroblock.md) |  | 
**CoinbasePayload** | [**CoinbaseTransactionMetadataCoinbasePayload**](CoinbaseTransactionMetadataCoinbasePayload.md) |  | 

## Methods

### NewTransaction

`func NewTransaction(txType string, tokenTransfer TokenTransferTransactionMetadataTokenTransfer, contractCall ContractCallTransactionMetadataContractCall, poisonMicroblock PoisonMicroblockTransactionMetadataPoisonMicroblock, coinbasePayload CoinbaseTransactionMetadataCoinbasePayload, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *Transaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *Transaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *Transaction) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *Transaction) GetTokenTransfer() TokenTransferTransactionMetadataTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *Transaction) GetTokenTransferOk() (*TokenTransferTransactionMetadataTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *Transaction) SetTokenTransfer(v TokenTransferTransactionMetadataTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.


### GetContractCall

`func (o *Transaction) GetContractCall() ContractCallTransactionMetadataContractCall`

GetContractCall returns the ContractCall field if non-nil, zero value otherwise.

### GetContractCallOk

`func (o *Transaction) GetContractCallOk() (*ContractCallTransactionMetadataContractCall, bool)`

GetContractCallOk returns a tuple with the ContractCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractCall

`func (o *Transaction) SetContractCall(v ContractCallTransactionMetadataContractCall)`

SetContractCall sets ContractCall field to given value.


### GetPoisonMicroblock

`func (o *Transaction) GetPoisonMicroblock() PoisonMicroblockTransactionMetadataPoisonMicroblock`

GetPoisonMicroblock returns the PoisonMicroblock field if non-nil, zero value otherwise.

### GetPoisonMicroblockOk

`func (o *Transaction) GetPoisonMicroblockOk() (*PoisonMicroblockTransactionMetadataPoisonMicroblock, bool)`

GetPoisonMicroblockOk returns a tuple with the PoisonMicroblock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoisonMicroblock

`func (o *Transaction) SetPoisonMicroblock(v PoisonMicroblockTransactionMetadataPoisonMicroblock)`

SetPoisonMicroblock sets PoisonMicroblock field to given value.


### GetCoinbasePayload

`func (o *Transaction) GetCoinbasePayload() CoinbaseTransactionMetadataCoinbasePayload`

GetCoinbasePayload returns the CoinbasePayload field if non-nil, zero value otherwise.

### GetCoinbasePayloadOk

`func (o *Transaction) GetCoinbasePayloadOk() (*CoinbaseTransactionMetadataCoinbasePayload, bool)`

GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbasePayload

`func (o *Transaction) SetCoinbasePayload(v CoinbaseTransactionMetadataCoinbasePayload)`

SetCoinbasePayload sets CoinbasePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


