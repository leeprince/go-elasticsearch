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

// TargetMeanEncodingPreprocessor type.
type TargetMeanEncodingPreprocessor struct {
	DefaultValue float64            `json:"default_value"`
	FeatureName  string             `json:"feature_name"`
	Field        string             `json:"field"`
	TargetMap    map[string]float64 `json:"target_map"`
}

// TargetMeanEncodingPreprocessorBuilder holds TargetMeanEncodingPreprocessor struct and provides a builder API.
type TargetMeanEncodingPreprocessorBuilder struct {
	v *TargetMeanEncodingPreprocessor
}

// NewTargetMeanEncodingPreprocessor provides a builder for the TargetMeanEncodingPreprocessor struct.
func NewTargetMeanEncodingPreprocessor() *TargetMeanEncodingPreprocessorBuilder {
	r := TargetMeanEncodingPreprocessorBuilder{
		&TargetMeanEncodingPreprocessor{
			TargetMap: make(map[string]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TargetMeanEncodingPreprocessor struct
func (rb *TargetMeanEncodingPreprocessorBuilder) Build() TargetMeanEncodingPreprocessor {
	return *rb.v
}

// DefaultValue set the DefaultValue property for TargetMeanEncodingPreprocessorBuilder.
func (rb *TargetMeanEncodingPreprocessorBuilder) DefaultValue(defaultvalue float64) *TargetMeanEncodingPreprocessorBuilder {
	rb.v.DefaultValue = defaultvalue
	return rb
}

// FeatureName set the FeatureName property for TargetMeanEncodingPreprocessorBuilder.
func (rb *TargetMeanEncodingPreprocessorBuilder) FeatureName(featurename string) *TargetMeanEncodingPreprocessorBuilder {
	rb.v.FeatureName = featurename
	return rb
}

// Field set the Field property for TargetMeanEncodingPreprocessorBuilder.
func (rb *TargetMeanEncodingPreprocessorBuilder) Field(field string) *TargetMeanEncodingPreprocessorBuilder {
	rb.v.Field = field
	return rb
}

// TargetMap set the TargetMap property for TargetMeanEncodingPreprocessorBuilder.
func (rb *TargetMeanEncodingPreprocessorBuilder) TargetMap(value map[string]float64) *TargetMeanEncodingPreprocessorBuilder {
	rb.v.TargetMap = value
	return rb
}
