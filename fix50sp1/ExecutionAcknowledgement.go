package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ExecutionAcknowledgement msg type = BN.
type ExecutionAcknowledgement struct {
	message.Message
}

//ExecutionAcknowledgementBuilder builds ExecutionAcknowledgement messages.
type ExecutionAcknowledgementBuilder struct {
	message.MessageBuilder
}

//NewExecutionAcknowledgementBuilder returns an initialized ExecutionAcknowledgementBuilder with specified required fields.
func NewExecutionAcknowledgementBuilder(
	orderid field.OrderID,
	execackstatus field.ExecAckStatus,
	execid field.ExecID,
	side field.Side) *ExecutionAcknowledgementBuilder {
	builder := new(ExecutionAcknowledgementBuilder)
	builder.Body.Set(orderid)
	builder.Body.Set(execackstatus)
	builder.Body.Set(execid)
	builder.Body.Set(side)
	return builder
}

//OrderID is a required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ExecAckStatus is a required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ExecAckStatus() (*field.ExecAckStatus, error) {
	f := new(field.ExecAckStatus)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//DKReason is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) DKReason() (*field.DKReason, error) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayAmount is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FuturesValuationMethod is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//LastQty is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//LastParPx is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) LastParPx() (*field.LastParPx, error) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}

//CumQty is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) CumQty() (*field.CumQty, error) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ExecutionAcknowledgement.
func (m *ExecutionAcknowledgement) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
