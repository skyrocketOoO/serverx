package err

type Code string

func (c Code) String() string {
	return string(c)
}

const (
	EmptyRequest Code = "400.0001"

	NewPasswordRequired Code = "401.0001"

	Unknown        Code = "500"
	NotImplemented Code = "501"
)

var CodeToMsg = map[Code]string{
	EmptyRequest: "empty request body",

	NewPasswordRequired: "new password required",

	Unknown:        "unknown",
	NotImplemented: "not implemented",
}
