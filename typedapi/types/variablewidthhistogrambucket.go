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

// VariableWidthHistogramBucket type.
type VariableWidthHistogramBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Key          float64                     `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
	Max          float64                     `json:"max"`
	MaxAsString  *string                     `json:"max_as_string,omitempty"`
	Min          float64                     `json:"min"`
	MinAsString  *string                     `json:"min_as_string,omitempty"`
}

// VariableWidthHistogramBucketBuilder holds VariableWidthHistogramBucket struct and provides a builder API.
type VariableWidthHistogramBucketBuilder struct {
	v *VariableWidthHistogramBucket
}

// NewVariableWidthHistogramBucket provides a builder for the VariableWidthHistogramBucket struct.
func NewVariableWidthHistogramBucket() *VariableWidthHistogramBucketBuilder {
	r := VariableWidthHistogramBucketBuilder{
		&VariableWidthHistogramBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the VariableWidthHistogramBucket struct
func (rb *VariableWidthHistogramBucketBuilder) Build() VariableWidthHistogramBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *VariableWidthHistogramBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) DocCount(doccount int64) *VariableWidthHistogramBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Key set the Key property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) Key(key float64) *VariableWidthHistogramBucketBuilder {
	rb.v.Key = key
	return rb
}

// KeyAsString set the KeyAsString property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) KeyAsString(keyasstring string) *VariableWidthHistogramBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}

// Max set the Max property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) Max(max float64) *VariableWidthHistogramBucketBuilder {
	rb.v.Max = max
	return rb
}

// MaxAsString set the MaxAsString property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) MaxAsString(maxasstring string) *VariableWidthHistogramBucketBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

// Min set the Min property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) Min(min float64) *VariableWidthHistogramBucketBuilder {
	rb.v.Min = min
	return rb
}

// MinAsString set the MinAsString property for VariableWidthHistogramBucketBuilder.
func (rb *VariableWidthHistogramBucketBuilder) MinAsString(minasstring string) *VariableWidthHistogramBucketBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}
