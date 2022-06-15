# SmartContractTransactionMetadataSmartContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractId** | **string** | Contract identifier formatted as &#x60;&lt;principaladdress&gt;.&lt;contract_name&gt;&#x60; | 
**SourceCode** | **string** | Clarity code of the smart contract being deployed | 

## Methods

### NewSmartContractTransactionMetadataSmartContract

`func NewSmartContractTransactionMetadataSmartContract(contractId string, sourceCode string, ) *SmartContractTransactionMetadataSmartContract`

NewSmartContractTransactionMetadataSmartContract instantiates a new SmartContractTransactionMetadataSmartContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractTransactionMetadataSmartContractWithDefaults

`func NewSmartContractTransactionMetadataSmartContractWithDefaults() *SmartContractTransactionMetadataSmartContract`

NewSmartContractTransactionMetadataSmartContractWithDefaults instantiates a new SmartContractTransactionMetadataSmartContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractId

`func (o *SmartContractTransactionMetadataSmartContract) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *SmartContractTransactionMetadataSmartContract) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *SmartContractTransactionMetadataSmartContract) SetContractId(v string)`

SetContractId sets ContractId field to given value.


### GetSourceCode

`func (o *SmartContractTransactionMetadataSmartContract) GetSourceCode() string`

GetSourceCode returns the SourceCode field if non-nil, zero value otherwise.

### GetSourceCodeOk

`func (o *SmartContractTransactionMetadataSmartContract) GetSourceCodeOk() (*string, bool)`

GetSourceCodeOk returns a tuple with the SourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCode

`func (o *SmartContractTransactionMetadataSmartContract) SetSourceCode(v string)`

SetSourceCode sets SourceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


