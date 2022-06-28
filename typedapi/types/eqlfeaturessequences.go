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

// EqlFeaturesSequences type.
type EqlFeaturesSequences struct {
	SequenceMaxspan           uint `json:"sequence_maxspan"`
	SequenceQueriesFiveOrMore uint `json:"sequence_queries_five_or_more"`
	SequenceQueriesFour       uint `json:"sequence_queries_four"`
	SequenceQueriesThree      uint `json:"sequence_queries_three"`
	SequenceQueriesTwo        uint `json:"sequence_queries_two"`
	SequenceUntil             uint `json:"sequence_until"`
}

// EqlFeaturesSequencesBuilder holds EqlFeaturesSequences struct and provides a builder API.
type EqlFeaturesSequencesBuilder struct {
	v *EqlFeaturesSequences
}

// NewEqlFeaturesSequences provides a builder for the EqlFeaturesSequences struct.
func NewEqlFeaturesSequences() *EqlFeaturesSequencesBuilder {
	r := EqlFeaturesSequencesBuilder{
		&EqlFeaturesSequences{},
	}

	return &r
}

// Build finalize the chain and returns the EqlFeaturesSequences struct
func (rb *EqlFeaturesSequencesBuilder) Build() EqlFeaturesSequences {
	return *rb.v
}

// SequenceMaxspan set the SequenceMaxspan property for EqlFeaturesSequencesBuilder.
func (rb *EqlFeaturesSequencesBuilder) SequenceMaxspan(sequencemaxspan uint) *EqlFeaturesSequencesBuilder {
	rb.v.SequenceMaxspan = sequencemaxspan
	return rb
}

// SequenceQueriesFiveOrMore set the SequenceQueriesFiveOrMore property for EqlFeaturesSequencesBuilder.
func (rb *EqlFeaturesSequencesBuilder) SequenceQueriesFiveOrMore(sequencequeriesfiveormore uint) *EqlFeaturesSequencesBuilder {
	rb.v.SequenceQueriesFiveOrMore = sequencequeriesfiveormore
	return rb
}

// SequenceQueriesFour set the SequenceQueriesFour property for EqlFeaturesSequencesBuilder.
func (rb *EqlFeaturesSequencesBuilder) SequenceQueriesFour(sequencequeriesfour uint) *EqlFeaturesSequencesBuilder {
	rb.v.SequenceQueriesFour = sequencequeriesfour
	return rb
}

// SequenceQueriesThree set the SequenceQueriesThree property for EqlFeaturesSequencesBuilder.
func (rb *EqlFeaturesSequencesBuilder) SequenceQueriesThree(sequencequeriesthree uint) *EqlFeaturesSequencesBuilder {
	rb.v.SequenceQueriesThree = sequencequeriesthree
	return rb
}

// SequenceQueriesTwo set the SequenceQueriesTwo property for EqlFeaturesSequencesBuilder.
func (rb *EqlFeaturesSequencesBuilder) SequenceQueriesTwo(sequencequeriestwo uint) *EqlFeaturesSequencesBuilder {
	rb.v.SequenceQueriesTwo = sequencequeriestwo
	return rb
}

// SequenceUntil set the SequenceUntil property for EqlFeaturesSequencesBuilder.
func (rb *EqlFeaturesSequencesBuilder) SequenceUntil(sequenceuntil uint) *EqlFeaturesSequencesBuilder {
	rb.v.SequenceUntil = sequenceuntil
	return rb
}
