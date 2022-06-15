# RunFaucetStxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | STX testnet address | [optional] 
**Stacking** | Pointer to **bool** | Use required number of tokens for stacking | [optional] 

## Methods

### NewRunFaucetStxRequest

`func NewRunFaucetStxRequest() *RunFaucetStxRequest`

NewRunFaucetStxRequest instantiates a new RunFaucetStxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunFaucetStxRequestWithDefaults

`func NewRunFaucetStxRequestWithDefaults() *RunFaucetStxRequest`

NewRunFaucetStxRequestWithDefaults instantiates a new RunFaucetStxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RunFaucetStxRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RunFaucetStxRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RunFaucetStxRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RunFaucetStxRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetStacking

`func (o *RunFaucetStxRequest) GetStacking() bool`

GetStacking returns the Stacking field if non-nil, zero value otherwise.

### GetStackingOk

`func (o *RunFaucetStxRequest) GetStackingOk() (*bool, bool)`

GetStackingOk returns a tuple with the Stacking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacking

`func (o *RunFaucetStxRequest) SetStacking(v bool)`

SetStacking sets Stacking field to given value.

### HasStacking

`func (o *RunFaucetStxRequest) HasStacking() bool`

HasStacking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


