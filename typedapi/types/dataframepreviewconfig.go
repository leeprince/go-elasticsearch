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

// DataframePreviewConfig type.
type DataframePreviewConfig struct {
	Analysis         DataframeAnalysisContainer       `json:"analysis"`
	AnalyzedFields   *DataframeAnalysisAnalyzedFields `json:"analyzed_fields,omitempty"`
	MaxNumThreads    *int                             `json:"max_num_threads,omitempty"`
	ModelMemoryLimit *string                          `json:"model_memory_limit,omitempty"`
	Source           DataframeAnalyticsSource         `json:"source"`
}

// DataframePreviewConfigBuilder holds DataframePreviewConfig struct and provides a builder API.
type DataframePreviewConfigBuilder struct {
	v *DataframePreviewConfig
}

// NewDataframePreviewConfig provides a builder for the DataframePreviewConfig struct.
func NewDataframePreviewConfig() *DataframePreviewConfigBuilder {
	r := DataframePreviewConfigBuilder{
		&DataframePreviewConfig{},
	}

	return &r
}

// Build finalize the chain and returns the DataframePreviewConfig struct
func (rb *DataframePreviewConfigBuilder) Build() DataframePreviewConfig {
	return *rb.v
}

// Analysis set the Analysis property for DataframePreviewConfigBuilder.
func (rb *DataframePreviewConfigBuilder) Analysis(analysis DataframeAnalysisContainer) *DataframePreviewConfigBuilder {
	rb.v.Analysis = analysis
	return rb
}

// AnalyzedFields set the AnalyzedFields property for DataframePreviewConfigBuilder.
func (rb *DataframePreviewConfigBuilder) AnalyzedFields(analyzedfields DataframeAnalysisAnalyzedFields) *DataframePreviewConfigBuilder {
	rb.v.AnalyzedFields = &analyzedfields
	return rb
}

// MaxNumThreads set the MaxNumThreads property for DataframePreviewConfigBuilder.
func (rb *DataframePreviewConfigBuilder) MaxNumThreads(maxnumthreads int) *DataframePreviewConfigBuilder {
	rb.v.MaxNumThreads = &maxnumthreads
	return rb
}

// ModelMemoryLimit set the ModelMemoryLimit property for DataframePreviewConfigBuilder.
func (rb *DataframePreviewConfigBuilder) ModelMemoryLimit(modelmemorylimit string) *DataframePreviewConfigBuilder {
	rb.v.ModelMemoryLimit = &modelmemorylimit
	return rb
}

// Source set the Source property for DataframePreviewConfigBuilder.
func (rb *DataframePreviewConfigBuilder) Source(source DataframeAnalyticsSource) *DataframePreviewConfigBuilder {
	rb.v.Source = source
	return rb
}
