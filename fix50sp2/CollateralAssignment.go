package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralAssignment msg type = AY.
type CollateralAssignment struct {
	message.Message
}

//CollateralAssignmentBuilder builds CollateralAssignment messages.
type CollateralAssignmentBuilder struct {
	message.MessageBuilder
}

//NewCollateralAssignmentBuilder returns an initialized CollateralAssignmentBuilder with specified required fields.
func NewCollateralAssignmentBuilder(
	collasgnid field.CollAsgnID,
	collasgnreason field.CollAsgnReason,
	collasgntranstype field.CollAsgnTransType,
	transacttime field.TransactTime) *CollateralAssignmentBuilder {
	builder := new(CollateralAssignmentBuilder)
	builder.Body.Set(collasgnid)
	builder.Body.Set(collasgnreason)
	builder.Body.Set(collasgntranstype)
	builder.Body.Set(transacttime)
	return builder
}

//CollAsgnID is a required field for CollateralAssignment.
func (m *CollateralAssignment) CollAsgnID() (*field.CollAsgnID, error) {
	f := new(field.CollAsgnID)
	err := m.Body.Get(f)
	return f, err
}

//CollReqID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CollReqID() (*field.CollReqID, error) {
	f := new(field.CollReqID)
	err := m.Body.Get(f)
	return f, err
}

//CollAsgnReason is a required field for CollateralAssignment.
func (m *CollateralAssignment) CollAsgnReason() (*field.CollAsgnReason, error) {
	f := new(field.CollAsgnReason)
	err := m.Body.Get(f)
	return f, err
}

//CollAsgnTransType is a required field for CollateralAssignment.
func (m *CollateralAssignment) CollAsgnTransType() (*field.CollAsgnTransType, error) {
	f := new(field.CollAsgnTransType)
	err := m.Body.Get(f)
	return f, err
}

//CollAsgnRefID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CollAsgnRefID() (*field.CollAsgnRefID, error) {
	f := new(field.CollAsgnRefID)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for CollateralAssignment.
func (m *CollateralAssignment) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrderID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//NoTrades is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoTrades() (*field.NoTrades, error) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutAmount is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ValuationMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//FlowScheduleType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//RestructuringType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//Seniority is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//AttachmentPoint is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//DetachmentPoint is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//NoComplexEvents is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//Quantity is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//MarginExcess is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) MarginExcess() (*field.MarginExcess, error) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}

//TotalNetValue is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) TotalNetValue() (*field.TotalNetValue, error) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}

//CashOutstanding is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) CashOutstanding() (*field.CashOutstanding, error) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//Side is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//NoMiscFees is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoMiscFees() (*field.NoMiscFees, error) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}

//Price is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//StartCash is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//EndCash is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//SettlDeliveryType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SettlDeliveryType() (*field.SettlDeliveryType, error) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbType is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbName is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//NoDlvyInst is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) NoDlvyInst() (*field.NoDlvyInst, error) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for CollateralAssignment.
func (m *CollateralAssignment) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
