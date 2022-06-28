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

// Job type.
type Job struct {
	AllowLazyOpen                        bool             `json:"allow_lazy_open"`
	AnalysisConfig                       AnalysisConfig   `json:"analysis_config"`
	AnalysisLimits                       *AnalysisLimits  `json:"analysis_limits,omitempty"`
	BackgroundPersistInterval            *Time            `json:"background_persist_interval,omitempty"`
	Blocked                              *JobBlocked      `json:"blocked,omitempty"`
	CreateTime                           *int             `json:"create_time,omitempty"`
	CustomSettings                       *CustomSettings  `json:"custom_settings,omitempty"`
	DailyModelSnapshotRetentionAfterDays *int64           `json:"daily_model_snapshot_retention_after_days,omitempty"`
	DataDescription                      DataDescription  `json:"data_description"`
	DatafeedConfig                       *Datafeed        `json:"datafeed_config,omitempty"`
	Deleting                             *bool            `json:"deleting,omitempty"`
	Description                          *string          `json:"description,omitempty"`
	FinishedTime                         *int             `json:"finished_time,omitempty"`
	Groups                               []string         `json:"groups,omitempty"`
	JobId                                Id               `json:"job_id"`
	JobType                              *string          `json:"job_type,omitempty"`
	JobVersion                           *VersionString   `json:"job_version,omitempty"`
	ModelPlotConfig                      *ModelPlotConfig `json:"model_plot_config,omitempty"`
	ModelSnapshotId                      *Id              `json:"model_snapshot_id,omitempty"`
	ModelSnapshotRetentionDays           int64            `json:"model_snapshot_retention_days"`
	RenormalizationWindowDays            *int64           `json:"renormalization_window_days,omitempty"`
	ResultsIndexName                     IndexName        `json:"results_index_name"`
	ResultsRetentionDays                 *int64           `json:"results_retention_days,omitempty"`
}

// JobBuilder holds Job struct and provides a builder API.
type JobBuilder struct {
	v *Job
}

// NewJob provides a builder for the Job struct.
func NewJob() *JobBuilder {
	r := JobBuilder{
		&Job{},
	}

	return &r
}

// Build finalize the chain and returns the Job struct
func (rb *JobBuilder) Build() Job {
	return *rb.v
}

// AllowLazyOpen set the AllowLazyOpen property for JobBuilder.
func (rb *JobBuilder) AllowLazyOpen(allowlazyopen bool) *JobBuilder {
	rb.v.AllowLazyOpen = allowlazyopen
	return rb
}

// AnalysisConfig set the AnalysisConfig property for JobBuilder.
func (rb *JobBuilder) AnalysisConfig(analysisconfig AnalysisConfig) *JobBuilder {
	rb.v.AnalysisConfig = analysisconfig
	return rb
}

// AnalysisLimits set the AnalysisLimits property for JobBuilder.
func (rb *JobBuilder) AnalysisLimits(analysislimits AnalysisLimits) *JobBuilder {
	rb.v.AnalysisLimits = &analysislimits
	return rb
}

// BackgroundPersistInterval set the BackgroundPersistInterval property for JobBuilder.
func (rb *JobBuilder) BackgroundPersistInterval(backgroundpersistinterval Time) *JobBuilder {
	rb.v.BackgroundPersistInterval = &backgroundpersistinterval
	return rb
}

// Blocked set the Blocked property for JobBuilder.
func (rb *JobBuilder) Blocked(blocked JobBlocked) *JobBuilder {
	rb.v.Blocked = &blocked
	return rb
}

// CreateTime set the CreateTime property for JobBuilder.
func (rb *JobBuilder) CreateTime(createtime int) *JobBuilder {
	rb.v.CreateTime = &createtime
	return rb
}

// CustomSettings set the CustomSettings property for JobBuilder.
func (rb *JobBuilder) CustomSettings(customsettings CustomSettings) *JobBuilder {
	rb.v.CustomSettings = &customsettings
	return rb
}

