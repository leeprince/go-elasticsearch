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

// ErrorCause type.
type ErrorCause struct {
	CausedBy *ErrorCause            `json:"caused_by,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Reason A human-readable explanation of the error, in english
	Reason    string       `json:"reason"`
	RootCause []ErrorCause `json:"root_cause,omitempty"`
	// StackTrace The server stack trace. Present only if the `error_trace=true` parameter was
	// sent with the request.
	StackTrace *string      `json:"stack_trace,omitempty"`
	Suppressed []ErrorCause `json:"suppressed,omitempty"`
	// Type The type of error
	Type string `json:"type"`
}

// ErrorCauseBuilder holds ErrorCause struct and provides a builder API.
type ErrorCauseBuilder struct {
	v *ErrorCause
}

// NewErrorCause provides a builder for the ErrorCause struct.
func NewErrorCause() *ErrorCauseBuilder {
	r := ErrorCauseBuilder{
		&ErrorCause{
			Metadata: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ErrorCause struct
func (rb *ErrorCauseBuilder) Build() ErrorCause {
	return *rb.v
}

// CausedBy set the CausedBy property for ErrorCauseBuilder.
func (rb *ErrorCauseBuilder) CausedBy(causedby ErrorCause) *ErrorCauseBuilder {
	rb.v.CausedBy = &causedby
	return rb
}

// Metadata set the Metadata property for ErrorCauseBuilder.
func (rb *ErrorCauseBuilder) Metadata(value map[string]interface{}) *ErrorCauseBuilder {
	rb.v.Metadata = value
	return rb
}

// Reason A human-readable explanation of the error, in english
func (rb *ErrorCauseBuilder) Reason(reason string) *ErrorCauseBuilder {
	rb.v.Reason = reason
	return rb
}

// RootCause set the RootCause property for ErrorCauseBuilder.
func (rb *ErrorCauseBuilder) RootCause(root_cause ...ErrorCause) *ErrorCauseBuilder {
	rb.v.RootCause = append(rb.v.RootCause, root_cause...)
	return rb
}

// StackTrace The server stack trace. Present only if the `error_trace=true` parameter was
// sent with the request.
func (rb *ErrorCauseBuilder) StackTrace(stacktrace string) *ErrorCauseBuilder {
	rb.v.StackTrace = &stacktrace
	return rb
}

// Suppressed set the Suppressed property for ErrorCauseBuilder.
func (rb *ErrorCauseBuilder) Suppressed(suppressed ...ErrorCause) *ErrorCauseBuilder {
	rb.v.Suppressed = append(rb.v.Suppressed, suppressed...)
	return rb
}

// Type The type of error
func (rb *ErrorCauseBuilder) Type_(type_ string) *ErrorCauseBuilder {
	rb.v.Type = type_
	return rb
}
