/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateExchangeAndURLAPIRequest struct for CreateExchangeAndURLAPIRequest
type CreateExchangeAndURLAPIRequest struct {
	// Provider name
	Provider *string `json:"provider,omitempty"`
	// Transaction type
	Type *string `json:"type,omitempty"`
	// Ethereum address of the user who wants to create transaction
	WalletAddress *string `json:"wallet_address,omitempty"`
	Widget *WidgetParams `json:"widget,omitempty"`
}

// NewCreateExchangeAndURLAPIRequest instantiates a new CreateExchangeAndURLAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExchangeAndURLAPIRequest() *CreateExchangeAndURLAPIRequest {
	this := CreateExchangeAndURLAPIRequest{}
	return &this
}

// NewCreateExchangeAndURLAPIRequestWithDefaults instantiates a new CreateExchangeAndURLAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExchangeAndURLAPIRequestWithDefaults() *CreateExchangeAndURLAPIRequest {
	this := CreateExchangeAndURLAPIRequest{}
	return &this
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CreateExchangeAndURLAPIRequest) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeAndURLAPIRequest) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CreateExchangeAndURLAPIRequest) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CreateExchangeAndURLAPIRequest) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateExchangeAndURLAPIRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeAndURLAPIRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateExchangeAndURLAPIRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateExchangeAndURLAPIRequest) SetType(v string) {
	o.Type = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *CreateExchangeAndURLAPIRequest) GetWalletAddress() string {
	if o == nil || o.WalletAddress == nil {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeAndURLAPIRequest) GetWalletAddressOk() (*string, bool) {
	if o == nil || o.WalletAddress == nil {
		return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *CreateExchangeAndURLAPIRequest) HasWalletAddress() bool {
	if o != nil && o.WalletAddress != nil {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *CreateExchangeAndURLAPIRequest) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

// GetWidget returns the Widget field value if set, zero value otherwise.
func (o *CreateExchangeAndURLAPIRequest) GetWidget() WidgetParams {
	if o == nil || o.Widget == nil {
		var ret WidgetParams
		return ret
	}
	return *o.Widget
}

// GetWidgetOk returns a tuple with the Widget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateExchangeAndURLAPIRequest) GetWidgetOk() (*WidgetParams, bool) {
	if o == nil || o.Widget == nil {
		return nil, false
	}
	return o.Widget, true
}

// HasWidget returns a boolean if a field has been set.
func (o *CreateExchangeAndURLAPIRequest) HasWidget() bool {
	if o != nil && o.Widget != nil {
		return true
	}

	return false
}

// SetWidget gets a reference to the given WidgetParams and assigns it to the Widget field.
func (o *CreateExchangeAndURLAPIRequest) SetWidget(v WidgetParams) {
	o.Widget = &v
}

func (o CreateExchangeAndURLAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.WalletAddress != nil {
		toSerialize["wallet_address"] = o.WalletAddress
	}
	if o.Widget != nil {
		toSerialize["widget"] = o.Widget
	}
	return json.Marshal(toSerialize)
}

type NullableCreateExchangeAndURLAPIRequest struct {
	value *CreateExchangeAndURLAPIRequest
	isSet bool
}

func (v NullableCreateExchangeAndURLAPIRequest) Get() *CreateExchangeAndURLAPIRequest {
	return v.value
}

func (v *NullableCreateExchangeAndURLAPIRequest) Set(val *CreateExchangeAndURLAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExchangeAndURLAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExchangeAndURLAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExchangeAndURLAPIRequest(val *CreateExchangeAndURLAPIRequest) *NullableCreateExchangeAndURLAPIRequest {
	return &NullableCreateExchangeAndURLAPIRequest{value: val, isSet: true}
}

func (v NullableCreateExchangeAndURLAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateExchangeAndURLAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


