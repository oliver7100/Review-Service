package proto

import (
	"context"

	"github.com/oliver7100/review-service/database"
)

type service struct {
	UnimplementedReviewServiceServer
	Conn *database.Connection
}

func (s *service) CreateReview(ctx context.Context, req *CreateReviewRequest) (*CreateReviewResponse, error) {
	return nil, nil
}

func (s *service) GetReviewByTask(ctx context.Context, req *GetReviewsByTaskRequest) (*GetReviewsByTaskResponse, error) {
	return nil, nil
}

func CreateNewService(conn *database.Connection) *service {
	return &service{
		Conn: conn,
	}
}
