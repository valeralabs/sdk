# GetStxSupplyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnlockedPercent** | **string** | String quoted decimal number of the percentage of STX that have unlocked | 
**TotalSTX** | **string** | String quoted decimal number of the total possible number of STX | 
**UnlockedSTX** | **string** | String quoted decimal number of the STX that have been mined or unlocked | 
**BlockHeight** | **int32** | The block height at which this information was queried | 

## Methods

### NewGetStxSupplyResponse

`func NewGetStxSupplyResponse(unlockedPercent string, totalSTX string, unlockedSTX string, blockHeight int32, ) *GetStxSupplyResponse`

NewGetStxSupplyResponse instantiates a new GetStxSupplyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStxSupplyResponseWithDefaults

`func NewGetStxSupplyResponseWithDefaults() *GetStxSupplyResponse`

NewGetStxSupplyResponseWithDefaults instantiates a new GetStxSupplyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnlockedPercent

`func (o *GetStxSupplyResponse) GetUnlockedPercent() string`

GetUnlockedPercent returns the UnlockedPercent field if non-nil, zero value otherwise.

### GetUnlockedPercentOk

`func (o *GetStxSupplyResponse) GetUnlockedPercentOk() (*string, bool)`

GetUnlockedPercentOk returns a tuple with the UnlockedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedPercent

`func (o *GetStxSupplyResponse) SetUnlockedPercent(v string)`

SetUnlockedPercent sets UnlockedPercent field to given value.


### GetTotalSTX

`func (o *GetStxSupplyResponse) GetTotalSTX() string`

GetTotalSTX returns the TotalSTX field if non-nil, zero value otherwise.

### GetTotalSTXOk

`func (o *GetStxSupplyResponse) GetTotalSTXOk() (*string, bool)`

GetTotalSTXOk returns a tuple with the TotalSTX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSTX

`func (o *GetStxSupplyResponse) SetTotalSTX(v string)`

SetTotalSTX sets TotalSTX field to given value.


### GetUnlockedSTX

`func (o *GetStxSupplyResponse) GetUnlockedSTX() string`

GetUnlockedSTX returns the UnlockedSTX field if non-nil, zero value otherwise.

### GetUnlockedSTXOk

`func (o *GetStxSupplyResponse) GetUnlockedSTXOk() (*string, bool)`

GetUnlockedSTXOk returns a tuple with the UnlockedSTX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedSTX

`func (o *GetStxSupplyResponse) SetUnlockedSTX(v string)`

SetUnlockedSTX sets UnlockedSTX field to given value.


### GetBlockHeight

`func (o *GetStxSupplyResponse) GetBlockHeight() int32`

GetBlockHeight returns the BlockHeight field if non-nil, zero value otherwise.

### GetBlockHeightOk

`func (o *GetStxSupplyResponse) GetBlockHeightOk() (*int32, bool)`

GetBlockHeightOk returns a tuple with the BlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHeight

`func (o *GetStxSupplyResponse) SetBlockHeight(v int32)`

SetBlockHeight sets BlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


