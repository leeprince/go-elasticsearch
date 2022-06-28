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

// ClusterProcessor type.
type ClusterProcessor struct {
	Count        int64 `json:"count"`
	Current      int64 `json:"current"`
	Failed       int64 `json:"failed"`
	TimeInMillis int64 `json:"time_in_millis"`
}

// ClusterProcessorBuilder holds ClusterProcessor struct and provides a builder API.
type ClusterProcessorBuilder struct {
	v *ClusterProcessor
}

// NewClusterProcessor provides a builder for the ClusterProcessor struct.
func NewClusterProcessor() *ClusterProcessorBuilder {
	r := ClusterProcessorBuilder{
		&ClusterProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterProcessor struct
func (rb *ClusterProcessorBuilder) Build() ClusterProcessor {
	return *rb.v
}

// Count set the Count property for ClusterProcessorBuilder.
func (rb *ClusterProcessorBuilder) Count(count int64) *ClusterProcessorBuilder {
	rb.v.Count = count
	return rb
}

// Current set the Current property for ClusterProcessorBuilder.
func (rb *ClusterProcessorBuilder) Current(current int64) *ClusterProcessorBuilder {
	rb.v.Current = current
	return rb
}

// Failed set the Failed property for ClusterProcessorBuilder.
func (rb *ClusterProcessorBuilder) Failed(failed int64) *ClusterProcessorBuilder {
	rb.v.Failed = failed
	return rb
}

// TimeInMillis set the TimeInMillis property for ClusterProcessorBuilder.
func (rb *ClusterProcessorBuilder) TimeInMillis(timeinmillis int64) *ClusterProcessorBuilder {
	rb.v.TimeInMillis = timeinmillis
	return rb
}
