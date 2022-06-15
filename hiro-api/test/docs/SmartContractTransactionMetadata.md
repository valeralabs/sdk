# SmartContractTransactionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** |  | 
**SmartContract** | [**SmartContractTransactionMetadataSmartContract**](SmartContractTransactionMetadataSmartContract.md) |  | 

## Methods

### NewSmartContractTransactionMetadata

`func NewSmartContractTransactionMetadata(txType string, smartContract SmartContractTransactionMetadataSmartContract, ) *SmartContractTransactionMetadata`

NewSmartContractTransactionMetadata instantiates a new SmartContractTransactionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractTransactionMetadataWithDefaults

`func NewSmartContractTransactionMetadataWithDefaults() *SmartContractTransactionMetadata`

NewSmartContractTransactionMetadataWithDefaults instantiates a new SmartContractTransactionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *SmartContractTransactionMetadata) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *SmartContractTransactionMetadata) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *SmartContractTransactionMetadata) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetSmartContract

`func (o *SmartContractTransactionMetadata) GetSmartContract() SmartContractTransactionMetadataSmartContract`

GetSmartContract returns the SmartContract field if non-nil, zero value otherwise.

### GetSmartContractOk

`func (o *SmartContractTransactionMetadata) GetSmartContractOk() (*SmartContractTransactionMetadataSmartContract, bool)`

GetSmartContractOk returns a tuple with the SmartContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartContract

`func (o *SmartContractTransactionMetadata) SetSmartContract(v SmartContractTransactionMetadataSmartContract)`

SetSmartContract sets SmartContract field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


