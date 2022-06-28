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

// JobConfig type.
type JobConfig struct {
	AllowLazyOpen                        *bool            `json:"allow_lazy_open,omitempty"`
	AnalysisConfig                       AnalysisConfig   `json:"analysis_config"`
	AnalysisLimits                       *AnalysisLimits  `json:"analysis_limits,omitempty"`
	BackgroundPersistInterval            *Time            `json:"background_persist_interval,omitempty"`
	CustomSettings                       *CustomSettings  `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays *int64           `json:"daily_model_snapshot_retention_after_days,omitempty"`
	DataDescription                      DataDescription  `json:"data_description"`
	DatafeedConfig                       *DatafeedConfig  `json:"datafeed_config,omitempty"`
	Description                          *string          `json:"description,omitempty"`
	Groups                               []string         `json:"groups,omitempty"`
	JobId                                *Id              `json:"job_id,omitempty"`
	JobType                              *string          `json:"job_type,omitempty"`
	ModelPlotConfig                      *ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelSnapshotRetentionDays           *int64           `json:"model_snapshot_retention_days,omitempty"`
	RenormalizationWindowDays            *int64           `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     *IndexName       `json:"results_index_name,omitempty"`
	ResultsRetentionDays                 *int64           `json:"results_retention_days,omitempty"`
}

// JobConfigBuilder holds JobConfig struct and provides a builder API.
type JobConfigBuilder struct {
	v *JobConfig
}

// NewJobConfig provides a builder for the JobConfig struct.
func NewJobConfig() *JobConfigBuilder {
	r := JobConfigBuilder{
		&JobConfig{},
	}

	return &r
}

// Build finalize the chain and returns the JobConfig struct
func (rb *JobConfigBuilder) Build() JobConfig {
	return *rb.v
}

// AllowLazyOpen set the AllowLazyOpen property for JobConfigBuilder.
func (rb *JobConfigBuilder) AllowLazyOpen(allowlazyopen bool) *JobConfigBuilder {
	rb.v.AllowLazyOpen = &allowlazyopen
	return rb
}

// AnalysisConfig set the AnalysisConfig property for JobConfigBuilder.
func (rb *JobConfigBuilder) AnalysisConfig(analysisconfig AnalysisConfig) *JobConfigBuilder {
	rb.v.AnalysisConfig = analysisconfig
	return rb
}

// AnalysisLimits set the AnalysisLimits property for JobConfigBuilder.
func (rb *JobConfigBuilder) AnalysisLimits(analysislimits AnalysisLimits) *JobConfigBuilder {
	rb.v.AnalysisLimits = &analysislimits
	return rb
}

// BackgroundPersistInterval set the BackgroundPersistInterval property for JobConfigBuilder.
func (rb *JobConfigBuilder) BackgroundPersistInterval(backgroundpersistinterval Time) *JobConfigBuilder {
	rb.v.BackgroundPersistInterval = &backgroundpersistinterval
	return rb
}

// CustomSettings set the CustomSettings property for JobConfigBuilder.
func (rb *JobConfigBuilder) CustomSettings(customsettings CustomSettings) *JobConfigBuilder {
	rb.v.CustomSettings = &customsettings
	return rb
}

// DailyModelSnapshotRetentionAfterDays set the DailyModelSnapshotRetentionAfterDays property for JobConfigBuilder.
func (rb *JobConfigBuilder) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *JobConfigBuilder {
	rb.v.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays
	return rb
}

// DataDescription set the DataDescription property for JobConfigBuilder.
func (rb *JobConfigBuilder) DataDescription(datadescription DataDescription) *JobConfigBuilder {
	rb.v.DataDescription = datadescription
	return rb
}

// DatafeedConfig set the DatafeedConfig property for JobConfigBuilder.
func (rb *JobConfigBuilder) DatafeedConfig(datafeedconfig DatafeedConfig) *JobConfigBuilder {
	rb.v.DatafeedConfig = &datafeedconfig
	return rb
}

// Description set the Description property for JobConfigBuilder.
func (rb *JobConfigBuilder) Description(description string) *JobConfigBuilder {
	rb.v.Description = &description
	return rb
}

// Groups set the Groups property for JobConfigBuilder.
func (rb *JobConfigBuilder) Groups(groups ...string) *JobConfigBuilder {
	rb.v.Groups = append(rb.v.Groups, groups...)
	return rb
}

// JobId set the JobId property for JobConfigBuilder.
func (rb *JobConfigBuilder) JobId(jobid Id) *JobConfigBuilder {
	rb.v.JobId = &jobid
	return rb
}

// JobType set the JobType property for JobConfigBuilder.
func (rb *JobConfigBuilder) JobType(jobtype string) *JobConfigBuilder {
	rb.v.JobType = &jobtype
	return rb
}

// ModelPlotConfig set the ModelPlotConfig property for JobConfigBuilder.
func (rb *JobConfigBuilder) ModelPlotConfig(modelplotconfig ModelPlotConfig) *JobConfigBuilder {
	rb.v.ModelPlotConfig = &modelplotconfig
	return rb
}

// ModelSnapshotRetentionDays set the ModelSnapshotRetentionDays property for JobConfigBuilder.
func (rb *JobConfigBuilder) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *JobConfigBuilder {
	rb.v.ModelSnapshotRetentionDays = &modelsnapshotretentiondays
	return rb
}

// RenormalizationWindowDays set the RenormalizationWindowDays property for JobConfigBuilder.
func (rb *JobConfigBuilder) RenormalizationWindowDays(renormalizationwindowdays int64) *JobConfigBuilder {
	rb.v.RenormalizationWindowDays = &renormalizationwindowdays
	return rb
}

// ResultsIndexName set the ResultsIndexName property for JobConfigBuilder.
func (rb *JobConfigBuilder) ResultsIndexName(resultsindexname IndexName) *JobConfigBuilder {
	rb.v.ResultsIndexName = &resultsindexname
	return rb
}

// ResultsRetentionDays set the ResultsRetentionDays property for JobConfigBuilder.
func (rb *JobConfigBuilder) ResultsRetentionDays(resultsretentiondays int64) *JobConfigBuilder {
	rb.v.ResultsRetentionDays = &resultsretentiondays
	return rb
}
