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

// Property holds the union for the following types:
//     AggregateMetricDoubleProperty
//     BinaryProperty
//     BooleanProperty
//     ByteNumberProperty
//     CompletionProperty
//     ConstantKeywordProperty
//     DateNanosProperty
//     DateProperty
//     DateRangeProperty
//     DenseVectorProperty
//     DoubleNumberProperty
//     DoubleRangeProperty
//     DynamicProperty
//     FieldAliasProperty
//     FlattenedProperty
//     FloatNumberProperty
//     FloatRangeProperty
//     GeoPointProperty
//     GeoShapeProperty
//     HalfFloatNumberProperty
//     HistogramProperty
//     IntegerNumberProperty
//     IntegerRangeProperty
//     IpProperty
//     IpRangeProperty
//     JoinProperty
//     KeywordProperty
//     LongNumberProperty
//     LongRangeProperty
//     MatchOnlyTextProperty
//     Murmur3HashProperty
//     NestedProperty
//     ObjectProperty
//     PercolatorProperty
//     PointProperty
//     RankFeatureProperty
//     RankFeaturesProperty
//     ScaledFloatNumberProperty
//     SearchAsYouTypeProperty
//     ShapeProperty
//     ShortNumberProperty
//     TextProperty
//     TokenCountProperty
//     UnsignedLongNumberProperty
//     VersionProperty
//     WildcardProperty
type Property interface{}

// PropertyBuilder holds Property struct and provides a builder API.
type PropertyBuilder struct {
	v Property
}

// NewProperty provides a builder for the Property struct.
func NewProperty() *PropertyBuilder {
	return &PropertyBuilder{}
}

// Build finalize the chain and returns the Property struct
func (u *PropertyBuilder) Build() Property {
	return u.v
}

// AggregateMetricDoubleProperty set the AggregateMetricDoubleProperty property for PropertyBuilder.
func (u *PropertyBuilder) AggregateMetricDoubleProperty(v AggregateMetricDoubleProperty) *PropertyBuilder {
	u.v = v
	return u
}

// BinaryProperty set the BinaryProperty property for PropertyBuilder.
func (u *PropertyBuilder) BinaryProperty(v BinaryProperty) *PropertyBuilder {
	u.v = v
	return u
}

// BooleanProperty set the BooleanProperty property for PropertyBuilder.
func (u *PropertyBuilder) BooleanProperty(v BooleanProperty) *PropertyBuilder {
	u.v = v
	return u
}

// ByteNumberProperty set the ByteNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) ByteNumberProperty(v ByteNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// CompletionProperty set the CompletionProperty property for PropertyBuilder.
func (u *PropertyBuilder) CompletionProperty(v CompletionProperty) *PropertyBuilder {
	u.v = v
	return u
}

