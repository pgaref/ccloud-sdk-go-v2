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

// SqlV1alpha1StatementStatus The status of the Statement
type SqlV1alpha1StatementStatus struct {
	// The lifecycle phase of the submitted SQL statement: PENDING: SQL statement is pending execution; RUNNING: SQL statement execution is in progress; COMPLETED: SQL statement is completed; DELETING: SQL statement deletion is in progress; FAILING: SQL statement is failing; FAILED: SQL statement execution has failed; 
	Phase string `json:"phase,omitempty"`
	ResultSchema *SqlV1alpha1ResultSchema `json:"result_schema,omitempty"`
	// Description of a SQL statement phase.
	Detail *string `json:"detail,omitempty"`
}

// NewSqlV1alpha1StatementStatus instantiates a new SqlV1alpha1StatementStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlV1alpha1StatementStatus(phase string) *SqlV1alpha1StatementStatus {
	this := SqlV1alpha1StatementStatus{}
	this.Phase = phase
	return &this
}

// NewSqlV1alpha1StatementStatusWithDefaults instantiates a new SqlV1alpha1StatementStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlV1alpha1StatementStatusWithDefaults() *SqlV1alpha1StatementStatus {
	this := SqlV1alpha1StatementStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *SqlV1alpha1StatementStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1StatementStatus) GetPhaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *SqlV1alpha1StatementStatus) SetPhase(v string) {
	o.Phase = v
}

// GetResultSchema returns the ResultSchema field value if set, zero value otherwise.
func (o *SqlV1alpha1StatementStatus) GetResultSchema() SqlV1alpha1ResultSchema {
	if o == nil || o.ResultSchema == nil {
		var ret SqlV1alpha1ResultSchema
		return ret
	}
	return *o.ResultSchema
}

// GetResultSchemaOk returns a tuple with the ResultSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1StatementStatus) GetResultSchemaOk() (*SqlV1alpha1ResultSchema, bool) {
	if o == nil || o.ResultSchema == nil {
		return nil, false
	}
	return o.ResultSchema, true
}

// HasResultSchema returns a boolean if a field has been set.
func (o *SqlV1alpha1StatementStatus) HasResultSchema() bool {
	if o != nil && o.ResultSchema != nil {
		return true
	}

	return false
}

// SetResultSchema gets a reference to the given SqlV1alpha1ResultSchema and assigns it to the ResultSchema field.
func (o *SqlV1alpha1StatementStatus) SetResultSchema(v SqlV1alpha1ResultSchema) {
	o.ResultSchema = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *SqlV1alpha1StatementStatus) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1StatementStatus) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *SqlV1alpha1StatementStatus) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *SqlV1alpha1StatementStatus) SetDetail(v string) {
	o.Detail = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SqlV1alpha1StatementStatus) Redact() {
    o.recurseRedact(&o.Phase)
    o.recurseRedact(o.ResultSchema)
    o.recurseRedact(o.Detail)
}

func (o *SqlV1alpha1StatementStatus) recurseRedact(v interface{}) {
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

func (o SqlV1alpha1StatementStatus) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SqlV1alpha1StatementStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	if o.ResultSchema != nil {
		toSerialize["result_schema"] = o.ResultSchema
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableSqlV1alpha1StatementStatus struct {
	value *SqlV1alpha1StatementStatus
	isSet bool
}

func (v NullableSqlV1alpha1StatementStatus) Get() *SqlV1alpha1StatementStatus {
	return v.value
}

func (v *NullableSqlV1alpha1StatementStatus) Set(val *SqlV1alpha1StatementStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlV1alpha1StatementStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlV1alpha1StatementStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlV1alpha1StatementStatus(val *SqlV1alpha1StatementStatus) *NullableSqlV1alpha1StatementStatus {
	return &NullableSqlV1alpha1StatementStatus{value: val, isSet: true}
}

func (v NullableSqlV1alpha1StatementStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSqlV1alpha1StatementStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


