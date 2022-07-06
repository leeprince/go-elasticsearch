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

// Process type.
type Process struct {
	Cpu                 *Cpu         `json:"cpu,omitempty"`
	MaxFileDescriptors  *int         `json:"max_file_descriptors,omitempty"`
	Mem                 *MemoryStats `json:"mem,omitempty"`
	OpenFileDescriptors *int         `json:"open_file_descriptors,omitempty"`
	Timestamp           *int64       `json:"timestamp,omitempty"`
}

// ProcessBuilder holds Process struct and provides a builder API.
type ProcessBuilder struct {
	v *Process
}

// NewProcess provides a builder for the Process struct.
func NewProcess() *ProcessBuilder {
	r := ProcessBuilder{
		&Process{},
	}

	return &r
}

// Build finalize the chain and returns the Process struct
func (rb *ProcessBuilder) Build() Process {
	return *rb.v
}

// Cpu set the Cpu property for ProcessBuilder.
func (rb *ProcessBuilder) Cpu(cpu Cpu) *ProcessBuilder {
	rb.v.Cpu = &cpu
	return rb
}

// MaxFileDescriptors set the MaxFileDescriptors property for ProcessBuilder.
func (rb *ProcessBuilder) MaxFileDescriptors(maxfiledescriptors int) *ProcessBuilder {
	rb.v.MaxFileDescriptors = &maxfiledescriptors
	return rb
}

// Mem set the Mem property for ProcessBuilder.
func (rb *ProcessBuilder) Mem(mem MemoryStats) *ProcessBuilder {
	rb.v.Mem = &mem
	return rb
}

// OpenFileDescriptors set the OpenFileDescriptors property for ProcessBuilder.
func (rb *ProcessBuilder) OpenFileDescriptors(openfiledescriptors int) *ProcessBuilder {
	rb.v.OpenFileDescriptors = &openfiledescriptors
	return rb
}

// Timestamp set the Timestamp property for ProcessBuilder.
func (rb *ProcessBuilder) Timestamp(timestamp int64) *ProcessBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}
