package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralReport msg type = BA.
type CollateralReport struct {
	message.Message
}

//CollateralReportBuilder builds CollateralReport messages.
type CollateralReportBuilder struct {
	message.MessageBuilder
}

//NewCollateralReportBuilder returns an initialized CollateralReportBuilder with specified required fields.
func NewCollateralReportBuilder(
	collrptid field.CollRptID,
	collstatus field.CollStatus) *CollateralReportBuilder {
	builder := new(CollateralReportBuilder)
	builder.Body.Set(collrptid)
	builder.Body.Set(collstatus)
	return builder
}

//CollRptID is a required field for CollateralReport.
func (m *CollateralReport) CollRptID() (*field.CollRptID, error) {
	f := new(field.CollRptID)
	err := m.Body.Get(f)
	return f, err
}

//CollInquiryID is a non-required field for CollateralReport.
func (m *CollateralReport) CollInquiryID() (*field.CollInquiryID, error) {
	f := new(field.CollInquiryID)
	err := m.Body.Get(f)
	return f, err
}

//CollStatus is a required field for CollateralReport.
func (m *CollateralReport) CollStatus() (*field.CollStatus, error) {
	f := new(field.CollStatus)
	err := m.Body.Get(f)
	return f, err
}

//TotNumReports is a non-required field for CollateralReport.
func (m *CollateralReport) TotNumReports() (*field.TotNumReports, error) {
	f := new(field.TotNumReports)
	err := m.Body.Get(f)
	return f, err
}

