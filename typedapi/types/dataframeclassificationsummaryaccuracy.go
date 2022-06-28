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

// DataframeClassificationSummaryAccuracy type.
type DataframeClassificationSummaryAccuracy struct {
	Classes         []DataframeEvaluationClass `json:"classes"`
	OverallAccuracy float64                    `json:"overall_accuracy"`
}

// DataframeClassificationSummaryAccuracyBuilder holds DataframeClassificationSummaryAccuracy struct and provides a builder API.
type DataframeClassificationSummaryAccuracyBuilder struct {
	v *DataframeClassificationSummaryAccuracy
}

// NewDataframeClassificationSummaryAccuracy provides a builder for the DataframeClassificationSummaryAccuracy struct.
func NewDataframeClassificationSummaryAccuracy() *DataframeClassificationSummaryAccuracyBuilder {
	r := DataframeClassificationSummaryAccuracyBuilder{
		&DataframeClassificationSummaryAccuracy{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummaryAccuracy struct
func (rb *DataframeClassificationSummaryAccuracyBuilder) Build() DataframeClassificationSummaryAccuracy {
	return *rb.v
}

// Classes set the Classes property for DataframeClassificationSummaryAccuracyBuilder.
func (rb *DataframeClassificationSummaryAccuracyBuilder) Classes(classes ...DataframeEvaluationClass) *DataframeClassificationSummaryAccuracyBuilder {
	rb.v.Classes = append(rb.v.Classes, classes...)
	return rb
}

// OverallAccuracy set the OverallAccuracy property for DataframeClassificationSummaryAccuracyBuilder.
func (rb *DataframeClassificationSummaryAccuracyBuilder) OverallAccuracy(overallaccuracy float64) *DataframeClassificationSummaryAccuracyBuilder {
	rb.v.OverallAccuracy = overallaccuracy
	return rb
}
