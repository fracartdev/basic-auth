package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fracartdev/basic-auth/graph/generated"
	"github.com/fracartdev/basic-auth/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	id := uuid.New()
	fmt.Println("<id: ", id, "username: ", input.Username, "password: ", input.Password, "email: ", input.Email)

	var err error

	return &model.User{
		ID:       id.String(),
		Username: input.Username,
		Password: input.Password,
		Email:    input.Email,
	}, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	fmt.Println("<Utente non trovato perche` non esiste un DB>")
	var err error
	var users []*model.User
	return users, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
