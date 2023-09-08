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

// CreateLinkRequestData struct for CreateLinkRequestData
type CreateLinkRequestData struct {
	SourceClusterId      *string `json:"source_cluster_id,omitempty"`
	DestinationClusterId *string `json:"destination_cluster_id,omitempty"`
	// The expected remote cluster ID.
	RemoteClusterId *string `json:"remote_cluster_id,omitempty"`
	// The expected cluster link ID. Can be provided when creating the second side of a bidirectional link for validating the link ID is as expected. If it's not provided, it's inferred from the remote cluster.
	ClusterLinkId *string       `json:"cluster_link_id,omitempty"`
	Configs       *[]ConfigData `json:"configs,omitempty"`
}

// NewCreateLinkRequestData instantiates a new CreateLinkRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateLinkRequestData() *CreateLinkRequestData {
	this := CreateLinkRequestData{}
	return &this
}

// NewCreateLinkRequestDataWithDefaults instantiates a new CreateLinkRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateLinkRequestDataWithDefaults() *CreateLinkRequestData {
	this := CreateLinkRequestData{}
	return &this
}

// GetSourceClusterId returns the SourceClusterId field value if set, zero value otherwise.
func (o *CreateLinkRequestData) GetSourceClusterId() string {
	if o == nil || o.SourceClusterId == nil {
		var ret string
		return ret
	}
	return *o.SourceClusterId
}

// GetSourceClusterIdOk returns a tuple with the SourceClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequestData) GetSourceClusterIdOk() (*string, bool) {
	if o == nil || o.SourceClusterId == nil {
		return nil, false
	}
	return o.SourceClusterId, true
}

// HasSourceClusterId returns a boolean if a field has been set.
func (o *CreateLinkRequestData) HasSourceClusterId() bool {
	if o != nil && o.SourceClusterId != nil {
		return true
	}

	return false
}

// SetSourceClusterId gets a reference to the given string and assigns it to the SourceClusterId field.
func (o *CreateLinkRequestData) SetSourceClusterId(v string) {
	o.SourceClusterId = &v
}

// GetDestinationClusterId returns the DestinationClusterId field value if set, zero value otherwise.
func (o *CreateLinkRequestData) GetDestinationClusterId() string {
	if o == nil || o.DestinationClusterId == nil {
		var ret string
		return ret
	}
	return *o.DestinationClusterId
}

// GetDestinationClusterIdOk returns a tuple with the DestinationClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequestData) GetDestinationClusterIdOk() (*string, bool) {
	if o == nil || o.DestinationClusterId == nil {
		return nil, false
	}
	return o.DestinationClusterId, true
}

// HasDestinationClusterId returns a boolean if a field has been set.
func (o *CreateLinkRequestData) HasDestinationClusterId() bool {
	if o != nil && o.DestinationClusterId != nil {
		return true
	}

	return false
}

// SetDestinationClusterId gets a reference to the given string and assigns it to the DestinationClusterId field.
func (o *CreateLinkRequestData) SetDestinationClusterId(v string) {
	o.DestinationClusterId = &v
}

// GetRemoteClusterId returns the RemoteClusterId field value if set, zero value otherwise.
func (o *CreateLinkRequestData) GetRemoteClusterId() string {
	if o == nil || o.RemoteClusterId == nil {
		var ret string
		return ret
	}
	return *o.RemoteClusterId
}

// GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequestData) GetRemoteClusterIdOk() (*string, bool) {
	if o == nil || o.RemoteClusterId == nil {
		return nil, false
	}
	return o.RemoteClusterId, true
}

// HasRemoteClusterId returns a boolean if a field has been set.
func (o *CreateLinkRequestData) HasRemoteClusterId() bool {
	if o != nil && o.RemoteClusterId != nil {
		return true
	}

	return false
}

// SetRemoteClusterId gets a reference to the given string and assigns it to the RemoteClusterId field.
func (o *CreateLinkRequestData) SetRemoteClusterId(v string) {
	o.RemoteClusterId = &v
}

// GetClusterLinkId returns the ClusterLinkId field value if set, zero value otherwise.
func (o *CreateLinkRequestData) GetClusterLinkId() string {
	if o == nil || o.ClusterLinkId == nil {
		var ret string
		return ret
	}
	return *o.ClusterLinkId
}

// GetClusterLinkIdOk returns a tuple with the ClusterLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequestData) GetClusterLinkIdOk() (*string, bool) {
	if o == nil || o.ClusterLinkId == nil {
		return nil, false
	}
	return o.ClusterLinkId, true
}

// HasClusterLinkId returns a boolean if a field has been set.
func (o *CreateLinkRequestData) HasClusterLinkId() bool {
	if o != nil && o.ClusterLinkId != nil {
		return true
	}

	return false
}

// SetClusterLinkId gets a reference to the given string and assigns it to the ClusterLinkId field.
func (o *CreateLinkRequestData) SetClusterLinkId(v string) {
	o.ClusterLinkId = &v
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *CreateLinkRequestData) GetConfigs() []ConfigData {
	if o == nil || o.Configs == nil {
		var ret []ConfigData
		return ret
	}
	return *o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateLinkRequestData) GetConfigsOk() (*[]ConfigData, bool) {
	if o == nil || o.Configs == nil {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *CreateLinkRequestData) HasConfigs() bool {
	if o != nil && o.Configs != nil {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []ConfigData and assigns it to the Configs field.
func (o *CreateLinkRequestData) SetConfigs(v []ConfigData) {
	o.Configs = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CreateLinkRequestData) Redact() {
	o.recurseRedact(o.SourceClusterId)
	o.recurseRedact(o.DestinationClusterId)
	o.recurseRedact(o.RemoteClusterId)
	o.recurseRedact(o.ClusterLinkId)
	o.recurseRedact(o.Configs)
}

func (o *CreateLinkRequestData) recurseRedact(v interface{}) {
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

func (o CreateLinkRequestData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CreateLinkRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceClusterId != nil {
		toSerialize["source_cluster_id"] = o.SourceClusterId
	}
	if o.DestinationClusterId != nil {
		toSerialize["destination_cluster_id"] = o.DestinationClusterId
	}
	if o.RemoteClusterId != nil {
		toSerialize["remote_cluster_id"] = o.RemoteClusterId
	}
	if o.ClusterLinkId != nil {
		toSerialize["cluster_link_id"] = o.ClusterLinkId
	}
	if o.Configs != nil {
		toSerialize["configs"] = o.Configs
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableCreateLinkRequestData struct {
	value *CreateLinkRequestData
	isSet bool
}

func (v NullableCreateLinkRequestData) Get() *CreateLinkRequestData {
	return v.value
}

func (v *NullableCreateLinkRequestData) Set(val *CreateLinkRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateLinkRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateLinkRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateLinkRequestData(val *CreateLinkRequestData) *NullableCreateLinkRequestData {
	return &NullableCreateLinkRequestData{value: val, isSet: true}
}

func (v NullableCreateLinkRequestData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCreateLinkRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
