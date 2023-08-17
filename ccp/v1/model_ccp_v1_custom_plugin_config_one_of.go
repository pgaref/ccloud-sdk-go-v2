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
Custom Plugin Management API

This is Custom Plugin Management API.

API version: 1.0.0
Contact: compute-platform-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CcpV1CustomPluginConfigOneOf - struct for CcpV1CustomPluginConfigOneOf
type CcpV1CustomPluginConfigOneOf struct {
	CcpV1CustomPluginConnect *CcpV1CustomPluginConnect
}

// CcpV1CustomPluginConnectAsCcpV1CustomPluginConfigOneOf is a convenience function that returns CcpV1CustomPluginConnect wrapped in CcpV1CustomPluginConfigOneOf
func CcpV1CustomPluginConnectAsCcpV1CustomPluginConfigOneOf(v *CcpV1CustomPluginConnect) CcpV1CustomPluginConfigOneOf {
	return CcpV1CustomPluginConfigOneOf{CcpV1CustomPluginConnect: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CcpV1CustomPluginConfigOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'CONNECT'
	if jsonDict["plugin_type"] == "CONNECT" {
		// try to unmarshal JSON data into CcpV1CustomPluginConnect
		err = json.Unmarshal(data, &dst.CcpV1CustomPluginConnect)
		if err == nil {
			return nil // data stored in dst.CcpV1CustomPluginConnect, return on the first match
		} else {
			dst.CcpV1CustomPluginConnect = nil
			return fmt.Errorf("Failed to unmarshal CcpV1CustomPluginConfigOneOf as CcpV1CustomPluginConnect: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ccp.v1.CustomPlugin.Connect'
	if jsonDict["plugin_type"] == "ccp.v1.CustomPlugin.Connect" {
		// try to unmarshal JSON data into CcpV1CustomPluginConnect
		err = json.Unmarshal(data, &dst.CcpV1CustomPluginConnect)
		if err == nil {
			return nil // data stored in dst.CcpV1CustomPluginConnect, return on the first match
		} else {
			dst.CcpV1CustomPluginConnect = nil
			return fmt.Errorf("Failed to unmarshal CcpV1CustomPluginConfigOneOf as CcpV1CustomPluginConnect: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CcpV1CustomPluginConfigOneOf) MarshalJSON() ([]byte, error) {
	if src.CcpV1CustomPluginConnect != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.CcpV1CustomPluginConnect)
		return buffer.Bytes(), err
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CcpV1CustomPluginConfigOneOf) GetActualInstance() interface{} {
	if obj.CcpV1CustomPluginConnect != nil {
		return obj.CcpV1CustomPluginConnect
	}

	// all schemas are nil
	return nil
}

type NullableCcpV1CustomPluginConfigOneOf struct {
	value *CcpV1CustomPluginConfigOneOf
	isSet bool
}

func (v NullableCcpV1CustomPluginConfigOneOf) Get() *CcpV1CustomPluginConfigOneOf {
	return v.value
}

func (v *NullableCcpV1CustomPluginConfigOneOf) Set(val *CcpV1CustomPluginConfigOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCcpV1CustomPluginConfigOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCcpV1CustomPluginConfigOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCcpV1CustomPluginConfigOneOf(val *CcpV1CustomPluginConfigOneOf) *NullableCcpV1CustomPluginConfigOneOf {
	return &NullableCcpV1CustomPluginConfigOneOf{value: val, isSet: true}
}

func (v NullableCcpV1CustomPluginConfigOneOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCcpV1CustomPluginConfigOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
