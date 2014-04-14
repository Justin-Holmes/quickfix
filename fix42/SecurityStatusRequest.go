package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityStatusRequest msg type = e.
type SecurityStatusRequest struct {
	message.Message
}

//SecurityStatusRequestBuilder builds SecurityStatusRequest messages.
type SecurityStatusRequestBuilder struct {
	message.MessageBuilder
}

//NewSecurityStatusRequestBuilder returns an initialized SecurityStatusRequestBuilder with specified required fields.
func NewSecurityStatusRequestBuilder(
	securitystatusreqid field.SecurityStatusReqID,
	symbol field.Symbol,
	subscriptionrequesttype field.SubscriptionRequestType) *SecurityStatusRequestBuilder {
	builder := new(SecurityStatusRequestBuilder)
	builder.Body.Set(securitystatusreqid)
	builder.Body.Set(symbol)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//SecurityStatusReqID is a required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SecurityStatusReqID() (*field.SecurityStatusReqID, error) {
	f := new(field.SecurityStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityStatusRequest.
func (m *SecurityStatusRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
