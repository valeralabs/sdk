# AddressNftListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Offset** | **int32** |  | 
**Total** | **int32** |  | 
**NftEvents** | [**[]NftEvent**](NftEvent.md) |  | 

## Methods

### NewAddressNftListResponse

`func NewAddressNftListResponse(limit int32, offset int32, total int32, nftEvents []NftEvent, ) *AddressNftListResponse`

NewAddressNftListResponse instantiates a new AddressNftListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressNftListResponseWithDefaults

`func NewAddressNftListResponseWithDefaults() *AddressNftListResponse`

NewAddressNftListResponseWithDefaults instantiates a new AddressNftListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *AddressNftListResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AddressNftListResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AddressNftListResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *AddressNftListResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *AddressNftListResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *AddressNftListResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *AddressNftListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AddressNftListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AddressNftListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetNftEvents

`func (o *AddressNftListResponse) GetNftEvents() []NftEvent`

GetNftEvents returns the NftEvents field if non-nil, zero value otherwise.

### GetNftEventsOk

`func (o *AddressNftListResponse) GetNftEventsOk() (*[]NftEvent, bool)`

GetNftEventsOk returns a tuple with the NftEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNftEvents

`func (o *AddressNftListResponse) SetNftEvents(v []NftEvent)`

SetNftEvents sets NftEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


