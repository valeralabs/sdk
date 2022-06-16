# AddressTransactionWithTransfersNFTTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetIdentifier** | **string** | Non Fungible Token asset identifier. | 
**Value** | [**AddressTransactionWithTransfersNFTTransfersInnerValue**](AddressTransactionWithTransfersNFTTransfersInnerValue.md) |  | 
**Sender** | Pointer to **string** | Principal that sent the asset. | [optional] 
**Recipient** | Pointer to **string** | Principal that received the asset. | [optional] 

## Methods

### NewAddressTransactionWithTransfersNFTTransfersInner

`func NewAddressTransactionWithTransfersNFTTransfersInner(assetIdentifier string, value AddressTransactionWithTransfersNFTTransfersInnerValue, ) *AddressTransactionWithTransfersNFTTransfersInner`

NewAddressTransactionWithTransfersNFTTransfersInner instantiates a new AddressTransactionWithTransfersNFTTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersNFTTransfersInnerWithDefaults

`func NewAddressTransactionWithTransfersNFTTransfersInnerWithDefaults() *AddressTransactionWithTransfersNFTTransfersInner`

NewAddressTransactionWithTransfersNFTTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersNFTTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetIdentifier

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetAssetIdentifier() string`

GetAssetIdentifier returns the AssetIdentifier field if non-nil, zero value otherwise.

### GetAssetIdentifierOk

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetAssetIdentifierOk() (*string, bool)`

GetAssetIdentifierOk returns a tuple with the AssetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdentifier

`func (o *AddressTransactionWithTransfersNFTTransfersInner) SetAssetIdentifier(v string)`

SetAssetIdentifier sets AssetIdentifier field to given value.


### GetValue

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetValue() AddressTransactionWithTransfersNFTTransfersInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetValueOk() (*AddressTransactionWithTransfersNFTTransfersInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddressTransactionWithTransfersNFTTransfersInner) SetValue(v AddressTransactionWithTransfersNFTTransfersInnerValue)`

SetValue sets Value field to given value.


### GetSender

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionWithTransfersNFTTransfersInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionWithTransfersNFTTransfersInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionWithTransfersNFTTransfersInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionWithTransfersNFTTransfersInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionWithTransfersNFTTransfersInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


