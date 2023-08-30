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
Custom Connector Plugin Management API

This is Custom Connector Plugin Management API.

API version: 1.0.0
Contact: compute-platform-team@confluent.io
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

// ConnectV1CustomConnectorPlugin `CustomConnectorPlugins` objects represent Custom Connector Plugins on Confluent Cloud.  The API allows you to list, create, read, update, and delete your Custom Connector Plugins.   ## The Custom Connector Plugins Model <SchemaDefinition schemaRef=\"#/components/schemas/connect.v1.CustomConnectorPlugin\" />
type ConnectV1CustomConnectorPlugin struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id *string `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// Display name of Custom Connector Plugin.
	DisplayName *string `json:"display_name,omitempty"`
	// Archive format of Custom Connector Plugin.
	ContentFormat *string `json:"content_format,omitempty"`
	// Description of Custom Connector Plugin.
	Description *string `json:"description,omitempty"`
	// Document link of Custom Connector Plugin.
	DocumentationLink *string `json:"documentation_link,omitempty"`
	// Java class or alias for connector. You can get connector class from connector documentation provided by developer.
	ConnectorClass *string `json:"connector_class,omitempty"`
	// Custom Connector type. 
	ConnectorType *string `json:"connector_type,omitempty"`
	// A sensitive property is a connector configuration property that must be hidden after a user enters property value when setting up connector.
	SensitiveConfigProperties *[]string `json:"sensitive_config_properties,omitempty"`
	// [immutable] Upload source of Custom Connector Plugin. Only required in `create` request, will be ignored in `read`, `update` or `list`.
	UploadSource *ConnectV1CustomConnectorPluginUploadSourceOneOf `json:"upload_source,omitempty"`
}

// NewConnectV1CustomConnectorPlugin instantiates a new ConnectV1CustomConnectorPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1CustomConnectorPlugin() *ConnectV1CustomConnectorPlugin {
	this := ConnectV1CustomConnectorPlugin{}
	return &this
}

// NewConnectV1CustomConnectorPluginWithDefaults instantiates a new ConnectV1CustomConnectorPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1CustomConnectorPluginWithDefaults() *ConnectV1CustomConnectorPlugin {
	this := ConnectV1CustomConnectorPlugin{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ConnectV1CustomConnectorPlugin) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ConnectV1CustomConnectorPlugin) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectV1CustomConnectorPlugin) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *ConnectV1CustomConnectorPlugin) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ConnectV1CustomConnectorPlugin) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetContentFormat returns the ContentFormat field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetContentFormat() string {
	if o == nil || o.ContentFormat == nil {
		var ret string
		return ret
	}
	return *o.ContentFormat
}

// GetContentFormatOk returns a tuple with the ContentFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetContentFormatOk() (*string, bool) {
	if o == nil || o.ContentFormat == nil {
		return nil, false
	}
	return o.ContentFormat, true
}

// HasContentFormat returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasContentFormat() bool {
	if o != nil && o.ContentFormat != nil {
		return true
	}

	return false
}

// SetContentFormat gets a reference to the given string and assigns it to the ContentFormat field.
func (o *ConnectV1CustomConnectorPlugin) SetContentFormat(v string) {
	o.ContentFormat = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectV1CustomConnectorPlugin) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentationLink returns the DocumentationLink field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetDocumentationLink() string {
	if o == nil || o.DocumentationLink == nil {
		var ret string
		return ret
	}
	return *o.DocumentationLink
}

// GetDocumentationLinkOk returns a tuple with the DocumentationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetDocumentationLinkOk() (*string, bool) {
	if o == nil || o.DocumentationLink == nil {
		return nil, false
	}
	return o.DocumentationLink, true
}

// HasDocumentationLink returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasDocumentationLink() bool {
	if o != nil && o.DocumentationLink != nil {
		return true
	}

	return false
}

// SetDocumentationLink gets a reference to the given string and assigns it to the DocumentationLink field.
func (o *ConnectV1CustomConnectorPlugin) SetDocumentationLink(v string) {
	o.DocumentationLink = &v
}

// GetConnectorClass returns the ConnectorClass field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetConnectorClass() string {
	if o == nil || o.ConnectorClass == nil {
		var ret string
		return ret
	}
	return *o.ConnectorClass
}

// GetConnectorClassOk returns a tuple with the ConnectorClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetConnectorClassOk() (*string, bool) {
	if o == nil || o.ConnectorClass == nil {
		return nil, false
	}
	return o.ConnectorClass, true
}

// HasConnectorClass returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasConnectorClass() bool {
	if o != nil && o.ConnectorClass != nil {
		return true
	}

	return false
}

// SetConnectorClass gets a reference to the given string and assigns it to the ConnectorClass field.
func (o *ConnectV1CustomConnectorPlugin) SetConnectorClass(v string) {
	o.ConnectorClass = &v
}

// GetConnectorType returns the ConnectorType field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetConnectorType() string {
	if o == nil || o.ConnectorType == nil {
		var ret string
		return ret
	}
	return *o.ConnectorType
}

// GetConnectorTypeOk returns a tuple with the ConnectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetConnectorTypeOk() (*string, bool) {
	if o == nil || o.ConnectorType == nil {
		return nil, false
	}
	return o.ConnectorType, true
}

// HasConnectorType returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasConnectorType() bool {
	if o != nil && o.ConnectorType != nil {
		return true
	}

	return false
}

// SetConnectorType gets a reference to the given string and assigns it to the ConnectorType field.
func (o *ConnectV1CustomConnectorPlugin) SetConnectorType(v string) {
	o.ConnectorType = &v
}

// GetSensitiveConfigProperties returns the SensitiveConfigProperties field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetSensitiveConfigProperties() []string {
	if o == nil || o.SensitiveConfigProperties == nil {
		var ret []string
		return ret
	}
	return *o.SensitiveConfigProperties
}

// GetSensitiveConfigPropertiesOk returns a tuple with the SensitiveConfigProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetSensitiveConfigPropertiesOk() (*[]string, bool) {
	if o == nil || o.SensitiveConfigProperties == nil {
		return nil, false
	}
	return o.SensitiveConfigProperties, true
}

// HasSensitiveConfigProperties returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasSensitiveConfigProperties() bool {
	if o != nil && o.SensitiveConfigProperties != nil {
		return true
	}

	return false
}

// SetSensitiveConfigProperties gets a reference to the given []string and assigns it to the SensitiveConfigProperties field.
func (o *ConnectV1CustomConnectorPlugin) SetSensitiveConfigProperties(v []string) {
	o.SensitiveConfigProperties = &v
}

// GetUploadSource returns the UploadSource field value if set, zero value otherwise.
func (o *ConnectV1CustomConnectorPlugin) GetUploadSource() ConnectV1CustomConnectorPluginUploadSourceOneOf {
	if o == nil || o.UploadSource == nil {
		var ret ConnectV1CustomConnectorPluginUploadSourceOneOf
		return ret
	}
	return *o.UploadSource
}

// GetUploadSourceOk returns a tuple with the UploadSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1CustomConnectorPlugin) GetUploadSourceOk() (*ConnectV1CustomConnectorPluginUploadSourceOneOf, bool) {
	if o == nil || o.UploadSource == nil {
		return nil, false
	}
	return o.UploadSource, true
}

// HasUploadSource returns a boolean if a field has been set.
func (o *ConnectV1CustomConnectorPlugin) HasUploadSource() bool {
	if o != nil && o.UploadSource != nil {
		return true
	}

	return false
}

// SetUploadSource gets a reference to the given ConnectV1CustomConnectorPluginUploadSourceOneOf and assigns it to the UploadSource field.
func (o *ConnectV1CustomConnectorPlugin) SetUploadSource(v ConnectV1CustomConnectorPluginUploadSourceOneOf) {
	o.UploadSource = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1CustomConnectorPlugin) Redact() {
    o.recurseRedact(o.ApiVersion)
    o.recurseRedact(o.Kind)
    o.recurseRedact(o.Id)
    o.recurseRedact(o.Metadata)
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.ContentFormat)
    o.recurseRedact(o.Description)
    o.recurseRedact(o.DocumentationLink)
    o.recurseRedact(o.ConnectorClass)
    o.recurseRedact(o.ConnectorType)
    o.recurseRedact(o.SensitiveConfigProperties)
    o.recurseRedact(o.UploadSource)
}

func (o *ConnectV1CustomConnectorPlugin) recurseRedact(v interface{}) {
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

func (o ConnectV1CustomConnectorPlugin) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1CustomConnectorPlugin) MarshalJSON() ([]byte, error) {
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
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.ContentFormat != nil {
		toSerialize["content_format"] = o.ContentFormat
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentationLink != nil {
		toSerialize["documentation_link"] = o.DocumentationLink
	}
	if o.ConnectorClass != nil {
		toSerialize["connector_class"] = o.ConnectorClass
	}
	if o.ConnectorType != nil {
		toSerialize["connector_type"] = o.ConnectorType
	}
	if o.SensitiveConfigProperties != nil {
		toSerialize["sensitive_config_properties"] = o.SensitiveConfigProperties
	}
	if o.UploadSource != nil {
		toSerialize["upload_source"] = o.UploadSource
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1CustomConnectorPlugin struct {
	value *ConnectV1CustomConnectorPlugin
	isSet bool
}

func (v NullableConnectV1CustomConnectorPlugin) Get() *ConnectV1CustomConnectorPlugin {
	return v.value
}

func (v *NullableConnectV1CustomConnectorPlugin) Set(val *ConnectV1CustomConnectorPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1CustomConnectorPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1CustomConnectorPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1CustomConnectorPlugin(val *ConnectV1CustomConnectorPlugin) *NullableConnectV1CustomConnectorPlugin {
	return &NullableConnectV1CustomConnectorPlugin{value: val, isSet: true}
}

func (v NullableConnectV1CustomConnectorPlugin) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1CustomConnectorPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


