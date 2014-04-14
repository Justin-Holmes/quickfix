package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationInstructionAlert msg type = BM.
type AllocationInstructionAlert struct {
	message.Message
}

//AllocationInstructionAlertBuilder builds AllocationInstructionAlert messages.
type AllocationInstructionAlertBuilder struct {
	message.MessageBuilder
}

//NewAllocationInstructionAlertBuilder returns an initialized AllocationInstructionAlertBuilder with specified required fields.
func NewAllocationInstructionAlertBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	alloctype field.AllocType,
	side field.Side,
	quantity field.Quantity,
	tradedate field.TradeDate) *AllocationInstructionAlertBuilder {
	builder := new(AllocationInstructionAlertBuilder)
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(alloctype)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(tradedate)
	return builder
}

//AllocID is a required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocTransType is a required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//AllocType is a required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocType() (*field.AllocType, error) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//RefAllocID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) RefAllocID() (*field.RefAllocID, error) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocCancReplaceReason is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocCancReplaceReason() (*field.AllocCancReplaceReason, error) {
	f := new(field.AllocCancReplaceReason)
	err := m.Body.Get(f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocIntermedReqType() (*field.AllocIntermedReqType, error) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocLinkID() (*field.AllocLinkID, error) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocLinkType() (*field.AllocLinkType, error) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//BookingRefID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BookingRefID() (*field.BookingRefID, error) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}

//AllocNoOrdersType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AllocNoOrdersType() (*field.AllocNoOrdersType, error) {
	f := new(field.AllocNoOrdersType)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//PreviouslyReported is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//ReversalIndicator is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ReversalIndicator() (*field.ReversalIndicator, error) {
	f := new(field.ReversalIndicator)
	err := m.Body.Get(f)
	return f, err
}

//MatchType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayAmount is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FuturesValuationMethod is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//Quantity is a required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradeOriginationDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//AvgParPx is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AvgParPx() (*field.AvgParPx, error) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxPrecision is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AvgPxPrecision() (*field.AvgPxPrecision, error) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//BookingType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//Concession is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//TotalTakedown is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//AutoAcceptIndicator is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AutoAcceptIndicator() (*field.AutoAcceptIndicator, error) {
	f := new(field.AutoAcceptIndicator)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NumDaysInterest is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//TotalAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, error) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//InterestAtMaturity is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) InterestAtMaturity() (*field.InterestAtMaturity, error) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//StartCash is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//EndCash is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//LegalConfirm is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//NoPosAmt is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//TotNoAllocs is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TotNoAllocs() (*field.TotNoAllocs, error) {
	f := new(field.TotNoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxIndicator is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//TrdType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//TrdSubType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//CustOrderCapacity is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//TradeInputSource is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) TradeInputSource() (*field.TradeInputSource, error) {
	f := new(field.TradeInputSource)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//MessageEventSource is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//RndPx is a non-required field for AllocationInstructionAlert.
func (m *AllocationInstructionAlert) RndPx() (*field.RndPx, error) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}
