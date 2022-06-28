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

// ReadException type.
type ReadException struct {
	Exception ErrorCause     `json:"exception"`
	FromSeqNo SequenceNumber `json:"from_seq_no"`
	Retries   int            `json:"retries"`
}

// ReadExceptionBuilder holds ReadException struct and provides a builder API.
type ReadExceptionBuilder struct {
	v *ReadException
}

// NewReadException provides a builder for the ReadException struct.
func NewReadException() *ReadExceptionBuilder {
	r := ReadExceptionBuilder{
		&ReadException{},
	}

	return &r
}

// Build finalize the chain and returns the ReadException struct
func (rb *ReadExceptionBuilder) Build() ReadException {
	return *rb.v
}

// Exception set the Exception property for ReadExceptionBuilder.
func (rb *ReadExceptionBuilder) Exception(exception ErrorCause) *ReadExceptionBuilder {
	rb.v.Exception = exception
	return rb
}

// FromSeqNo set the FromSeqNo property for ReadExceptionBuilder.
func (rb *ReadExceptionBuilder) FromSeqNo(fromseqno SequenceNumber) *ReadExceptionBuilder {
	rb.v.FromSeqNo = fromseqno
	return rb
}

// Retries set the Retries property for ReadExceptionBuilder.
func (rb *ReadExceptionBuilder) Retries(retries int) *ReadExceptionBuilder {
	rb.v.Retries = retries
	return rb
}
