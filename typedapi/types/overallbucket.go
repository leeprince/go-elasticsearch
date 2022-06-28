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

// OverallBucket type.
type OverallBucket struct {
	// BucketSpan The length of the bucket in seconds. Matches the job with the longest
	// bucket_span value.
	BucketSpan int64 `json:"bucket_span"`
	// IsInterim If true, this is an interim result. In other words, the results are
	// calculated based on partial input data.
	IsInterim bool `json:"is_interim"`
	// Jobs An array of objects that contain the max_anomaly_score per job_id.
	Jobs []OverallBucketJob `json:"jobs"`
	// OverallScore The top_n average of the maximum bucket anomaly_score per job.
	OverallScore float64 `json:"overall_score"`
	// ResultType Internal. This is always set to overall_bucket.
	ResultType string `json:"result_type"`
	// Timestamp The start time of the bucket for which these results were calculated.
	Timestamp Time `json:"timestamp"`
}

// OverallBucketBuilder holds OverallBucket struct and provides a builder API.
type OverallBucketBuilder struct {
	v *OverallBucket
}

// NewOverallBucket provides a builder for the OverallBucket struct.
func NewOverallBucket() *OverallBucketBuilder {
	r := OverallBucketBuilder{
		&OverallBucket{},
	}

	return &r
}

// Build finalize the chain and returns the OverallBucket struct
func (rb *OverallBucketBuilder) Build() OverallBucket {
	return *rb.v
}

// BucketSpan The length of the bucket in seconds. Matches the job with the longest
// bucket_span value.
func (rb *OverallBucketBuilder) BucketSpan(bucketspan int64) *OverallBucketBuilder {
	rb.v.BucketSpan = bucketspan
	return rb
}

// IsInterim If true, this is an interim result. In other words, the results are
// calculated based on partial input data.
func (rb *OverallBucketBuilder) IsInterim(isinterim bool) *OverallBucketBuilder {
	rb.v.IsInterim = isinterim
	return rb
}

// Jobs An array of objects that contain the max_anomaly_score per job_id.
func (rb *OverallBucketBuilder) Jobs(jobs ...OverallBucketJob) *OverallBucketBuilder {
	rb.v.Jobs = append(rb.v.Jobs, jobs...)
	return rb
}

// OverallScore The top_n average of the maximum bucket anomaly_score per job.
func (rb *OverallBucketBuilder) OverallScore(overallscore float64) *OverallBucketBuilder {
	rb.v.OverallScore = overallscore
	return rb
}

// ResultType Internal. This is always set to overall_bucket.
func (rb *OverallBucketBuilder) ResultType(resulttype string) *OverallBucketBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Timestamp The start time of the bucket for which these results were calculated.
func (rb *OverallBucketBuilder) Timestamp(timestamp Time) *OverallBucketBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
