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

// FileSystem type.
type FileSystem struct {
	Data      []DataPathStats  `json:"data,omitempty"`
	IoStats   *IoStats         `json:"io_stats,omitempty"`
	Timestamp *int64           `json:"timestamp,omitempty"`
	Total     *FileSystemTotal `json:"total,omitempty"`
}

// FileSystemBuilder holds FileSystem struct and provides a builder API.
type FileSystemBuilder struct {
	v *FileSystem
}

// NewFileSystem provides a builder for the FileSystem struct.
func NewFileSystem() *FileSystemBuilder {
	r := FileSystemBuilder{
		&FileSystem{},
	}

	return &r
}

// Build finalize the chain and returns the FileSystem struct
func (rb *FileSystemBuilder) Build() FileSystem {
	return *rb.v
}

// Data set the Data property for FileSystemBuilder.
func (rb *FileSystemBuilder) Data(data ...DataPathStats) *FileSystemBuilder {
	rb.v.Data = append(rb.v.Data, data...)
	return rb
}

// IoStats set the IoStats property for FileSystemBuilder.
func (rb *FileSystemBuilder) IoStats(iostats IoStats) *FileSystemBuilder {
	rb.v.IoStats = &iostats
	return rb
}

// Timestamp set the Timestamp property for FileSystemBuilder.
func (rb *FileSystemBuilder) Timestamp(timestamp int64) *FileSystemBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// Total set the Total property for FileSystemBuilder.
func (rb *FileSystemBuilder) Total(total FileSystemTotal) *FileSystemBuilder {
	rb.v.Total = &total
	return rb
}
