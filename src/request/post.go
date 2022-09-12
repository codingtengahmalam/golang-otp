package request

type (
	PostRequest struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"body"  validate:"required"`
	}
)
