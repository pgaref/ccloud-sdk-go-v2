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
Network API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// NetworkingV1PrivateLinkAccessList Add or remove access to PrivateLink endpoints by AWS account, Azure subscription and GCP project ID.  Related guide: [Private Links Overview](https://docs.confluent.io/cloud/current/networking/private-links/index.html).  ## The Private Link Accesses Model <SchemaDefinition schemaRef=\"#/components/schemas/networking.v1.PrivateLinkAccess\" />  ## Quotas and Limits This resource is subject to the following quotas:  | Quota | Description | | --- | --- | | `private_link_accounts_per_network` | Number of AWS accounts per network | | `private_link_subscriptions_per_network` | Number of Azure subscriptions per network | | `private_service_connect_projects_per_network` | Number of GCP projects per network |
type NetworkingV1PrivateLinkAccessList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind string `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []NetworkingV1PrivateLinkAccess `json:"data"`
}

// NewNetworkingV1PrivateLinkAccessList instantiates a new NetworkingV1PrivateLinkAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1PrivateLinkAccessList(apiVersion string, kind string, metadata ListMeta, data []NetworkingV1PrivateLinkAccess) *NetworkingV1PrivateLinkAccessList {
	this := NetworkingV1PrivateLinkAccessList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewNetworkingV1PrivateLinkAccessListWithDefaults instantiates a new NetworkingV1PrivateLinkAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1PrivateLinkAccessListWithDefaults() *NetworkingV1PrivateLinkAccessList {
	this := NetworkingV1PrivateLinkAccessList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *NetworkingV1PrivateLinkAccessList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAccessList) GetApiVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *NetworkingV1PrivateLinkAccessList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *NetworkingV1PrivateLinkAccessList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAccessList) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1PrivateLinkAccessList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *NetworkingV1PrivateLinkAccessList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAccessList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *NetworkingV1PrivateLinkAccessList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *NetworkingV1PrivateLinkAccessList) GetData() []NetworkingV1PrivateLinkAccess {
	if o == nil {
		var ret []NetworkingV1PrivateLinkAccess
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAccessList) GetDataOk() (*[]NetworkingV1PrivateLinkAccess, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *NetworkingV1PrivateLinkAccessList) SetData(v []NetworkingV1PrivateLinkAccess) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1PrivateLinkAccessList) Redact() {
    o.recurseRedact(&o.ApiVersion)
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Metadata)
    o.recurseRedact(&o.Data)
}

func (o *NetworkingV1PrivateLinkAccessList) recurseRedact(v interface{}) {
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

func (o NetworkingV1PrivateLinkAccessList) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1PrivateLinkAccessList) MarshalJSON() ([]byte, error) {
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

type NullableNetworkingV1PrivateLinkAccessList struct {
	value *NetworkingV1PrivateLinkAccessList
	isSet bool
}

func (v NullableNetworkingV1PrivateLinkAccessList) Get() *NetworkingV1PrivateLinkAccessList {
	return v.value
}

func (v *NullableNetworkingV1PrivateLinkAccessList) Set(val *NetworkingV1PrivateLinkAccessList) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PrivateLinkAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PrivateLinkAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PrivateLinkAccessList(val *NetworkingV1PrivateLinkAccessList) *NullableNetworkingV1PrivateLinkAccessList {
	return &NullableNetworkingV1PrivateLinkAccessList{value: val, isSet: true}
}

func (v NullableNetworkingV1PrivateLinkAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1PrivateLinkAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


