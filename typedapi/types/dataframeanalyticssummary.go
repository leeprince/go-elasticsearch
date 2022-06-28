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

// DataframeAnalyticsSummary type.
type DataframeAnalyticsSummary struct {
	AllowLazyStart   *bool                            `json:"allow_lazy_start,omitempty"`
	Analysis         DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	CreateTime       *int64                           `json:"create_time,omitempty"`
	Description      *string                          `json:"description,omitempty"`
	Dest             DataframeAnalyticsDestination    `json:"dest"`
	Id               Id                               `json:"id"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
	Version          *VersionString                   `json:"version,omitempty"`
}

// DataframeAnalyticsSummaryBuilder holds DataframeAnalyticsSummary struct and provides a builder API.
type DataframeAnalyticsSummaryBuilder struct {
	v *DataframeAnalyticsSummary
}

// NewDataframeAnalyticsSummary provides a builder for the DataframeAnalyticsSummary struct.
func NewDataframeAnalyticsSummary() *DataframeAnalyticsSummaryBuilder {
	r := DataframeAnalyticsSummaryBuilder{
		&DataframeAnalyticsSummary{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsSummary struct
func (rb *DataframeAnalyticsSummaryBuilder) Build() DataframeAnalyticsSummary {
	return *rb.v
}

// AllowLazyStart set the AllowLazyStart property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) AllowLazyStart(allowlazystart bool) *DataframeAnalyticsSummaryBuilder {
	rb.v.AllowLazyStart = &allowlazystart
	return rb
}

// Analysis set the Analysis property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) Analysis(analysis DataframeAnalysisContainer) *DataframeAnalyticsSummaryBuilder {
	rb.v.Analysis = analysis
	return rb
}

// AnalyzedFields set the AnalyzedFields property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) AnalyzedFields(analyzedfields DataframeAnalysisAnalyzedFields) *DataframeAnalyticsSummaryBuilder {
	rb.v.AnalyzedFields = &analyzedfields
	return rb
}

// CreateTime set the CreateTime property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) CreateTime(createtime int64) *DataframeAnalyticsSummaryBuilder {
	rb.v.CreateTime = &createtime
	return rb
}

// Description set the Description property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) Description(description string) *DataframeAnalyticsSummaryBuilder {
	rb.v.Description = &description
	return rb
}

// Dest set the Dest property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) Dest(dest DataframeAnalyticsDestination) *DataframeAnalyticsSummaryBuilder {
	rb.v.Dest = dest
	return rb
}

// Id set the Id property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) Id(id Id) *DataframeAnalyticsSummaryBuilder {
	rb.v.Id = id
	return rb
}

// MaxNumThreads set the MaxNumThreads property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) MaxNumThreads(maxnumthreads int) *DataframeAnalyticsSummaryBuilder {
	rb.v.MaxNumThreads = &maxnumthreads
	return rb
}

// ModelMemoryLimit set the ModelMemoryLimit property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) ModelMemoryLimit(modelmemorylimit string) *DataframeAnalyticsSummaryBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

// Source set the Source property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) Source(source DataframeAnalyticsSource) *DataframeAnalyticsSummaryBuilder {
	rb.v.Source = source
	return rb
}

// Version set the Version property for DataframeAnalyticsSummaryBuilder.
func (rb *DataframeAnalyticsSummaryBuilder) Version(version VersionString) *DataframeAnalyticsSummaryBuilder {
	rb.v.Version = &version
	return rb
}
