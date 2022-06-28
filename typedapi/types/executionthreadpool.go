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

// ExecutionThreadPool type.
type ExecutionThreadPool struct {
	MaxSize   int64 `json:"max_size"`
	QueueSize int64 `json:"queue_size"`
}

// ExecutionThreadPoolBuilder holds ExecutionThreadPool struct and provides a builder API.
type ExecutionThreadPoolBuilder struct {
	v *ExecutionThreadPool
}

// NewExecutionThreadPool provides a builder for the ExecutionThreadPool struct.
func NewExecutionThreadPool() *ExecutionThreadPoolBuilder {
	r := ExecutionThreadPoolBuilder{
		&ExecutionThreadPool{},
	}

	return &r
}

// Build finalize the chain and returns the ExecutionThreadPool struct
func (rb *ExecutionThreadPoolBuilder) Build() ExecutionThreadPool {
	return *rb.v
}

// MaxSize set the MaxSize property for ExecutionThreadPoolBuilder.
func (rb *ExecutionThreadPoolBuilder) MaxSize(maxsize int64) *ExecutionThreadPoolBuilder {
	rb.v.MaxSize = maxsize
	return rb
}

// QueueSize set the QueueSize property for ExecutionThreadPoolBuilder.
func (rb *ExecutionThreadPoolBuilder) QueueSize(queuesize int64) *ExecutionThreadPoolBuilder {
	rb.v.QueueSize = queuesize
	return rb
}
