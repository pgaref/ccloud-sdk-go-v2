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
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ClusterDataListAllOf struct for ClusterDataListAllOf
type ClusterDataListAllOf struct {
	Data []ClusterData `json:"data,omitempty"`
}

// NewClusterDataListAllOf instantiates a new ClusterDataListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDataListAllOf(data []ClusterData) *ClusterDataListAllOf {
	this := ClusterDataListAllOf{}
	this.Data = data
	return &this
}

// NewClusterDataListAllOfWithDefaults instantiates a new ClusterDataListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDataListAllOfWithDefaults() *ClusterDataListAllOf {
	this := ClusterDataListAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *ClusterDataListAllOf) GetData() []ClusterData {
	if o == nil {
		var ret []ClusterData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ClusterDataListAllOf) GetDataOk() (*[]ClusterData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ClusterDataListAllOf) SetData(v []ClusterData) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ClusterDataListAllOf) Redact() {
	o.recurseRedact(&o.Data)
}

func (o *ClusterDataListAllOf) recurseRedact(v interface{}) {
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

func (o ClusterDataListAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ClusterDataListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableClusterDataListAllOf struct {
	value *ClusterDataListAllOf
	isSet bool
}

func (v NullableClusterDataListAllOf) Get() *ClusterDataListAllOf {
	return v.value
}

func (v *NullableClusterDataListAllOf) Set(val *ClusterDataListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDataListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDataListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDataListAllOf(val *ClusterDataListAllOf) *NullableClusterDataListAllOf {
	return &NullableClusterDataListAllOf{value: val, isSet: true}
}

func (v NullableClusterDataListAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableClusterDataListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
