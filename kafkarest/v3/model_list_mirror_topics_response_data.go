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

// ListMirrorTopicsResponseData struct for ListMirrorTopicsResponseData
type ListMirrorTopicsResponseData struct {
	Kind                        string            `json:"kind,omitempty"`
	Metadata                    ResourceMetadata  `json:"metadata,omitempty"`
	LinkName                    string            `json:"link_name,omitempty"`
	MirrorTopicName             string            `json:"mirror_topic_name,omitempty"`
	SourceTopicName             string            `json:"source_topic_name,omitempty"`
	NumPartitions               int32             `json:"num_partitions,omitempty"`
	MirrorLags                  MirrorLags        `json:"mirror_lags,omitempty"`
	MirrorStatus                MirrorTopicStatus `json:"mirror_status,omitempty"`
	StateTimeMs                 int64             `json:"state_time_ms,omitempty"`
	MirrorStateTransitionErrors *[]LinkTaskError  `json:"mirror_state_transition_errors,omitempty"`
}

// NewListMirrorTopicsResponseData instantiates a new ListMirrorTopicsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMirrorTopicsResponseData(kind string, metadata ResourceMetadata, linkName string, mirrorTopicName string, sourceTopicName string, numPartitions int32, mirrorLags MirrorLags, mirrorStatus MirrorTopicStatus, stateTimeMs int64) *ListMirrorTopicsResponseData {
	this := ListMirrorTopicsResponseData{}
	this.Kind = kind
	this.Metadata = metadata
	this.LinkName = linkName
	this.MirrorTopicName = mirrorTopicName
	this.SourceTopicName = sourceTopicName
	this.NumPartitions = numPartitions
	this.MirrorLags = mirrorLags
	this.MirrorStatus = mirrorStatus
	this.StateTimeMs = stateTimeMs
	return &this
}

// NewListMirrorTopicsResponseDataWithDefaults instantiates a new ListMirrorTopicsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMirrorTopicsResponseDataWithDefaults() *ListMirrorTopicsResponseData {
	this := ListMirrorTopicsResponseData{}
	return &this
}

// GetKind returns the Kind field value
func (o *ListMirrorTopicsResponseData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ListMirrorTopicsResponseData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ListMirrorTopicsResponseData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ListMirrorTopicsResponseData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetLinkName returns the LinkName field value
func (o *ListMirrorTopicsResponseData) GetLinkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkName
}

// GetLinkNameOk returns a tuple with the LinkName field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetLinkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkName, true
}

// SetLinkName sets field value
func (o *ListMirrorTopicsResponseData) SetLinkName(v string) {
	o.LinkName = v
}

// GetMirrorTopicName returns the MirrorTopicName field value
func (o *ListMirrorTopicsResponseData) GetMirrorTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MirrorTopicName
}

// GetMirrorTopicNameOk returns a tuple with the MirrorTopicName field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetMirrorTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorTopicName, true
}

// SetMirrorTopicName sets field value
func (o *ListMirrorTopicsResponseData) SetMirrorTopicName(v string) {
	o.MirrorTopicName = v
}

// GetSourceTopicName returns the SourceTopicName field value
func (o *ListMirrorTopicsResponseData) GetSourceTopicName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceTopicName
}

// GetSourceTopicNameOk returns a tuple with the SourceTopicName field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetSourceTopicNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceTopicName, true
}

// SetSourceTopicName sets field value
func (o *ListMirrorTopicsResponseData) SetSourceTopicName(v string) {
	o.SourceTopicName = v
}

// GetNumPartitions returns the NumPartitions field value
func (o *ListMirrorTopicsResponseData) GetNumPartitions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumPartitions
}

// GetNumPartitionsOk returns a tuple with the NumPartitions field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetNumPartitionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumPartitions, true
}

// SetNumPartitions sets field value
func (o *ListMirrorTopicsResponseData) SetNumPartitions(v int32) {
	o.NumPartitions = v
}

// GetMirrorLags returns the MirrorLags field value
func (o *ListMirrorTopicsResponseData) GetMirrorLags() MirrorLags {
	if o == nil {
		var ret MirrorLags
		return ret
	}

	return o.MirrorLags
}

