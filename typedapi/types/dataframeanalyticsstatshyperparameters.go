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

// DataframeAnalyticsStatsHyperparameters type.
type DataframeAnalyticsStatsHyperparameters struct {
	Hyperparameters Hyperparameters `json:"hyperparameters"`
	// Iteration The number of iterations on the analysis.
	Iteration      int            `json:"iteration"`
	Timestamp      DateString     `json:"timestamp"`
	TimingStats    TimingStats    `json:"timing_stats"`
	ValidationLoss ValidationLoss `json:"validation_loss"`
}

// DataframeAnalyticsStatsHyperparametersBuilder holds DataframeAnalyticsStatsHyperparameters struct and provides a builder API.
type DataframeAnalyticsStatsHyperparametersBuilder struct {
	v *DataframeAnalyticsStatsHyperparameters
}

// NewDataframeAnalyticsStatsHyperparameters provides a builder for the DataframeAnalyticsStatsHyperparameters struct.
func NewDataframeAnalyticsStatsHyperparameters() *DataframeAnalyticsStatsHyperparametersBuilder {
	r := DataframeAnalyticsStatsHyperparametersBuilder{
		&DataframeAnalyticsStatsHyperparameters{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsHyperparameters struct
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Build() DataframeAnalyticsStatsHyperparameters {
	return *rb.v
}

// Hyperparameters set the Hyperparameters property for DataframeAnalyticsStatsHyperparametersBuilder.
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Hyperparameters(hyperparameters Hyperparameters) *DataframeAnalyticsStatsHyperparametersBuilder {
	rb.v.Hyperparameters = hyperparameters
	return rb
}

// Iteration The number of iterations on the analysis.
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Iteration(iteration int) *DataframeAnalyticsStatsHyperparametersBuilder {
	rb.v.Iteration = iteration
	return rb
}

// Timestamp set the Timestamp property for DataframeAnalyticsStatsHyperparametersBuilder.
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) Timestamp(timestamp DateString) *DataframeAnalyticsStatsHyperparametersBuilder {
	rb.v.Timestamp = timestamp
	return rb
}

// TimingStats set the TimingStats property for DataframeAnalyticsStatsHyperparametersBuilder.
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) TimingStats(timingstats TimingStats) *DataframeAnalyticsStatsHyperparametersBuilder {
	rb.v.TimingStats = timingstats
	return rb
}

// ValidationLoss set the ValidationLoss property for DataframeAnalyticsStatsHyperparametersBuilder.
func (rb *DataframeAnalyticsStatsHyperparametersBuilder) ValidationLoss(validationloss ValidationLoss) *DataframeAnalyticsStatsHyperparametersBuilder {
	rb.v.ValidationLoss = validationloss
	return rb
}
