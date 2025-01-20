package cm

type Pager struct {
	Number int `json:"number" validate:"required,min=1" example:"1"`
	Size   int `json:"size" validate:"required,min=1" example:"10"`
}

type Sorter struct {
	Field string `json:"field" validate:"required" example:"Time"`
	Asc   bool   `json:"asc" validate:"required" example:"false"`
}
