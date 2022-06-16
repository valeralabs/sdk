# RunFaucetSTXRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | STX testnet address | [optional] 
**Stacking** | Pointer to **bool** | Use required number of tokens for stacking | [optional] 

## Methods

### NewRunFaucetSTXRequest

`func NewRunFaucetSTXRequest() *RunFaucetSTXRequest`

NewRunFaucetSTXRequest instantiates a new RunFaucetSTXRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFaucetSTXRequestWithDefaults

`func NewRunFaucetSTXRequestWithDefaults() *RunFaucetSTXRequest`

NewRunFaucetSTXRequestWithDefaults instantiates a new RunFaucetSTXRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RunFaucetSTXRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RunFaucetSTXRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RunFaucetSTXRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RunFaucetSTXRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStacking

`func (o *RunFaucetSTXRequest) GetStacking() bool`

GetStacking returns the Stacking field if non-nil, zero value otherwise.

### GetStackingOk

`func (o *RunFaucetSTXRequest) GetStackingOk() (*bool, bool)`

GetStackingOk returns a tuple with the Stacking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacking

`func (o *RunFaucetSTXRequest) SetStacking(v bool)`

SetStacking sets Stacking field to given value.

### HasStacking

`func (o *RunFaucetSTXRequest) HasStacking() bool`

HasStacking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


