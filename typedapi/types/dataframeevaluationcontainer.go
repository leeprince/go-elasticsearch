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

// DataframeEvaluationContainer type.
type DataframeEvaluationContainer struct {
	// Classification Classification evaluation evaluates the results of a classification analysis
	// which outputs a prediction that identifies to which of the classes each
	// document belongs.
	Classification *DataframeEvaluationClassification `json:"classification,omitempty"`
	// OutlierDetection Outlier detection evaluates the results of an outlier detection analysis
	// which outputs the probability that each document is an outlier.
	OutlierDetection *DataframeEvaluationOutlierDetection `json:"outlier_detection,omitempty"`
	// Regression Regression evaluation evaluates the results of a regression analysis which
	// outputs a prediction of values.
	Regression *DataframeEvaluationRegression `json:"regression,omitempty"`
}

// DataframeEvaluationContainerBuilder holds DataframeEvaluationContainer struct and provides a builder API.
type DataframeEvaluationContainerBuilder struct {
	v *DataframeEvaluationContainer
}

// NewDataframeEvaluationContainer provides a builder for the DataframeEvaluationContainer struct.
func NewDataframeEvaluationContainer() *DataframeEvaluationContainerBuilder {
	r := DataframeEvaluationContainerBuilder{
		&DataframeEvaluationContainer{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationContainer struct
func (rb *DataframeEvaluationContainerBuilder) Build() DataframeEvaluationContainer {
	return *rb.v
}

// Classification Classification evaluation evaluates the results of a classification analysis
// which outputs a prediction that identifies to which of the classes each
// document belongs.
func (rb *DataframeEvaluationContainerBuilder) Classification(classification DataframeEvaluationClassification) *DataframeEvaluationContainerBuilder {
	rb.v.Classification = &classification
	return rb
}

// OutlierDetection Outlier detection evaluates the results of an outlier detection analysis
// which outputs the probability that each document is an outlier.
func (rb *DataframeEvaluationContainerBuilder) OutlierDetection(outlierdetection DataframeEvaluationOutlierDetection) *DataframeEvaluationContainerBuilder {
	rb.v.OutlierDetection = &outlierdetection
	return rb
}

// Regression Regression evaluation evaluates the results of a regression analysis which
// outputs a prediction of values.
func (rb *DataframeEvaluationContainerBuilder) Regression(regression DataframeEvaluationRegression) *DataframeEvaluationContainerBuilder {
	rb.v.Regression = &regression
	return rb
}
