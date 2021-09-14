package repositories

import (
	"encoding/json"

	"github.com/alex0206/workplace-accounting/internal/model"
	"github.com/alex0206/workplace-accounting/internal/pg"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// WorkplaceRepository describe workplace repository
type WorkplaceRepository struct {
	db *pg.DB
}

// NewWorkplaceRepository create a new repository
func NewWorkplaceRepository(db *pg.DB) *WorkplaceRepository {
	return &WorkplaceRepository{db: db}
}

// Add adding or updating a workplace
func (r *WorkplaceRepository) Add(ctx context.Context, entity *model.WorkplaceInfo) error {
	data, err := json.Marshal(entity)
	if err != nil {
		return errors.Wrap(err, "error to marshal workplace info")
	}
	_, err = r.db.Conn.Exec(ctx, "INSERT INTO workplace(info) VALUES($1) ON CONFLICT (info) DO UPDATE SET updated_at=now()", data)
	return err
}

// Delete deleting a workplace
func (r *WorkplaceRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Conn.Exec(ctx, "DELETE FROM workplace WHERE id=$1", id)
	return err
}
