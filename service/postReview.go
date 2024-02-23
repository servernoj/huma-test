package service

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type ReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review message"`
	}
}

func (*Implementation) RegisterPostReview(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID:   "post-review",
		Summary:       "Post a review",
		Method:        http.MethodPost,
		Path:          "/reviews",
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, i *ReviewInput) (*struct{}, error) {
		// TODO: save review in data store.
		return nil, nil
	})
}
