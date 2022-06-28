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

// Percentiles holds the union for the following types:
//     []ArrayPercentilesItem
//     KeyedPercentiles
type Percentiles interface{}

// PercentilesBuilder holds Percentiles struct and provides a builder API.
type PercentilesBuilder struct {
	v Percentiles
}

// NewPercentiles provides a builder for the Percentiles struct.
func NewPercentiles() *PercentilesBuilder {
	return &PercentilesBuilder{}
}

// Build finalize the chain and returns the Percentiles struct
func (u *PercentilesBuilder) Build() Percentiles {
	return u.v
}

// ArrayPercentilesItems set the ArrayPercentilesItems property for PercentilesBuilder.
func (u *PercentilesBuilder) ArrayPercentilesItems(v []ArrayPercentilesItem) *PercentilesBuilder {
	u.v = v
	return u
}

// KeyedPercentiles set the KeyedPercentiles property for PercentilesBuilder.
func (u *PercentilesBuilder) KeyedPercentiles(v KeyedPercentiles) *PercentilesBuilder {
	u.v = v
	return u
}
