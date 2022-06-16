# PostConditionNonFungibleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionCode** | **string** | A non-fungible condition code encodes a statement being made about a non-fungible token, with respect to whether or not the particular non-fungible token is owned by the account. | 
**Type** | **string** |  | 
**AssetValue** | [**PostConditionNonFungibleAllOfAssetValue**](PostConditionNonFungibleAllOfAssetValue.md) |  | 
**Asset** | [**PostConditionFungibleAllOfAsset**](PostConditionFungibleAllOfAsset.md) |  | 

## Methods

### NewPostConditionNonFungibleAllOf

`func NewPostConditionNonFungibleAllOf(conditionCode string, type_ string, assetValue PostConditionNonFungibleAllOfAssetValue, asset PostConditionFungibleAllOfAsset, ) *PostConditionNonFungibleAllOf`

NewPostConditionNonFungibleAllOf instantiates a new PostConditionNonFungibleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostConditionNonFungibleAllOfWithDefaults

`func NewPostConditionNonFungibleAllOfWithDefaults() *PostConditionNonFungibleAllOf`

NewPostConditionNonFungibleAllOfWithDefaults instantiates a new PostConditionNonFungibleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionCode

`func (o *PostConditionNonFungibleAllOf) GetConditionCode() string`

GetConditionCode returns the ConditionCode field if non-nil, zero value otherwise.

### GetConditionCodeOk

`func (o *PostConditionNonFungibleAllOf) GetConditionCodeOk() (*string, bool)`

GetConditionCodeOk returns a tuple with the ConditionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionCode

`func (o *PostConditionNonFungibleAllOf) SetConditionCode(v string)`

SetConditionCode sets ConditionCode field to given value.


### GetType

`func (o *PostConditionNonFungibleAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostConditionNonFungibleAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostConditionNonFungibleAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetAssetValue

`func (o *PostConditionNonFungibleAllOf) GetAssetValue() PostConditionNonFungibleAllOfAssetValue`

GetAssetValue returns the AssetValue field if non-nil, zero value otherwise.

### GetAssetValueOk

`func (o *PostConditionNonFungibleAllOf) GetAssetValueOk() (*PostConditionNonFungibleAllOfAssetValue, bool)`

GetAssetValueOk returns a tuple with the AssetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetValue

`func (o *PostConditionNonFungibleAllOf) SetAssetValue(v PostConditionNonFungibleAllOfAssetValue)`

SetAssetValue sets AssetValue field to given value.


### GetAsset

`func (o *PostConditionNonFungibleAllOf) GetAsset() PostConditionFungibleAllOfAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *PostConditionNonFungibleAllOf) GetAssetOk() (*PostConditionFungibleAllOfAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *PostConditionNonFungibleAllOf) SetAsset(v PostConditionFungibleAllOfAsset)`

SetAsset sets Asset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


