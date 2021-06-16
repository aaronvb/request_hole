package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aaronvb/request_hole/graph/generated"
	"github.com/aaronvb/request_hole/pkg/protocol"
)

func (r *queryResolver) Requests(ctx context.Context) ([]protocol.RequestPayload, error) {
	return *r.RequestPayloads, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
