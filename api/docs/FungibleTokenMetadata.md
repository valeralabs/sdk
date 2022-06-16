# FungibleTokenMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenUri** | **string** | An optional string that is a valid URI which resolves to this token&#39;s metadata. Can be empty. | 
**Name** | **string** | Identifies the asset to which this token represents | 
**Description** | **string** | Describes the asset to which this token represents | 
**ImageUri** | **string** | A URI pointing to a resource with mime type image/_* representing the asset to which this token represents. The API may provide a URI to a cached resource, dependending on configuration. Otherwise, this can be the same value as the canonical image URI. | 
**ImageCanonicalUri** | **string** | The original image URI specified by the contract. A URI pointing to a resource with mime type image/_* representing the asset to which this token represents. Consider making any images at a width between 320 and 1080 pixels and aspect ratio between 1.91:1 and 4:5 inclusive. | 
**Symbol** | **string** | A shorter representation of a token. This is sometimes referred to as a \&quot;ticker\&quot;. Examples: \&quot;STX\&quot;, \&quot;COOL\&quot;, etc. Typically, a token could be referred to as $SYMBOL when referencing it in writing. | 
**Decimals** | **float32** | The number of decimal places in a token. | 
**TxId** | **string** | Tx id that deployed the contract | 
**SenderAddress** | **string** | principle that deployed the contract | 

## Methods

### NewFungibleTokenMetadata

`func NewFungibleTokenMetadata(tokenUri string, name string, description string, imageUri string, imageCanonicalUri string, symbol string, decimals float32, txId string, senderAddress string, ) *FungibleTokenMetadata`

NewFungibleTokenMetadata instantiates a new FungibleTokenMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFungibleTokenMetadataWithDefaults

`func NewFungibleTokenMetadataWithDefaults() *FungibleTokenMetadata`

NewFungibleTokenMetadataWithDefaults instantiates a new FungibleTokenMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenUri

`func (o *FungibleTokenMetadata) GetTokenUri() string`

GetTokenUri returns the TokenUri field if non-nil, zero value otherwise.

### GetTokenUriOk

`func (o *FungibleTokenMetadata) GetTokenUriOk() (*string, bool)`

GetTokenUriOk returns a tuple with the TokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUri

`func (o *FungibleTokenMetadata) SetTokenUri(v string)`

SetTokenUri sets TokenUri field to given value.


### GetName

`func (o *FungibleTokenMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FungibleTokenMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FungibleTokenMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FungibleTokenMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FungibleTokenMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FungibleTokenMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImageUri

`func (o *FungibleTokenMetadata) GetImageUri() string`

GetImageUri returns the ImageUri field if non-nil, zero value otherwise.

### GetImageUriOk

`func (o *FungibleTokenMetadata) GetImageUriOk() (*string, bool)`

GetImageUriOk returns a tuple with the ImageUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUri

`func (o *FungibleTokenMetadata) SetImageUri(v string)`

SetImageUri sets ImageUri field to given value.


### GetImageCanonicalUri

`func (o *FungibleTokenMetadata) GetImageCanonicalUri() string`

GetImageCanonicalUri returns the ImageCanonicalUri field if non-nil, zero value otherwise.

### GetImageCanonicalUriOk

`func (o *FungibleTokenMetadata) GetImageCanonicalUriOk() (*string, bool)`

GetImageCanonicalUriOk returns a tuple with the ImageCanonicalUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCanonicalUri

`func (o *FungibleTokenMetadata) SetImageCanonicalUri(v string)`

SetImageCanonicalUri sets ImageCanonicalUri field to given value.


### GetSymbol

`func (o *FungibleTokenMetadata) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FungibleTokenMetadata) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FungibleTokenMetadata) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *FungibleTokenMetadata) GetDecimals() float32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *FungibleTokenMetadata) GetDecimalsOk() (*float32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *FungibleTokenMetadata) SetDecimals(v float32)`

SetDecimals sets Decimals field to given value.


### GetTxId

`func (o *FungibleTokenMetadata) GetTxId() string`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *FungibleTokenMetadata) GetTxIdOk() (*string, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *FungibleTokenMetadata) SetTxId(v string)`

SetTxId sets TxId field to given value.


### GetSenderAddress

`func (o *FungibleTokenMetadata) GetSenderAddress() string`

GetSenderAddress returns the SenderAddress field if non-nil, zero value otherwise.

### GetSenderAddressOk

`func (o *FungibleTokenMetadata) GetSenderAddressOk() (*string, bool)`

GetSenderAddressOk returns a tuple with the SenderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAddress

`func (o *FungibleTokenMetadata) SetSenderAddress(v string)`

SetSenderAddress sets SenderAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