// ConstantKeywordProperty set the ConstantKeywordProperty property for PropertyBuilder.
func (u *PropertyBuilder) ConstantKeywordProperty(v ConstantKeywordProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DateNanosProperty set the DateNanosProperty property for PropertyBuilder.
func (u *PropertyBuilder) DateNanosProperty(v DateNanosProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DateProperty set the DateProperty property for PropertyBuilder.
func (u *PropertyBuilder) DateProperty(v DateProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DateRangeProperty set the DateRangeProperty property for PropertyBuilder.
func (u *PropertyBuilder) DateRangeProperty(v DateRangeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DenseVectorProperty set the DenseVectorProperty property for PropertyBuilder.
func (u *PropertyBuilder) DenseVectorProperty(v DenseVectorProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DoubleNumberProperty set the DoubleNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) DoubleNumberProperty(v DoubleNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DoubleRangeProperty set the DoubleRangeProperty property for PropertyBuilder.
func (u *PropertyBuilder) DoubleRangeProperty(v DoubleRangeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// DynamicProperty set the DynamicProperty property for PropertyBuilder.
func (u *PropertyBuilder) DynamicProperty(v DynamicProperty) *PropertyBuilder {
	u.v = v
	return u
}

// FieldAliasProperty set the FieldAliasProperty property for PropertyBuilder.
func (u *PropertyBuilder) FieldAliasProperty(v FieldAliasProperty) *PropertyBuilder {
	u.v = v
	return u
}

// FlattenedProperty set the FlattenedProperty property for PropertyBuilder.
func (u *PropertyBuilder) FlattenedProperty(v FlattenedProperty) *PropertyBuilder {
	u.v = v
	return u
}

// FloatNumberProperty set the FloatNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) FloatNumberProperty(v FloatNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// FloatRangeProperty set the FloatRangeProperty property for PropertyBuilder.
func (u *PropertyBuilder) FloatRangeProperty(v FloatRangeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// GeoPointProperty set the GeoPointProperty property for PropertyBuilder.
func (u *PropertyBuilder) GeoPointProperty(v GeoPointProperty) *PropertyBuilder {
	u.v = v
	return u
}

// GeoShapeProperty set the GeoShapeProperty property for PropertyBuilder.
func (u *PropertyBuilder) GeoShapeProperty(v GeoShapeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// HalfFloatNumberProperty set the HalfFloatNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) HalfFloatNumberProperty(v HalfFloatNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// HistogramProperty set the HistogramProperty property for PropertyBuilder.
func (u *PropertyBuilder) HistogramProperty(v HistogramProperty) *PropertyBuilder {
	u.v = v
	return u
}

// IntegerNumberProperty set the IntegerNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) IntegerNumberProperty(v IntegerNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// IntegerRangeProperty set the IntegerRangeProperty property for PropertyBuilder.
func (u *PropertyBuilder) IntegerRangeProperty(v IntegerRangeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// IpProperty set the IpProperty property for PropertyBuilder.
func (u *PropertyBuilder) IpProperty(v IpProperty) *PropertyBuilder {
	u.v = v
	return u
}

// IpRangeProperty set the IpRangeProperty property for PropertyBuilder.
func (u *PropertyBuilder) IpRangeProperty(v IpRangeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// JoinProperty set the JoinProperty property for PropertyBuilder.
func (u *PropertyBuilder) JoinProperty(v JoinProperty) *PropertyBuilder {
	u.v = v
	return u
}

// KeywordProperty set the KeywordProperty property for PropertyBuilder.
func (u *PropertyBuilder) KeywordProperty(v KeywordProperty) *PropertyBuilder {
	u.v = v
	return u
}

// LongNumberProperty set the LongNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) LongNumberProperty(v LongNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// LongRangeProperty set the LongRangeProperty property for PropertyBuilder.
func (u *PropertyBuilder) LongRangeProperty(v LongRangeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// MatchOnlyTextProperty set the MatchOnlyTextProperty property for PropertyBuilder.
func (u *PropertyBuilder) MatchOnlyTextProperty(v MatchOnlyTextProperty) *PropertyBuilder {
	u.v = v
	return u
}

// Murmur3HashProperty set the Murmur3HashProperty property for PropertyBuilder.
func (u *PropertyBuilder) Murmur3HashProperty(v Murmur3HashProperty) *PropertyBuilder {
	u.v = v
	return u
}

// NestedProperty set the NestedProperty property for PropertyBuilder.
func (u *PropertyBuilder) NestedProperty(v NestedProperty) *PropertyBuilder {
	u.v = v
	return u
}

// ObjectProperty set the ObjectProperty property for PropertyBuilder.
func (u *PropertyBuilder) ObjectProperty(v ObjectProperty) *PropertyBuilder {
	u.v = v
	return u
}

// PercolatorProperty set the PercolatorProperty property for PropertyBuilder.
func (u *PropertyBuilder) PercolatorProperty(v PercolatorProperty) *PropertyBuilder {
	u.v = v
	return u
}

// PointProperty set the PointProperty property for PropertyBuilder.
func (u *PropertyBuilder) PointProperty(v PointProperty) *PropertyBuilder {
	u.v = v
	return u
}

// RankFeatureProperty set the RankFeatureProperty property for PropertyBuilder.
func (u *PropertyBuilder) RankFeatureProperty(v RankFeatureProperty) *PropertyBuilder {
	u.v = v
	return u
}

// RankFeaturesProperty set the RankFeaturesProperty property for PropertyBuilder.
func (u *PropertyBuilder) RankFeaturesProperty(v RankFeaturesProperty) *PropertyBuilder {
	u.v = v
	return u
}

// ScaledFloatNumberProperty set the ScaledFloatNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) ScaledFloatNumberProperty(v ScaledFloatNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// SearchAsYouTypeProperty set the SearchAsYouTypeProperty property for PropertyBuilder.
func (u *PropertyBuilder) SearchAsYouTypeProperty(v SearchAsYouTypeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// ShapeProperty set the ShapeProperty property for PropertyBuilder.
func (u *PropertyBuilder) ShapeProperty(v ShapeProperty) *PropertyBuilder {
	u.v = v
	return u
}

// ShortNumberProperty set the ShortNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) ShortNumberProperty(v ShortNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// TextProperty set the TextProperty property for PropertyBuilder.
func (u *PropertyBuilder) TextProperty(v TextProperty) *PropertyBuilder {
	u.v = v
	return u
}

// TokenCountProperty set the TokenCountProperty property for PropertyBuilder.
func (u *PropertyBuilder) TokenCountProperty(v TokenCountProperty) *PropertyBuilder {
	u.v = v
	return u
}

// UnsignedLongNumberProperty set the UnsignedLongNumberProperty property for PropertyBuilder.
func (u *PropertyBuilder) UnsignedLongNumberProperty(v UnsignedLongNumberProperty) *PropertyBuilder {
	u.v = v
	return u
}

// VersionProperty set the VersionProperty property for PropertyBuilder.
func (u *PropertyBuilder) VersionProperty(v VersionProperty) *PropertyBuilder {
	u.v = v
	return u
}

// WildcardProperty set the WildcardProperty property for PropertyBuilder.
func (u *PropertyBuilder) WildcardProperty(v WildcardProperty) *PropertyBuilder {
	u.v = v
	return u
}
