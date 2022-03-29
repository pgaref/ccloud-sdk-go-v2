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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"encoding/json"
)

import (
	"reflect"
)

// AlterMirrorStatusResponseDataListAllOf struct for AlterMirrorStatusResponseDataListAllOf
type AlterMirrorStatusResponseDataListAllOf struct {
	Data []AlterMirrorStatusResponseData `json:"data"`
}

// NewAlterMirrorStatusResponseDataListAllOf instantiates a new AlterMirrorStatusResponseDataListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlterMirrorStatusResponseDataListAllOf(data []AlterMirrorStatusResponseData) *AlterMirrorStatusResponseDataListAllOf {
	this := AlterMirrorStatusResponseDataListAllOf{}
	this.Data = data
	return &this
}

// NewAlterMirrorStatusResponseDataListAllOfWithDefaults instantiates a new AlterMirrorStatusResponseDataListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlterMirrorStatusResponseDataListAllOfWithDefaults() *AlterMirrorStatusResponseDataListAllOf {
	this := AlterMirrorStatusResponseDataListAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *AlterMirrorStatusResponseDataListAllOf) GetData() []AlterMirrorStatusResponseData {
	if o == nil {
		var ret []AlterMirrorStatusResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AlterMirrorStatusResponseDataListAllOf) GetDataOk() (*[]AlterMirrorStatusResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AlterMirrorStatusResponseDataListAllOf) SetData(v []AlterMirrorStatusResponseData) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *AlterMirrorStatusResponseDataListAllOf) Redact() {
	o.recurseRedact(&o.Data)
}

func (o *AlterMirrorStatusResponseDataListAllOf) recurseRedact(v interface{}) {
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

func (o AlterMirrorStatusResponseDataListAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o AlterMirrorStatusResponseDataListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAlterMirrorStatusResponseDataListAllOf struct {
	value *AlterMirrorStatusResponseDataListAllOf
	isSet bool
}

func (v NullableAlterMirrorStatusResponseDataListAllOf) Get() *AlterMirrorStatusResponseDataListAllOf {
	return v.value
}

func (v *NullableAlterMirrorStatusResponseDataListAllOf) Set(val *AlterMirrorStatusResponseDataListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlterMirrorStatusResponseDataListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlterMirrorStatusResponseDataListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlterMirrorStatusResponseDataListAllOf(val *AlterMirrorStatusResponseDataListAllOf) *NullableAlterMirrorStatusResponseDataListAllOf {
	return &NullableAlterMirrorStatusResponseDataListAllOf{value: val, isSet: true}
}

func (v NullableAlterMirrorStatusResponseDataListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlterMirrorStatusResponseDataListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
