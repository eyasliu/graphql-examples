package resolver

import (
	"context"
	"server/gql"
	"server/gql/gen"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type resolver struct{}

func (r *resolver) Mutation() gql.MutationResolver {
	return &mutationResolver{r}
}
func (r *resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, input *gen.UserBase) (*gen.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreatePost(ctx context.Context, input *gen.PostBase) (*gen.Post, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateComment(ctx context.Context, input *gen.CommentBase) (*gen.Comment, error) {
	panic("not implemented")
}

type queryResolver struct{ *resolver }

func (r *queryResolver) Me(ctx context.Context) (*gen.User, error) {
	panic("not implemented")
}
