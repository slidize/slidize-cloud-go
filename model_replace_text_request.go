/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Smallize">
 *   Copyright (c) 2024 Slidize for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */
package slidizecloud

import (
	"encoding/json"
	"os"
)

// checks if the ReplaceTextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceTextRequest{}

// ReplaceTextRequest struct for ReplaceTextRequest
type ReplaceTextRequest struct {
	Documents []*os.File `json:"documents,omitempty"`
	Options *ReplaceTextOptions `json:"options,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReplaceTextRequest ReplaceTextRequest

// NewReplaceTextRequest instantiates a new ReplaceTextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceTextRequest() *ReplaceTextRequest {
	this := ReplaceTextRequest{}
	return &this
}

// NewReplaceTextRequestWithDefaults instantiates a new ReplaceTextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceTextRequestWithDefaults() *ReplaceTextRequest {
	this := ReplaceTextRequest{}
	return &this
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *ReplaceTextRequest) GetDocuments() []*os.File {
	if o == nil || IsNil(o.Documents) {
		var ret []*os.File
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceTextRequest) GetDocumentsOk() ([]*os.File, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *ReplaceTextRequest) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []*os.File and assigns it to the Documents field.
func (o *ReplaceTextRequest) SetDocuments(v []*os.File) {
	o.Documents = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ReplaceTextRequest) GetOptions() ReplaceTextOptions {
	if o == nil || IsNil(o.Options) {
		var ret ReplaceTextOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReplaceTextRequest) GetOptionsOk() (*ReplaceTextOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ReplaceTextRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ReplaceTextOptions and assigns it to the Options field.
func (o *ReplaceTextRequest) SetOptions(v ReplaceTextOptions) {
	o.Options = &v
}

func (o ReplaceTextRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceTextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReplaceTextRequest) UnmarshalJSON(bytes []byte) (err error) {
	varReplaceTextRequest := _ReplaceTextRequest{}

	if err = json.Unmarshal(bytes, &varReplaceTextRequest); err == nil {
		*o = ReplaceTextRequest(varReplaceTextRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "documents")
		delete(additionalProperties, "options")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReplaceTextRequest struct {
	value *ReplaceTextRequest
	isSet bool
}

func (v NullableReplaceTextRequest) Get() *ReplaceTextRequest {
	return v.value
}

func (v *NullableReplaceTextRequest) Set(val *ReplaceTextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceTextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceTextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceTextRequest(val *ReplaceTextRequest) *NullableReplaceTextRequest {
	return &NullableReplaceTextRequest{value: val, isSet: true}
}

func (v NullableReplaceTextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceTextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