//LastRptRequested is a non-required field for CollateralReport.
func (m *CollateralReport) LastRptRequested() (*field.LastRptRequested, error) {
	f := new(field.LastRptRequested)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralReport.
func (m *CollateralReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for CollateralReport.
func (m *CollateralReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for CollateralReport.
func (m *CollateralReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for CollateralReport.
func (m *CollateralReport) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrderID is a non-required field for CollateralReport.
func (m *CollateralReport) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralReport.
func (m *CollateralReport) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralReport.
func (m *CollateralReport) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for CollateralReport.
func (m *CollateralReport) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//NoTrades is a non-required field for CollateralReport.
func (m *CollateralReport) NoTrades() (*field.NoTrades, error) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for CollateralReport.
func (m *CollateralReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralReport.
func (m *CollateralReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for CollateralReport.
func (m *CollateralReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralReport.
func (m *CollateralReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralReport.
func (m *CollateralReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for CollateralReport.
func (m *CollateralReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for CollateralReport.
func (m *CollateralReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for CollateralReport.
func (m *CollateralReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralReport.
func (m *CollateralReport) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralReport.
func (m *CollateralReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for CollateralReport.
func (m *CollateralReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralReport.
func (m *CollateralReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for CollateralReport.
func (m *CollateralReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralReport.
func (m *CollateralReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralReport.
func (m *CollateralReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralReport.
func (m *CollateralReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for CollateralReport.
func (m *CollateralReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for CollateralReport.
func (m *CollateralReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralReport.
func (m *CollateralReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralReport.
func (m *CollateralReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralReport.
func (m *CollateralReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralReport.
func (m *CollateralReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralReport.
func (m *CollateralReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for CollateralReport.
func (m *CollateralReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralReport.
func (m *CollateralReport) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for CollateralReport.
func (m *CollateralReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralReport.
func (m *CollateralReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for CollateralReport.
func (m *CollateralReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralReport.
func (m *CollateralReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for CollateralReport.
func (m *CollateralReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralReport.
func (m *CollateralReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralReport.
func (m *CollateralReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralReport.
func (m *CollateralReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralReport.
func (m *CollateralReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralReport.
func (m *CollateralReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for CollateralReport.
func (m *CollateralReport) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralReport.
func (m *CollateralReport) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for CollateralReport.
func (m *CollateralReport) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for CollateralReport.
func (m *CollateralReport) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for CollateralReport.
func (m *CollateralReport) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for CollateralReport.
func (m *CollateralReport) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralReport.
func (m *CollateralReport) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralReport.
func (m *CollateralReport) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralReport.
func (m *CollateralReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralReport.
func (m *CollateralReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralReport.
func (m *CollateralReport) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for CollateralReport.
func (m *CollateralReport) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralReport.
func (m *CollateralReport) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for CollateralReport.
func (m *CollateralReport) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralReport.
func (m *CollateralReport) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralReport.
func (m *CollateralReport) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralReport.
func (m *CollateralReport) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for CollateralReport.
func (m *CollateralReport) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for CollateralReport.
func (m *CollateralReport) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralReport.
func (m *CollateralReport) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for CollateralReport.
func (m *CollateralReport) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for CollateralReport.
func (m *CollateralReport) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralReport.
func (m *CollateralReport) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for CollateralReport.
func (m *CollateralReport) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for CollateralReport.
func (m *CollateralReport) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for CollateralReport.
func (m *CollateralReport) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for CollateralReport.
func (m *CollateralReport) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for CollateralReport.
func (m *CollateralReport) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for CollateralReport.
func (m *CollateralReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//Quantity is a non-required field for CollateralReport.
func (m *CollateralReport) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for CollateralReport.
func (m *CollateralReport) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for CollateralReport.
func (m *CollateralReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for CollateralReport.
func (m *CollateralReport) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralReport.
func (m *CollateralReport) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//MarginExcess is a non-required field for CollateralReport.
func (m *CollateralReport) MarginExcess() (*field.MarginExcess, error) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}

//TotalNetValue is a non-required field for CollateralReport.
func (m *CollateralReport) TotalNetValue() (*field.TotalNetValue, error) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}

//CashOutstanding is a non-required field for CollateralReport.
func (m *CollateralReport) CashOutstanding() (*field.CashOutstanding, error) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for CollateralReport.
func (m *CollateralReport) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//Side is a non-required field for CollateralReport.
func (m *CollateralReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//NoMiscFees is a non-required field for CollateralReport.
func (m *CollateralReport) NoMiscFees() (*field.NoMiscFees, error) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for CollateralReport.
func (m *CollateralReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for CollateralReport.
func (m *CollateralReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for CollateralReport.
func (m *CollateralReport) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for CollateralReport.
func (m *CollateralReport) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//StartCash is a non-required field for CollateralReport.
func (m *CollateralReport) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//EndCash is a non-required field for CollateralReport.
func (m *CollateralReport) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for CollateralReport.
func (m *CollateralReport) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CollateralReport.
func (m *CollateralReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for CollateralReport.
func (m *CollateralReport) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//SettlDeliveryType is a non-required field for CollateralReport.
func (m *CollateralReport) SettlDeliveryType() (*field.SettlDeliveryType, error) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbType is a non-required field for CollateralReport.
func (m *CollateralReport) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbName is a non-required field for CollateralReport.
func (m *CollateralReport) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbID is a non-required field for CollateralReport.
func (m *CollateralReport) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//NoDlvyInst is a non-required field for CollateralReport.
func (m *CollateralReport) NoDlvyInst() (*field.NoDlvyInst, error) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralReport.
func (m *CollateralReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralReport.
func (m *CollateralReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for CollateralReport.
func (m *CollateralReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralReport.
func (m *CollateralReport) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralReport.
func (m *CollateralReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for CollateralReport.
func (m *CollateralReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralReport.
func (m *CollateralReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for CollateralReport.
func (m *CollateralReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for CollateralReport.
func (m *CollateralReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//CollApplType is a non-required field for CollateralReport.
func (m *CollateralReport) CollApplType() (*field.CollApplType, error) {
	f := new(field.CollApplType)
	err := m.Body.Get(f)
	return f, err
}

//FinancialStatus is a non-required field for CollateralReport.
func (m *CollateralReport) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}
