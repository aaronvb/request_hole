package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/aaronvb/request_hole/graph/generated"
	"github.com/aaronvb/request_hole/graph/model"
	"github.com/aaronvb/request_hole/pkg/protocol"
	"github.com/google/uuid"
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

func (r *subscriptionResolver) Request(ctx context.Context) (<-chan *protocol.RequestPayload, error) {
	// Generate UUID for browser connection
	id := uuid.New().String()

	// Go routine to handle deleting channel if browser connection is closed.
	go func() {
		<-ctx.Done()
		r.mu.Lock()
		delete(*r.RequestPayloadObserver, id)
		r.mu.Unlock()
	}()

	requests := make(chan *protocol.RequestPayload, 1)
	r.mu.Lock()
	(*r.RequestPayloadObserver)[id] = requests
	r.mu.Unlock()

	return requests, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
