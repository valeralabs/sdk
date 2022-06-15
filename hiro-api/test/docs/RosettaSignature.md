# RosettaSignature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningPayload** | [**Object**](Object.md) |  | 
**PublicKey** | [**Object**](Object.md) |  | 
**SignatureType** | **string** | SignatureType is the type of a cryptographic signature. | 
**HexBytes** | **string** |  | 

## Methods

### NewRosettaSignature

`func NewRosettaSignature(signingPayload Object, publicKey Object, signatureType string, hexBytes string, ) *RosettaSignature`

NewRosettaSignature instantiates a new RosettaSignature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosettaSignatureWithDefaults

`func NewRosettaSignatureWithDefaults() *RosettaSignature`

NewRosettaSignatureWithDefaults instantiates a new RosettaSignature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningPayload

`func (o *RosettaSignature) GetSigningPayload() Object`

GetSigningPayload returns the SigningPayload field if non-nil, zero value otherwise.

### GetSigningPayloadOk

`func (o *RosettaSignature) GetSigningPayloadOk() (*Object, bool)`

GetSigningPayloadOk returns a tuple with the SigningPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningPayload

`func (o *RosettaSignature) SetSigningPayload(v Object)`

SetSigningPayload sets SigningPayload field to given value.


### GetPublicKey

`func (o *RosettaSignature) GetPublicKey() Object`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RosettaSignature) GetPublicKeyOk() (*Object, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RosettaSignature) SetPublicKey(v Object)`

SetPublicKey sets PublicKey field to given value.


### GetSignatureType

`func (o *RosettaSignature) GetSignatureType() string`

GetSignatureType returns the SignatureType field if non-nil, zero value otherwise.

### GetSignatureTypeOk

`func (o *RosettaSignature) GetSignatureTypeOk() (*string, bool)`

GetSignatureTypeOk returns a tuple with the SignatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureType

`func (o *RosettaSignature) SetSignatureType(v string)`

SetSignatureType sets SignatureType field to given value.


### GetHexBytes

`func (o *RosettaSignature) GetHexBytes() string`

GetHexBytes returns the HexBytes field if non-nil, zero value otherwise.

### GetHexBytesOk

`func (o *RosettaSignature) GetHexBytesOk() (*string, bool)`

GetHexBytesOk returns a tuple with the HexBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexBytes

`func (o *RosettaSignature) SetHexBytes(v string)`

SetHexBytes sets HexBytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


