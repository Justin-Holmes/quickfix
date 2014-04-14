package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityListUpdateReport msg type = BK.
type SecurityListUpdateReport struct {
	message.Message
}

//SecurityListUpdateReportBuilder builds SecurityListUpdateReport messages.
type SecurityListUpdateReportBuilder struct {
	message.MessageBuilder
}

//NewSecurityListUpdateReportBuilder returns an initialized SecurityListUpdateReportBuilder with specified required fields.
func NewSecurityListUpdateReportBuilder() *SecurityListUpdateReportBuilder {
	builder := new(SecurityListUpdateReportBuilder)
	return builder
}

//SecurityReportID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityReportID() (*field.SecurityReportID, error) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityReqID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityRequestResult is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//TotNoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityUpdateAction is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, error) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//CorporateAction is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityListID() (*field.SecurityListID, error) {
	f := new(field.SecurityListID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListRefID is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityListRefID() (*field.SecurityListRefID, error) {
	f := new(field.SecurityListRefID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListDesc is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityListDesc() (*field.SecurityListDesc, error) {
	f := new(field.SecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityListDescLen is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLen, error) {
	f := new(field.EncodedSecurityListDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityListDesc is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) EncodedSecurityListDesc() (*field.EncodedSecurityListDesc, error) {
	f := new(field.EncodedSecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListType is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityListType() (*field.SecurityListType, error) {
	f := new(field.SecurityListType)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListTypeSource is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) SecurityListTypeSource() (*field.SecurityListTypeSource, error) {
	f := new(field.SecurityListTypeSource)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for SecurityListUpdateReport.
func (m *SecurityListUpdateReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
