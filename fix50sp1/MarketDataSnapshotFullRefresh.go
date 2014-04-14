package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataSnapshotFullRefresh msg type = W.
type MarketDataSnapshotFullRefresh struct {
	message.Message
}

//MarketDataSnapshotFullRefreshBuilder builds MarketDataSnapshotFullRefresh messages.
type MarketDataSnapshotFullRefreshBuilder struct {
	message.MessageBuilder
}

//NewMarketDataSnapshotFullRefreshBuilder returns an initialized MarketDataSnapshotFullRefreshBuilder with specified required fields.
func NewMarketDataSnapshotFullRefreshBuilder(
	nomdentries field.NoMDEntries) *MarketDataSnapshotFullRefreshBuilder {
	builder := new(MarketDataSnapshotFullRefreshBuilder)
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayAmount is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FuturesValuationMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//NetChgPrevDay is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDay, error) {
	f := new(field.NetChgPrevDay)
	err := m.Body.Get(f)
	return f, err
}

//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoMDEntries() (*field.NoMDEntries, error) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}

//ApplQueueDepth is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ApplQueueDepth() (*field.ApplQueueDepth, error) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}

//ApplQueueResolution is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ApplQueueResolution() (*field.ApplQueueResolution, error) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}

//MDReportID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MDReportID() (*field.MDReportID, error) {
	f := new(field.MDReportID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//MDBookType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MDBookType() (*field.MDBookType, error) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}

//MDFeedType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MDFeedType() (*field.MDFeedType, error) {
	f := new(field.MDFeedType)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//NoRoutingIDs is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//MDSubBookType is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MDSubBookType() (*field.MDSubBookType, error) {
	f := new(field.MDSubBookType)
	err := m.Body.Get(f)
	return f, err
}

//MarketDepth is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) MarketDepth() (*field.MarketDepth, error) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}

//TotNumReports is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) TotNumReports() (*field.TotNumReports, error) {
	f := new(field.TotNumReports)
	err := m.Body.Get(f)
	return f, err
}

//RefreshIndicator is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) RefreshIndicator() (*field.RefreshIndicator, error) {
	f := new(field.RefreshIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for MarketDataSnapshotFullRefresh.
func (m *MarketDataSnapshotFullRefresh) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
