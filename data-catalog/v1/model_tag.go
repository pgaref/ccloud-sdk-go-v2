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
Confluent Data Catalog

REST API for the Data Catalog

API version: 1.0.0
Contact: data-governance@confluent.io
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

// Tag struct for Tag
type Tag struct {
	// The tag name
	TypeName *string `json:"typeName,omitempty"`
	// The tag attributes
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// The internal entity guid
	EntityGuid *string `json:"entityGuid,omitempty"`
	// The entity status
	EntityStatus *string `json:"entityStatus,omitempty"`
	// Whether to propagate the tag
	Propagate *bool `json:"propagate,omitempty"`
	// The validity periods
	ValidityPeriods *[]TimeBoundary `json:"validityPeriods,omitempty"`
	// Whether to remove propagations on entity delete
	RemovePropagationsOnEntityDelete *bool `json:"removePropagationsOnEntityDelete,omitempty"`
	// The entity type
	EntityType *string `json:"entityType,omitempty"`
	// The qualified name of the entity
	EntityName *string `json:"entityName,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag() *Tag {
	this := Tag{}
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *Tag) GetTypeName() string {
	if o == nil || o.TypeName == nil {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetTypeNameOk() (*string, bool) {
	if o == nil || o.TypeName == nil {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *Tag) HasTypeName() bool {
	if o != nil && o.TypeName != nil {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *Tag) SetTypeName(v string) {
	o.TypeName = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Tag) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Tag) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Tag) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetEntityGuid returns the EntityGuid field value if set, zero value otherwise.
func (o *Tag) GetEntityGuid() string {
	if o == nil || o.EntityGuid == nil {
		var ret string
		return ret
	}
	return *o.EntityGuid
}

// GetEntityGuidOk returns a tuple with the EntityGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetEntityGuidOk() (*string, bool) {
	if o == nil || o.EntityGuid == nil {
		return nil, false
	}
	return o.EntityGuid, true
}

// HasEntityGuid returns a boolean if a field has been set.
func (o *Tag) HasEntityGuid() bool {
	if o != nil && o.EntityGuid != nil {
		return true
	}

	return false
}

// SetEntityGuid gets a reference to the given string and assigns it to the EntityGuid field.
func (o *Tag) SetEntityGuid(v string) {
	o.EntityGuid = &v
}

// GetEntityStatus returns the EntityStatus field value if set, zero value otherwise.
func (o *Tag) GetEntityStatus() string {
	if o == nil || o.EntityStatus == nil {
		var ret string
		return ret
	}
	return *o.EntityStatus
}

// GetEntityStatusOk returns a tuple with the EntityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetEntityStatusOk() (*string, bool) {
	if o == nil || o.EntityStatus == nil {
		return nil, false
	}
	return o.EntityStatus, true
}

// HasEntityStatus returns a boolean if a field has been set.
func (o *Tag) HasEntityStatus() bool {
	if o != nil && o.EntityStatus != nil {
		return true
	}

	return false
}

// SetEntityStatus gets a reference to the given string and assigns it to the EntityStatus field.
func (o *Tag) SetEntityStatus(v string) {
	o.EntityStatus = &v
}

// GetPropagate returns the Propagate field value if set, zero value otherwise.
func (o *Tag) GetPropagate() bool {
	if o == nil || o.Propagate == nil {
		var ret bool
		return ret
	}
	return *o.Propagate
}

// GetPropagateOk returns a tuple with the Propagate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetPropagateOk() (*bool, bool) {
	if o == nil || o.Propagate == nil {
		return nil, false
	}
	return o.Propagate, true
}

// HasPropagate returns a boolean if a field has been set.
func (o *Tag) HasPropagate() bool {
	if o != nil && o.Propagate != nil {
		return true
	}

	return false
}

// SetPropagate gets a reference to the given bool and assigns it to the Propagate field.
func (o *Tag) SetPropagate(v bool) {
	o.Propagate = &v
}

// GetValidityPeriods returns the ValidityPeriods field value if set, zero value otherwise.
func (o *Tag) GetValidityPeriods() []TimeBoundary {
	if o == nil || o.ValidityPeriods == nil {
		var ret []TimeBoundary
		return ret
	}
	return *o.ValidityPeriods
}

// GetValidityPeriodsOk returns a tuple with the ValidityPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetValidityPeriodsOk() (*[]TimeBoundary, bool) {
	if o == nil || o.ValidityPeriods == nil {
		return nil, false
	}
	return o.ValidityPeriods, true
}

// HasValidityPeriods returns a boolean if a field has been set.
func (o *Tag) HasValidityPeriods() bool {
	if o != nil && o.ValidityPeriods != nil {
		return true
	}

	return false
}

// SetValidityPeriods gets a reference to the given []TimeBoundary and assigns it to the ValidityPeriods field.
func (o *Tag) SetValidityPeriods(v []TimeBoundary) {
	o.ValidityPeriods = &v
}

// GetRemovePropagationsOnEntityDelete returns the RemovePropagationsOnEntityDelete field value if set, zero value otherwise.
func (o *Tag) GetRemovePropagationsOnEntityDelete() bool {
	if o == nil || o.RemovePropagationsOnEntityDelete == nil {
		var ret bool
		return ret
	}
	return *o.RemovePropagationsOnEntityDelete
}

// GetRemovePropagationsOnEntityDeleteOk returns a tuple with the RemovePropagationsOnEntityDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetRemovePropagationsOnEntityDeleteOk() (*bool, bool) {
	if o == nil || o.RemovePropagationsOnEntityDelete == nil {
		return nil, false
	}
	return o.RemovePropagationsOnEntityDelete, true
}

// HasRemovePropagationsOnEntityDelete returns a boolean if a field has been set.
func (o *Tag) HasRemovePropagationsOnEntityDelete() bool {
	if o != nil && o.RemovePropagationsOnEntityDelete != nil {
		return true
	}

	return false
}

// SetRemovePropagationsOnEntityDelete gets a reference to the given bool and assigns it to the RemovePropagationsOnEntityDelete field.
func (o *Tag) SetRemovePropagationsOnEntityDelete(v bool) {
	o.RemovePropagationsOnEntityDelete = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *Tag) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *Tag) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *Tag) SetEntityType(v string) {
	o.EntityType = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *Tag) GetEntityName() string {
	if o == nil || o.EntityName == nil {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetEntityNameOk() (*string, bool) {
	if o == nil || o.EntityName == nil {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *Tag) HasEntityName() bool {
	if o != nil && o.EntityName != nil {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *Tag) SetEntityName(v string) {
	o.EntityName = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *Tag) Redact() {
	o.recurseRedact(o.TypeName)
	o.recurseRedact(o.Attributes)
	o.recurseRedact(o.EntityGuid)
	o.recurseRedact(o.EntityStatus)
	o.recurseRedact(o.Propagate)
	o.recurseRedact(o.ValidityPeriods)
	o.recurseRedact(o.RemovePropagationsOnEntityDelete)
	o.recurseRedact(o.EntityType)
	o.recurseRedact(o.EntityName)
}

func (o *Tag) recurseRedact(v interface{}) {
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

func (o Tag) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeName != nil {
		toSerialize["typeName"] = o.TypeName
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.EntityGuid != nil {
		toSerialize["entityGuid"] = o.EntityGuid
	}
	if o.EntityStatus != nil {
		toSerialize["entityStatus"] = o.EntityStatus
	}
	if o.Propagate != nil {
		toSerialize["propagate"] = o.Propagate
	}
	if o.ValidityPeriods != nil {
		toSerialize["validityPeriods"] = o.ValidityPeriods
	}
	if o.RemovePropagationsOnEntityDelete != nil {
		toSerialize["removePropagationsOnEntityDelete"] = o.RemovePropagationsOnEntityDelete
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.EntityName != nil {
		toSerialize["entityName"] = o.EntityName
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
