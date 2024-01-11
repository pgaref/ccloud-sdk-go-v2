// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Network API

Network API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// NetworkingV1AzurePrivateLinkAttachmentStatus Azure PrivateLink attachment represents reserved capacity in a PrivateLink service that can be used to establish PrivateLink
type NetworkingV1AzurePrivateLinkAttachmentStatus struct {
	// PrivateLinkAttachmentStatus kind.
	Kind string `json:"kind,omitempty"`
	// Azure PrivateLink service that can be used to connect to a PrivateEndpoint.
	PrivateLinkService NetworkingV1AzurePrivateLinkService `json:"private_link_service,omitempty"`
}

// NewNetworkingV1AzurePrivateLinkAttachmentStatus instantiates a new NetworkingV1AzurePrivateLinkAttachmentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AzurePrivateLinkAttachmentStatus(kind string, privateLinkService NetworkingV1AzurePrivateLinkService) *NetworkingV1AzurePrivateLinkAttachmentStatus {
	this := NetworkingV1AzurePrivateLinkAttachmentStatus{}
	this.Kind = kind
	this.PrivateLinkService = privateLinkService
	return &this
}

// NewNetworkingV1AzurePrivateLinkAttachmentStatusWithDefaults instantiates a new NetworkingV1AzurePrivateLinkAttachmentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AzurePrivateLinkAttachmentStatusWithDefaults() *NetworkingV1AzurePrivateLinkAttachmentStatus {
	this := NetworkingV1AzurePrivateLinkAttachmentStatus{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) SetKind(v string) {
	o.Kind = v
}

// GetPrivateLinkService returns the PrivateLinkService field value
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetPrivateLinkService() NetworkingV1AzurePrivateLinkService {
	if o == nil {
		var ret NetworkingV1AzurePrivateLinkService
		return ret
	}

	return o.PrivateLinkService
}

// GetPrivateLinkServiceOk returns a tuple with the PrivateLinkService field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) GetPrivateLinkServiceOk() (*NetworkingV1AzurePrivateLinkService, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateLinkService, true
}

// SetPrivateLinkService sets field value
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) SetPrivateLinkService(v NetworkingV1AzurePrivateLinkService) {
	o.PrivateLinkService = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.PrivateLinkService)
}

func (o *NetworkingV1AzurePrivateLinkAttachmentStatus) recurseRedact(v interface{}) {
	type redactor interface {
		Redact()
	}
	if r, ok := v.(redactor); ok {
		r.Redact()
	} else {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < val.Len(); i++ {
				// support data types declared without pointers
				o.recurseRedact(val.Index(i).Interface())
				// ... and data types that were declared without but need pointers (for Redact)
				if val.Index(i).CanAddr() {
					o.recurseRedact(val.Index(i).Addr().Interface())
				}
			}
		}
	}
}

func (o NetworkingV1AzurePrivateLinkAttachmentStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AzurePrivateLinkAttachmentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["private_link_service"] = o.PrivateLinkService
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1AzurePrivateLinkAttachmentStatus struct {
	value *NetworkingV1AzurePrivateLinkAttachmentStatus
	isSet bool
}

func (v NullableNetworkingV1AzurePrivateLinkAttachmentStatus) Get() *NetworkingV1AzurePrivateLinkAttachmentStatus {
	return v.value
}

func (v *NullableNetworkingV1AzurePrivateLinkAttachmentStatus) Set(val *NetworkingV1AzurePrivateLinkAttachmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AzurePrivateLinkAttachmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AzurePrivateLinkAttachmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AzurePrivateLinkAttachmentStatus(val *NetworkingV1AzurePrivateLinkAttachmentStatus) *NullableNetworkingV1AzurePrivateLinkAttachmentStatus {
	return &NullableNetworkingV1AzurePrivateLinkAttachmentStatus{value: val, isSet: true}
}

func (v NullableNetworkingV1AzurePrivateLinkAttachmentStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1AzurePrivateLinkAttachmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
