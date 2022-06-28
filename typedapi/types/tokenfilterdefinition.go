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

// TokenFilterDefinition holds the union for the following types:
//     AsciiFoldingTokenFilter
//     CommonGramsTokenFilter
//     ConditionTokenFilter
//     DelimitedPayloadTokenFilter
//     DictionaryDecompounderTokenFilter
//     EdgeNGramTokenFilter
//     ElisionTokenFilter
//     FingerprintTokenFilter
//     HunspellTokenFilter
//     HyphenationDecompounderTokenFilter
//     IcuCollationTokenFilter
//     IcuFoldingTokenFilter
//     IcuNormalizationTokenFilter
//     IcuTokenizer
//     IcuTransformTokenFilter
//     KStemTokenFilter
//     KeepTypesTokenFilter
//     KeepWordsTokenFilter
//     KeywordMarkerTokenFilter
//     KuromojiPartOfSpeechTokenFilter
//     KuromojiReadingFormTokenFilter
//     KuromojiStemmerTokenFilter
//     LengthTokenFilter
//     LimitTokenCountTokenFilter
//     LowercaseTokenFilter
//     MultiplexerTokenFilter
//     NGramTokenFilter
//     NoriPartOfSpeechTokenFilter
//     PatternCaptureTokenFilter
//     PatternReplaceTokenFilter
//     PhoneticTokenFilter
//     PorterStemTokenFilter
//     PredicateTokenFilter
//     RemoveDuplicatesTokenFilter
//     ReverseTokenFilter
//     ShingleTokenFilter
//     SnowballTokenFilter
//     StemmerOverrideTokenFilter
//     StemmerTokenFilter
//     StopTokenFilter
//     SynonymGraphTokenFilter
//     SynonymTokenFilter
//     TrimTokenFilter
//     TruncateTokenFilter
//     UniqueTokenFilter
//     UppercaseTokenFilter
//     WordDelimiterGraphTokenFilter
//     WordDelimiterTokenFilter
type TokenFilterDefinition interface{}

// TokenFilterDefinitionBuilder holds TokenFilterDefinition struct and provides a builder API.
type TokenFilterDefinitionBuilder struct {
	v TokenFilterDefinition
}

// NewTokenFilterDefinition provides a builder for the TokenFilterDefinition struct.
func NewTokenFilterDefinition() *TokenFilterDefinitionBuilder {
	return &TokenFilterDefinitionBuilder{}
}

// Build finalize the chain and returns the TokenFilterDefinition struct
func (u *TokenFilterDefinitionBuilder) Build() TokenFilterDefinition {
	return u.v
}

