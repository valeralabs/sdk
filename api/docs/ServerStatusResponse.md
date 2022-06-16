# ServerStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerVersion** | Pointer to **string** | the server version that is currently running | [optional] 
**Status** | **string** | the current server status | 
**ChainTip** | Pointer to [**ChainTip**](ChainTip.md) |  | [optional] 

## Methods

### NewServerStatusResponse

`func NewServerStatusResponse(status string, ) *ServerStatusResponse`

NewServerStatusResponse instantiates a new ServerStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerStatusResponseWithDefaults

`func NewServerStatusResponseWithDefaults() *ServerStatusResponse`

NewServerStatusResponseWithDefaults instantiates a new ServerStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerVersion

`func (o *ServerStatusResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *ServerStatusResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *ServerStatusResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *ServerStatusResponse) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ServerStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChainTip

`func (o *ServerStatusResponse) GetChainTip() ChainTip`

GetChainTip returns the ChainTip field if non-nil, zero value otherwise.

### GetChainTipOk

`func (o *ServerStatusResponse) GetChainTipOk() (*ChainTip, bool)`

GetChainTipOk returns a tuple with the ChainTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainTip

`func (o *ServerStatusResponse) SetChainTip(v ChainTip)`

SetChainTip sets ChainTip field to given value.

### HasChainTip

`func (o *ServerStatusResponse) HasChainTip() bool`

HasChainTip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


