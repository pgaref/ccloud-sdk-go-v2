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

// ListMirrorTopicsResponseDataAllOf struct for ListMirrorTopicsResponseDataAllOf
type ListMirrorTopicsResponseDataAllOf struct {
	LinkName        string            `json:"link_name,omitempty"`
	MirrorTopicName string            `json:"mirror_topic_name,omitempty"`
	SourceTopicName string            `json:"source_topic_name,omitempty"`
	NumPartitions   int32             `json:"num_partitions,omitempty"`
	MirrorLags      MirrorLags        `json:"mirror_lags,omitempty"`
	MirrorStatus    MirrorTopicStatus `json:"mirror_status,omitempty"`
	StateTimeMs     int64             `json:"state_time_ms,omitempty"`
}

// NewListMirrorTopicsResponseDataAllOf instantiates a new ListMirrorTopicsResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMirrorTopicsResponseDataAllOf(linkName string, mirrorTopicName string, sourceTopicName string, numPartitions int32, mirrorLags MirrorLags, mirrorStatus MirrorTopicStatus, stateTimeMs int64) *ListMirrorTopicsResponseDataAllOf {
	this := ListMirrorTopicsResponseDataAllOf{}
	this.LinkName = linkName
	this.MirrorTopicName = mirrorTopicName
	this.SourceTopicName = sourceTopicName
	this.NumPartitions = numPartitions
	this.MirrorLags = mirrorLags
	this.MirrorStatus = mirrorStatus
	this.StateTimeMs = stateTimeMs
	return &this
}

// NewListMirrorTopicsResponseDataAllOfWithDefaults instantiates a new ListMirrorTopicsResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMirrorTopicsResponseDataAllOfWithDefaults() *ListMirrorTopicsResponseDataAllOf {
	this := ListMirrorTopicsResponseDataAllOf{}
	return &this
}

// GetLinkName returns the LinkName field value
func (o *ListMirrorTopicsResponseDataAllOf) GetLinkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkName
}

// GetLinkNameOk returns a tuple with the LinkName field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetLinkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkName, true
}

// SetLinkName sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetLinkName(v string) {
	o.LinkName = v
}

// GetMirrorTopicName returns the MirrorTopicName field value
func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MirrorTopicName
}

// GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorTopicName, true
}

// SetMirrorTopicName sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetMirrorTopicName(v string) {
	o.MirrorTopicName = v
}

// GetSourceTopicName returns the SourceTopicName field value
func (o *ListMirrorTopicsResponseDataAllOf) GetSourceTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceTopicName
}

// GetSourceTopicNameOk returns a tuple with the SourceTopicName field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetSourceTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceTopicName, true
}

// SetSourceTopicName sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetSourceTopicName(v string) {
	o.SourceTopicName = v
}

// GetNumPartitions returns the NumPartitions field value
func (o *ListMirrorTopicsResponseDataAllOf) GetNumPartitions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPartitions
}

// GetNumPartitionsOk returns a tuple with the NumPartitions field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetNumPartitionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPartitions, true
}

// SetNumPartitions sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetNumPartitions(v int32) {
	o.NumPartitions = v
}

// GetMirrorLags returns the MirrorLags field value
func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorLags() MirrorLags {
	if o == nil {
		var ret MirrorLags
		return ret
	}

	return o.MirrorLags
}

// GetMirrorLagsOk returns a tuple with the MirrorLags field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorLagsOk() (*MirrorLags, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorLags, true
}

// SetMirrorLags sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetMirrorLags(v MirrorLags) {
	o.MirrorLags = v
}

// GetMirrorStatus returns the MirrorStatus field value
func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorStatus() MirrorTopicStatus {
	if o == nil {
		var ret MirrorTopicStatus
		return ret
	}

	return o.MirrorStatus
}

// GetMirrorStatusOk returns a tuple with the MirrorStatus field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetMirrorStatusOk() (*MirrorTopicStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorStatus, true
}

// SetMirrorStatus sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetMirrorStatus(v MirrorTopicStatus) {
	o.MirrorStatus = v
}

// GetStateTimeMs returns the StateTimeMs field value
func (o *ListMirrorTopicsResponseDataAllOf) GetStateTimeMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StateTimeMs
}

// GetStateTimeMsOk returns a tuple with the StateTimeMs field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseDataAllOf) GetStateTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateTimeMs, true
}

// SetStateTimeMs sets field value
func (o *ListMirrorTopicsResponseDataAllOf) SetStateTimeMs(v int64) {
	o.StateTimeMs = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ListMirrorTopicsResponseDataAllOf) Redact() {
	o.recurseRedact(&o.LinkName)
	o.recurseRedact(&o.MirrorTopicName)
	o.recurseRedact(&o.SourceTopicName)
	o.recurseRedact(&o.NumPartitions)
	o.recurseRedact(&o.MirrorLags)
	o.recurseRedact(&o.MirrorStatus)
	o.recurseRedact(&o.StateTimeMs)
}

func (o *ListMirrorTopicsResponseDataAllOf) recurseRedact(v interface{}) {
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

func (o ListMirrorTopicsResponseDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ListMirrorTopicsResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["link_name"] = o.LinkName
	}
	if true {
		toSerialize["mirror_topic_name"] = o.MirrorTopicName
	}
	if true {
		toSerialize["source_topic_name"] = o.SourceTopicName
	}
	if true {
		toSerialize["num_partitions"] = o.NumPartitions
	}
	if true {
		toSerialize["mirror_lags"] = o.MirrorLags
	}
	if true {
		toSerialize["mirror_status"] = o.MirrorStatus
	}
	if true {
		toSerialize["state_time_ms"] = o.StateTimeMs
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableListMirrorTopicsResponseDataAllOf struct {
	value *ListMirrorTopicsResponseDataAllOf
	isSet bool
}

func (v NullableListMirrorTopicsResponseDataAllOf) Get() *ListMirrorTopicsResponseDataAllOf {
	return v.value
}

func (v *NullableListMirrorTopicsResponseDataAllOf) Set(val *ListMirrorTopicsResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListMirrorTopicsResponseDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListMirrorTopicsResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMirrorTopicsResponseDataAllOf(val *ListMirrorTopicsResponseDataAllOf) *NullableListMirrorTopicsResponseDataAllOf {
	return &NullableListMirrorTopicsResponseDataAllOf{value: val, isSet: true}
}

func (v NullableListMirrorTopicsResponseDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableListMirrorTopicsResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
