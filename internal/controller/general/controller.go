package general

import "context"

type Usecase interface {
	Healthy(c context.Context) error
}

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
