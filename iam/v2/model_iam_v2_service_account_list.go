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
Identity & Access Management API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: paas-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// IamV2ServiceAccountList `ServiceAccount` objects are typically used to represent applications and other non-human principals that may access your Confluent resources.  The API allows you to create, retrieve, update, and delete individual service accounts, as well as list all your service accounts.   Related guide: [Service Accounts in Confluent Cloud](https://docs.confluent.io/cloud/current/access-management/service-account.html).  ## The Service Accounts Model <SchemaDefinition schemaRef=\"#/components/schemas/iam.v2.ServiceAccount\" />  ## Quotas and Limits This resource is subject to the following quotas:  | Quota | Description | | --- | --- | | `service_accounts_per_org` | Service Accounts in one Confluent Cloud organization |
type IamV2ServiceAccountList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []IamV2ServiceAccount `json:"data"`
}

// NewIamV2ServiceAccountList instantiates a new IamV2ServiceAccountList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2ServiceAccountList(apiVersion string, kind string, metadata ListMeta, data []IamV2ServiceAccount) *IamV2ServiceAccountList {
	this := IamV2ServiceAccountList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewIamV2ServiceAccountListWithDefaults instantiates a new IamV2ServiceAccountList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2ServiceAccountListWithDefaults() *IamV2ServiceAccountList {
	this := IamV2ServiceAccountList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *IamV2ServiceAccountList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *IamV2ServiceAccountList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *IamV2ServiceAccountList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *IamV2ServiceAccountList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *IamV2ServiceAccountList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *IamV2ServiceAccountList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *IamV2ServiceAccountList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *IamV2ServiceAccountList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *IamV2ServiceAccountList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *IamV2ServiceAccountList) GetData() []IamV2ServiceAccount {
	if o == nil {
		var ret []IamV2ServiceAccount
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IamV2ServiceAccountList) GetDataOk() (*[]IamV2ServiceAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *IamV2ServiceAccountList) SetData(v []IamV2ServiceAccount) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2ServiceAccountList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *IamV2ServiceAccountList) recurseRedact(v interface{}) {
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

func (o IamV2ServiceAccountList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o IamV2ServiceAccountList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_version"] = o.ApiVersion
	}
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

type NullableIamV2ServiceAccountList struct {
	value *IamV2ServiceAccountList
	isSet bool
}

func (v NullableIamV2ServiceAccountList) Get() *IamV2ServiceAccountList {
	return v.value
}

func (v *NullableIamV2ServiceAccountList) Set(val *IamV2ServiceAccountList) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2ServiceAccountList) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2ServiceAccountList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2ServiceAccountList(val *IamV2ServiceAccountList) *NullableIamV2ServiceAccountList {
	return &NullableIamV2ServiceAccountList{value: val, isSet: true}
}

func (v NullableIamV2ServiceAccountList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamV2ServiceAccountList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
