package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityListRequest msg type = x.
type SecurityListRequest struct {
	message.Message
}

//SecurityListRequestBuilder builds SecurityListRequest messages.
type SecurityListRequestBuilder struct {
	message.MessageBuilder
}

//NewSecurityListRequestBuilder returns an initialized SecurityListRequestBuilder with specified required fields.
func NewSecurityListRequestBuilder(
	securityreqid field.SecurityReqID,
	securitylistrequesttype field.SecurityListRequestType) *SecurityListRequestBuilder {
	builder := new(SecurityListRequestBuilder)
	builder.Body.Set(securityreqid)
	builder.Body.Set(securitylistrequesttype)
	return builder
}

//SecurityReqID is a required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListRequestType is a required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityListRequestType() (*field.SecurityListRequestType, error) {
	f := new(field.SecurityListRequestType)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for SecurityListRequest.
func (m *SecurityListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
