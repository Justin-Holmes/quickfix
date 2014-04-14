package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteStatusReport msg type = AI.
type QuoteStatusReport struct {
	message.Message
}

//QuoteStatusReportBuilder builds QuoteStatusReport messages.
type QuoteStatusReportBuilder struct {
	message.MessageBuilder
}

//NewQuoteStatusReportBuilder returns an initialized QuoteStatusReportBuilder with specified required fields.
func NewQuoteStatusReportBuilder(
	quoteid field.QuoteID) *QuoteStatusReportBuilder {
	builder := new(QuoteStatusReportBuilder)
	builder.Body.Set(quoteid)
	return builder
}

//QuoteStatusReqID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteStatusReqID() (*field.QuoteStatusReqID, error) {
	f := new(field.QuoteStatusReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteReqID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteID is a required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteRespID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteRespID() (*field.QuoteRespID, error) {
	f := new(field.QuoteRespID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//Side is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlDate2() (*field.SettlDate2, error) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoQuoteQualifiers is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) NoQuoteQualifiers() (*field.NoQuoteQualifiers, error) {
	f := new(field.NoQuoteQualifiers)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BidPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidPx() (*field.BidPx, error) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}

//OfferPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferPx() (*field.OfferPx, error) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}

//MktBidPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MktBidPx() (*field.MktBidPx, error) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}

//MktOfferPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MktOfferPx() (*field.MktOfferPx, error) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}

//MinBidSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MinBidSize() (*field.MinBidSize, error) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}

//BidSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidSize() (*field.BidSize, error) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}

//MinOfferSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MinOfferSize() (*field.MinOfferSize, error) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//OfferSize is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferSize() (*field.OfferSize, error) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}

//ValidUntilTime is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//BidSpotRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidSpotRate() (*field.BidSpotRate, error) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//OfferSpotRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferSpotRate() (*field.OfferSpotRate, error) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//BidForwardPoints is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidForwardPoints() (*field.BidForwardPoints, error) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//OfferForwardPoints is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferForwardPoints() (*field.OfferForwardPoints, error) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//MidPx is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MidPx() (*field.MidPx, error) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}

//BidYield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidYield() (*field.BidYield, error) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}

//MidYield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) MidYield() (*field.MidYield, error) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}

//OfferYield is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferYield() (*field.OfferYield, error) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//OrdType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//BidForwardPoints2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) BidForwardPoints2() (*field.BidForwardPoints2, error) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//OfferForwardPoints2 is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) OfferForwardPoints2() (*field.OfferForwardPoints2, error) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrBidFxRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, error) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrOfferFxRate is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, error) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//ExDestination is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//QuoteStatus is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) QuoteStatus() (*field.QuoteStatus, error) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//ExDestinationIDSource is a non-required field for QuoteStatusReport.
func (m *QuoteStatusReport) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
