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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/numericfielddataformat"
)

// NumericFielddata type.
type NumericFielddata struct {
	Format numericfielddataformat.NumericFielddataFormat `json:"format"`
}

// NumericFielddataBuilder holds NumericFielddata struct and provides a builder API.
type NumericFielddataBuilder struct {
	v *NumericFielddata
}

// NewNumericFielddata provides a builder for the NumericFielddata struct.
func NewNumericFielddata() *NumericFielddataBuilder {
	r := NumericFielddataBuilder{
		&NumericFielddata{},
	}

	return &r
}

// Build finalize the chain and returns the NumericFielddata struct
func (rb *NumericFielddataBuilder) Build() NumericFielddata {
	return *rb.v
}

// Format set the Format property for NumericFielddataBuilder.
func (rb *NumericFielddataBuilder) Format(format numericfielddataformat.NumericFielddataFormat) *NumericFielddataBuilder {
	rb.v.Format = format
	return rb
}
