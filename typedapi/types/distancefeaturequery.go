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

// DistanceFeatureQuery holds the union for the following types:
//     DateDistanceFeatureQuery
//     GeoDistanceFeatureQuery
type DistanceFeatureQuery interface{}

// DistanceFeatureQueryBuilder holds DistanceFeatureQuery struct and provides a builder API.
type DistanceFeatureQueryBuilder struct {
	v DistanceFeatureQuery
}

// NewDistanceFeatureQuery provides a builder for the DistanceFeatureQuery struct.
func NewDistanceFeatureQuery() *DistanceFeatureQueryBuilder {
	return &DistanceFeatureQueryBuilder{}
}

// Build finalize the chain and returns the DistanceFeatureQuery struct
func (u *DistanceFeatureQueryBuilder) Build() DistanceFeatureQuery {
	return u.v
}

// DateDistanceFeatureQuery set the DateDistanceFeatureQuery property for DistanceFeatureQueryBuilder.
func (u *DistanceFeatureQueryBuilder) DateDistanceFeatureQuery(v DateDistanceFeatureQuery) *DistanceFeatureQueryBuilder {
	u.v = v
	return u
}

// GeoDistanceFeatureQuery set the GeoDistanceFeatureQuery property for DistanceFeatureQueryBuilder.
func (u *DistanceFeatureQueryBuilder) GeoDistanceFeatureQuery(v GeoDistanceFeatureQuery) *DistanceFeatureQueryBuilder {
	u.v = v
	return u
}
