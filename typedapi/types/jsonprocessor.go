// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification


package types

// JsonProcessor type.
type JsonProcessor struct {
	AddToRoot     bool                 `json:"add_to_root"`
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   Field                `json:"target_field"`
}

// JsonProcessorBuilder holds JsonProcessor struct and provides a builder API.
type JsonProcessorBuilder struct {
	v *JsonProcessor
}

// NewJsonProcessor provides a builder for the JsonProcessor struct.
func NewJsonProcessor() *JsonProcessorBuilder {
	r := JsonProcessorBuilder{
		&JsonProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the JsonProcessor struct
func (rb *JsonProcessorBuilder) Build() JsonProcessor {
	return *rb.v
}

// AddToRoot set the AddToRoot property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) AddToRoot(addtoroot bool) *JsonProcessorBuilder {
	rb.v.AddToRoot = addtoroot
	return rb
}

// Field set the Field property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) Field(field Field) *JsonProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) If_(if_ string) *JsonProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) IgnoreFailure(ignorefailure bool) *JsonProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *JsonProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) Tag(tag string) *JsonProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for JsonProcessorBuilder.
func (rb *JsonProcessorBuilder) TargetField(targetfield Field) *JsonProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