// DailyModelSnapshotRetentionAfterDays set the DailyModelSnapshotRetentionAfterDays property for JobBuilder.
func (rb *JobBuilder) DailyModelSnapshotRetentionAfterDays(dailymodelsnapshotretentionafterdays int64) *JobBuilder {
	rb.v.DailyModelSnapshotRetentionAfterDays = &dailymodelsnapshotretentionafterdays
	return rb
}

// DataDescription set the DataDescription property for JobBuilder.
func (rb *JobBuilder) DataDescription(datadescription DataDescription) *JobBuilder {
	rb.v.DataDescription = datadescription
	return rb
}

// DatafeedConfig set the DatafeedConfig property for JobBuilder.
func (rb *JobBuilder) DatafeedConfig(datafeedconfig Datafeed) *JobBuilder {
	rb.v.DatafeedConfig = &datafeedconfig
	return rb
}

// Deleting set the Deleting property for JobBuilder.
func (rb *JobBuilder) Deleting(deleting bool) *JobBuilder {
	rb.v.Deleting = &deleting
	return rb
}

// Description set the Description property for JobBuilder.
func (rb *JobBuilder) Description(description string) *JobBuilder {
	rb.v.Description = &description
	return rb
}

// FinishedTime set the FinishedTime property for JobBuilder.
func (rb *JobBuilder) FinishedTime(finishedtime int) *JobBuilder {
	rb.v.FinishedTime = &finishedtime
	return rb
}

// Groups set the Groups property for JobBuilder.
func (rb *JobBuilder) Groups(groups ...string) *JobBuilder {
	rb.v.Groups = append(rb.v.Groups, groups...)
	return rb
}

// JobId set the JobId property for JobBuilder.
func (rb *JobBuilder) JobId(jobid Id) *JobBuilder {
	rb.v.JobId = jobid
	return rb
}

// JobType set the JobType property for JobBuilder.
func (rb *JobBuilder) JobType(jobtype string) *JobBuilder {
	rb.v.JobType = &jobtype
	return rb
}

// JobVersion set the JobVersion property for JobBuilder.
func (rb *JobBuilder) JobVersion(jobversion VersionString) *JobBuilder {
	rb.v.JobVersion = &jobversion
	return rb
}

// ModelPlotConfig set the ModelPlotConfig property for JobBuilder.
func (rb *JobBuilder) ModelPlotConfig(modelplotconfig ModelPlotConfig) *JobBuilder {
	rb.v.ModelPlotConfig = &modelplotconfig
	return rb
}

// ModelSnapshotId set the ModelSnapshotId property for JobBuilder.
func (rb *JobBuilder) ModelSnapshotId(modelsnapshotid Id) *JobBuilder {
	rb.v.ModelSnapshotId = &modelsnapshotid
	return rb
}

// ModelSnapshotRetentionDays set the ModelSnapshotRetentionDays property for JobBuilder.
func (rb *JobBuilder) ModelSnapshotRetentionDays(modelsnapshotretentiondays int64) *JobBuilder {
	rb.v.ModelSnapshotRetentionDays = modelsnapshotretentiondays
	return rb
}

// RenormalizationWindowDays set the RenormalizationWindowDays property for JobBuilder.
func (rb *JobBuilder) RenormalizationWindowDays(renormalizationwindowdays int64) *JobBuilder {
	rb.v.RenormalizationWindowDays = &renormalizationwindowdays
	return rb
}

// ResultsIndexName set the ResultsIndexName property for JobBuilder.
func (rb *JobBuilder) ResultsIndexName(resultsindexname IndexName) *JobBuilder {
	rb.v.ResultsIndexName = resultsindexname
	return rb
}

// ResultsRetentionDays set the ResultsRetentionDays property for JobBuilder.
func (rb *JobBuilder) ResultsRetentionDays(resultsretentiondays int64) *JobBuilder {
	rb.v.ResultsRetentionDays = &resultsretentiondays
	return rb
}
