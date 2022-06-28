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

// TopClassEntry type.
type TopClassEntry struct {
	ClassName        string  `json:"class_name"`
	ClassProbability float64 `json:"class_probability"`
	ClassScore       float64 `json:"class_score"`
}

// TopClassEntryBuilder holds TopClassEntry struct and provides a builder API.
type TopClassEntryBuilder struct {
	v *TopClassEntry
}

// NewTopClassEntry provides a builder for the TopClassEntry struct.
func NewTopClassEntry() *TopClassEntryBuilder {
	r := TopClassEntryBuilder{
		&TopClassEntry{},
	}

	return &r
}

// Build finalize the chain and returns the TopClassEntry struct
func (rb *TopClassEntryBuilder) Build() TopClassEntry {
	return *rb.v
}

// ClassName set the ClassName property for TopClassEntryBuilder.
func (rb *TopClassEntryBuilder) ClassName(classname string) *TopClassEntryBuilder {
	rb.v.ClassName = classname
	return rb
}

// ClassProbability set the ClassProbability property for TopClassEntryBuilder.
func (rb *TopClassEntryBuilder) ClassProbability(classprobability float64) *TopClassEntryBuilder {
	rb.v.ClassProbability = classprobability
	return rb
}

// ClassScore set the ClassScore property for TopClassEntryBuilder.
func (rb *TopClassEntryBuilder) ClassScore(classscore float64) *TopClassEntryBuilder {
	rb.v.ClassScore = classscore
	return rb
}
