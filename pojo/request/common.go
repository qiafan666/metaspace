package request

import "context"

type BaseRequest struct {
	Ctx      context.Context `json:"ctx"`
	Language string          `json:"language"`
}

type BasePagination struct {
	CurrentPage  int `json:"currentPage" validate:"required,min=1"`
	PrePageCount int `json:"prePageCount" validate:"required,max=50"`
}
