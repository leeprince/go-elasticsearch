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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/categorizationstatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/memorystatus"
)

// ModelSizeStats type.
type ModelSizeStats struct {
	AssignmentMemoryBasis         *string                                   `json:"assignment_memory_basis,omitempty"`
	BucketAllocationFailuresCount int64                                     `json:"bucket_allocation_failures_count"`
	CategorizationStatus          categorizationstatus.CategorizationStatus `json:"categorization_status"`
	CategorizedDocCount           int                                       `json:"categorized_doc_count"`
	DeadCategoryCount             int                                       `json:"dead_category_count"`
	FailedCategoryCount           int                                       `json:"failed_category_count"`
	FrequentCategoryCount         int                                       `json:"frequent_category_count"`
	JobId                         Id                                        `json:"job_id"`
	LogTime                       Time                                      `json:"log_time"`
	MemoryStatus                  memorystatus.MemoryStatus                 `json:"memory_status"`
	ModelBytes                    ByteSize                                  `json:"model_bytes"`
	ModelBytesExceeded            *ByteSize                                 `json:"model_bytes_exceeded,omitempty"`
	ModelBytesMemoryLimit         *ByteSize                                 `json:"model_bytes_memory_limit,omitempty"`
	PeakModelBytes                *ByteSize                                 `json:"peak_model_bytes,omitempty"`
	RareCategoryCount             int                                       `json:"rare_category_count"`
	ResultType                    string                                    `json:"result_type"`
	Timestamp                     *int64                                    `json:"timestamp,omitempty"`
	TotalByFieldCount             int64                                     `json:"total_by_field_count"`
	TotalCategoryCount            int                                       `json:"total_category_count"`
	TotalOverFieldCount           int64                                     `json:"total_over_field_count"`
	TotalPartitionFieldCount      int64                                     `json:"total_partition_field_count"`
}

// ModelSizeStatsBuilder holds ModelSizeStats struct and provides a builder API.
type ModelSizeStatsBuilder struct {
	v *ModelSizeStats
}

// NewModelSizeStats provides a builder for the ModelSizeStats struct.
func NewModelSizeStats() *ModelSizeStatsBuilder {
	r := ModelSizeStatsBuilder{
		&ModelSizeStats{},
	}

	return &r
}

// Build finalize the chain and returns the ModelSizeStats struct
func (rb *ModelSizeStatsBuilder) Build() ModelSizeStats {
	return *rb.v
}

// AssignmentMemoryBasis set the AssignmentMemoryBasis property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) AssignmentMemoryBasis(assignmentmemorybasis string) *ModelSizeStatsBuilder {
	rb.v.AssignmentMemoryBasis = &assignmentmemorybasis
	return rb
}

// BucketAllocationFailuresCount set the BucketAllocationFailuresCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) BucketAllocationFailuresCount(bucketallocationfailurescount int64) *ModelSizeStatsBuilder {
	rb.v.BucketAllocationFailuresCount = bucketallocationfailurescount
	return rb
}

// CategorizationStatus set the CategorizationStatus property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) CategorizationStatus(categorizationstatus categorizationstatus.CategorizationStatus) *ModelSizeStatsBuilder {
	rb.v.CategorizationStatus = categorizationstatus
	return rb
}

// CategorizedDocCount set the CategorizedDocCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) CategorizedDocCount(categorizeddoccount int) *ModelSizeStatsBuilder {
	rb.v.CategorizedDocCount = categorizeddoccount
	return rb
}

// DeadCategoryCount set the DeadCategoryCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) DeadCategoryCount(deadcategorycount int) *ModelSizeStatsBuilder {
	rb.v.DeadCategoryCount = deadcategorycount
	return rb
}

// FailedCategoryCount set the FailedCategoryCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) FailedCategoryCount(failedcategorycount int) *ModelSizeStatsBuilder {
	rb.v.FailedCategoryCount = failedcategorycount
	return rb
}

// FrequentCategoryCount set the FrequentCategoryCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) FrequentCategoryCount(frequentcategorycount int) *ModelSizeStatsBuilder {
	rb.v.FrequentCategoryCount = frequentcategorycount
	return rb
}

// JobId set the JobId property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) JobId(jobid Id) *ModelSizeStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

// LogTime set the LogTime property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) LogTime(logtime Time) *ModelSizeStatsBuilder {
	rb.v.LogTime = logtime
	return rb
}

// MemoryStatus set the MemoryStatus property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) MemoryStatus(memorystatus memorystatus.MemoryStatus) *ModelSizeStatsBuilder {
	rb.v.MemoryStatus = memorystatus
	return rb
}

// ModelBytes set the ModelBytes property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) ModelBytes(modelbytes ByteSize) *ModelSizeStatsBuilder {
	rb.v.ModelBytes = modelbytes
	return rb
}

// ModelBytesExceeded set the ModelBytesExceeded property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) ModelBytesExceeded(modelbytesexceeded ByteSize) *ModelSizeStatsBuilder {
	rb.v.ModelBytesExceeded = &modelbytesexceeded
	return rb
}

// ModelBytesMemoryLimit set the ModelBytesMemoryLimit property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) ModelBytesMemoryLimit(modelbytesmemorylimit ByteSize) *ModelSizeStatsBuilder {
	rb.v.ModelBytesMemoryLimit = &modelbytesmemorylimit
	return rb
}

// PeakModelBytes set the PeakModelBytes property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) PeakModelBytes(peakmodelbytes ByteSize) *ModelSizeStatsBuilder {
	rb.v.PeakModelBytes = &peakmodelbytes
	return rb
}

// RareCategoryCount set the RareCategoryCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) RareCategoryCount(rarecategorycount int) *ModelSizeStatsBuilder {
	rb.v.RareCategoryCount = rarecategorycount
	return rb
}

// ResultType set the ResultType property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) ResultType(resulttype string) *ModelSizeStatsBuilder {
	rb.v.ResultType = resulttype
	return rb
}

// Timestamp set the Timestamp property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) Timestamp(timestamp int64) *ModelSizeStatsBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// TotalByFieldCount set the TotalByFieldCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) TotalByFieldCount(totalbyfieldcount int64) *ModelSizeStatsBuilder {
	rb.v.TotalByFieldCount = totalbyfieldcount
	return rb
}

// TotalCategoryCount set the TotalCategoryCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) TotalCategoryCount(totalcategorycount int) *ModelSizeStatsBuilder {
	rb.v.TotalCategoryCount = totalcategorycount
	return rb
}

// TotalOverFieldCount set the TotalOverFieldCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) TotalOverFieldCount(totaloverfieldcount int64) *ModelSizeStatsBuilder {
	rb.v.TotalOverFieldCount = totaloverfieldcount
	return rb
}

// TotalPartitionFieldCount set the TotalPartitionFieldCount property for ModelSizeStatsBuilder.
func (rb *ModelSizeStatsBuilder) TotalPartitionFieldCount(totalpartitionfieldcount int64) *ModelSizeStatsBuilder {
	rb.v.TotalPartitionFieldCount = totalpartitionfieldcount
	return rb
}
