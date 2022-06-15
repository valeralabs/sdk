# NftEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | **string** |  | 
**Recipient** | **string** |  | 
**AssetIdentifier** | **string** |  | 
**Value** | [**NftEventValue**](NftEventValue.md) |  | 
**TxId** | **string** |  | 
**BlockHeight** | **float32** |  | 

## Methods

### NewNftEvent

`func NewNftEvent(sender string, recipient string, assetIdentifier string, value NftEventValue, txId string, blockHeight float32, ) *NftEvent`

NewNftEvent instantiates a new NftEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftEventWithDefaults

`func NewNftEventWithDefaults() *NftEvent`

NewNftEventWithDefaults instantiates a new NftEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *NftEvent) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *NftEvent) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *NftEvent) SetSender(v string)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *NftEvent) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NftEvent) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NftEvent) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetAssetIdentifier

`func (o *NftEvent) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *NftEvent) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *NftEvent) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetValue

`func (o *NftEvent) GetValue() NftEventValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NftEvent) GetValueOk() (*NftEventValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NftEvent) SetValue(v NftEventValue)`

SetValue sets Value field to given value.


### GetTxId

`func (o *NftEvent) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *NftEvent) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *NftEvent) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetBlockHeight

`func (o *NftEvent) GetBlockHeight() float32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *NftEvent) GetBlockHeightOk() (*float32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *NftEvent) SetBlockHeight(v float32)`

SetBlockHeight sets BlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


