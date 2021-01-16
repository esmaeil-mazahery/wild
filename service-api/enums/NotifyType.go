package enums

//NotifyType verify status of Notify   x
type NotifyType uint32

//Types verify status of member
const (
	NotifyTypeUnknown NotifyType = iota
	NotifyTypeFollow
	NotifyTypeLike
	NotifyTypeComment
)

//NotifyTypes list of Enum
var NotifyTypes = [...]string{
	"Unknown",
	"Follow",
	"Like",
	"Comment",
}

// String() function will return the english name
// that we want out constant Day be recognized as
func (notifyType NotifyType) String() string {
	return NotifyTypes[notifyType]
}
