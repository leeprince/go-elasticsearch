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

// TransformIndexerStats type.
type TransformIndexerStats struct {
	DeleteTimeInMs                     *EpochMillis `json:"delete_time_in_ms,omitempty"`
	DocumentsDeleted                   *int64       `json:"documents_deleted,omitempty"`
	DocumentsIndexed                   int64        `json:"documents_indexed"`
	DocumentsProcessed                 int64        `json:"documents_processed"`
	ExponentialAvgCheckpointDurationMs float64      `json:"exponential_avg_checkpoint_duration_ms"`
	ExponentialAvgDocumentsIndexed     float64      `json:"exponential_avg_documents_indexed"`
	ExponentialAvgDocumentsProcessed   float64      `json:"exponential_avg_documents_processed"`
	IndexFailures                      int64        `json:"index_failures"`
	IndexTimeInMs                      int64        `json:"index_time_in_ms"`
	IndexTotal                         int64        `json:"index_total"`
	PagesProcessed                     int64        `json:"pages_processed"`
	ProcessingTimeInMs                 int64        `json:"processing_time_in_ms"`
	ProcessingTotal                    int64        `json:"processing_total"`
	SearchFailures                     int64        `json:"search_failures"`
	SearchTimeInMs                     int64        `json:"search_time_in_ms"`
	SearchTotal                        int64        `json:"search_total"`
	TriggerCount                       int64        `json:"trigger_count"`
}

// TransformIndexerStatsBuilder holds TransformIndexerStats struct and provides a builder API.
type TransformIndexerStatsBuilder struct {
	v *TransformIndexerStats
}

// NewTransformIndexerStats provides a builder for the TransformIndexerStats struct.
func NewTransformIndexerStats() *TransformIndexerStatsBuilder {
	r := TransformIndexerStatsBuilder{
		&TransformIndexerStats{},
	}

	return &r
}

// Build finalize the chain and returns the TransformIndexerStats struct
func (rb *TransformIndexerStatsBuilder) Build() TransformIndexerStats {
	return *rb.v
}

// DeleteTimeInMs set the DeleteTimeInMs property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) DeleteTimeInMs(deletetimeinms EpochMillis) *TransformIndexerStatsBuilder {
	rb.v.DeleteTimeInMs = &deletetimeinms
	return rb
}

// DocumentsDeleted set the DocumentsDeleted property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) DocumentsDeleted(documentsdeleted int64) *TransformIndexerStatsBuilder {
	rb.v.DocumentsDeleted = &documentsdeleted
	return rb
}

// DocumentsIndexed set the DocumentsIndexed property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) DocumentsIndexed(documentsindexed int64) *TransformIndexerStatsBuilder {
	rb.v.DocumentsIndexed = documentsindexed
	return rb
}

// DocumentsProcessed set the DocumentsProcessed property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) DocumentsProcessed(documentsprocessed int64) *TransformIndexerStatsBuilder {
	rb.v.DocumentsProcessed = documentsprocessed
	return rb
}

// ExponentialAvgCheckpointDurationMs set the ExponentialAvgCheckpointDurationMs property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) ExponentialAvgCheckpointDurationMs(exponentialavgcheckpointdurationms float64) *TransformIndexerStatsBuilder {
	rb.v.ExponentialAvgCheckpointDurationMs = exponentialavgcheckpointdurationms
	return rb
}

// ExponentialAvgDocumentsIndexed set the ExponentialAvgDocumentsIndexed property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) ExponentialAvgDocumentsIndexed(exponentialavgdocumentsindexed float64) *TransformIndexerStatsBuilder {
	rb.v.ExponentialAvgDocumentsIndexed = exponentialavgdocumentsindexed
	return rb
}

// ExponentialAvgDocumentsProcessed set the ExponentialAvgDocumentsProcessed property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) ExponentialAvgDocumentsProcessed(exponentialavgdocumentsprocessed float64) *TransformIndexerStatsBuilder {
	rb.v.ExponentialAvgDocumentsProcessed = exponentialavgdocumentsprocessed
	return rb
}

// IndexFailures set the IndexFailures property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) IndexFailures(indexfailures int64) *TransformIndexerStatsBuilder {
	rb.v.IndexFailures = indexfailures
	return rb
}

// IndexTimeInMs set the IndexTimeInMs property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) IndexTimeInMs(indextimeinms int64) *TransformIndexerStatsBuilder {
	rb.v.IndexTimeInMs = indextimeinms
	return rb
}

// IndexTotal set the IndexTotal property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) IndexTotal(indextotal int64) *TransformIndexerStatsBuilder {
	rb.v.IndexTotal = indextotal
	return rb
}

// PagesProcessed set the PagesProcessed property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) PagesProcessed(pagesprocessed int64) *TransformIndexerStatsBuilder {
	rb.v.PagesProcessed = pagesprocessed
	return rb
}

// ProcessingTimeInMs set the ProcessingTimeInMs property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) ProcessingTimeInMs(processingtimeinms int64) *TransformIndexerStatsBuilder {
	rb.v.ProcessingTimeInMs = processingtimeinms
	return rb
}

// ProcessingTotal set the ProcessingTotal property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) ProcessingTotal(processingtotal int64) *TransformIndexerStatsBuilder {
	rb.v.ProcessingTotal = processingtotal
	return rb
}

// SearchFailures set the SearchFailures property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) SearchFailures(searchfailures int64) *TransformIndexerStatsBuilder {
	rb.v.SearchFailures = searchfailures
	return rb
}

// SearchTimeInMs set the SearchTimeInMs property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) SearchTimeInMs(searchtimeinms int64) *TransformIndexerStatsBuilder {
	rb.v.SearchTimeInMs = searchtimeinms
	return rb
}

// SearchTotal set the SearchTotal property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) SearchTotal(searchtotal int64) *TransformIndexerStatsBuilder {
	rb.v.SearchTotal = searchtotal
	return rb
}

// TriggerCount set the TriggerCount property for TransformIndexerStatsBuilder.
func (rb *TransformIndexerStatsBuilder) TriggerCount(triggercount int64) *TransformIndexerStatsBuilder {
	rb.v.TriggerCount = triggercount
	return rb
}
