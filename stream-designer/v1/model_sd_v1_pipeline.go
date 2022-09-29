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
Stream Designer API

# Introduction  Stream Designer API provides resources/API for defining stream processing pipelines. Each pipeline describes a set of stream processing components, including connectors, topics, streams, tables, queries and schemas. The components in a pipeline need not exist as Confluent Cloud resources until the pipeline is activated.  This API defines operations to create, list, modify, manage and delete pipelines. 

API version: 0.0.1-alpha0
Contact: stream-designer@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// SdV1Pipeline `Pipeline` objects represent information about a user-defined pipeline of Confluent Cloud components. The pipeline's content is available separately.  The API allows you to create, retrieve, update, and delete your pipelines, as well as list all of your pipelines for the particular environment and Kafka cluster.   Related guide: [Pipelines in Confluent Cloud](https://docs.confluent.io/cloud/current/api.html).  ## The Pipelines Model <SchemaDefinition schemaRef=\"#/components/schemas/sd.v1.Pipeline\" />  ## Quotas and Limits This resource is subject to the following quotas:  | Quota | Description | | --- | --- | | `pipelines_per_org` | Pipelines in one Confluent Cloud organization | | `pipelines_per_cluster` | Pipelines in one Confluent Cloud Kafka cluster |
type SdV1Pipeline struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id *string `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	Spec *SdV1PipelineSpec `json:"spec,omitempty"`
	Status *SdV1PipelineStatus `json:"status,omitempty"`
}

// NewSdV1Pipeline instantiates a new SdV1Pipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdV1Pipeline() *SdV1Pipeline {
	this := SdV1Pipeline{}
	return &this
}

// NewSdV1PipelineWithDefaults instantiates a new SdV1Pipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdV1PipelineWithDefaults() *SdV1Pipeline {
	this := SdV1Pipeline{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *SdV1Pipeline) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdV1Pipeline) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *SdV1Pipeline) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *SdV1Pipeline) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SdV1Pipeline) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdV1Pipeline) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SdV1Pipeline) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SdV1Pipeline) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SdV1Pipeline) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdV1Pipeline) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SdV1Pipeline) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SdV1Pipeline) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SdV1Pipeline) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdV1Pipeline) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SdV1Pipeline) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *SdV1Pipeline) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *SdV1Pipeline) GetSpec() SdV1PipelineSpec {
	if o == nil || o.Spec == nil {
		var ret SdV1PipelineSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdV1Pipeline) GetSpecOk() (*SdV1PipelineSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *SdV1Pipeline) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given SdV1PipelineSpec and assigns it to the Spec field.
func (o *SdV1Pipeline) SetSpec(v SdV1PipelineSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SdV1Pipeline) GetStatus() SdV1PipelineStatus {
	if o == nil || o.Status == nil {
		var ret SdV1PipelineStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdV1Pipeline) GetStatusOk() (*SdV1PipelineStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SdV1Pipeline) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SdV1PipelineStatus and assigns it to the Status field.
func (o *SdV1Pipeline) SetStatus(v SdV1PipelineStatus) {
	o.Status = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SdV1Pipeline) Redact() {
    o.recurseRedact(o.ApiVersion)
    o.recurseRedact(o.Kind)
    o.recurseRedact(o.Id)
    o.recurseRedact(o.Metadata)
    o.recurseRedact(o.Spec)
    o.recurseRedact(o.Status)
}

func (o *SdV1Pipeline) recurseRedact(v interface{}) {
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

func (o SdV1Pipeline) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SdV1Pipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableSdV1Pipeline struct {
	value *SdV1Pipeline
	isSet bool
}

func (v NullableSdV1Pipeline) Get() *SdV1Pipeline {
	return v.value
}

func (v *NullableSdV1Pipeline) Set(val *SdV1Pipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableSdV1Pipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableSdV1Pipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdV1Pipeline(val *SdV1Pipeline) *NullableSdV1Pipeline {
	return &NullableSdV1Pipeline{value: val, isSet: true}
}

func (v NullableSdV1Pipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdV1Pipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


