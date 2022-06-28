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

// Anomaly type.
type Anomaly struct {
	Actual              []float64      `json:"actual,omitempty"`
	BucketSpan          Time           `json:"bucket_span"`
	ByFieldName         *string        `json:"by_field_name,omitempty"`
	ByFieldValue        *string        `json:"by_field_value,omitempty"`
	Causes              []AnomalyCause `json:"causes,omitempty"`
	DetectorIndex       int            `json:"detector_index"`
	FieldName           *string        `json:"field_name,omitempty"`
	Function            *string        `json:"function,omitempty"`
	FunctionDescription *string        `json:"function_description,omitempty"`
	Influencers         []Influence    `json:"influencers,omitempty"`
	InitialRecordScore  float64        `json:"initial_record_score"`
	IsInterim           bool           `json:"is_interim"`
	JobId               string         `json:"job_id"`
	OverFieldName       *string        `json:"over_field_name,omitempty"`
	OverFieldValue      *string        `json:"over_field_value,omitempty"`
	PartitionFieldName  *string        `json:"partition_field_name,omitempty"`
	PartitionFieldValue *string        `json:"partition_field_value,omitempty"`
	Probability         float64        `json:"probability"`
	RecordScore         float64        `json:"record_score"`
	ResultType          string         `json:"result_type"`
	Timestamp           EpochMillis    `json:"timestamp"`
	Typical             []float64      `json:"typical,omitempty"`
}

// AnomalyBuilder holds Anomaly struct and provides a builder API.
type AnomalyBuilder struct {
	v *Anomaly
}

// NewAnomaly provides a builder for the Anomaly struct.
func NewAnomaly() *AnomalyBuilder {
	r := AnomalyBuilder{
		&Anomaly{},
	}

	return &r
}

// Build finalize the chain and returns the Anomaly struct
func (rb *AnomalyBuilder) Build() Anomaly {
	return *rb.v
}

// Actual set the Actual property for AnomalyBuilder.
func (rb *AnomalyBuilder) Actual(actual ...float64) *AnomalyBuilder {
	rb.v.Actual = append(rb.v.Actual, actual...)
	return rb
}

// BucketSpan set the BucketSpan property for AnomalyBuilder.
func (rb *AnomalyBuilder) BucketSpan(bucketspan Time) *AnomalyBuilder {
	rb.v.BucketSpan = bucketspan
	return rb
}

// ByFieldName set the ByFieldName property for AnomalyBuilder.
func (rb *AnomalyBuilder) ByFieldName(byfieldname string) *AnomalyBuilder {
	rb.v.ByFieldName = &byfieldname
	return rb
}

// ByFieldValue set the ByFieldValue property for AnomalyBuilder.
func (rb *AnomalyBuilder) ByFieldValue(byfieldvalue string) *AnomalyBuilder {
	rb.v.ByFieldValue = &byfieldvalue
	return rb
}

// Causes set the Causes property for AnomalyBuilder.
func (rb *AnomalyBuilder) Causes(causes ...AnomalyCause) *AnomalyBuilder {
	rb.v.Causes = append(rb.v.Causes, causes...)
	return rb
}

// DetectorIndex set the DetectorIndex property for AnomalyBuilder.
func (rb *AnomalyBuilder) DetectorIndex(detectorindex int) *AnomalyBuilder {
	rb.v.DetectorIndex = detectorindex
	return rb
}

// FieldName set the FieldName property for AnomalyBuilder.
func (rb *AnomalyBuilder) FieldName(fieldname string) *AnomalyBuilder {
	rb.v.FieldName = &fieldname
	return rb
}

// Function set the Function property for AnomalyBuilder.
func (rb *AnomalyBuilder) Function(function string) *AnomalyBuilder {
	rb.v.Function = &function
	return rb
}

// FunctionDescription set the FunctionDescription property for AnomalyBuilder.
func (rb *AnomalyBuilder) FunctionDescription(functiondescription string) *AnomalyBuilder {
	rb.v.FunctionDescription = &functiondescription
	return rb
}

// Influencers set the Influencers property for AnomalyBuilder.
func (rb *AnomalyBuilder) Influencers(influencers ...Influence) *AnomalyBuilder {
	rb.v.Influencers = append(rb.v.Influencers, influencers...)
	return rb
}

// InitialRecordScore set the InitialRecordScore property for AnomalyBuilder.
func (rb *AnomalyBuilder) InitialRecordScore(initialrecordscore float64) *AnomalyBuilder {
	rb.v.InitialRecordScore = initialrecordscore
	return rb
}

// IsInterim set the IsInterim property for AnomalyBuilder.
func (rb *AnomalyBuilder) IsInterim(isinterim bool) *AnomalyBuilder {
	rb.v.IsInterim = isinterim
	return rb
}

// JobId set the JobId property for AnomalyBuilder.
func (rb *AnomalyBuilder) JobId(jobid string) *AnomalyBuilder {
	rb.v.JobId = jobid
	return rb
}

// OverFieldName set the OverFieldName property for AnomalyBuilder.
func (rb *AnomalyBuilder) OverFieldName(overfieldname string) *AnomalyBuilder {
	rb.v.OverFieldName = &overfieldname
	return rb
}

// OverFieldValue set the OverFieldValue property for AnomalyBuilder.
func (rb *AnomalyBuilder) OverFieldValue(overfieldvalue string) *AnomalyBuilder {
	rb.v.OverFieldValue = &overfieldvalue
	return rb
}

// PartitionFieldName set the PartitionFieldName property for AnomalyBuilder.
func (rb *AnomalyBuilder) PartitionFieldName(partitionfieldname string) *AnomalyBuilder {
	rb.v.PartitionFieldName = &partitionfieldname
	return rb
}

// PartitionFieldValue set the PartitionFieldValue property for AnomalyBuilder.
func (rb *AnomalyBuilder) PartitionFieldValue(partitionfieldvalue string) *AnomalyBuilder {
	rb.v.PartitionFieldValue = &partitionfieldvalue
	return rb
}

// Probability set the Probability property for AnomalyBuilder.
func (rb *AnomalyBuilder) Probability(probability float64) *AnomalyBuilder {
	rb.v.Probability = probability
	return rb
}

// RecordScore set the RecordScore property for AnomalyBuilder.
func (rb *AnomalyBuilder) RecordScore(recordscore float64) *AnomalyBuilder {
	rb.v.RecordScore = recordscore
	return rb
}

// ResultType set the ResultType property for AnomalyBuilder.
func (rb *AnomalyBuilder) ResultType(resulttype string) *AnomalyBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Timestamp set the Timestamp property for AnomalyBuilder.
func (rb *AnomalyBuilder) Timestamp(timestamp EpochMillis) *AnomalyBuilder {
	rb.v.Timestamp = timestamp
	return rb
}

// Typical set the Typical property for AnomalyBuilder.
func (rb *AnomalyBuilder) Typical(typical ...float64) *AnomalyBuilder {
	rb.v.Typical = append(rb.v.Typical, typical...)
	return rb
}
