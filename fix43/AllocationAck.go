package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationAck msg type = P.
type AllocationAck struct {
	message.Message
}

//AllocationAckBuilder builds AllocationAck messages.
type AllocationAckBuilder struct {
	message.MessageBuilder
}

//NewAllocationAckBuilder returns an initialized AllocationAckBuilder with specified required fields.
func NewAllocationAckBuilder(
	allocid field.AllocID,
	tradedate field.TradeDate,
	allocstatus field.AllocStatus) *AllocationAckBuilder {
	builder := new(AllocationAckBuilder)
	builder.Body.Set(allocid)
	builder.Body.Set(tradedate)
	builder.Body.Set(allocstatus)
	return builder
}

//NoPartyIDs is a non-required field for AllocationAck.
func (m *AllocationAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//AllocID is a required field for AllocationAck.
func (m *AllocationAck) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for AllocationAck.
func (m *AllocationAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for AllocationAck.
func (m *AllocationAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//AllocStatus is a required field for AllocationAck.
func (m *AllocationAck) AllocStatus() (*field.AllocStatus, error) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationAck.
func (m *AllocationAck) AllocRejCode() (*field.AllocRejCode, error) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for AllocationAck.
func (m *AllocationAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationAck.
func (m *AllocationAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for AllocationAck.
func (m *AllocationAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//LegalConfirm is a non-required field for AllocationAck.
func (m *AllocationAck) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}
