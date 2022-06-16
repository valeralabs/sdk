# AddressTransactionWithTransfersSTXTransfersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount transferred in micro-STX as an integer string. | 
**Sender** | Pointer to **string** | Principal that sent STX. This is unspecified if the STX were minted. | [optional] 
**Recipient** | Pointer to **string** | Principal that received STX. This is unspecified if the STX were burned. | [optional] 

## Methods

### NewAddressTransactionWithTransfersSTXTransfersInner

`func NewAddressTransactionWithTransfersSTXTransfersInner(amount string, ) *AddressTransactionWithTransfersSTXTransfersInner`

NewAddressTransactionWithTransfersSTXTransfersInner instantiates a new AddressTransactionWithTransfersSTXTransfersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersSTXTransfersInnerWithDefaults

`func NewAddressTransactionWithTransfersSTXTransfersInnerWithDefaults() *AddressTransactionWithTransfersSTXTransfersInner`

NewAddressTransactionWithTransfersSTXTransfersInnerWithDefaults instantiates a new AddressTransactionWithTransfersSTXTransfersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *AddressTransactionWithTransfersSTXTransfersInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AddressTransactionWithTransfersSTXTransfersInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AddressTransactionWithTransfersSTXTransfersInner) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetSender

`func (o *AddressTransactionWithTransfersSTXTransfersInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AddressTransactionWithTransfersSTXTransfersInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AddressTransactionWithTransfersSTXTransfersInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AddressTransactionWithTransfersSTXTransfersInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *AddressTransactionWithTransfersSTXTransfersInner) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AddressTransactionWithTransfersSTXTransfersInner) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AddressTransactionWithTransfersSTXTransfersInner) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AddressTransactionWithTransfersSTXTransfersInner) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


