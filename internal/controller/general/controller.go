package general

import "context"

type usecase interface {
	Healthy(c context.Context) error
}

type Handler struct {
	usecase usecase
}

func NewHandler(usecase usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
