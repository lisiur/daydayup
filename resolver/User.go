package resolver

import (
	context "context"
	"fmt"

	models "github.com/lisiur/daydayup/models"
)

// CreatedAt .
func (r *UserResolver) CreatedAt(ctx context.Context, obj *models.User) (*string, error) {
	createdAt := fmt.Sprint(obj.CreatedAt)
	return &createdAt, nil
}

// UpdatedAt .
func (r *UserResolver) UpdatedAt(ctx context.Context, obj *models.User) (*string, error) {
	updatedAt := fmt.Sprint(obj.UpdatedAt)
	return &updatedAt, nil
}

// DeletedAt .
func (r *UserResolver) DeletedAt(ctx context.Context, obj *models.User) (*string, error) {
	if obj.DeletedAt != nil {
		deletedAt := fmt.Sprint(*obj.DeletedAt)
		return &deletedAt, nil
	}
	return nil, nil
}
