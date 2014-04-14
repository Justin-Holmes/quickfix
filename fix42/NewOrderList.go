package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderList msg type = E.
type NewOrderList struct {
	message.Message
}

//NewOrderListBuilder builds NewOrderList messages.
type NewOrderListBuilder struct {
	message.MessageBuilder
}

//NewNewOrderListBuilder returns an initialized NewOrderListBuilder with specified required fields.
func NewNewOrderListBuilder(
	listid field.ListID,
	bidtype field.BidType,
	totnoorders field.TotNoOrders,
	noorders field.NoOrders) *NewOrderListBuilder {
	builder := new(NewOrderListBuilder)
	builder.Body.Set(listid)
	builder.Body.Set(bidtype)
	builder.Body.Set(totnoorders)
	builder.Body.Set(noorders)
	return builder
}

//ListID is a required field for NewOrderList.
func (m *NewOrderList) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//BidID is a non-required field for NewOrderList.
func (m *NewOrderList) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//ClientBidID is a non-required field for NewOrderList.
func (m *NewOrderList) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m *NewOrderList) ProgRptReqs() (*field.ProgRptReqs, error) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}

//BidType is a required field for NewOrderList.
func (m *NewOrderList) BidType() (*field.BidType, error) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m *NewOrderList) ProgPeriodInterval() (*field.ProgPeriodInterval, error) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}

//ListExecInstType is a non-required field for NewOrderList.
func (m *NewOrderList) ListExecInstType() (*field.ListExecInstType, error) {
	f := new(field.ListExecInstType)
	err := m.Body.Get(f)
	return f, err
}

//ListExecInst is a non-required field for NewOrderList.
func (m *NewOrderList) ListExecInst() (*field.ListExecInst, error) {
	f := new(field.ListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m *NewOrderList) EncodedListExecInstLen() (*field.EncodedListExecInstLen, error) {
	f := new(field.EncodedListExecInstLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m *NewOrderList) EncodedListExecInst() (*field.EncodedListExecInst, error) {
	f := new(field.EncodedListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//TotNoOrders is a required field for NewOrderList.
func (m *NewOrderList) TotNoOrders() (*field.TotNoOrders, error) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a required field for NewOrderList.
func (m *NewOrderList) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
