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

// GeoHashPrecision holds the union for the following types:
//     int
//     string
type GeoHashPrecision interface{}

// GeoHashPrecisionBuilder holds GeoHashPrecision struct and provides a builder API.
type GeoHashPrecisionBuilder struct {
	v GeoHashPrecision
}

// NewGeoHashPrecision provides a builder for the GeoHashPrecision struct.
func NewGeoHashPrecision() *GeoHashPrecisionBuilder {
	return &GeoHashPrecisionBuilder{}
}

// Build finalize the chain and returns the GeoHashPrecision struct
func (u *GeoHashPrecisionBuilder) Build() GeoHashPrecision {
	return u.v
}

// Int set the Int property for GeoHashPrecisionBuilder.
func (u *GeoHashPrecisionBuilder) Int(v int) *GeoHashPrecisionBuilder {
	u.v = v
	return u
}

// String set the String property for GeoHashPrecisionBuilder.
func (u *GeoHashPrecisionBuilder) String(v string) *GeoHashPrecisionBuilder {
	u.v = v
	return u
}
