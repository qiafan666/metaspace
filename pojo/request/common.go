package request

import "context"

type BaseRequest struct {
	Ctx      context.Context `json:"ctx"`
	Language string          `json:"language"`
}
