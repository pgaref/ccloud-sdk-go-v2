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
Flink Artifact Management API

This is the Flink Artifact Management API.

API version: 0.0.1
Contact: flink-runtime@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ArtifactV1PresignedUrlRequest Request schema of the presigned upload URL.
type ArtifactV1PresignedUrlRequest struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// Archive format of the Flink Artifact.
	ContentFormat *string `json:"content_format,omitempty"`
	// Cloud provider where the Flink Artifact archive is uploaded.
	Cloud *string `json:"cloud,omitempty"`
	// The Cloud provider region the Flink Artifact archive is uploaded.
	Region *string `json:"region,omitempty"`
}

// NewArtifactV1PresignedUrlRequest instantiates a new ArtifactV1PresignedUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactV1PresignedUrlRequest() *ArtifactV1PresignedUrlRequest {
	this := ArtifactV1PresignedUrlRequest{}
	return &this
}

// NewArtifactV1PresignedUrlRequestWithDefaults instantiates a new ArtifactV1PresignedUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactV1PresignedUrlRequestWithDefaults() *ArtifactV1PresignedUrlRequest {
	this := ArtifactV1PresignedUrlRequest{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ArtifactV1PresignedUrlRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ArtifactV1PresignedUrlRequest) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArtifactV1PresignedUrlRequest) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *ArtifactV1PresignedUrlRequest) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetContentFormat returns the ContentFormat field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetContentFormat() string {
	if o == nil || o.ContentFormat == nil {
		var ret string
		return ret
	}
	return *o.ContentFormat
}

// GetContentFormatOk returns a tuple with the ContentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetContentFormatOk() (*string, bool) {
	if o == nil || o.ContentFormat == nil {
		return nil, false
	}
	return o.ContentFormat, true
}

// HasContentFormat returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasContentFormat() bool {
	if o != nil && o.ContentFormat != nil {
		return true
	}

	return false
}

// SetContentFormat gets a reference to the given string and assigns it to the ContentFormat field.
func (o *ArtifactV1PresignedUrlRequest) SetContentFormat(v string) {
	o.ContentFormat = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *ArtifactV1PresignedUrlRequest) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ArtifactV1PresignedUrlRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1PresignedUrlRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ArtifactV1PresignedUrlRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ArtifactV1PresignedUrlRequest) SetRegion(v string) {
	o.Region = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ArtifactV1PresignedUrlRequest) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.ContentFormat)
	o.recurseRedact(o.Cloud)
	o.recurseRedact(o.Region)
}

func (o *ArtifactV1PresignedUrlRequest) recurseRedact(v interface{}) {
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

func (o ArtifactV1PresignedUrlRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ArtifactV1PresignedUrlRequest) MarshalJSON() ([]byte, error) {
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
	if o.ContentFormat != nil {
		toSerialize["content_format"] = o.ContentFormat
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableArtifactV1PresignedUrlRequest struct {
	value *ArtifactV1PresignedUrlRequest
	isSet bool
}

func (v NullableArtifactV1PresignedUrlRequest) Get() *ArtifactV1PresignedUrlRequest {
	return v.value
}

func (v *NullableArtifactV1PresignedUrlRequest) Set(val *ArtifactV1PresignedUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactV1PresignedUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactV1PresignedUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactV1PresignedUrlRequest(val *ArtifactV1PresignedUrlRequest) *NullableArtifactV1PresignedUrlRequest {
	return &NullableArtifactV1PresignedUrlRequest{value: val, isSet: true}
}

func (v NullableArtifactV1PresignedUrlRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableArtifactV1PresignedUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
