# TransactionEventStxLockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**STXLockEvent** | [**TransactionEventStxLockAllOfSTXLockEvent**](TransactionEventStxLockAllOfSTXLockEvent.md) |  | 

## Methods

### NewTransactionEventStxLockAllOf

`func NewTransactionEventStxLockAllOf(eventType string, txId string, sTXLockEvent TransactionEventStxLockAllOfSTXLockEvent, ) *TransactionEventStxLockAllOf`

NewTransactionEventStxLockAllOf instantiates a new TransactionEventStxLockAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventStxLockAllOfWithDefaults

`func NewTransactionEventStxLockAllOfWithDefaults() *TransactionEventStxLockAllOf`

NewTransactionEventStxLockAllOfWithDefaults instantiates a new TransactionEventStxLockAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TransactionEventStxLockAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransactionEventStxLockAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransactionEventStxLockAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TransactionEventStxLockAllOf) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionEventStxLockAllOf) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionEventStxLockAllOf) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetSTXLockEvent

`func (o *TransactionEventStxLockAllOf) GetSTXLockEvent() TransactionEventStxLockAllOfSTXLockEvent`

GetSTXLockEvent returns the STXLockEvent field if non-nil, zero value otherwise.

### GetSTXLockEventOk

`func (o *TransactionEventStxLockAllOf) GetSTXLockEventOk() (*TransactionEventStxLockAllOfSTXLockEvent, bool)`

GetSTXLockEventOk returns a tuple with the STXLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTXLockEvent

`func (o *TransactionEventStxLockAllOf) SetSTXLockEvent(v TransactionEventStxLockAllOfSTXLockEvent)`

SetSTXLockEvent sets STXLockEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


