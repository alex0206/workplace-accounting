package services

import (
	"context"

	"github.com/alex0206/workplace-accounting/internal/e"
	"github.com/alex0206/workplace-accounting/internal/model"
)

// WorkplaceRepository describe workplace repository
type WorkplaceRepository interface {
	Add(ctx context.Context, entity *model.WorkplaceInfo) error
	Delete(ctx context.Context, ID int) error
}

// WorkplaceService describe workplace service
type WorkplaceService struct {
	repository WorkplaceRepository
}

// NewWorkplaceService get workplace service
func NewWorkplaceService(r WorkplaceRepository) *WorkplaceService {
	return &WorkplaceService{repository: r}
}

// Add adds a new workplace
func (s WorkplaceService) Add(ctx context.Context, workplaceInfo *model.WorkplaceInfo) e.Error {
	if workplaceInfo == nil {
		return e.NewInternal("no workplace info provided")
	}

	if workplaceInfo.ComputerName == "" || workplaceInfo.IP == "" || workplaceInfo.Username == "" {
		return e.NewBadRequest("workplace must not have empty fields")
	}

	err := s.repository.Add(ctx, workplaceInfo)
	if err != nil {
		return e.NewInternalf("failed creating workplace: %s", err)
	}

	return nil
}

// Delete remove a workplace by id
func (s *WorkplaceService) Delete(ctx context.Context, id int) e.Error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return e.NewInternalf("failed deleting workplace: %s", err)
	}

	return nil
}
