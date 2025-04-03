package user

type usecase interface{}

type Handler struct {
	usecase usecase
}

func NewHandler(usecase usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}
