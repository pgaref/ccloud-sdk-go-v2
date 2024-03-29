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

// ConnectV1AlterOffsetStatus Status of the alter offset operation. The previous offsets in the response  is the offsets that the connector last processed, before the offsets were altered, via a patch or delete operation.
type ConnectV1AlterOffsetStatus struct {
	Request ConnectV1AlterOffsetRequestInfo  `json:"request,omitempty"`
	Status  ConnectV1AlterOffsetStatusStatus `json:"status,omitempty"`
	// Array of offsets which are categorised into partitions.
	PreviousOffsets *[]map[string]interface{} `json:"previous_offsets,omitempty"`
	// The time at which the offsets were applied. The time is in UTC, ISO 8601 format.
	AppliedAt NullableTime `json:"applied_at,omitempty"`
}

// NewConnectV1AlterOffsetStatus instantiates a new ConnectV1AlterOffsetStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1AlterOffsetStatus(request ConnectV1AlterOffsetRequestInfo, status ConnectV1AlterOffsetStatusStatus) *ConnectV1AlterOffsetStatus {
	this := ConnectV1AlterOffsetStatus{}
	this.Request = request
	this.Status = status
	return &this
}

// NewConnectV1AlterOffsetStatusWithDefaults instantiates a new ConnectV1AlterOffsetStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1AlterOffsetStatusWithDefaults() *ConnectV1AlterOffsetStatus {
	this := ConnectV1AlterOffsetStatus{}
	return &this
}

// GetRequest returns the Request field value
func (o *ConnectV1AlterOffsetStatus) GetRequest() ConnectV1AlterOffsetRequestInfo {
	if o == nil {
		var ret ConnectV1AlterOffsetRequestInfo
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetStatus) GetRequestOk() (*ConnectV1AlterOffsetRequestInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *ConnectV1AlterOffsetStatus) SetRequest(v ConnectV1AlterOffsetRequestInfo) {
	o.Request = v
}

// GetStatus returns the Status field value
func (o *ConnectV1AlterOffsetStatus) GetStatus() ConnectV1AlterOffsetStatusStatus {
	if o == nil {
		var ret ConnectV1AlterOffsetStatusStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetStatus) GetStatusOk() (*ConnectV1AlterOffsetStatusStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ConnectV1AlterOffsetStatus) SetStatus(v ConnectV1AlterOffsetStatusStatus) {
	o.Status = v
}

// GetPreviousOffsets returns the PreviousOffsets field value if set, zero value otherwise.
func (o *ConnectV1AlterOffsetStatus) GetPreviousOffsets() []map[string]interface{} {
	if o == nil || o.PreviousOffsets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.PreviousOffsets
}

// GetPreviousOffsetsOk returns a tuple with the PreviousOffsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetStatus) GetPreviousOffsetsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.PreviousOffsets == nil {
		return nil, false
	}
	return o.PreviousOffsets, true
}

// HasPreviousOffsets returns a boolean if a field has been set.
func (o *ConnectV1AlterOffsetStatus) HasPreviousOffsets() bool {
	if o != nil && o.PreviousOffsets != nil {
		return true
	}

	return false
}

// SetPreviousOffsets gets a reference to the given []map[string]interface{} and assigns it to the PreviousOffsets field.
func (o *ConnectV1AlterOffsetStatus) SetPreviousOffsets(v []map[string]interface{}) {
	o.PreviousOffsets = &v
}

// GetAppliedAt returns the AppliedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectV1AlterOffsetStatus) GetAppliedAt() time.Time {
	if o == nil || o.AppliedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.AppliedAt.Get()
}

// GetAppliedAtOk returns a tuple with the AppliedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectV1AlterOffsetStatus) GetAppliedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppliedAt.Get(), o.AppliedAt.IsSet()
}

// HasAppliedAt returns a boolean if a field has been set.
func (o *ConnectV1AlterOffsetStatus) HasAppliedAt() bool {
	if o != nil && o.AppliedAt.IsSet() {
		return true
	}

	return false
}

// SetAppliedAt gets a reference to the given NullableTime and assigns it to the AppliedAt field.
func (o *ConnectV1AlterOffsetStatus) SetAppliedAt(v time.Time) {
	o.AppliedAt.Set(&v)
}

// SetAppliedAtNil sets the value for AppliedAt to be an explicit nil
func (o *ConnectV1AlterOffsetStatus) SetAppliedAtNil() {
	o.AppliedAt.Set(nil)
}

// UnsetAppliedAt ensures that no value is present for AppliedAt, not even an explicit nil
func (o *ConnectV1AlterOffsetStatus) UnsetAppliedAt() {
	o.AppliedAt.Unset()
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1AlterOffsetStatus) Redact() {
	o.recurseRedact(&o.Request)
	o.recurseRedact(&o.Status)
	o.recurseRedact(o.PreviousOffsets)
	o.recurseRedact(o.AppliedAt)
}

func (o *ConnectV1AlterOffsetStatus) recurseRedact(v interface{}) {
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

func (o ConnectV1AlterOffsetStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1AlterOffsetStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request"] = o.Request
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.PreviousOffsets != nil {
		toSerialize["previous_offsets"] = o.PreviousOffsets
	}
	if o.AppliedAt.IsSet() {
		toSerialize["applied_at"] = o.AppliedAt.Get()
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1AlterOffsetStatus struct {
	value *ConnectV1AlterOffsetStatus
	isSet bool
}

func (v NullableConnectV1AlterOffsetStatus) Get() *ConnectV1AlterOffsetStatus {
	return v.value
}

func (v *NullableConnectV1AlterOffsetStatus) Set(val *ConnectV1AlterOffsetStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1AlterOffsetStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1AlterOffsetStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1AlterOffsetStatus(val *ConnectV1AlterOffsetStatus) *NullableConnectV1AlterOffsetStatus {
	return &NullableConnectV1AlterOffsetStatus{value: val, isSet: true}
}

func (v NullableConnectV1AlterOffsetStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1AlterOffsetStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
