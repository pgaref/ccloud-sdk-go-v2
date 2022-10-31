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

// NetworkingV1NetworkSpec The desired state of the Network
type NetworkingV1NetworkSpec struct {
	// The name of the network
	DisplayName *string `json:"display_name,omitempty"`
	// The cloud service provider in which the network exists.
	Cloud *string `json:"cloud,omitempty"`
	// The cloud service provider region in which the network exists.
	Region *string `json:"region,omitempty"`
	ConnectionTypes *NetworkingV1ConnectionTypes `json:"connection_types,omitempty"`
	// The IPv4 [CIDR block](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) to used for this network. Must be `/16`. Required for VPC peering and AWS TransitGateway. 
	Cidr *string `json:"cidr,omitempty"`
	// The 3 availability zones for this network. They can optionally be specified for AWS networks used with PrivateLink, for GCP networks used with Private Service Connect, and for AWS and GCP networks used with Peering. Otherwise, they are automatically chosen by Confluent Cloud.  On AWS, zones are AWS [AZ IDs](https://docs.aws.amazon.com/ram/latest/userguide/working-with-az-ids.html)  (e.g. use1-az3)  On GCP, zones are GCP [zones](https://cloud.google.com/compute/docs/regions-zones)  (e.g. us-central1-c).  On Azure, zones are Confluent-chosen names (e.g. 1, 2, 3) since Azure does not  have universal zone identifiers. 
	Zones *[]string `json:"zones,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
}

// NewNetworkingV1NetworkSpec instantiates a new NetworkingV1NetworkSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1NetworkSpec() *NetworkingV1NetworkSpec {
	this := NetworkingV1NetworkSpec{}
	return &this
}

// NewNetworkingV1NetworkSpecWithDefaults instantiates a new NetworkingV1NetworkSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1NetworkSpecWithDefaults() *NetworkingV1NetworkSpec {
	this := NetworkingV1NetworkSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1NetworkSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *NetworkingV1NetworkSpec) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *NetworkingV1NetworkSpec) SetRegion(v string) {
	o.Region = &v
}

// GetConnectionTypes returns the ConnectionTypes field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetConnectionTypes() NetworkingV1ConnectionTypes {
	if o == nil || o.ConnectionTypes == nil {
		var ret NetworkingV1ConnectionTypes
		return ret
	}
	return *o.ConnectionTypes
}

// GetConnectionTypesOk returns a tuple with the ConnectionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetConnectionTypesOk() (*NetworkingV1ConnectionTypes, bool) {
	if o == nil || o.ConnectionTypes == nil {
		return nil, false
	}
	return o.ConnectionTypes, true
}

// HasConnectionTypes returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasConnectionTypes() bool {
	if o != nil && o.ConnectionTypes != nil {
		return true
	}

	return false
}

// SetConnectionTypes gets a reference to the given NetworkingV1ConnectionTypes and assigns it to the ConnectionTypes field.
func (o *NetworkingV1NetworkSpec) SetConnectionTypes(v NetworkingV1ConnectionTypes) {
	o.ConnectionTypes = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *NetworkingV1NetworkSpec) SetCidr(v string) {
	o.Cidr = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetZones() []string {
	if o == nil || o.Zones == nil {
		var ret []string
		return ret
	}
	return *o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetZonesOk() (*[]string, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasZones() bool {
	if o != nil && o.Zones != nil {
		return true
	}

	return false
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *NetworkingV1NetworkSpec) SetZones(v []string) {
	o.Zones = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1NetworkSpec) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkSpec) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1NetworkSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1NetworkSpec) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1NetworkSpec) Redact() {
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.Cloud)
    o.recurseRedact(o.Region)
    o.recurseRedact(o.ConnectionTypes)
    o.recurseRedact(o.Cidr)
    o.recurseRedact(o.Zones)
    o.recurseRedact(o.Environment)
}

func (o *NetworkingV1NetworkSpec) recurseRedact(v interface{}) {
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

func (o NetworkingV1NetworkSpec) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1NetworkSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.ConnectionTypes != nil {
		toSerialize["connection_types"] = o.ConnectionTypes
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1NetworkSpec struct {
	value *NetworkingV1NetworkSpec
	isSet bool
}

func (v NullableNetworkingV1NetworkSpec) Get() *NetworkingV1NetworkSpec {
	return v.value
}

func (v *NullableNetworkingV1NetworkSpec) Set(val *NetworkingV1NetworkSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1NetworkSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1NetworkSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1NetworkSpec(val *NetworkingV1NetworkSpec) *NullableNetworkingV1NetworkSpec {
	return &NullableNetworkingV1NetworkSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1NetworkSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1NetworkSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


