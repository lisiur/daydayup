//go:generate gorunpkg github.com/99designs/gqlgen

package resolver

import (
	generated "github.com/lisiur/daydayup/graph/generated"
)

// Resolver .
type Resolver struct{}

// QueryResolver .
type QueryResolver struct{ *Resolver }

// MutationResolver .
type MutationResolver struct{ *Resolver }

// UserResolver .
type UserResolver struct{ *Resolver }

// LoginUserResolver .
type LoginUserResolver struct{ *Resolver }

// Mutation .
func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolver{r}
}

// Query .
func (r *Resolver) Query() generated.QueryResolver {
	return &QueryResolver{r}
}

// User .
func (r *Resolver) User() generated.UserResolver {
	return &UserResolver{r}
}

// LoginUser .
func (r *Resolver) LoginUser() generated.LoginUserResolver {
	return &LoginUserResolver{r}
}
