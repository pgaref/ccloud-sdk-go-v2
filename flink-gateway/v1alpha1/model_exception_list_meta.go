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
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ExceptionListMeta ListMeta describes metadata that resource collections may have
type ExceptionListMeta struct {
	// Self is a Uniform Resource Locator (URL) at which an object can be addressed. This URL encodes the service location, API version, and other particulars necessary to locate the resource at a point in time
	Self *string `json:"self,omitempty"`
}

// NewExceptionListMeta instantiates a new ExceptionListMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionListMeta() *ExceptionListMeta {
	this := ExceptionListMeta{}
	return &this
}

// NewExceptionListMetaWithDefaults instantiates a new ExceptionListMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionListMetaWithDefaults() *ExceptionListMeta {
	this := ExceptionListMeta{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ExceptionListMeta) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionListMeta) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ExceptionListMeta) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ExceptionListMeta) SetSelf(v string) {
	o.Self = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ExceptionListMeta) Redact() {
	o.recurseRedact(o.Self)
}

func (o *ExceptionListMeta) recurseRedact(v interface{}) {
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

func (o ExceptionListMeta) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ExceptionListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableExceptionListMeta struct {
	value *ExceptionListMeta
	isSet bool
}

func (v NullableExceptionListMeta) Get() *ExceptionListMeta {
	return v.value
}

func (v *NullableExceptionListMeta) Set(val *ExceptionListMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionListMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionListMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionListMeta(val *ExceptionListMeta) *NullableExceptionListMeta {
	return &NullableExceptionListMeta{value: val, isSet: true}
}

func (v NullableExceptionListMeta) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableExceptionListMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
