package resolver

import (
	"context"

	"github.com/lisiur/daydayup/models"
)

// CreateUser .
func (r *MutationResolver) CreateUser(ctx context.Context, input models.Register) (models.User, error) {
	var user = models.User{}
	if err := user.Create(input); err != nil {
		return user, err
	}
	return user, nil
}
