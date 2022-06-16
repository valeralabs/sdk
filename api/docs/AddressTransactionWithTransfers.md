# AddressTransactionWithTransfers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | [**Transaction**](Transaction.md) |  | 
**STXSent** | **string** | Total sent from the given address, including the tx fee, in micro-STX as an integer string. | 
**STXReceived** | **string** | Total received by the given address in micro-STX as an integer string. | 
**STXTransfers** | [**[]AddressTransactionWithTransfersSTXTransfersInner**](AddressTransactionWithTransfersSTXTransfersInner.md) |  | 
**FtTransfers** | Pointer to [**[]AddressTransactionWithTransfersFtTransfersInner**](AddressTransactionWithTransfersFtTransfersInner.md) |  | [optional] 
**NFTTransfers** | Pointer to [**[]AddressTransactionWithTransfersNFTTransfersInner**](AddressTransactionWithTransfersNFTTransfersInner.md) |  | [optional] 

## Methods

### NewAddressTransactionWithTransfers

`func NewAddressTransactionWithTransfers(tx Transaction, sTXSent string, sTXReceived string, sTXTransfers []AddressTransactionWithTransfersSTXTransfersInner, ) *AddressTransactionWithTransfers`

NewAddressTransactionWithTransfers instantiates a new AddressTransactionWithTransfers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressTransactionWithTransfersWithDefaults

`func NewAddressTransactionWithTransfersWithDefaults() *AddressTransactionWithTransfers`

NewAddressTransactionWithTransfersWithDefaults instantiates a new AddressTransactionWithTransfers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *AddressTransactionWithTransfers) GetTx() Transaction`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *AddressTransactionWithTransfers) GetTxOk() (*Transaction, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *AddressTransactionWithTransfers) SetTx(v Transaction)`

SetTx sets Tx field to given value.


### GetSTXSent

`func (o *AddressTransactionWithTransfers) GetSTXSent() string`

GetSTXSent returns the STXSent field if non-nil, zero value otherwise.

### GetSTXSentOk

`func (o *AddressTransactionWithTransfers) GetSTXSentOk() (*string, bool)`

GetSTXSentOk returns a tuple with the STXSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTXSent

`func (o *AddressTransactionWithTransfers) SetSTXSent(v string)`

SetSTXSent sets STXSent field to given value.


### GetSTXReceived

`func (o *AddressTransactionWithTransfers) GetSTXReceived() string`

GetSTXReceived returns the STXReceived field if non-nil, zero value otherwise.

### GetSTXReceivedOk

`func (o *AddressTransactionWithTransfers) GetSTXReceivedOk() (*string, bool)`

GetSTXReceivedOk returns a tuple with the STXReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTXReceived

`func (o *AddressTransactionWithTransfers) SetSTXReceived(v string)`

SetSTXReceived sets STXReceived field to given value.


### GetSTXTransfers

`func (o *AddressTransactionWithTransfers) GetSTXTransfers() []AddressTransactionWithTransfersSTXTransfersInner`

GetSTXTransfers returns the STXTransfers field if non-nil, zero value otherwise.

### GetSTXTransfersOk

`func (o *AddressTransactionWithTransfers) GetSTXTransfersOk() (*[]AddressTransactionWithTransfersSTXTransfersInner, bool)`

GetSTXTransfersOk returns a tuple with the STXTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTXTransfers

`func (o *AddressTransactionWithTransfers) SetSTXTransfers(v []AddressTransactionWithTransfersSTXTransfersInner)`

SetSTXTransfers sets STXTransfers field to given value.


### GetFtTransfers

`func (o *AddressTransactionWithTransfers) GetFtTransfers() []AddressTransactionWithTransfersFtTransfersInner`

GetFtTransfers returns the FtTransfers field if non-nil, zero value otherwise.

### GetFtTransfersOk

`func (o *AddressTransactionWithTransfers) GetFtTransfersOk() (*[]AddressTransactionWithTransfersFtTransfersInner, bool)`

GetFtTransfersOk returns a tuple with the FtTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtTransfers

`func (o *AddressTransactionWithTransfers) SetFtTransfers(v []AddressTransactionWithTransfersFtTransfersInner)`

SetFtTransfers sets FtTransfers field to given value.

### HasFtTransfers

`func (o *AddressTransactionWithTransfers) HasFtTransfers() bool`

HasFtTransfers returns a boolean if a field has been set.

### GetNFTTransfers

`func (o *AddressTransactionWithTransfers) GetNFTTransfers() []AddressTransactionWithTransfersNFTTransfersInner`

GetNFTTransfers returns the NFTTransfers field if non-nil, zero value otherwise.

### GetNFTTransfersOk

`func (o *AddressTransactionWithTransfers) GetNFTTransfersOk() (*[]AddressTransactionWithTransfersNFTTransfersInner, bool)`

GetNFTTransfersOk returns a tuple with the NFTTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFTTransfers

`func (o *AddressTransactionWithTransfers) SetNFTTransfers(v []AddressTransactionWithTransfersNFTTransfersInner)`

SetNFTTransfers sets NFTTransfers field to given value.

### HasNFTTransfers

`func (o *AddressTransactionWithTransfers) HasNFTTransfers() bool`

HasNFTTransfers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


