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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/childscoremode"
)

// HasChildQuery type.
type HasChildQuery struct {
	Boost          *float32                       `json:"boost,omitempty"`
	IgnoreUnmapped *bool                          `json:"ignore_unmapped,omitempty"`
	InnerHits      *InnerHits                     `json:"inner_hits,omitempty"`
	MaxChildren    *int                           `json:"max_children,omitempty"`
	MinChildren    *int                           `json:"min_children,omitempty"`
	Query          *QueryContainer                `json:"query,omitempty"`
	QueryName_     *string                        `json:"_name,omitempty"`
	ScoreMode      *childscoremode.ChildScoreMode `json:"score_mode,omitempty"`
	Type           RelationName                   `json:"type"`
}

// HasChildQueryBuilder holds HasChildQuery struct and provides a builder API.
type HasChildQueryBuilder struct {
	v *HasChildQuery
}

// NewHasChildQuery provides a builder for the HasChildQuery struct.
func NewHasChildQuery() *HasChildQueryBuilder {
	r := HasChildQueryBuilder{
		&HasChildQuery{},
	}

	return &r
}

// Build finalize the chain and returns the HasChildQuery struct
func (rb *HasChildQueryBuilder) Build() HasChildQuery {
	return *rb.v
}

// Boost set the Boost property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) Boost(boost float32) *HasChildQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// IgnoreUnmapped set the IgnoreUnmapped property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *HasChildQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

// InnerHits set the InnerHits property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) InnerHits(innerhits InnerHits) *HasChildQueryBuilder {
	rb.v.InnerHits = &innerhits
	return rb
}

// MaxChildren set the MaxChildren property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) MaxChildren(maxchildren int) *HasChildQueryBuilder {
	rb.v.MaxChildren = &maxchildren
	return rb
}

// MinChildren set the MinChildren property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) MinChildren(minchildren int) *HasChildQueryBuilder {
	rb.v.MinChildren = &minchildren
	return rb
}

// Query set the Query property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) Query(query QueryContainer) *HasChildQueryBuilder {
	rb.v.Query = &query
	return rb
}

// QueryName_ set the QueryName_ property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) QueryName_(queryname_ string) *HasChildQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// ScoreMode set the ScoreMode property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) ScoreMode(scoremode childscoremode.ChildScoreMode) *HasChildQueryBuilder {
	rb.v.ScoreMode = &scoremode
	return rb
}

// Type set the Type property for HasChildQueryBuilder.
func (rb *HasChildQueryBuilder) Type_(type_ RelationName) *HasChildQueryBuilder {
	rb.v.Type = type_
	return rb
}
