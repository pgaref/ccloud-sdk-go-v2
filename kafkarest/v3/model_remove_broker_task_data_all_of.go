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

// RemoveBrokerTaskDataAllOf struct for RemoveBrokerTaskDataAllOf
type RemoveBrokerTaskDataAllOf struct {
	ClusterId                    string         `json:"cluster_id,omitempty"`
	BrokerId                     int32          `json:"broker_id,omitempty"`
	ShutdownScheduled            bool           `json:"shutdown_scheduled,omitempty"`
	BrokerReplicaExclusionStatus string         `json:"broker_replica_exclusion_status,omitempty"`
	PartitionReassignmentStatus  string         `json:"partition_reassignment_status,omitempty"`
	BrokerShutdownStatus         string         `json:"broker_shutdown_status,omitempty"`
	ErrorCode                    NullableInt32  `json:"error_code,omitempty"`
	ErrorMessage                 NullableString `json:"error_message,omitempty"`
	Broker                       Relationship   `json:"broker,omitempty"`
}

// NewRemoveBrokerTaskDataAllOf instantiates a new RemoveBrokerTaskDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveBrokerTaskDataAllOf(clusterId string, brokerId int32, shutdownScheduled bool, brokerReplicaExclusionStatus string, partitionReassignmentStatus string, brokerShutdownStatus string, broker Relationship) *RemoveBrokerTaskDataAllOf {
	this := RemoveBrokerTaskDataAllOf{}
	this.ClusterId = clusterId
	this.BrokerId = brokerId
	this.ShutdownScheduled = shutdownScheduled
	this.BrokerReplicaExclusionStatus = brokerReplicaExclusionStatus
	this.PartitionReassignmentStatus = partitionReassignmentStatus
	this.BrokerShutdownStatus = brokerShutdownStatus
	this.Broker = broker
	return &this
}

// NewRemoveBrokerTaskDataAllOfWithDefaults instantiates a new RemoveBrokerTaskDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveBrokerTaskDataAllOfWithDefaults() *RemoveBrokerTaskDataAllOf {
	this := RemoveBrokerTaskDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *RemoveBrokerTaskDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *RemoveBrokerTaskDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetBrokerId returns the BrokerId field value
func (o *RemoveBrokerTaskDataAllOf) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *RemoveBrokerTaskDataAllOf) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetShutdownScheduled returns the ShutdownScheduled field value
func (o *RemoveBrokerTaskDataAllOf) GetShutdownScheduled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ShutdownScheduled
}

// GetShutdownScheduledOk returns a tuple with the ShutdownScheduled field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetShutdownScheduledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShutdownScheduled, true
}

// SetShutdownScheduled sets field value
func (o *RemoveBrokerTaskDataAllOf) SetShutdownScheduled(v bool) {
	o.ShutdownScheduled = v
}

// GetBrokerReplicaExclusionStatus returns the BrokerReplicaExclusionStatus field value
func (o *RemoveBrokerTaskDataAllOf) GetBrokerReplicaExclusionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrokerReplicaExclusionStatus
}

// GetBrokerReplicaExclusionStatusOk returns a tuple with the BrokerReplicaExclusionStatus field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetBrokerReplicaExclusionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerReplicaExclusionStatus, true
}

// SetBrokerReplicaExclusionStatus sets field value
func (o *RemoveBrokerTaskDataAllOf) SetBrokerReplicaExclusionStatus(v string) {
	o.BrokerReplicaExclusionStatus = v
}

// GetPartitionReassignmentStatus returns the PartitionReassignmentStatus field value
func (o *RemoveBrokerTaskDataAllOf) GetPartitionReassignmentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartitionReassignmentStatus
}

// GetPartitionReassignmentStatusOk returns a tuple with the PartitionReassignmentStatus field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetPartitionReassignmentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionReassignmentStatus, true
}

// SetPartitionReassignmentStatus sets field value
func (o *RemoveBrokerTaskDataAllOf) SetPartitionReassignmentStatus(v string) {
	o.PartitionReassignmentStatus = v
}

// GetBrokerShutdownStatus returns the BrokerShutdownStatus field value
func (o *RemoveBrokerTaskDataAllOf) GetBrokerShutdownStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrokerShutdownStatus
}

// GetBrokerShutdownStatusOk returns a tuple with the BrokerShutdownStatus field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetBrokerShutdownStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerShutdownStatus, true
}

// SetBrokerShutdownStatus sets field value
func (o *RemoveBrokerTaskDataAllOf) SetBrokerShutdownStatus(v string) {
	o.BrokerShutdownStatus = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoveBrokerTaskDataAllOf) GetErrorCode() int32 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoveBrokerTaskDataAllOf) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *RemoveBrokerTaskDataAllOf) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *RemoveBrokerTaskDataAllOf) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}

// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *RemoveBrokerTaskDataAllOf) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *RemoveBrokerTaskDataAllOf) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoveBrokerTaskDataAllOf) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoveBrokerTaskDataAllOf) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *RemoveBrokerTaskDataAllOf) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *RemoveBrokerTaskDataAllOf) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *RemoveBrokerTaskDataAllOf) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *RemoveBrokerTaskDataAllOf) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetBroker returns the Broker field value
func (o *RemoveBrokerTaskDataAllOf) GetBroker() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value
// and a boolean to check if the value has been set.
func (o *RemoveBrokerTaskDataAllOf) GetBrokerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broker, true
}

// SetBroker sets field value
func (o *RemoveBrokerTaskDataAllOf) SetBroker(v Relationship) {
	o.Broker = v
}

// Redact resets all sensitive fields to their zero value.
func (o *RemoveBrokerTaskDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(&o.ShutdownScheduled)
	o.recurseRedact(&o.BrokerReplicaExclusionStatus)
	o.recurseRedact(&o.PartitionReassignmentStatus)
	o.recurseRedact(&o.BrokerShutdownStatus)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(&o.Broker)
}

func (o *RemoveBrokerTaskDataAllOf) recurseRedact(v interface{}) {
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

func (o RemoveBrokerTaskDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o RemoveBrokerTaskDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["broker_id"] = o.BrokerId
	}
	if true {
		toSerialize["shutdown_scheduled"] = o.ShutdownScheduled
	}
	if true {
		toSerialize["broker_replica_exclusion_status"] = o.BrokerReplicaExclusionStatus
	}
	if true {
		toSerialize["partition_reassignment_status"] = o.PartitionReassignmentStatus
	}
	if true {
		toSerialize["broker_shutdown_status"] = o.BrokerShutdownStatus
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	if true {
		toSerialize["broker"] = o.Broker
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableRemoveBrokerTaskDataAllOf struct {
	value *RemoveBrokerTaskDataAllOf
	isSet bool
}

func (v NullableRemoveBrokerTaskDataAllOf) Get() *RemoveBrokerTaskDataAllOf {
	return v.value
}

func (v *NullableRemoveBrokerTaskDataAllOf) Set(val *RemoveBrokerTaskDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveBrokerTaskDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveBrokerTaskDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveBrokerTaskDataAllOf(val *RemoveBrokerTaskDataAllOf) *NullableRemoveBrokerTaskDataAllOf {
	return &NullableRemoveBrokerTaskDataAllOf{value: val, isSet: true}
}

func (v NullableRemoveBrokerTaskDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableRemoveBrokerTaskDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
