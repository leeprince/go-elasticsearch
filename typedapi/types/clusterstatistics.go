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

// ClusterStatistics type.
type ClusterStatistics struct {
	Skipped    int `json:"skipped"`
	Successful int `json:"successful"`
	Total      int `json:"total"`
}

// ClusterStatisticsBuilder holds ClusterStatistics struct and provides a builder API.
type ClusterStatisticsBuilder struct {
	v *ClusterStatistics
}

// NewClusterStatistics provides a builder for the ClusterStatistics struct.
func NewClusterStatistics() *ClusterStatisticsBuilder {
	r := ClusterStatisticsBuilder{
		&ClusterStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterStatistics struct
func (rb *ClusterStatisticsBuilder) Build() ClusterStatistics {
	return *rb.v
}

// Skipped set the Skipped property for ClusterStatisticsBuilder.
func (rb *ClusterStatisticsBuilder) Skipped(skipped int) *ClusterStatisticsBuilder {
	rb.v.Skipped = skipped
	return rb
}

// Successful set the Successful property for ClusterStatisticsBuilder.
func (rb *ClusterStatisticsBuilder) Successful(successful int) *ClusterStatisticsBuilder {
	rb.v.Successful = successful
	return rb
}

// Total set the Total property for ClusterStatisticsBuilder.
func (rb *ClusterStatisticsBuilder) Total(total int) *ClusterStatisticsBuilder {
	rb.v.Total = total
	return rb
}
