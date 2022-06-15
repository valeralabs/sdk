# MempoolTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxStatus** | **string** | Status of the transaction | 
**ReceiptTime** | **float32** | A unix timestamp (in seconds) indicating when the transaction broadcast was received by the node. | 
**ReceiptTimeIso** | **string** | An ISO 8601 (YYYY-MM-DDTHH:mm:ss.sssZ) timestamp indicating when the transaction broadcast was received by the node. | 

## Methods

### NewMempoolTransaction

`func NewMempoolTransaction(txStatus string, receiptTime float32, receiptTimeIso string, ) *MempoolTransaction`

NewMempoolTransaction instantiates a new MempoolTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMempoolTransactionWithDefaults

`func NewMempoolTransactionWithDefaults() *MempoolTransaction`

NewMempoolTransactionWithDefaults instantiates a new MempoolTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxStatus

`func (o *MempoolTransaction) GetTxStatus() string`

GetTxStatus returns the TxStatus field if non-nil, zero value otherwise.

### GetTxStatusOk

`func (o *MempoolTransaction) GetTxStatusOk() (*string, bool)`

GetTxStatusOk returns a tuple with the TxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxStatus

`func (o *MempoolTransaction) SetTxStatus(v string)`

SetTxStatus sets TxStatus field to given value.


### GetReceiptTime

`func (o *MempoolTransaction) GetReceiptTime() float32`

GetReceiptTime returns the ReceiptTime field if non-nil, zero value otherwise.

### GetReceiptTimeOk

`func (o *MempoolTransaction) GetReceiptTimeOk() (*float32, bool)`

GetReceiptTimeOk returns a tuple with the ReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTime

`func (o *MempoolTransaction) SetReceiptTime(v float32)`

SetReceiptTime sets ReceiptTime field to given value.


### GetReceiptTimeIso

`func (o *MempoolTransaction) GetReceiptTimeIso() string`

GetReceiptTimeIso returns the ReceiptTimeIso field if non-nil, zero value otherwise.

### GetReceiptTimeIsoOk

`func (o *MempoolTransaction) GetReceiptTimeIsoOk() (*string, bool)`

GetReceiptTimeIsoOk returns a tuple with the ReceiptTimeIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptTimeIso

`func (o *MempoolTransaction) SetReceiptTimeIso(v string)`

SetReceiptTimeIso sets ReceiptTimeIso field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


