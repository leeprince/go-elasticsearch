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

// AnomalyCause type.
type AnomalyCause struct {
	Actual                 []float64   `json:"actual"`
	ByFieldName            Name        `json:"by_field_name"`
	ByFieldValue           string      `json:"by_field_value"`
	CorrelatedByFieldValue string      `json:"correlated_by_field_value"`
	FieldName              Field       `json:"field_name"`
	Function               string      `json:"function"`
	FunctionDescription    string      `json:"function_description"`
	Influencers            []Influence `json:"influencers"`
	OverFieldName          Name        `json:"over_field_name"`
	OverFieldValue         string      `json:"over_field_value"`
	PartitionFieldName     string      `json:"partition_field_name"`
	PartitionFieldValue    string      `json:"partition_field_value"`
	Probability            float64     `json:"probability"`
	Typical                []float64   `json:"typical"`
}

// AnomalyCauseBuilder holds AnomalyCause struct and provides a builder API.
type AnomalyCauseBuilder struct {
	v *AnomalyCause
}

// NewAnomalyCause provides a builder for the AnomalyCause struct.
func NewAnomalyCause() *AnomalyCauseBuilder {
	r := AnomalyCauseBuilder{
		&AnomalyCause{},
	}

	return &r
}

// Build finalize the chain and returns the AnomalyCause struct
func (rb *AnomalyCauseBuilder) Build() AnomalyCause {
	return *rb.v
}

// Actual set the Actual property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) Actual(actual ...float64) *AnomalyCauseBuilder {
	rb.v.Actual = append(rb.v.Actual, actual...)
	return rb
}

// ByFieldName set the ByFieldName property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) ByFieldName(byfieldname Name) *AnomalyCauseBuilder {
	rb.v.ByFieldName = byfieldname
	return rb
}

// ByFieldValue set the ByFieldValue property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) ByFieldValue(byfieldvalue string) *AnomalyCauseBuilder {
	rb.v.ByFieldValue = byfieldvalue
	return rb
}

// CorrelatedByFieldValue set the CorrelatedByFieldValue property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) CorrelatedByFieldValue(correlatedbyfieldvalue string) *AnomalyCauseBuilder {
	rb.v.CorrelatedByFieldValue = correlatedbyfieldvalue
	return rb
}

// FieldName set the FieldName property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) FieldName(fieldname Field) *AnomalyCauseBuilder {
	rb.v.FieldName = fieldname
	return rb
}

// Function set the Function property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) Function(function string) *AnomalyCauseBuilder {
	rb.v.Function = function
	return rb
}

// FunctionDescription set the FunctionDescription property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) FunctionDescription(functiondescription string) *AnomalyCauseBuilder {
	rb.v.FunctionDescription = functiondescription
	return rb
}

// Influencers set the Influencers property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) Influencers(influencers ...Influence) *AnomalyCauseBuilder {
	rb.v.Influencers = append(rb.v.Influencers, influencers...)
	return rb
}

// OverFieldName set the OverFieldName property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) OverFieldName(overfieldname Name) *AnomalyCauseBuilder {
	rb.v.OverFieldName = overfieldname
	return rb
}

// OverFieldValue set the OverFieldValue property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) OverFieldValue(overfieldvalue string) *AnomalyCauseBuilder {
	rb.v.OverFieldValue = overfieldvalue
	return rb
}

// PartitionFieldName set the PartitionFieldName property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) PartitionFieldName(partitionfieldname string) *AnomalyCauseBuilder {
	rb.v.PartitionFieldName = partitionfieldname
	return rb
}

// PartitionFieldValue set the PartitionFieldValue property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) PartitionFieldValue(partitionfieldvalue string) *AnomalyCauseBuilder {
	rb.v.PartitionFieldValue = partitionfieldvalue
	return rb
}

// Probability set the Probability property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) Probability(probability float64) *AnomalyCauseBuilder {
	rb.v.Probability = probability
	return rb
}

// Typical set the Typical property for AnomalyCauseBuilder.
func (rb *AnomalyCauseBuilder) Typical(typical ...float64) *AnomalyCauseBuilder {
	rb.v.Typical = append(rb.v.Typical, typical...)
	return rb
}
