# TransactionEventFungibleAssetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** |  | 
**TxId** | **string** |  | 
**Asset** | [**TransactionEventFungibleAssetAllOfAsset**](TransactionEventFungibleAssetAllOfAsset.md) |  | 

## Methods

### NewTransactionEventFungibleAssetAllOf

`func NewTransactionEventFungibleAssetAllOf(eventType string, txId string, asset TransactionEventFungibleAssetAllOfAsset, ) *TransactionEventFungibleAssetAllOf`

NewTransactionEventFungibleAssetAllOf instantiates a new TransactionEventFungibleAssetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEventFungibleAssetAllOfWithDefaults

`func NewTransactionEventFungibleAssetAllOfWithDefaults() *TransactionEventFungibleAssetAllOf`

NewTransactionEventFungibleAssetAllOfWithDefaults instantiates a new TransactionEventFungibleAssetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TransactionEventFungibleAssetAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransactionEventFungibleAssetAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransactionEventFungibleAssetAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTxId

`func (o *TransactionEventFungibleAssetAllOf) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *TransactionEventFungibleAssetAllOf) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *TransactionEventFungibleAssetAllOf) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetAsset

`func (o *TransactionEventFungibleAssetAllOf) GetAsset() TransactionEventFungibleAssetAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *TransactionEventFungibleAssetAllOf) GetAssetOk() (*TransactionEventFungibleAssetAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *TransactionEventFungibleAssetAllOf) SetAsset(v TransactionEventFungibleAssetAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


