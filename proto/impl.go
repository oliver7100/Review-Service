package proto

import "context"

type IService interface {
	CreateReview(context.Context, CreateReviewRequest) (CreateReviewResponse, error)
	GetReviewsByTask(context.Context, GetReviewsByTaskRequest) (GetReviewsByTaskResponse, error)
}
