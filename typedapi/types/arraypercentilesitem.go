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

// ArrayPercentilesItem type.
type ArrayPercentilesItem struct {
	Key           string  `json:"key"`
	Value         float64 `json:"value,omitempty"`
	ValueAsString *string `json:"value_as_string,omitempty"`
}

// ArrayPercentilesItemBuilder holds ArrayPercentilesItem struct and provides a builder API.
type ArrayPercentilesItemBuilder struct {
	v *ArrayPercentilesItem
}

// NewArrayPercentilesItem provides a builder for the ArrayPercentilesItem struct.
func NewArrayPercentilesItem() *ArrayPercentilesItemBuilder {
	r := ArrayPercentilesItemBuilder{
		&ArrayPercentilesItem{},
	}

	return &r
}

// Build finalize the chain and returns the ArrayPercentilesItem struct
func (rb *ArrayPercentilesItemBuilder) Build() ArrayPercentilesItem {
	return *rb.v
}

// Key set the Key property for ArrayPercentilesItemBuilder.
func (rb *ArrayPercentilesItemBuilder) Key(key string) *ArrayPercentilesItemBuilder {
	rb.v.Key = key
	return rb
}

// Value set the Value property for ArrayPercentilesItemBuilder.
func (rb *ArrayPercentilesItemBuilder) Value(value float64) *ArrayPercentilesItemBuilder {
	rb.v.Value = value
	return rb
}

// ValueAsString set the ValueAsString property for ArrayPercentilesItemBuilder.
func (rb *ArrayPercentilesItemBuilder) ValueAsString(valueasstring string) *ArrayPercentilesItemBuilder {
	rb.v.ValueAsString = &valueasstring
	return rb
}
