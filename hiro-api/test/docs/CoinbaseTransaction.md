# CoinbaseTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxType** | **string** |  | 
**CoinbasePayload** | [**CoinbaseTransactionMetadataCoinbasePayload**](CoinbaseTransactionMetadataCoinbasePayload.md) |  | 

## Methods

### NewCoinbaseTransaction

`func NewCoinbaseTransaction(txType string, coinbasePayload CoinbaseTransactionMetadataCoinbasePayload, ) *CoinbaseTransaction`

NewCoinbaseTransaction instantiates a new CoinbaseTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCoinbaseTransactionWithDefaults

`func NewCoinbaseTransactionWithDefaults() *CoinbaseTransaction`

NewCoinbaseTransactionWithDefaults instantiates a new CoinbaseTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxType

`func (o *CoinbaseTransaction) GetTxType() string`

GetTxType returns the TxType field if non-nil, zero value otherwise.

### GetTxTypeOk

`func (o *CoinbaseTransaction) GetTxTypeOk() (*string, bool)`

GetTxTypeOk returns a tuple with the TxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxType

`func (o *CoinbaseTransaction) SetTxType(v string)`

SetTxType sets TxType field to given value.


### GetCoinbasePayload

`func (o *CoinbaseTransaction) GetCoinbasePayload() CoinbaseTransactionMetadataCoinbasePayload`

GetCoinbasePayload returns the CoinbasePayload field if non-nil, zero value otherwise.

### GetCoinbasePayloadOk

`func (o *CoinbaseTransaction) GetCoinbasePayloadOk() (*CoinbaseTransactionMetadataCoinbasePayload, bool)`

GetCoinbasePayloadOk returns a tuple with the CoinbasePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinbasePayload

`func (o *CoinbaseTransaction) SetCoinbasePayload(v CoinbaseTransactionMetadataCoinbasePayload)`

SetCoinbasePayload sets CoinbasePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