// GetMirrorLagsOk returns a tuple with the MirrorLags field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetMirrorLagsOk() (*MirrorLags, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorLags, true
}

// SetMirrorLags sets field value
func (o *ListMirrorTopicsResponseData) SetMirrorLags(v MirrorLags) {
	o.MirrorLags = v
}

// GetMirrorStatus returns the MirrorStatus field value
func (o *ListMirrorTopicsResponseData) GetMirrorStatus() MirrorTopicStatus {
	if o == nil {
		var ret MirrorTopicStatus
		return ret
	}

	return o.MirrorStatus
}

// GetMirrorStatusOk returns a tuple with the MirrorStatus field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetMirrorStatusOk() (*MirrorTopicStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirrorStatus, true
}

// SetMirrorStatus sets field value
func (o *ListMirrorTopicsResponseData) SetMirrorStatus(v MirrorTopicStatus) {
	o.MirrorStatus = v
}

// GetStateTimeMs returns the StateTimeMs field value
func (o *ListMirrorTopicsResponseData) GetStateTimeMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StateTimeMs
}

// GetStateTimeMsOk returns a tuple with the StateTimeMs field value
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetStateTimeMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateTimeMs, true
}

// SetStateTimeMs sets field value
func (o *ListMirrorTopicsResponseData) SetStateTimeMs(v int64) {
	o.StateTimeMs = v
}

// GetMirrorStateTransitionErrors returns the MirrorStateTransitionErrors field value if set, zero value otherwise.
func (o *ListMirrorTopicsResponseData) GetMirrorStateTransitionErrors() []LinkTaskError {
	if o == nil || o.MirrorStateTransitionErrors == nil {
		var ret []LinkTaskError
		return ret
	}
	return *o.MirrorStateTransitionErrors
}

// GetMirrorStateTransitionErrorsOk returns a tuple with the MirrorStateTransitionErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMirrorTopicsResponseData) GetMirrorStateTransitionErrorsOk() (*[]LinkTaskError, bool) {
	if o == nil || o.MirrorStateTransitionErrors == nil {
		return nil, false
	}
	return o.MirrorStateTransitionErrors, true
}

// HasMirrorStateTransitionErrors returns a boolean if a field has been set.
func (o *ListMirrorTopicsResponseData) HasMirrorStateTransitionErrors() bool {
	if o != nil && o.MirrorStateTransitionErrors != nil {
		return true
	}

	return false
}

// SetMirrorStateTransitionErrors gets a reference to the given []LinkTaskError and assigns it to the MirrorStateTransitionErrors field.
func (o *ListMirrorTopicsResponseData) SetMirrorStateTransitionErrors(v []LinkTaskError) {
	o.MirrorStateTransitionErrors = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ListMirrorTopicsResponseData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.LinkName)
	o.recurseRedact(&o.MirrorTopicName)
	o.recurseRedact(&o.SourceTopicName)
	o.recurseRedact(&o.NumPartitions)
	o.recurseRedact(&o.MirrorLags)
	o.recurseRedact(&o.MirrorStatus)
	o.recurseRedact(&o.StateTimeMs)
	o.recurseRedact(o.MirrorStateTransitionErrors)
}

func (o *ListMirrorTopicsResponseData) recurseRedact(v interface{}) {
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

func (o ListMirrorTopicsResponseData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ListMirrorTopicsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
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
	if o.MirrorStateTransitionErrors != nil {
		toSerialize["mirror_state_transition_errors"] = o.MirrorStateTransitionErrors
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableListMirrorTopicsResponseData struct {
	value *ListMirrorTopicsResponseData
	isSet bool
}

func (v NullableListMirrorTopicsResponseData) Get() *ListMirrorTopicsResponseData {
	return v.value
}

func (v *NullableListMirrorTopicsResponseData) Set(val *ListMirrorTopicsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListMirrorTopicsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListMirrorTopicsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMirrorTopicsResponseData(val *ListMirrorTopicsResponseData) *NullableListMirrorTopicsResponseData {
	return &NullableListMirrorTopicsResponseData{value: val, isSet: true}
}

func (v NullableListMirrorTopicsResponseData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableListMirrorTopicsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
