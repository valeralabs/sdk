# TransactionEventStxLockAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**StxLockEvent** | [**TransactionEventStxLockAllOfStxLockEvent**](TransactionEventStxLockAllOfStxLockEvent.md) |  | 

## Methods

### NewTransactionEventStxLockAllOf

`func NewTransactionEventStxLockAllOf(eventType string, txId string, stxLockEvent TransactionEventStxLockAllOfStxLockEvent, ) *TransactionEventStxLockAllOf`

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


### GetStxLockEvent

`func (o *TransactionEventStxLockAllOf) GetStxLockEvent() TransactionEventStxLockAllOfStxLockEvent`

GetStxLockEvent returns the StxLockEvent field if non-nil, zero value otherwise.

### GetStxLockEventOk

`func (o *TransactionEventStxLockAllOf) GetStxLockEventOk() (*TransactionEventStxLockAllOfStxLockEvent, bool)`

GetStxLockEventOk returns a tuple with the StxLockEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStxLockEvent

`func (o *TransactionEventStxLockAllOf) SetStxLockEvent(v TransactionEventStxLockAllOfStxLockEvent)`

SetStxLockEvent sets StxLockEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


