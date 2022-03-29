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

// TopicConfigDataList struct for TopicConfigDataList
type TopicConfigDataList struct {
	Kind     string                     `json:"kind"`
	Metadata ResourceCollectionMetadata `json:"metadata"`
	Data     []TopicConfigData          `json:"data"`
}

// NewTopicConfigDataList instantiates a new TopicConfigDataList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicConfigDataList(kind string, metadata ResourceCollectionMetadata, data []TopicConfigData) *TopicConfigDataList {
	this := TopicConfigDataList{}
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewTopicConfigDataListWithDefaults instantiates a new TopicConfigDataList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicConfigDataListWithDefaults() *TopicConfigDataList {
	this := TopicConfigDataList{}
	return &this
}

// GetKind returns the Kind field value
func (o *TopicConfigDataList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *TopicConfigDataList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *TopicConfigDataList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *TopicConfigDataList) GetMetadata() ResourceCollectionMetadata {
	if o == nil {
		var ret ResourceCollectionMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *TopicConfigDataList) GetMetadataOk() (*ResourceCollectionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *TopicConfigDataList) SetMetadata(v ResourceCollectionMetadata) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *TopicConfigDataList) GetData() []TopicConfigData {
	if o == nil {
		var ret []TopicConfigData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TopicConfigDataList) GetDataOk() (*[]TopicConfigData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TopicConfigDataList) SetData(v []TopicConfigData) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *TopicConfigDataList) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *TopicConfigDataList) recurseRedact(v interface{}) {
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

func (o TopicConfigDataList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o TopicConfigDataList) MarshalJSON() ([]byte, error) {
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

type NullableTopicConfigDataList struct {
	value *TopicConfigDataList
	isSet bool
}

func (v NullableTopicConfigDataList) Get() *TopicConfigDataList {
	return v.value
}

func (v *NullableTopicConfigDataList) Set(val *TopicConfigDataList) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicConfigDataList) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicConfigDataList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicConfigDataList(val *TopicConfigDataList) *NullableTopicConfigDataList {
	return &NullableTopicConfigDataList{value: val, isSet: true}
}

func (v NullableTopicConfigDataList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicConfigDataList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
