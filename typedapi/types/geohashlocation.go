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

// GeoHashLocation type.
type GeoHashLocation struct {
	Geohash GeoHash `json:"geohash"`
}

// GeoHashLocationBuilder holds GeoHashLocation struct and provides a builder API.
type GeoHashLocationBuilder struct {
	v *GeoHashLocation
}

// NewGeoHashLocation provides a builder for the GeoHashLocation struct.
func NewGeoHashLocation() *GeoHashLocationBuilder {
	r := GeoHashLocationBuilder{
		&GeoHashLocation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoHashLocation struct
func (rb *GeoHashLocationBuilder) Build() GeoHashLocation {
	return *rb.v
}

// Geohash set the Geohash property for GeoHashLocationBuilder.
func (rb *GeoHashLocationBuilder) Geohash(geohash GeoHash) *GeoHashLocationBuilder {
	rb.v.Geohash = geohash
	return rb
}
