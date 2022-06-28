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

// RollupJobConfiguration type.
type RollupJobConfiguration struct {
	Cron         string        `json:"cron"`
	Groups       Groupings     `json:"groups"`
	Id           Id            `json:"id"`
	IndexPattern string        `json:"index_pattern"`
	Metrics      []FieldMetric `json:"metrics"`
	PageSize     int64         `json:"page_size"`
	RollupIndex  IndexName     `json:"rollup_index"`
	Timeout      Time          `json:"timeout"`
}

// RollupJobConfigurationBuilder holds RollupJobConfiguration struct and provides a builder API.
type RollupJobConfigurationBuilder struct {
	v *RollupJobConfiguration
}

// NewRollupJobConfiguration provides a builder for the RollupJobConfiguration struct.
func NewRollupJobConfiguration() *RollupJobConfigurationBuilder {
	r := RollupJobConfigurationBuilder{
		&RollupJobConfiguration{},
	}

	return &r
}

// Build finalize the chain and returns the RollupJobConfiguration struct
func (rb *RollupJobConfigurationBuilder) Build() RollupJobConfiguration {
	return *rb.v
}

// Cron set the Cron property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) Cron(cron string) *RollupJobConfigurationBuilder {
	rb.v.Cron = cron
	return rb
}

// Groups set the Groups property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) Groups(groups Groupings) *RollupJobConfigurationBuilder {
	rb.v.Groups = groups
	return rb
}

// Id set the Id property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) Id(id Id) *RollupJobConfigurationBuilder {
	rb.v.Id = id
	return rb
}

// IndexPattern set the IndexPattern property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) IndexPattern(indexpattern string) *RollupJobConfigurationBuilder {
	rb.v.IndexPattern = indexpattern
	return rb
}

// Metrics set the Metrics property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) Metrics(metrics ...FieldMetric) *RollupJobConfigurationBuilder {
	rb.v.Metrics = append(rb.v.Metrics, metrics...)
	return rb
}

// PageSize set the PageSize property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) PageSize(pagesize int64) *RollupJobConfigurationBuilder {
	rb.v.PageSize = pagesize
	return rb
}

// RollupIndex set the RollupIndex property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) RollupIndex(rollupindex IndexName) *RollupJobConfigurationBuilder {
	rb.v.RollupIndex = rollupindex
	return rb
}

// Timeout set the Timeout property for RollupJobConfigurationBuilder.
func (rb *RollupJobConfigurationBuilder) Timeout(timeout Time) *RollupJobConfigurationBuilder {
	rb.v.Timeout = timeout
	return rb
}
