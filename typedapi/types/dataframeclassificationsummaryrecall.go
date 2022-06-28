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

// DataframeClassificationSummaryRecall type.
type DataframeClassificationSummaryRecall struct {
	AvgRecall float64                    `json:"avg_recall"`
	Classes   []DataframeEvaluationClass `json:"classes"`
}

// DataframeClassificationSummaryRecallBuilder holds DataframeClassificationSummaryRecall struct and provides a builder API.
type DataframeClassificationSummaryRecallBuilder struct {
	v *DataframeClassificationSummaryRecall
}

// NewDataframeClassificationSummaryRecall provides a builder for the DataframeClassificationSummaryRecall struct.
func NewDataframeClassificationSummaryRecall() *DataframeClassificationSummaryRecallBuilder {
	r := DataframeClassificationSummaryRecallBuilder{
		&DataframeClassificationSummaryRecall{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummaryRecall struct
func (rb *DataframeClassificationSummaryRecallBuilder) Build() DataframeClassificationSummaryRecall {
	return *rb.v
}

// AvgRecall set the AvgRecall property for DataframeClassificationSummaryRecallBuilder.
func (rb *DataframeClassificationSummaryRecallBuilder) AvgRecall(avgrecall float64) *DataframeClassificationSummaryRecallBuilder {
	rb.v.AvgRecall = avgrecall
	return rb
}

// Classes set the Classes property for DataframeClassificationSummaryRecallBuilder.
func (rb *DataframeClassificationSummaryRecallBuilder) Classes(classes ...DataframeEvaluationClass) *DataframeClassificationSummaryRecallBuilder {
	rb.v.Classes = append(rb.v.Classes, classes...)
	return rb
}
