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

// ConsumerGroupDataList struct for ConsumerGroupDataList
type ConsumerGroupDataList struct {
	Kind     string                     `json:"kind"`
	Metadata ResourceCollectionMetadata `json:"metadata"`
	Data     []ConsumerGroupData        `json:"data"`
}

// NewConsumerGroupDataList instantiates a new ConsumerGroupDataList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupDataList(kind string, metadata ResourceCollectionMetadata, data []ConsumerGroupData) *ConsumerGroupDataList {
	this := ConsumerGroupDataList{}
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewConsumerGroupDataListWithDefaults instantiates a new ConsumerGroupDataList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupDataListWithDefaults() *ConsumerGroupDataList {
	this := ConsumerGroupDataList{}
	return &this
}

// GetKind returns the Kind field value
func (o *ConsumerGroupDataList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupDataList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConsumerGroupDataList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ConsumerGroupDataList) GetMetadata() ResourceCollectionMetadata {
	if o == nil {
		var ret ResourceCollectionMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ConsumerGroupDataList) SetMetadata(v ResourceCollectionMetadata) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *ConsumerGroupDataList) GetData() []ConsumerGroupData {
	if o == nil {
		var ret []ConsumerGroupData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupDataList) GetDataOk() (*[]ConsumerGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConsumerGroupDataList) SetData(v []ConsumerGroupData) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerGroupDataList) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *ConsumerGroupDataList) recurseRedact(v interface{}) {
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

func (o ConsumerGroupDataList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerGroupDataList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableConsumerGroupDataList struct {
	value *ConsumerGroupDataList
	isSet bool
}

func (v NullableConsumerGroupDataList) Get() *ConsumerGroupDataList {
	return v.value
}

func (v *NullableConsumerGroupDataList) Set(val *ConsumerGroupDataList) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupDataList) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupDataList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupDataList(val *ConsumerGroupDataList) *NullableConsumerGroupDataList {
	return &NullableConsumerGroupDataList{value: val, isSet: true}
}

func (v NullableConsumerGroupDataList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumerGroupDataList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
