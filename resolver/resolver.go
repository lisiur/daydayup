//go:generate gorunpkg github.com/99designs/gqlgen

package resolver

import (
	context "context"
	"fmt"
	"math/rand"

	"github.com/lisiur/daydayup/graph/generated"
	"github.com/lisiur/daydayup/models"
)

// Resolver .
type Resolver struct {
	todos []models.Todo
}

// Mutation .
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query .
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// Todo .
func (r *Resolver) Todo() generated.TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (models.Todo, error) {
	todo := models.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	return r.todos, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (models.User, error) {
	return models.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}
