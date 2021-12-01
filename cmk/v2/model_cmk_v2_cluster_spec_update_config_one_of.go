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
Cluster Management for Apache Kafka API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: orchestrator-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
	"fmt"
)

// CmkV2ClusterSpecUpdateConfigOneOf - struct for CmkV2ClusterSpecUpdateConfigOneOf
type CmkV2ClusterSpecUpdateConfigOneOf struct {
	CmkV2Basic *CmkV2Basic
	CmkV2Dedicated *CmkV2Dedicated
	CmkV2Standard *CmkV2Standard
}

// CmkV2BasicAsCmkV2ClusterSpecUpdateConfigOneOf is a convenience function that returns CmkV2Basic wrapped in CmkV2ClusterSpecUpdateConfigOneOf
func CmkV2BasicAsCmkV2ClusterSpecUpdateConfigOneOf(v *CmkV2Basic) CmkV2ClusterSpecUpdateConfigOneOf {
	return CmkV2ClusterSpecUpdateConfigOneOf{ CmkV2Basic: v}
}

// CmkV2DedicatedAsCmkV2ClusterSpecUpdateConfigOneOf is a convenience function that returns CmkV2Dedicated wrapped in CmkV2ClusterSpecUpdateConfigOneOf
func CmkV2DedicatedAsCmkV2ClusterSpecUpdateConfigOneOf(v *CmkV2Dedicated) CmkV2ClusterSpecUpdateConfigOneOf {
	return CmkV2ClusterSpecUpdateConfigOneOf{ CmkV2Dedicated: v}
}

// CmkV2StandardAsCmkV2ClusterSpecUpdateConfigOneOf is a convenience function that returns CmkV2Standard wrapped in CmkV2ClusterSpecUpdateConfigOneOf
func CmkV2StandardAsCmkV2ClusterSpecUpdateConfigOneOf(v *CmkV2Standard) CmkV2ClusterSpecUpdateConfigOneOf {
	return CmkV2ClusterSpecUpdateConfigOneOf{ CmkV2Standard: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CmkV2ClusterSpecUpdateConfigOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'Basic'
	if jsonDict["kind"] == "Basic" {
		// try to unmarshal JSON data into CmkV2Basic
		err = json.Unmarshal(data, &dst.CmkV2Basic)
		if err == nil {
			return nil // data stored in dst.CmkV2Basic, return on the first match
		} else {
			dst.CmkV2Basic = nil
			return fmt.Errorf("Failed to unmarshal CmkV2ClusterSpecUpdateConfigOneOf as CmkV2Basic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Dedicated'
	if jsonDict["kind"] == "Dedicated" {
		// try to unmarshal JSON data into CmkV2Dedicated
		err = json.Unmarshal(data, &dst.CmkV2Dedicated)
		if err == nil {
			return nil // data stored in dst.CmkV2Dedicated, return on the first match
		} else {
			dst.CmkV2Dedicated = nil
			return fmt.Errorf("Failed to unmarshal CmkV2ClusterSpecUpdateConfigOneOf as CmkV2Dedicated: %s", err.Error())
		}
	}

	// check if the discriminator value is 'Standard'
	if jsonDict["kind"] == "Standard" {
		// try to unmarshal JSON data into CmkV2Standard
		err = json.Unmarshal(data, &dst.CmkV2Standard)
		if err == nil {
			return nil // data stored in dst.CmkV2Standard, return on the first match
		} else {
			dst.CmkV2Standard = nil
			return fmt.Errorf("Failed to unmarshal CmkV2ClusterSpecUpdateConfigOneOf as CmkV2Standard: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cmk.v2.Basic'
	if jsonDict["kind"] == "cmk.v2.Basic" {
		// try to unmarshal JSON data into CmkV2Basic
		err = json.Unmarshal(data, &dst.CmkV2Basic)
		if err == nil {
			return nil // data stored in dst.CmkV2Basic, return on the first match
		} else {
			dst.CmkV2Basic = nil
			return fmt.Errorf("Failed to unmarshal CmkV2ClusterSpecUpdateConfigOneOf as CmkV2Basic: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cmk.v2.Dedicated'
	if jsonDict["kind"] == "cmk.v2.Dedicated" {
		// try to unmarshal JSON data into CmkV2Dedicated
		err = json.Unmarshal(data, &dst.CmkV2Dedicated)
		if err == nil {
			return nil // data stored in dst.CmkV2Dedicated, return on the first match
		} else {
			dst.CmkV2Dedicated = nil
			return fmt.Errorf("Failed to unmarshal CmkV2ClusterSpecUpdateConfigOneOf as CmkV2Dedicated: %s", err.Error())
		}
	}

	// check if the discriminator value is 'cmk.v2.Standard'
	if jsonDict["kind"] == "cmk.v2.Standard" {
		// try to unmarshal JSON data into CmkV2Standard
		err = json.Unmarshal(data, &dst.CmkV2Standard)
		if err == nil {
			return nil // data stored in dst.CmkV2Standard, return on the first match
		} else {
			dst.CmkV2Standard = nil
			return fmt.Errorf("Failed to unmarshal CmkV2ClusterSpecUpdateConfigOneOf as CmkV2Standard: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CmkV2ClusterSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	if src.CmkV2Basic != nil {
		return json.Marshal(&src.CmkV2Basic)
	}

	if src.CmkV2Dedicated != nil {
		return json.Marshal(&src.CmkV2Dedicated)
	}

	if src.CmkV2Standard != nil {
		return json.Marshal(&src.CmkV2Standard)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CmkV2ClusterSpecUpdateConfigOneOf) GetActualInstance() (interface{}) {
	if obj.CmkV2Basic != nil {
		return obj.CmkV2Basic
	}

	if obj.CmkV2Dedicated != nil {
		return obj.CmkV2Dedicated
	}

	if obj.CmkV2Standard != nil {
		return obj.CmkV2Standard
	}

	// all schemas are nil
	return nil
}

type NullableCmkV2ClusterSpecUpdateConfigOneOf struct {
	value *CmkV2ClusterSpecUpdateConfigOneOf
	isSet bool
}

func (v NullableCmkV2ClusterSpecUpdateConfigOneOf) Get() *CmkV2ClusterSpecUpdateConfigOneOf {
	return v.value
}

func (v *NullableCmkV2ClusterSpecUpdateConfigOneOf) Set(val *CmkV2ClusterSpecUpdateConfigOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCmkV2ClusterSpecUpdateConfigOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCmkV2ClusterSpecUpdateConfigOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmkV2ClusterSpecUpdateConfigOneOf(val *CmkV2ClusterSpecUpdateConfigOneOf) *NullableCmkV2ClusterSpecUpdateConfigOneOf {
	return &NullableCmkV2ClusterSpecUpdateConfigOneOf{value: val, isSet: true}
}

func (v NullableCmkV2ClusterSpecUpdateConfigOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmkV2ClusterSpecUpdateConfigOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


