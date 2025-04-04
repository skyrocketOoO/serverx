package domain

const (
	ApiVersion = "v1"
)

var (
	// env | flag
	Database    string
	AutoMigrate bool = false
	Env         string
)

type ErrResp struct {
	ID    string `json:"id"`
	Error string `json:"error"`
}
