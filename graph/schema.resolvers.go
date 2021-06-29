package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aaronvb/request_hole/graph/generated"
	"github.com/aaronvb/request_hole/graph/model"
	"github.com/aaronvb/request_hole/pkg/protocol"
)

func (r *mutationResolver) ClearRequests(ctx context.Context) (bool, error) {
	reqs := *r.RequestPayloads
	*r.RequestPayloads = reqs[:0]
	return true, nil
}

func (r *queryResolver) Requests(ctx context.Context) ([]*protocol.RequestPayload, error) {
	return *r.RequestPayloads, nil
}

func (r *queryResolver) ServerInfo(ctx context.Context) (*model.ServerInfo, error) {
	return r.Info, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
