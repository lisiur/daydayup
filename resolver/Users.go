package resolver

import (
	"context"

	"github.com/lisiur/daydayup/models"
)

// Users .
func (r *QueryResolver) Users(ctx context.Context) ([]models.User, error) {
	user := models.User{}
	return user.All()
}
