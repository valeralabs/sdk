# TokenTransferTransactionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** |  | 
**TokenTransfer** | [**TokenTransferTransactionMetadataTokenTransfer**](TokenTransferTransactionMetadataTokenTransfer.md) |  | 

## Methods

### NewTokenTransferTransactionMetadata

`func NewTokenTransferTransactionMetadata(txType string, tokenTransfer TokenTransferTransactionMetadataTokenTransfer, ) *TokenTransferTransactionMetadata`

NewTokenTransferTransactionMetadata instantiates a new TokenTransferTransactionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTransferTransactionMetadataWithDefaults

`func NewTokenTransferTransactionMetadataWithDefaults() *TokenTransferTransactionMetadata`

NewTokenTransferTransactionMetadataWithDefaults instantiates a new TokenTransferTransactionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *TokenTransferTransactionMetadata) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *TokenTransferTransactionMetadata) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *TokenTransferTransactionMetadata) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetTokenTransfer

`func (o *TokenTransferTransactionMetadata) GetTokenTransfer() TokenTransferTransactionMetadataTokenTransfer`

GetTokenTransfer returns the TokenTransfer field if non-nil, zero value otherwise.

### GetTokenTransferOk

`func (o *TokenTransferTransactionMetadata) GetTokenTransferOk() (*TokenTransferTransactionMetadataTokenTransfer, bool)`

GetTokenTransferOk returns a tuple with the TokenTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTransfer

`func (o *TokenTransferTransactionMetadata) SetTokenTransfer(v TokenTransferTransactionMetadataTokenTransfer)`

SetTokenTransfer sets TokenTransfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


