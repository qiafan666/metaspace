package request

import "context"

type BaseRequest struct {
	Ctx      context.Context `json:"ctx"`
	Language string          `json:"language"`
}

type BasePagination struct {
	CurrentPage int `json:"current_page" validate:"required,min=1"`
	PageCount   int `json:"page_count" validate:"required,max=50"`
}