// AsciiFoldingTokenFilter set the AsciiFoldingTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) AsciiFoldingTokenFilter(v AsciiFoldingTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// CommonGramsTokenFilter set the CommonGramsTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) CommonGramsTokenFilter(v CommonGramsTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// ConditionTokenFilter set the ConditionTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) ConditionTokenFilter(v ConditionTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// DelimitedPayloadTokenFilter set the DelimitedPayloadTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) DelimitedPayloadTokenFilter(v DelimitedPayloadTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// DictionaryDecompounderTokenFilter set the DictionaryDecompounderTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) DictionaryDecompounderTokenFilter(v DictionaryDecompounderTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// EdgeNGramTokenFilter set the EdgeNGramTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) EdgeNGramTokenFilter(v EdgeNGramTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// ElisionTokenFilter set the ElisionTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) ElisionTokenFilter(v ElisionTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// FingerprintTokenFilter set the FingerprintTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) FingerprintTokenFilter(v FingerprintTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// HunspellTokenFilter set the HunspellTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) HunspellTokenFilter(v HunspellTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// HyphenationDecompounderTokenFilter set the HyphenationDecompounderTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) HyphenationDecompounderTokenFilter(v HyphenationDecompounderTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// IcuCollationTokenFilter set the IcuCollationTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) IcuCollationTokenFilter(v IcuCollationTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// IcuFoldingTokenFilter set the IcuFoldingTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) IcuFoldingTokenFilter(v IcuFoldingTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// IcuNormalizationTokenFilter set the IcuNormalizationTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) IcuNormalizationTokenFilter(v IcuNormalizationTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// IcuTokenizer set the IcuTokenizer property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) IcuTokenizer(v IcuTokenizer) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// IcuTransformTokenFilter set the IcuTransformTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) IcuTransformTokenFilter(v IcuTransformTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KStemTokenFilter set the KStemTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KStemTokenFilter(v KStemTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KeepTypesTokenFilter set the KeepTypesTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KeepTypesTokenFilter(v KeepTypesTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KeepWordsTokenFilter set the KeepWordsTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KeepWordsTokenFilter(v KeepWordsTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KeywordMarkerTokenFilter set the KeywordMarkerTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KeywordMarkerTokenFilter(v KeywordMarkerTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KuromojiPartOfSpeechTokenFilter set the KuromojiPartOfSpeechTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KuromojiPartOfSpeechTokenFilter(v KuromojiPartOfSpeechTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KuromojiReadingFormTokenFilter set the KuromojiReadingFormTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KuromojiReadingFormTokenFilter(v KuromojiReadingFormTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// KuromojiStemmerTokenFilter set the KuromojiStemmerTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) KuromojiStemmerTokenFilter(v KuromojiStemmerTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// LengthTokenFilter set the LengthTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) LengthTokenFilter(v LengthTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// LimitTokenCountTokenFilter set the LimitTokenCountTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) LimitTokenCountTokenFilter(v LimitTokenCountTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// LowercaseTokenFilter set the LowercaseTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) LowercaseTokenFilter(v LowercaseTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// MultiplexerTokenFilter set the MultiplexerTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) MultiplexerTokenFilter(v MultiplexerTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// NGramTokenFilter set the NGramTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) NGramTokenFilter(v NGramTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// NoriPartOfSpeechTokenFilter set the NoriPartOfSpeechTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) NoriPartOfSpeechTokenFilter(v NoriPartOfSpeechTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// PatternCaptureTokenFilter set the PatternCaptureTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) PatternCaptureTokenFilter(v PatternCaptureTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// PatternReplaceTokenFilter set the PatternReplaceTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) PatternReplaceTokenFilter(v PatternReplaceTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// PhoneticTokenFilter set the PhoneticTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) PhoneticTokenFilter(v PhoneticTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// PorterStemTokenFilter set the PorterStemTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) PorterStemTokenFilter(v PorterStemTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// PredicateTokenFilter set the PredicateTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) PredicateTokenFilter(v PredicateTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// RemoveDuplicatesTokenFilter set the RemoveDuplicatesTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) RemoveDuplicatesTokenFilter(v RemoveDuplicatesTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// ReverseTokenFilter set the ReverseTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) ReverseTokenFilter(v ReverseTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// ShingleTokenFilter set the ShingleTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) ShingleTokenFilter(v ShingleTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// SnowballTokenFilter set the SnowballTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) SnowballTokenFilter(v SnowballTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// StemmerOverrideTokenFilter set the StemmerOverrideTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) StemmerOverrideTokenFilter(v StemmerOverrideTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// StemmerTokenFilter set the StemmerTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) StemmerTokenFilter(v StemmerTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// StopTokenFilter set the StopTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) StopTokenFilter(v StopTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// SynonymGraphTokenFilter set the SynonymGraphTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) SynonymGraphTokenFilter(v SynonymGraphTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// SynonymTokenFilter set the SynonymTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) SynonymTokenFilter(v SynonymTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// TrimTokenFilter set the TrimTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) TrimTokenFilter(v TrimTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// TruncateTokenFilter set the TruncateTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) TruncateTokenFilter(v TruncateTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// UniqueTokenFilter set the UniqueTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) UniqueTokenFilter(v UniqueTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// UppercaseTokenFilter set the UppercaseTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) UppercaseTokenFilter(v UppercaseTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// WordDelimiterGraphTokenFilter set the WordDelimiterGraphTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) WordDelimiterGraphTokenFilter(v WordDelimiterGraphTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}

// WordDelimiterTokenFilter set the WordDelimiterTokenFilter property for TokenFilterDefinitionBuilder.
func (u *TokenFilterDefinitionBuilder) WordDelimiterTokenFilter(v WordDelimiterTokenFilter) *TokenFilterDefinitionBuilder {
	u.v = v
	return u
}
