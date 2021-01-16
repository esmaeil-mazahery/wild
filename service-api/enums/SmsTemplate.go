package enums

//SmsTemplate verify status of member
type SmsTemplate uint32

//Types verify status of member
const (
	SmsTemplateUnknown SmsTemplate = iota
	SmsTemplateRegister
	SmsTemplateForgetpassword
	SmsTemplateVerifyMobile
)

//SmsTemplates list of Enum
var SmsTemplates = [...]string{
	"Unknown",
	"Register",
	"Forgetpassword",
}

// String() function will return the english name
// that we want out constant Day be recognized as
func (smsTemplate SmsTemplate) String() string {
	return SmsTemplates[smsTemplate]
}
