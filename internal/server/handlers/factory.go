package handlers

import (
	"github.com/alex0206/workplace-accounting/internal/pg"
	"github.com/alex0206/workplace-accounting/internal/repositories"
	"github.com/alex0206/workplace-accounting/internal/services"
)

// Factory  describe factory for handlers
type Factory struct {
	dbConn *pg.DB
}

// WorkPlaceHandler getting workplace handler
func (f *Factory) WorkPlaceHandler() *WorkplaceHandler {
	return NewWorkplaceHandler(services.NewWorkplaceService(repositories.NewWorkplaceRepository(f.dbConn)))
}

// NewFactory creates a new handler factory
func NewFactory(dbConn *pg.DB) *Factory {
	return &Factory{dbConn: dbConn}
}
