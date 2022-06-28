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

// DataframeClassificationSummaryMulticlassConfusionMatrix type.
type DataframeClassificationSummaryMulticlassConfusionMatrix struct {
	ConfusionMatrix       []ConfusionMatrixItem `json:"confusion_matrix"`
	OtherActualClassCount int                   `json:"other_actual_class_count"`
}

// DataframeClassificationSummaryMulticlassConfusionMatrixBuilder holds DataframeClassificationSummaryMulticlassConfusionMatrix struct and provides a builder API.
type DataframeClassificationSummaryMulticlassConfusionMatrixBuilder struct {
	v *DataframeClassificationSummaryMulticlassConfusionMatrix
}

// NewDataframeClassificationSummaryMulticlassConfusionMatrix provides a builder for the DataframeClassificationSummaryMulticlassConfusionMatrix struct.
func NewDataframeClassificationSummaryMulticlassConfusionMatrix() *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder {
	r := DataframeClassificationSummaryMulticlassConfusionMatrixBuilder{
		&DataframeClassificationSummaryMulticlassConfusionMatrix{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeClassificationSummaryMulticlassConfusionMatrix struct
func (rb *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) Build() DataframeClassificationSummaryMulticlassConfusionMatrix {
	return *rb.v
}

// ConfusionMatrix set the ConfusionMatrix property for DataframeClassificationSummaryMulticlassConfusionMatrixBuilder.
func (rb *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) ConfusionMatrix(confusion_matrix ...ConfusionMatrixItem) *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder {
	rb.v.ConfusionMatrix = append(rb.v.ConfusionMatrix, confusion_matrix...)
	return rb
}

// OtherActualClassCount set the OtherActualClassCount property for DataframeClassificationSummaryMulticlassConfusionMatrixBuilder.
func (rb *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder) OtherActualClassCount(otheractualclasscount int) *DataframeClassificationSummaryMulticlassConfusionMatrixBuilder {
	rb.v.OtherActualClassCount = otheractualclasscount
	return rb
}
