package resolver

import (
	context "context"
	"fmt"

	"github.com/lisiur/daydayup/models"
)

// CreatedAt .
func (r *LoginUserResolver) CreatedAt(ctx context.Context, obj *models.LoginUser) (*string, error) {
	createdAt := fmt.Sprint(obj.CreatedAt)
	return &createdAt, nil
}

// UpdatedAt .
func (r *LoginUserResolver) UpdatedAt(ctx context.Context, obj *models.LoginUser) (*string, error) {
	updatedAt := fmt.Sprint(obj.UpdatedAt)
	return &updatedAt, nil
}

// DeletedAt .
func (r *LoginUserResolver) DeletedAt(ctx context.Context, obj *models.LoginUser) (*string, error) {
	if obj.DeletedAt != nil {
		deletedAt := fmt.Sprint(*obj.DeletedAt)
		return &deletedAt, nil
	}
	return nil, nil
}
