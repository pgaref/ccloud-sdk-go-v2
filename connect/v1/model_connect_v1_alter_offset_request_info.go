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
Kafka Connect APIs

REST API for managing connectors

API version: 1.0
Contact: connect@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"time"
)

import (
	"reflect"
)

// ConnectV1AlterOffsetRequestInfo The request made to alter offsets.
type ConnectV1AlterOffsetRequestInfo struct {
	// The ID of the connector.
	Id string `json:"id,omitempty"`
	// The name of the connector.
	Name string `json:"name,omitempty"`
	// Array of offsets which are categorised into partitions.
	Offsets *[]map[string]interface{} `json:"offsets,omitempty"`
	// The time at which the request was made. The time is in UTC, ISO 8601 format.
	RequestedAt time.Time                       `json:"requested_at,omitempty"`
	Type        ConnectV1AlterOffsetRequestType `json:"type,omitempty"`
}

// NewConnectV1AlterOffsetRequestInfo instantiates a new ConnectV1AlterOffsetRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1AlterOffsetRequestInfo(id string, name string, requestedAt time.Time, type_ ConnectV1AlterOffsetRequestType) *ConnectV1AlterOffsetRequestInfo {
	this := ConnectV1AlterOffsetRequestInfo{}
	this.Id = id
	this.Name = name
	this.RequestedAt = requestedAt
	this.Type = type_
	return &this
}

// NewConnectV1AlterOffsetRequestInfoWithDefaults instantiates a new ConnectV1AlterOffsetRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1AlterOffsetRequestInfoWithDefaults() *ConnectV1AlterOffsetRequestInfo {
	this := ConnectV1AlterOffsetRequestInfo{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectV1AlterOffsetRequestInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequestInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectV1AlterOffsetRequestInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ConnectV1AlterOffsetRequestInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequestInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectV1AlterOffsetRequestInfo) SetName(v string) {
	o.Name = v
}

// GetOffsets returns the Offsets field value if set, zero value otherwise.
func (o *ConnectV1AlterOffsetRequestInfo) GetOffsets() []map[string]interface{} {
	if o == nil || o.Offsets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Offsets
}

// GetOffsetsOk returns a tuple with the Offsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequestInfo) GetOffsetsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Offsets == nil {
		return nil, false
	}
	return o.Offsets, true
}

// HasOffsets returns a boolean if a field has been set.
func (o *ConnectV1AlterOffsetRequestInfo) HasOffsets() bool {
	if o != nil && o.Offsets != nil {
		return true
	}

	return false
}

// SetOffsets gets a reference to the given []map[string]interface{} and assigns it to the Offsets field.
func (o *ConnectV1AlterOffsetRequestInfo) SetOffsets(v []map[string]interface{}) {
	o.Offsets = &v
}

// GetRequestedAt returns the RequestedAt field value
func (o *ConnectV1AlterOffsetRequestInfo) GetRequestedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequestInfo) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedAt, true
}

// SetRequestedAt sets field value
func (o *ConnectV1AlterOffsetRequestInfo) SetRequestedAt(v time.Time) {
	o.RequestedAt = v
}

// GetType returns the Type field value
func (o *ConnectV1AlterOffsetRequestInfo) GetType() ConnectV1AlterOffsetRequestType {
	if o == nil {
		var ret ConnectV1AlterOffsetRequestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequestInfo) GetTypeOk() (*ConnectV1AlterOffsetRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectV1AlterOffsetRequestInfo) SetType(v ConnectV1AlterOffsetRequestType) {
	o.Type = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1AlterOffsetRequestInfo) Redact() {
	o.recurseRedact(&o.Id)
	o.recurseRedact(&o.Name)
	o.recurseRedact(o.Offsets)
	o.recurseRedact(&o.RequestedAt)
	o.recurseRedact(&o.Type)
}

func (o *ConnectV1AlterOffsetRequestInfo) recurseRedact(v interface{}) {
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

func (o ConnectV1AlterOffsetRequestInfo) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1AlterOffsetRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Offsets != nil {
		toSerialize["offsets"] = o.Offsets
	}
	if true {
		toSerialize["requested_at"] = o.RequestedAt
	}
	if true {
		toSerialize["type"] = o.Type
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1AlterOffsetRequestInfo struct {
	value *ConnectV1AlterOffsetRequestInfo
	isSet bool
}

func (v NullableConnectV1AlterOffsetRequestInfo) Get() *ConnectV1AlterOffsetRequestInfo {
	return v.value
}

func (v *NullableConnectV1AlterOffsetRequestInfo) Set(val *ConnectV1AlterOffsetRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1AlterOffsetRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1AlterOffsetRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1AlterOffsetRequestInfo(val *ConnectV1AlterOffsetRequestInfo) *NullableConnectV1AlterOffsetRequestInfo {
	return &NullableConnectV1AlterOffsetRequestInfo{value: val, isSet: true}
}

func (v NullableConnectV1AlterOffsetRequestInfo) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1AlterOffsetRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
