package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//News msg type = B.
type News struct {
	message.Message
}

//NewsBuilder builds News messages.
type NewsBuilder struct {
	message.MessageBuilder
}

//NewNewsBuilder returns an initialized NewsBuilder with specified required fields.
func NewNewsBuilder(
	headline field.Headline,
	nolinesoftext field.NoLinesOfText) *NewsBuilder {
	builder := new(NewsBuilder)
	builder.Body.Set(headline)
	builder.Body.Set(nolinesoftext)
	return builder
}

//OrigTime is a non-required field for News.
func (m *News) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}

//Urgency is a non-required field for News.
func (m *News) Urgency() (*field.Urgency, error) {
	f := new(field.Urgency)
	err := m.Body.Get(f)
	return f, err
}

//Headline is a required field for News.
func (m *News) Headline() (*field.Headline, error) {
	f := new(field.Headline)
	err := m.Body.Get(f)
	return f, err
}

//EncodedHeadlineLen is a non-required field for News.
func (m *News) EncodedHeadlineLen() (*field.EncodedHeadlineLen, error) {
	f := new(field.EncodedHeadlineLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedHeadline is a non-required field for News.
func (m *News) EncodedHeadline() (*field.EncodedHeadline, error) {
	f := new(field.EncodedHeadline)
	err := m.Body.Get(f)
	return f, err
}

//NoRoutingIDs is a non-required field for News.
func (m *News) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for News.
func (m *News) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for News.
func (m *News) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for News.
func (m *News) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLinesOfText is a required field for News.
func (m *News) NoLinesOfText() (*field.NoLinesOfText, error) {
	f := new(field.NoLinesOfText)
	err := m.Body.Get(f)
	return f, err
}

//URLLink is a non-required field for News.
func (m *News) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}

//RawDataLength is a non-required field for News.
func (m *News) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//RawData is a non-required field for News.
func (m *News) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
